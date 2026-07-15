package umeng

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/ent/userdevice"
	"github.com/woocoos/msgcenter/notify"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/pkg/profile"
	"github.com/woocoos/msgcenter/template"
	"go.uber.org/zap"
)

var logger = log.Component("umeng")

const (
	// DefaultAPIURL is the default Umeng Push API endpoint.
	DefaultAPIURL = "https://msgapi.umeng.com/api/send"

	// Push types.
	TypeUnicast   = "unicast"
	TypeListcast  = "listcast"
	TypeGroupcast = "groupcast"
	TypeBroadcast = "broadcast"

	// App names.
	AppAndroid   = "android"
	AppIOS       = "ios"
	AppHarmonyOS = "harmonyos"
)

// Request is the Umeng Push API request body.
type Request struct {
	AppKey       string         `json:"app_key"`
	Timestamp    string         `json:"timestamp"`
	Sign         string         `json:"sign"`
	Type         string         `json:"type"`
	DeviceTokens string         `json:"device_tokens,omitempty"`
	Alias        string         `json:"alias,omitempty"`
	AliasType    string         `json:"alias_type,omitempty"`
	Filter       map[string]any `json:"filter,omitempty"`
	Payload      map[string]any `json:"payload"`
	Policy       map[string]any `json:"policy,omitempty"`
	Production   *bool          `json:"production_mode,omitempty"`
	Description  string         `json:"description,omitempty"`
	// Android-specific fields.
	Category          int            `json:"category,omitempty"`
	ChannelProperties map[string]any `json:"channel_properties,omitempty"`
	LocalProperties   map[string]any `json:"local_properties,omitempty"`
	CallbackParams    map[string]any `json:"callback_params,omitempty"`
}

// Response is the Umeng Push API response body.
type Response struct {
	Ret       string `json:"ret"`
	ErrorCode string `json:"error_code,omitempty"`
	ErrorMsg  string `json:"error_msg,omitempty"`
	Data      struct {
		MsgID string `json:"msg_id"`
	} `json:"data,omitempty"`
}

// Message wraps the template data for Umeng payload rendering.
type Message struct {
	*template.Data

	GroupKey        string   `json:"groupKey"`
	TruncatedAlerts uint64   `json:"truncatedAlerts"`
	DeviceTokens    []string `json:"deviceTokens,omitempty"`
	UserIDs         []string `json:"userIDs,omitempty"`
}

// Notifier sends notifications via Umeng Push API.
type Notifier struct {
	config        *profile.UmengConfig
	tmpl          *template.Template
	customTplFunc notify.CustomerConfigFunc[profile.UmengConfig]
	client        *http.Client
	retrier       *notify.Retrier
	db            *ent.Client
}

// New returns a new Umeng notifier.
func New(cfg *profile.UmengConfig, tmpl *template.Template,
	fn notify.CustomerConfigFunc[profile.UmengConfig], db *ent.Client,
) (*Notifier, error) {
	nf := &Notifier{
		config:        cfg,
		tmpl:          tmpl,
		customTplFunc: fn,
		db:            db,
		retrier: &notify.Retrier{
			CustomDetailsFunc: func(_ int, body io.Reader) string {
				return errDetails(body)
			},
		},
	}
	if cfg.HTTPConfig == nil {
		nf.client = http.DefaultClient
	} else {
		httpClient, err := cfg.HTTPConfig.Client(context.Background(), nil)
		if err != nil {
			return nil, err
		}
		nf.client = httpClient
	}
	return nf, nil
}

func (n *Notifier) SendResolved() bool {
	return n.config.SendResolved
}

