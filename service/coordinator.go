package service

import (
	"context"
	"crypto/md5"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/woocoos/entco/pkg/identity"
	"github.com/woocoos/entco/schemax/typex"
	"github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/ent/msgchannel"
	"github.com/woocoos/msgcenter/ent/msgevent"
	"github.com/woocoos/msgcenter/metrics"
	"github.com/woocoos/msgcenter/notify"
	"github.com/woocoos/msgcenter/notify/email"
	"github.com/woocoos/msgcenter/notify/message"
	"github.com/woocoos/msgcenter/notify/webhook"
	"github.com/woocoos/msgcenter/pkg/profile"
	"github.com/woocoos/msgcenter/template"
	"go.uber.org/zap"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

var logger = log.Component("config")

const (
	TempRelativePathTplTemp    = "tplTemp"
	TempRelativePathTplData    = "tplData"
	TempRelativePathAttachment = "attachment"
)

// Coordinator helps the Alert Manager collaborate with external components
type Coordinator struct {
	configuration *conf.Configuration
	// Protects profile and reloadHooks
	mutex       sync.RWMutex
	profile     *profile.Config
	reloadHooks []func(*profile.Config) error

	ActiveReceivers map[string]int // receiver name -> number of Notifier
	Template        *template.Template

	db        *ent.Client
	Subscribe *UserSubscribe
	// knockout http client
	KOClient *http.Client

	TempOptions TempOptions
}

type TempOptions struct {
	Path         string `yaml:"path"`
	FileBaseUrl  string `yaml:"fileBaseUrl"`
	RelativePath struct {
		TplTemp    string `yaml:"tplTemp"`
		TplData    string `yaml:"tplData"`
		Attachment string `yaml:"attachment"`
	} `yaml:"relativePath"`
}

// NewCoordinator returns a new coordinator with the given configuration for alert manager.
// It does not yet load the configuration from file. This is done in
// `Reload()`.
func NewCoordinator(cnf *conf.Configuration) *Coordinator {
	c := &Coordinator{
		configuration:   cnf,
		ActiveReceivers: make(map[string]int),
		Subscribe:       &UserSubscribe{},
	}

	tempOptions := TempOptions{}
	err := cnf.Sub("template").Unmarshal(&tempOptions)
	if err != nil {
		log.Fatal(err)
	}
	c.TempOptions = tempOptions
	return c
}

func (c *Coordinator) SetDBClient(db *ent.Client) {
	c.db = db
	c.Subscribe.DB = db
}

func (c *Coordinator) SetHttpClient(httpClient *http.Client) {
	c.KOClient = httpClient
}

func (c *Coordinator) ProfileString() string {
	return c.profile.String()
}

func (c *Coordinator) GetReceivers() []profile.Receiver {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.profile.Receivers
}

func (c *Coordinator) ResolveTimeout() time.Duration {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.profile.Global.ResolveTimeout
}

// ReloadHooks subscribes the given Subscribers to configuration changes.
func (c *Coordinator) ReloadHooks(ss ...func(*profile.Config) error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.reloadHooks = append(c.reloadHooks, ss...)
}

func (c *Coordinator) notifySubscribers() error {
	for _, s := range c.reloadHooks {
		if err := s(c.profile); err != nil {
			return err
		}
	}

	return nil
}

// Reload triggers a configuration reload from file and notifies all
// configuration change reloadHooks.
func (c *Coordinator) Reload() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	// clear active receivers,if error occurs, also clear it
	c.ActiveReceivers = make(map[string]int)

	configFilePath := c.configuration.String("config.file")
	logger.Info("loading configuration file", zap.String("file", configFilePath))

	if err := c.loadFromFile(); err != nil {
		logger.Error("Error loading file configuration file", zap.Error(err))
		metrics.Coordinator.ConfigSuccess.Set(0)
		return err
	}
	if err := c.loadFromDB(); err != nil {
		logger.Error("Error loading db configuration file", zap.Error(err))
		metrics.Coordinator.ConfigSuccess.Set(0)
		return err
	}
	logger.Info("completed loading configuration file", zap.String("file", configFilePath))

	if err := c.loadTemplates(); err != nil {
		logger.Error("Error loading template file", zap.Error(err))
		metrics.Coordinator.ConfigSuccess.Set(0)
		return err
	}

	if err := c.notifySubscribers(); err != nil {
		logger.Error("one or more config change reloadHooks failed to apply new config", zap.Error(err))
		metrics.Coordinator.ConfigSuccess.Set(0)
		return err
	}
	// Set metrics.
	var integrationsNum int
	for _, cs := range c.ActiveReceivers {
		integrationsNum += cs
	}
	metrics.Coordinator.ConfiguredReceivers.Set(float64(len(c.profile.Receivers)))
	metrics.Coordinator.ConfiguredIntegrations.Set(float64(integrationsNum))
	metrics.Coordinator.ConfigSuccess.Set(1)
	metrics.Coordinator.ConfigSuccessTime.SetToCurrentTime()

	return nil
}

