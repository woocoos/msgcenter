package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tsingsun/woocoo/pkg/conf"
	"github.com/tsingsun/woocoo/pkg/log"
	"github.com/tsingsun/woocoo/pkg/store/redisx"
	"github.com/woocoos/knockout-go/api"
	"github.com/woocoos/knockout-go/api/fs"
	"github.com/woocoos/knockout-go/ent/schemax/typex"
	"github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/ent/msgchannel"
	"github.com/woocoos/msgcenter/ent/msgevent"
	"github.com/woocoos/msgcenter/ent/msgtemplate"
	"github.com/woocoos/msgcenter/notify"
	"github.com/woocoos/msgcenter/notify/email"
	"github.com/woocoos/msgcenter/notify/message"
	"github.com/woocoos/msgcenter/notify/webhook"
	"github.com/woocoos/msgcenter/pkg/metrics"
	"github.com/woocoos/msgcenter/pkg/profile"
	"github.com/woocoos/msgcenter/service/kosdk"
	"github.com/woocoos/msgcenter/template"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

var logger = log.Component("config")

// Coordinator helps the Alert Manager collaborate with external components.
type Coordinator struct {
	cnf *conf.Configuration
	// Protects profile and reloadHooks
	mutex       sync.RWMutex
	profile     *profile.Config
	reloadHooks []func(*profile.Config) error

	ActiveReceivers map[string]int // receiver name -> number of Notifiers
	Template        *template.Template

	db *ent.Client
	// knockout sdk
	KOSdk *api.SDK
}

// NewCoordinator returns a new coordinator with the given configuration for alert manager.
// It does not yet load the configuration from file. This is done in
// `Reload()`.
func NewCoordinator(cnf *conf.Configuration) *Coordinator {
	c := &Coordinator{
		cnf:             cnf,
		ActiveReceivers: make(map[string]int),
	}

	return c
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

	configFilePath := c.cnf.String("config.file")
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

// loadTemplates loading template files
func (c *Coordinator) loadTemplates() error {
	tmpl, err := template.New()
	if err != nil {
		return fmt.Errorf("failed to parse templates %w", err)
	}
	tmpl.KOSdk = c.KOSdk
	c.Template = tmpl
	if err = c.cnf.Sub("template").Unmarshal(&c.Template); err != nil {
		return err
	}
	c.Template.BaseDir, err = filepath.Abs(c.Template.BaseDir)
	if err != nil {
		return err
	}
	// 远程下载文件
	if err = c.downloadTempFromRemote(); err != nil {
		logger.Error("Error loading remote template file", zap.Error(err))
		metrics.Coordinator.ConfigSuccess.Set(0)
		return err
	}
	for _, ptn := range c.profile.Templates {
		// 如果根目录未创建则跳过
		if _, err = os.Stat(c.Template.BaseDir); os.IsNotExist(err) {
			continue
		}
		if _, err = c.Template.ParseGlob(c.cnf.Abs(ptn)); err != nil {
			return err
		}
	}
	return nil
}

// loadTempFromRemote 从远程下载模板文件
func (c *Coordinator) downloadTempFromRemote() error {
	// 获取租户的模板文件
	tpls, err := c.db.MsgTemplate.Query().Where(msgtemplate.StatusEQ(typex.SimpleStatusActive)).All(context.Background())
	if err != nil {
		return err
	}
	for _, tpl := range tpls {
		// 下载模板文件
		if tpl.Tpl != "" {
			localFile, err := kosdk.DefaultFilePath(tpl.TenantID, tpl.Tpl, c.Template.BaseDir, c.Template.DataDir)
			if err != nil {
				return err
			}
			err = c.KOSdk.Fs().DownloadObjectByKey(strconv.Itoa(tpl.TenantID), tpl.Tpl, localFile, fs.WithOverwrittenFile(false))
			if err != nil {
				return err
			}
		}
		// 下载附件
		if len(tpl.Attachments) != 0 {
			for _, att := range tpl.Attachments {
				localFile, err := kosdk.DefaultFilePath(tpl.TenantID, att, c.Template.BaseDir, c.Template.AttachmentDir)
				if err != nil {
					return err
				}
				err = c.KOSdk.Fs().DownloadObjectByKey(strconv.Itoa(tpl.TenantID), att, localFile, fs.WithOverwrittenFile(false))
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// loadFromFile triggers a configuration load, discarding the old configuration.
func (c *Coordinator) loadFromFile() error {
	config, err := profile.NewConfig(c.cnf)
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
	basedir := c.Template.BaseDir
	attdir := c.Template.AttachmentDir
	for i, cfg := range nc.EmailConfigs {
		add("email", i, func() (notify.Notifier, error) {
			return email.New(cfg, tmpl, overrideEmailConfig(basedir, attdir, c.db))
		})
	}
	for i, cfg := range nc.WebhookConfigs {
		add("webhook", i, func() (notify.Notifier, error) {
			return webhook.New(cfg, tmpl, overrideWebHookConfig(basedir, attdir, c.db))
		})
	}
	if nc.MessageConfig != nil {
		cli, err := redisx.NewClient(c.cnf.Root().Sub("store.redis"))
		if err != nil {
			return nil, err
		}
		add("message", 0, func() (notify.Notifier, error) {
			return message.New(nc.MessageConfig, tmpl, c.db, cli.UniversalClient,
				overrideMessageConfig(basedir, attdir, c.db))
		})
	}
	if errs != nil {
		return nil, errs
	}
	return integrations, nil
}
