// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/woocoos/msgcenter/codegen/entgen/schema"
	"github.com/woocoos/msgcenter/ent/msgchannel"
	"github.com/woocoos/msgcenter/ent/msgevent"
	"github.com/woocoos/msgcenter/ent/msgsubscriber"
	"github.com/woocoos/msgcenter/ent/msgtemplate"
	"github.com/woocoos/msgcenter/ent/msgtype"
	"github.com/woocoos/msgcenter/ent/orgroleuser"
	"github.com/woocoos/msgcenter/ent/silence"
	"github.com/woocoos/msgcenter/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	msgchannelMixin := schema.MsgChannel{}.Mixin()
	msgchannelMixinHooks1 := msgchannelMixin[1].Hooks()
	msgchannelHooks := schema.MsgChannel{}.Hooks()
	msgchannel.Hooks[0] = msgchannelMixinHooks1[0]
	msgchannel.Hooks[1] = msgchannelHooks[0]
	msgchannelMixinFields1 := msgchannelMixin[1].Fields()
	_ = msgchannelMixinFields1
	msgchannelFields := schema.MsgChannel{}.Fields()
	_ = msgchannelFields
	// msgchannelDescCreatedAt is the schema descriptor for created_at field.
	msgchannelDescCreatedAt := msgchannelMixinFields1[1].Descriptor()
	// msgchannel.DefaultCreatedAt holds the default value on creation for the created_at field.
	msgchannel.DefaultCreatedAt = msgchannelDescCreatedAt.Default.(func() time.Time)
	// msgchannelDescName is the schema descriptor for name field.
	msgchannelDescName := msgchannelFields[0].Descriptor()
	// msgchannel.NameValidator is a validator for the "name" field. It is called by the builders before save.
	msgchannel.NameValidator = msgchannelDescName.Validators[0].(func(string) error)
	msgeventMixin := schema.MsgEvent{}.Mixin()
	msgeventMixinHooks1 := msgeventMixin[1].Hooks()
	msgeventHooks := schema.MsgEvent{}.Hooks()
	msgevent.Hooks[0] = msgeventMixinHooks1[0]
	msgevent.Hooks[1] = msgeventHooks[0]
	msgevent.Hooks[2] = msgeventHooks[1]
	msgevent.Hooks[3] = msgeventHooks[2]
	msgeventMixinFields1 := msgeventMixin[1].Fields()
	_ = msgeventMixinFields1
	msgeventFields := schema.MsgEvent{}.Fields()
	_ = msgeventFields
	// msgeventDescCreatedAt is the schema descriptor for created_at field.
	msgeventDescCreatedAt := msgeventMixinFields1[1].Descriptor()
	// msgevent.DefaultCreatedAt holds the default value on creation for the created_at field.
	msgevent.DefaultCreatedAt = msgeventDescCreatedAt.Default.(func() time.Time)
	// msgeventDescName is the schema descriptor for name field.
	msgeventDescName := msgeventFields[1].Descriptor()
	// msgevent.NameValidator is a validator for the "name" field. It is called by the builders before save.
	msgevent.NameValidator = func() func(string) error {
		validators := msgeventDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	msgsubscriberMixin := schema.MsgSubscriber{}.Mixin()
	msgsubscriberMixinHooks1 := msgsubscriberMixin[1].Hooks()
	msgsubscriberHooks := schema.MsgSubscriber{}.Hooks()
	msgsubscriber.Hooks[0] = msgsubscriberMixinHooks1[0]
	msgsubscriber.Hooks[1] = msgsubscriberHooks[0]
	msgsubscriberMixinFields1 := msgsubscriberMixin[1].Fields()
	_ = msgsubscriberMixinFields1
	msgsubscriberFields := schema.MsgSubscriber{}.Fields()
	_ = msgsubscriberFields
	// msgsubscriberDescCreatedAt is the schema descriptor for created_at field.
	msgsubscriberDescCreatedAt := msgsubscriberMixinFields1[1].Descriptor()
	// msgsubscriber.DefaultCreatedAt holds the default value on creation for the created_at field.
	msgsubscriber.DefaultCreatedAt = msgsubscriberDescCreatedAt.Default.(func() time.Time)
	// msgsubscriberDescExclude is the schema descriptor for exclude field.
	msgsubscriberDescExclude := msgsubscriberFields[4].Descriptor()
	// msgsubscriber.DefaultExclude holds the default value on creation for the exclude field.
	msgsubscriber.DefaultExclude = msgsubscriberDescExclude.Default.(bool)
	msgtemplateMixin := schema.MsgTemplate{}.Mixin()
	msgtemplateMixinHooks1 := msgtemplateMixin[1].Hooks()
	msgtemplate.Hooks[0] = msgtemplateMixinHooks1[0]
	msgtemplateMixinFields1 := msgtemplateMixin[1].Fields()
	_ = msgtemplateMixinFields1
	msgtemplateFields := schema.MsgTemplate{}.Fields()
	_ = msgtemplateFields
	// msgtemplateDescCreatedAt is the schema descriptor for created_at field.
	msgtemplateDescCreatedAt := msgtemplateMixinFields1[1].Descriptor()
	// msgtemplate.DefaultCreatedAt holds the default value on creation for the created_at field.
	msgtemplate.DefaultCreatedAt = msgtemplateDescCreatedAt.Default.(func() time.Time)
	// msgtemplateDescName is the schema descriptor for name field.
	msgtemplateDescName := msgtemplateFields[3].Descriptor()
	// msgtemplate.NameValidator is a validator for the "name" field. It is called by the builders before save.
	msgtemplate.NameValidator = msgtemplateDescName.Validators[0].(func(string) error)
	msgtypeMixin := schema.MsgType{}.Mixin()
	msgtypeMixinHooks1 := msgtypeMixin[1].Hooks()
	msgtype.Hooks[0] = msgtypeMixinHooks1[0]
	msgtypeMixinFields1 := msgtypeMixin[1].Fields()
	_ = msgtypeMixinFields1
	msgtypeFields := schema.MsgType{}.Fields()
	_ = msgtypeFields
	// msgtypeDescCreatedAt is the schema descriptor for created_at field.
	msgtypeDescCreatedAt := msgtypeMixinFields1[1].Descriptor()
	// msgtype.DefaultCreatedAt holds the default value on creation for the created_at field.
	msgtype.DefaultCreatedAt = msgtypeDescCreatedAt.Default.(func() time.Time)
	// msgtypeDescCategory is the schema descriptor for category field.
	msgtypeDescCategory := msgtypeFields[1].Descriptor()
	// msgtype.CategoryValidator is a validator for the "category" field. It is called by the builders before save.
	msgtype.CategoryValidator = msgtypeDescCategory.Validators[0].(func(string) error)
	// msgtypeDescName is the schema descriptor for name field.
	msgtypeDescName := msgtypeFields[2].Descriptor()
	// msgtype.NameValidator is a validator for the "name" field. It is called by the builders before save.
	msgtype.NameValidator = msgtypeDescName.Validators[0].(func(string) error)
	// msgtypeDescCanSubs is the schema descriptor for can_subs field.
	msgtypeDescCanSubs := msgtypeFields[5].Descriptor()
	// msgtype.DefaultCanSubs holds the default value on creation for the can_subs field.
	msgtype.DefaultCanSubs = msgtypeDescCanSubs.Default.(bool)
	// msgtypeDescCanCustom is the schema descriptor for can_custom field.
	msgtypeDescCanCustom := msgtypeFields[6].Descriptor()
	// msgtype.DefaultCanCustom holds the default value on creation for the can_custom field.
	msgtype.DefaultCanCustom = msgtypeDescCanCustom.Default.(bool)
	orgroleuserHooks := schema.OrgRoleUser{}.Hooks()
	orgroleuser.Hooks[0] = orgroleuserHooks[0]
	silenceMixin := schema.Silence{}.Mixin()
	silenceMixinHooks1 := silenceMixin[1].Hooks()
	silenceMixinHooks2 := silenceMixin[2].Hooks()
	silence.Hooks[0] = silenceMixinHooks1[0]
	silence.Hooks[1] = silenceMixinHooks2[0]
	silenceMixinInters2 := silenceMixin[2].Interceptors()
	silence.Interceptors[0] = silenceMixinInters2[0]
	silenceMixinFields1 := silenceMixin[1].Fields()
	_ = silenceMixinFields1
	silenceFields := schema.Silence{}.Fields()
	_ = silenceFields
	// silenceDescCreatedAt is the schema descriptor for created_at field.
	silenceDescCreatedAt := silenceMixinFields1[1].Descriptor()
	// silence.DefaultCreatedAt holds the default value on creation for the created_at field.
	silence.DefaultCreatedAt = silenceDescCreatedAt.Default.(func() time.Time)
	userHooks := schema.User{}.Hooks()
	user.Hooks[0] = userHooks[0]
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[3].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescMobile is the schema descriptor for mobile field.
	userDescMobile := userFields[4].Descriptor()
	// user.MobileValidator is a validator for the "mobile" field. It is called by the builders before save.
	user.MobileValidator = userDescMobile.Validators[0].(func(string) error)
}

const (
	Version = "v0.12.3"                                         // Version of ent codegen.
	Sum     = "h1:N5lO2EOrHpCH5HYfiMOCHYbo+oh5M8GjT0/cx5x6xkk=" // Sum of ent codegen.
)