func (c *Coordinator) WalkReceivers(visit func(receiver profile.Receiver) error) error {
	for _, rcv := range c.profile.Receivers {
		if _, found := c.ActiveReceivers[rcv.Name]; !found {
			// No need to build a receiver if no route is using it.
			logger.Info("skipping creation of receiver not referenced by any route", zap.String("receiver", rcv.Name))
			continue
		}
		if err := visit(rcv); err != nil {
			return err
		}
	}
	return nil
}

func (c *Coordinator) TempParseFiles(filenames ...string) error {
	_, err := c.Template.ParseFiles(filenames...)
	return err
}

// loadTemplates loading template files
func (c *Coordinator) loadTemplates() error {
	if tmpl, err := template.New(); err != nil {
		return fmt.Errorf("failed to parse templates %w", err)
	} else {
		c.Template = tmpl
		for _, ptn := range c.profile.Templates {
			if _, err := c.Template.ParseGlob(c.configuration.Abs(ptn)); err != nil {
				return err
			}
		}
	}
	return nil
}

// loadFromFile triggers a configuration load, discarding the old configuration.
func (c *Coordinator) loadFromFile() error {
	config, err := profile.NewConfig(c.configuration)
	if err != nil {
		return err
	}
	c.profile = config
	return nil
}

func (c *Coordinator) loadFromDB() error {
	if c.db == nil {
		return nil
	}

	// load receivers first
	cns, err := c.db.MsgChannel.Query().Where(msgchannel.StatusEQ(typex.SimpleStatusActive)).All(context.Background())
	if err != nil {
		return err
	}
	receivers := make([]*profile.Receiver, len(cns))
	for i, cn := range cns {
		cn.Receiver.Name = profile.TenantReceiverName(strconv.Itoa(cn.TenantID), cn.Receiver.Name)
		receivers[i] = cn.Receiver
	}
	// load routes
	evs, err := c.db.MsgEvent.Query().Where(msgevent.StatusEQ(typex.SimpleStatusActive)).WithMsgType().All(context.Background())
	if err != nil {
		return err
	}
	routes := make([]*profile.Route, 0, len(evs))
	for _, ev := range evs {
		if ev.Route == nil {
			continue
		}
		ev.Route.Name = profile.AppRouteName(strconv.Itoa(ev.Edges.MsgType.AppID), ev.Name)
		routes = append(routes, ev.Route)
	}
	if err = c.addTenantReceiver(receivers); err != nil {
		return err
	}
	if err = c.addNamedRoute(routes); err != nil {
		return err
	}
	return nil
}

func (c *Coordinator) AddNamedRoute(input []*profile.Route) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.addNamedRoute(input)
}

