package main

import (
	"context"
	"flag"
	"github.com/woocoos/entco/schemax/typex"
	"github.com/woocoos/msgcenter/ent"
	"github.com/woocoos/msgcenter/ent/msgtemplate"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/pkg/profile"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/woocoos/msgcenter/ent/runtime"
)

// receive two arguments: the migration name and the database dsn.
var (
	dsn  = flag.String("dsn", "root:@tcp(localhost:3306)/msgcenter?parseTime=true&loc=Local", "")
	name = flag.String("name", "mysql", "driver name")
)

func main() {
	flag.Parse()
	client, err := ent.Open(*name, *dsn, ent.Debug())
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	tx, err := client.Tx(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			panic(err)
		} else {
			tx.Commit()
		}
	}()
	initmsg(tx)
}

func initmsg(tx *ent.Tx) {
	alterPassWordEventName := "AlterPassword"
	subEventName := "SubEvent"
	ctx := context.Background()
	tx.MsgType.Create().SetName("alert").SetID(1).SetStatus(typex.SimpleStatusActive).SetCreatedBy(1).
		SetAppID(1).SetCategory("账户安全").SetCanSubs(true).SetCanCustom(true).SaveX(ctx)
	tx.MsgEvent.Create().SetID(1).SetMsgTypeID(1).SetName("AlterPassword").SetStatus(typex.SimpleStatusActive).
		SetCreatedBy(1).SetModes("email,internal").
		SetRoute(&profile.Route{
			Name:     alterPassWordEventName,
			Receiver: "email",
			Matchers: label.Matchers{
				{Type: label.MatchEqual, Name: "app", Value: "1"},
				{Name: "alertname", Value: alterPassWordEventName, Type: label.MatchEqual},
			},
			Routes: []*profile.Route{
				{
					Matchers: label.Matchers{
						{Name: "receiver", Value: "email", Type: label.MatchRegexp},
					},
					Receiver: "email",
					Continue: true,
				},
				{
					Matchers: label.Matchers{
						{Name: "receiver", Value: "internal", Type: label.MatchRegexp},
					},
					Receiver: "internal",
				},
			},
		}).SaveX(ctx)
	tx.MsgTemplate.Create().SetMsgTypeID(1).SetEventID(1).SetTenantID(1).SetName(alterPassWordEventName).SetCreatedBy(1).
		SetStatus(typex.SimpleStatusActive).SetFormat(msgtemplate.FormatTxt).SetReceiverType(profile.ReceiverEmail).SetTo(`{{ template "email.to" . }}`).
		SetSubject(`{{ with .CommonAnnotations }}{{.uid}}{{end}}密码到期提醒`).SetCc(`{{ template "email.cc" . }}`).
		SetBcc(`{{ template "email.bcc" . }}`).SetFrom(`custom <test@localhost>`).
		SetBody(`{{ template "1000.alterpwd.txt" . }}`).SetAttachments("1000/alterpwd.tmpl").
		SetTpl("1000/alterpwd.tmpl").SaveX(ctx)

	tx.MsgChannel.Create().SetName("email").SetStatus(typex.SimpleStatusActive).SetCreatedBy(1).
		SetTenantID(1000).SetReceiverType(profile.ReceiverEmail).
		SetReceiver(&profile.Receiver{
			Name: "email",
			EmailConfigs: []*profile.EmailConfig{
				{
					SmartHost: profile.HostPort{Host: "localhost", Port: "1025"},
					To:        `{{ template "email.to" . }}`,
					From:      "serviceSuite@localhost",
				},
			},
		}).SaveX(ctx)
	tx.MsgType.Create().SetName("alert").SetID(2).SetStatus(typex.SimpleStatusActive).SetCreatedBy(1).
		SetAppID(1).SetCategory("订阅类型").SetCanSubs(true).SetCanCustom(true).SaveX(ctx)
	tx.MsgEvent.Create().SetID(2).SetMsgTypeID(2).SetName(subEventName).SetStatus(typex.SimpleStatusActive).
		SetCreatedBy(1).SetModes("email,internal").
		SetRoute(&profile.Route{
			Name:     subEventName,
			Receiver: "email",
			Matchers: label.Matchers{
				{Type: label.MatchEqual, Name: "app", Value: "1"},
				{Name: "alertname", Value: subEventName, Type: label.MatchEqual},
			},
			Routes: []*profile.Route{
				{
					Matchers: label.Matchers{
						{Name: "receiver", Value: "email", Type: label.MatchRegexp},
					},
					Receiver: "email",
					Continue: true,
				},
				{
					Matchers: label.Matchers{
						{Name: "receiver", Value: "internal", Type: label.MatchRegexp},
					},
					Receiver: "internal",
				},
			},
		}).SaveX(ctx)
	tx.MsgTemplate.Create().SetMsgTypeID(2).SetEventID(2).SetTenantID(1).SetName(subEventName).SetCreatedBy(1).
		SetStatus(typex.SimpleStatusActive).SetFormat(msgtemplate.FormatTxt).SetReceiverType(profile.ReceiverEmail).SetTo(`{{ template "email.to" . }}`).
		SetSubject(`订阅事件提醒`).SetCc(`{{ template "email.cc" . }}`).
		SetBcc(`{{ template "email.bcc" . }}`).SetFrom(`custom <test@localhost>`).
		SetBody(`{{ template "1000.subevent.txt" . }}`).
		SetTpl("1000/subevent.tmpl").SaveX(ctx)

	tx.MsgSubscriber.Create().SetID(1).SetMsgTypeID(2).SetTenantID(1).SetUserID(1).SetCreatedBy(1).
		SaveX(ctx)
	tx.MsgSubscriber.Create().SetID(2).SetMsgTypeID(2).SetTenantID(1).SetOrgRoleID(1).SetCreatedBy(1).
		SaveX(ctx)
	tx.MsgSubscriber.Create().SetID(3).SetMsgTypeID(2).SetTenantID(1).SetUserID(3).SetExclude(true).
		SetCreatedBy(1).SaveX(ctx)

}