// CustomConfig returns a custom config for the notifier.
func (n *Notifier) CustomConfig(ctx context.Context) (*profile.UmengConfig, error) {
	if n.customTplFunc == nil {
		return n.config, nil
	}
	labels, ok := notify.GroupLabels(ctx)
	if !ok {
		return n.config, nil
	}
	cfg := n.config.Clone()
	err := n.customTplFunc(ctx, cfg, labels)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// Notify implements the Notifier interface.
func (n *Notifier) Notify(ctx context.Context, alerts ...*alert.Alert) (bool, error) {
	data := notify.GetTemplateData(ctx, n.tmpl, alerts)

	config, err := n.CustomConfig(ctx)
	if err != nil {
		return false, err
	}

	groupKey, err := notify.ExtractGroupKey(ctx)
	if err != nil {
		return false, err
	}

	msg := &Message{
		Data:     data,
		GroupKey: groupKey.String(),
	}

	// If apps are configured, query user devices and send per-app.
	if len(config.Apps) > 0 {
		return n.notifyMultiApp(ctx, config, msg)
	}

	// Single app mode: use first app credentials.
	return n.notifySingleApp(ctx, config, msg, "")
}

// notifyMultiApp sends push notifications to multiple apps based on user devices.
func (n *Notifier) notifyMultiApp(ctx context.Context, config *profile.UmengConfig, msg *Message) (bool, error) {
	// Extract app set from alert labels.
	appSet := extractAppSet(msg)

	// Extract user IDs from alerts.
	userIDs := extractUserIDs(msg)
	if len(userIDs) == 0 {
		// No user IDs, fall back to broadcast on all apps (or filtered by appSet).
		var lastErr error
		shouldRetry := false
		for app, appConfig := range config.Apps {
			// Filter by appSet if specified.
			if appSet != "" && appConfig.AppSet != appSet {
				continue
			}
			retry, err := n.notifySingleApp(ctx, config, msg, app)
			if err != nil {
				lastErr = err
				shouldRetry = shouldRetry || retry
				logger.Warn("umeng push failed", zap.String("app", app), zap.Error(err))
			}
		}
		return shouldRetry, lastErr
	}

	// Query user devices to determine which apps to push to.
	appUsers, err := n.queryUserApps(ctx, userIDs)
	if err != nil {
		return true, fmt.Errorf("querying user devices: %w", err)
	}

	// Log users without devices.
	for _, uid := range userIDs {
		if _, ok := appUsers[uid]; !ok {
			logger.Warn("user has no device registered, skipping", zap.String("user_id", uid))
		}
	}

	// Group users by app based on device platform and appSet.
	appUserMap := make(map[string][]string)
	for uid, devices := range appUsers {
		for _, dev := range devices {
			platform := normalizePlatform(dev.SystemName)
			if platform == "" {
				continue
			}
			// Find the app that matches this platform and appSet.
			for appName, appConfig := range config.Apps {
				if appConfig.Platform == platform {
					// Filter by appSet if specified.
					if appSet != "" && appConfig.AppSet != appSet {
						continue
					}
					// Add user ID to this app's list (avoid duplicates).
					found := false
					for _, existingUID := range appUserMap[appName] {
						if existingUID == uid {
							found = true
							break
						}
					}
					if !found {
						appUserMap[appName] = append(appUserMap[appName], uid)
					}
					break
				}
			}
		}
	}

	// Send push to each app using alias (user IDs).
	var lastErr error
	successCount := 0
	shouldRetry := false
	for app, uids := range appUserMap {
		if _, ok := config.Apps[app]; !ok {
			logger.Warn("app not configured, skipping", zap.String("app", app))
			continue
		}
		// Create a message with user IDs as alias.
		appMsg := n.filterMessageForUsers(msg, uids)
		retry, err := n.notifySingleApp(ctx, config, appMsg, app)
		if err != nil {
			lastErr = err
			shouldRetry = shouldRetry || retry
			logger.Warn("umeng push failed", zap.String("app", app), zap.Error(err))
		} else {
			successCount++
		}
	}

	if successCount == 0 && len(appUserMap) > 0 {
		return shouldRetry, lastErr
	}
	return false, nil
}

// notifySingleApp sends a push notification to a single app.
func (n *Notifier) notifySingleApp(ctx context.Context, config *profile.UmengConfig, msg *Message, app string) (bool, error) {
	// Build the Umeng request.
	req, err := n.buildRequestForApp(config, msg, app)
	if err != nil {
		return false, err
	}

	// Determine API URL.
	apiURL := config.APIURL
	if apiURL == "" {
		apiURL = DefaultAPIURL
	}

	// Encode request body.
	body, err := json.Marshal(req)
	if err != nil {
		return false, fmt.Errorf("encoding umeng request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, apiURL, bytes.NewReader(body))
	if err != nil {
		return true, err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := n.client.Do(httpReq)
	if err != nil {
		return true, redactURLError(err)
	}
	defer drain(resp)

	// Parse Umeng response.
	var umengResp Response
	if err := json.NewDecoder(resp.Body).Decode(&umengResp); err != nil {
		return true, fmt.Errorf("decoding umeng response: %w", err)
	}

	if umengResp.Ret != "SUCCESS" {
		return false, fmt.Errorf("umeng push failed: [%s] %s", umengResp.ErrorCode, umengResp.ErrorMsg)
	}

	return false, nil
}

// buildRequestForApp constructs the Umeng Push API request for a specific app.
func (n *Notifier) buildRequestForApp(config *profile.UmengConfig, msg *Message, app string) (*Request, error) {
	// Get platform from app config.
	var platform string
	if app != "" {
		if ac, ok := config.Apps[app]; ok {
			platform = ac.Platform
		}
	}

	// Dispatch to platform-specific builder.
	switch platform {
	case AppIOS:
		return n.buildRequestForIOS(config, msg, app)
	case AppHarmonyOS:
		return n.buildRequestForHarmonyOS(config, msg, app)
	default:
		// Default to Android for unknown platforms.
		return n.buildRequestForAndroid(config, msg, app)
	}
}

// buildRequestCommon builds common request fields shared across all platforms.
func (n *Notifier) buildRequestCommon(config *profile.UmengConfig, msg *Message, app string) (*Request, string, error) {
	// Resolve app credentials.
	appKey, appSecret := n.resolveCredentialsForApp(config, app)
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	sign := md5Hex(appKey + timestamp + appSecret)

	// Determine push type based on user IDs.
	// Prefer msg.UserIDs (from multi-app mode) over extracted userIDs.
	targetUserIDs := msg.UserIDs
	if len(targetUserIDs) == 0 {
		targetUserIDs = extractUserIDs(msg)
	}
	pushType := TypeBroadcast

	// Override push type based on user count.
	// When using alias (user IDs), use customizedcast instead of unicast/listcast.
	switch len(targetUserIDs) {
	case 0:
		// No users: use broadcast.
	default:
		// Use customizedcast when targeting users by alias.
		pushType = "customizedcast"
	}

	req := &Request{
		AppKey:     appKey,
		Timestamp:  timestamp,
		Sign:       sign,
		Type:       pushType,
		Production: config.ProductionMode,
	}

	// Set target based on push type and user IDs.
	// If device tokens are provided, use them directly.
	if len(msg.DeviceTokens) > 0 {
		req.DeviceTokens = strings.Join(msg.DeviceTokens, ",")
		if len(msg.DeviceTokens) == 1 {
			req.Type = TypeUnicast
		} else {
			req.Type = TypeListcast
		}
	} else if len(targetUserIDs) > 0 {
		// Use alias (user IDs) for customizedcast targeting.
		req.Alias = strings.Join(targetUserIDs, ",")
		// Get alias_type from app config, default to "uid".
		if app != "" {
			if ac, ok := config.Apps[app]; ok && ac.AliasType != "" {
				req.AliasType = ac.AliasType
			}
		}
		if req.AliasType == "" {
			req.AliasType = "uid"
		}
	}

	return req, appKey + timestamp + appSecret, nil
}

// buildRequestForAndroid builds Android-specific request.
// Android has additional fields: category, channel_properties, local_properties.
func (n *Notifier) buildRequestForAndroid(config *profile.UmengConfig, msg *Message, app string) (*Request, error) {
	req, _, err := n.buildRequestCommon(config, msg, app)
	if err != nil {
		return nil, err
	}

	// Build Android-specific payload inline.
	// Android uses a flat body structure with title, text, ticker, etc.
	title := n.renderString(`{{ template "umeng.default.title" . }}`, msg)
	text := n.renderString(`{{ template "umeng.default.text" . }}`, msg)
	body := map[string]any{
		"title":  title,
		"ticker": title,
		"text":   text,
	}

	// Android-specific: after_open and activity from app config.
	if app != "" {
		if ac, ok := config.Apps[app]; ok {
			if ac.AfterOpen != "" {
				body["after_open"] = ac.AfterOpen
			}
			if ac.Activity != "" {
				body["activity"] = ac.Activity
			}
		}
	}

	req.Payload = map[string]any{
		"display_type": "notification",
		"body":         body,
	}

	// Android-specific: category for message classification.
	// 0: 资讯营销类消息, 1: 服务与通讯类消息
	req.Category = 1

	return req, nil
}

// buildRequestForIOS builds iOS-specific request.
// iOS follows APNs standard and has different policy fields.
func (n *Notifier) buildRequestForIOS(config *profile.UmengConfig, msg *Message, app string) (*Request, error) {
	req, _, err := n.buildRequestCommon(config, msg, app)
	if err != nil {
		return nil, err
	}

	// Build iOS-specific payload inline.
	// iOS follows APNs standard with aps dictionary structure.
	title := n.renderString(`{{ template "umeng.default.title" . }}`, msg)
	text := n.renderString(`{{ template "umeng.default.text" . }}`, msg)
	req.Payload = map[string]any{
		"display_type": "notification",
		"aps": map[string]any{
			"alert": map[string]any{
				"title": title,
				"body":  text,
			},
			"sound": "default",
		},
	}

	// iOS does not have category, channel_properties, or local_properties.
	// iOS policy can have apns_collapse_id for message collapsing.

	return req, nil
}

// buildRequestForHarmonyOS builds HarmonyOS-specific request.
// HarmonyOS has channel_properties but different from Android.
func (n *Notifier) buildRequestForHarmonyOS(config *profile.UmengConfig, msg *Message, app string) (*Request, error) {
	req, _, err := n.buildRequestCommon(config, msg, app)
	if err != nil {
		return nil, err
	}

	// Build HarmonyOS-specific payload inline.
	// HarmonyOS uses a structure similar to Android but with some differences.
	title := n.renderString(`{{ template "umeng.default.title" . }}`, msg)
	text := n.renderString(`{{ template "umeng.default.text" . }}`, msg)
	req.Payload = map[string]any{
		"display_type": "notification",
		"body": map[string]any{
			"title": title,
			"body":  text,
		},
	}

	// HarmonyOS-specific: channel_properties with harmony_channel_category.
	req.ChannelProperties = map[string]any{
		"harmony_channel_category": "MARKETING",
	}

	// HarmonyOS policy can have channel_strategy.
	req.Policy = map[string]any{
		"channel_strategy": map[string]any{
			"default": 2, // 在线时通过友盟通道下发，离线尝试通过厂商下发
		},
	}

	return req, nil
}

// renderString renders a Go template string with the message data.
func (n *Notifier) renderString(tpl string, msg *Message) string {
	if tpl == "" {
		return ""
	}
	result, err := n.tmpl.ExecuteTextString(tpl, msg)
	if err != nil {
		return tpl
	}
	return result
}

func md5Hex(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func truncateAlerts(maxAlerts uint64, alerts []*alert.Alert) ([]*alert.Alert, uint64) {
	if maxAlerts != 0 && uint64(len(alerts)) > maxAlerts {
		return alerts[:maxAlerts], uint64(len(alerts)) - maxAlerts
	}
	return alerts, 0
}

func errDetails(body io.Reader) string {
	if body == nil {
		return ""
	}
	bs, err := io.ReadAll(body)
	if err != nil {
		return ""
	}
	return string(bs)
}

func drain(resp *http.Response) {
	io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
}

func redactURLError(err error) error {
	var urlErr *url.Error
	if errors.As(err, &urlErr) {
		urlErr.URL = "<redacted>"
	}
	return err
}

// extractUserIDs extracts unique user IDs from alert labels.
func extractUserIDs(msg *Message) []string {
	if msg == nil || msg.Data == nil {
		return nil
	}
	seen := make(map[string]struct{})
	var ids []string
	for _, a := range msg.Data.Alerts {
		if v, ok := a.Labels[string(label.ToUserIDLabel)]; ok && v != "" {
			for _, id := range strings.Split(v, ",") {
				id = strings.TrimSpace(id)
				if id == "" {
					continue
				}
				if _, dup := seen[id]; !dup {
					seen[id] = struct{}{}
					ids = append(ids, id)
				}
			}
		}
	}
	return ids
}

// resolveCredentialsForApp returns the appKey and appMasterSecret for the specified app.
func (n *Notifier) resolveCredentialsForApp(config *profile.UmengConfig, app string) (string, string) {
	if app != "" {
		if ac, ok := config.Apps[app]; ok {
			return ac.AppKey, ac.AppMasterSecret
		}
	}
	// Fallback to first app if available.
	for _, ac := range config.Apps {
		return ac.AppKey, ac.AppMasterSecret
	}
	return "", ""
}

// queryUserApps queries user devices from the database and groups by user ID.
func (n *Notifier) queryUserApps(ctx context.Context, userIDs []string) (map[string][]*ent.UserDevice, error) {
	if n.db == nil {
		return nil, nil
	}
	// Convert user IDs to integers.
	uids := make([]int, 0, len(userIDs))
	for _, uid := range userIDs {
		if id, err := strconv.Atoi(uid); err == nil {
			uids = append(uids, id)
		}
	}
	if len(uids) == 0 {
		return nil, nil
	}

	devices, err := n.db.UserDevice.Query().
		Where(userdevice.UserIDIn(uids...)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	result := make(map[string][]*ent.UserDevice)
	for _, dev := range devices {
		uid := strconv.Itoa(dev.UserID)
		result[uid] = append(result[uid], dev)
	}
	return result, nil
}

// normalizePlatform normalizes system name to platform constant.
func normalizePlatform(systemName string) string {
	lower := strings.ToLower(systemName)
	switch {
	case strings.Contains(lower, "android"):
		return AppAndroid
	case strings.Contains(lower, "ios"), strings.Contains(lower, "iphone"), strings.Contains(lower, "ipad"):
		return AppIOS
	case strings.Contains(lower, "harmony"), strings.Contains(lower, "鸿蒙"):
		return AppHarmonyOS
	default:
		return ""
	}
}

// filterMessageForDevices creates a new message with only the specified device tokens.
func (n *Notifier) filterMessageForDevices(msg *Message, deviceTokens []string) *Message {
	return &Message{
		Data:            msg.Data,
		GroupKey:        msg.GroupKey,
		TruncatedAlerts: msg.TruncatedAlerts,
		DeviceTokens:    deviceTokens,
	}
}

// filterMessageForUsers creates a new message with only the specified user IDs.
func (n *Notifier) filterMessageForUsers(msg *Message, userIDs []string) *Message {
	return &Message{
		Data:            msg.Data,
		GroupKey:        msg.GroupKey,
		TruncatedAlerts: msg.TruncatedAlerts,
		UserIDs:         userIDs,
	}
}

// extractAppSet extracts the application set from alert labels.
// Returns the first non-empty appSet value found across all alerts.
func extractAppSet(msg *Message) string {
	if msg == nil || msg.Data == nil {
		return ""
	}
	for _, a := range msg.Data.Alerts {
		if v, ok := a.Labels[string(label.AppSetLabel)]; ok && v != "" {
			return strings.TrimSpace(v)
		}
	}
	return ""
}