// AddNamedRoute adds a route to the current configuration. Notice: it is unsafe to call this method concurrently.
func (c *Coordinator) addNamedRoute(input []*profile.Route) error {
	rs, err := json.Marshal(input)
	if err != nil {
		return err
	}
	// use json validate. TODO: use other validate
	var routes []*profile.Route
	if err := json.Unmarshal(rs, &routes); err != nil {
		return err
	}
	// try to check
	vc, err := c.profile.DeepClone()
	if err != nil {
		return err
	}
	for _, route := range input {
		index := -1
		for i, r := range vc.Route.Routes {
			if r.Name == route.Name {
				vc.Route.Routes[i] = route
				index = i
			}
		}
		if index == -1 {
			vc.Route.Routes = append(vc.Route.Routes, route)
		}
		if err := vc.Validate(); err != nil {
			return err
		}
		if index != -1 {
			c.profile.Route.Routes[index] = route
		} else {
			c.profile.Route.Routes = append(c.profile.Route.Routes, route)
		}
	}
	return nil
}

func (c *Coordinator) RemoveNamedRoute(routeNames []string) error {
	rs := c.profile.Route.Routes
	for _, v := range routeNames {
		if v == "" {
			continue
		}
		nrs := make([]*profile.Route, 0)
		for _, r := range rs {
			if r.Name != v {
				nrs = append(nrs, r)
			}
		}
		rs = nrs
	}
	c.profile.Route.Routes = rs
	return nil
}

func (c *Coordinator) AddTenantReceiver(input []*profile.Receiver) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.addTenantReceiver(input)
}

func (c *Coordinator) addTenantReceiver(input []*profile.Receiver) error {
	rs, err := json.Marshal(input)
	if err != nil {
		return err
	}
	// use json validate. TODO: use other validate
	var receivers []*profile.Receiver
	if err := json.Unmarshal(rs, &receivers); err != nil {
		return err
	}
	// try to check
	vc, err := c.profile.DeepClone()
	if err != nil {
		return err
	}
	for _, receiver := range receivers {
		index := -1
		for i, r := range vc.Receivers {
			if r.Name == receiver.Name {
				vc.Receivers[i] = *receiver
				index = i
			}
		}
		if index == -1 {
			vc.Receivers = append(vc.Receivers, *receiver)
		}
		if err := vc.Validate(); err != nil {
			return err
		}
		if index != -1 {
			c.profile.Receivers[index] = *receiver
		} else {
			c.profile.Receivers = append(c.profile.Receivers, *receiver)
			// custom receiver default enable
			c.ActiveReceivers[receiver.Name] = 1
		}
	}
	return nil
}

func (c *Coordinator) RemoveTenantReceiver(receiverNames []string) error {
	rs := c.profile.Receivers
	for _, v := range receiverNames {
		if v == "" {
			continue
		}
		nrs := make([]profile.Receiver, 0)
		for _, r := range rs {
			if r.Name != v {
				nrs = append(nrs, r)
			}
		}
		rs = nrs
	}
	c.profile.Receivers = rs
	return nil
}

// buildReceiverIntegrations builds a list of integration notifiers off of a
// receiver config.
func (c *Coordinator) buildReceiverIntegrations(nc profile.Receiver, tmpl *template.Template) (integrations []notify.Integration, errs error) {
	add := func(name string, i int, f func() (notify.Notifier, error)) {
		n, err := f()
		if err != nil {
			errs = errors.Join(err)
			return
		}
		integrations = append(integrations, notify.NewIntegration(n, name, i))
	}
	var ()
	tpldir := c.configuration.Abs(c.TempOptions.Path)
	for i, cfg := range nc.EmailConfigs {
		add("email", i, func() (notify.Notifier, error) {
			return email.New(cfg, tmpl, overrideEmailConfig(tpldir, c.db))
		})
	}
	for i, cfg := range nc.WebhookConfigs {
		add("webhook", i, func() (notify.Notifier, error) {
			return webhook.New(cfg, tmpl, overrideWebHookConfig(tpldir, c.db))
		})
	}
	if nc.MessageConfig != nil {
		add("message", 0, func() (notify.Notifier, error) {
			return message.New(nc.MessageConfig, tmpl, c.db, overrideMessageConfig(tpldir, c.db))
		})
	}
	if errs != nil {
		return nil, errs
	}
	return integrations, nil
}

func md5HashAsMetricValue(data []byte) float64 {
	sum := md5.Sum(data)
	// We only want 48 bits as a float64 only has a 53 bit mantissa.
	smallSum := sum[0:6]
	bytes := make([]byte, 8)
	copy(bytes, smallSum)
	return float64(binary.LittleEndian.Uint64(bytes))
}

// ValidateFilePath 验证路径是否符合规则
// dir:tplData、tplTemp、attachment
func (c *Coordinator) ValidateFilePath(ctx context.Context, path, dir string) error {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return err
	}
	path = filepath.Join(path)
	p := strings.TrimPrefix(path, "/")
	rp := ""
	if dir == TempRelativePathTplData {
		rp = c.TempOptions.RelativePath.TplData
	} else if dir == TempRelativePathTplTemp {
		rp = c.TempOptions.RelativePath.TplTemp
	} else if dir == TempRelativePathAttachment {
		rp = c.TempOptions.RelativePath.Attachment
	}
	prefixPath := filepath.Join(rp, strconv.Itoa(tid))
	if !strings.HasPrefix(p, strings.TrimPrefix(prefixPath, "/")) {
		return fmt.Errorf("invalid path: %s,must be like:%s/xxx", path, prefixPath)
	}
	return nil
}

func (c *Coordinator) CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_CREATE|os.O_EXCL|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

// GetTplDataPath 将tpl临时文件路径转为正式存储路径
func (c *Coordinator) GetTplDataPath(tempPath string) string {
	tpldir := c.TempOptions.Path
	data := c.TempOptions.RelativePath.TplData
	temp := c.TempOptions.RelativePath.TplTemp
	return c.configuration.Abs(
		filepath.Join(
			tpldir,
			data,
			strings.TrimPrefix(
				strings.TrimPrefix(tempPath, "/"),
				strings.TrimPrefix(temp, "/"),
			),
		),
	)
}

// GetTplTempPath 获取tpl正式文件路径
func (c *Coordinator) GetTplTempPath(tempPath string) string {
	tpldir := c.configuration.Abs(c.TempOptions.Path)
	return c.configuration.Abs(filepath.Join(tpldir, tempPath))
}

// ReportFileRefCount 文件引用上报
func (c *Coordinator) ReportFileRefCount(ctx context.Context, newFileIDs, oldFileIDs []int) error {
	tid, err := identity.TenantIDFromContext(ctx)
	if err != nil {
		return err
	}
	params := ""
	for _, v := range newFileIDs {
		params = params + fmt.Sprintf(`{ "fileId": %d, "opType": "plus" },`, v)
	}
	for _, v := range oldFileIDs {
		params = params + fmt.Sprintf(`{ "fileId": %d, "opType": "minus" },`, v)
	}
	if params == "" {
		return nil
	}
	params = strings.TrimSuffix(params, ",")
	body := strings.NewReader(fmt.Sprintf(`{ "inputs": [%s] }`, params))
	req, err := http.NewRequest("POST", c.TempOptions.FileBaseUrl+"/files/report-ref-count", body)
	if err != nil {
		return err
	}
	req.Header.Add("X-Tenant-ID", strconv.Itoa(tid))
	req.Header.Add("Content-Type", "application/json")
	_, err = c.KOClient.Do(req)
	return err
}

// EnableTplDataFile 启用模板文件
// tplPath 为temp文件路径
func (c *Coordinator) EnableTplDataFile(tplPath string) error {
	if tplPath == "" {
		return nil
	}
	// 将temp文件复制到data目录下
	distName := c.GetTplDataPath(tplPath)
	_, err := c.CopyFile(distName, c.GetTplTempPath(tplPath))
	if err != nil {
		return err
	}
	// 加载模板
	err = c.TempParseFiles(distName)
	if err != nil {
		return err
	}
	return nil
}

// RemoveTplDataFile 移除data目录模板
// tplPath 为temp文件路径
func (c *Coordinator) RemoveTplDataFile(tplPath string) error {
	if tplPath == "" {
		return nil
	}
	// 将文件从data目录下删除
	dataPath := c.GetTplDataPath(tplPath)
	_, err := os.Stat(dataPath)
	if err == nil {
		err = os.Remove(dataPath)
		if err != nil {
			return err
		}
	}
	return nil
}
