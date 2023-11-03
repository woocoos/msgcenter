// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/woocoos/msgcenter/codegen/entgen/schema"
	"github.com/woocoos/msgcenter/ent/msgalert"
	"github.com/woocoos/msgcenter/ent/msgchannel"
	"github.com/woocoos/msgcenter/ent/msgevent"
	"github.com/woocoos/msgcenter/ent/msginternal"
	"github.com/woocoos/msgcenter/ent/msginternalto"
	"github.com/woocoos/msgcenter/ent/msgsubscriber"
	"github.com/woocoos/msgcenter/ent/msgtemplate"
	"github.com/woocoos/msgcenter/ent/msgtype"
	"github.com/woocoos/msgcenter/ent/nlog"
	"github.com/woocoos/msgcenter/ent/nlogalert"
	"github.com/woocoos/msgcenter/ent/orgroleuser"
	"github.com/woocoos/msgcenter/ent/silence"
	"github.com/woocoos/msgcenter/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	msgalertMixin := schema.MsgAlert{}.Mixin()
	msgalertMixinHooks1 := msgalertMixin[1].Hooks()
	msgalert.Hooks[0] = msgalertMixinHooks1[0]
	msgalertMixinInters1 := msgalertMixin[1].Interceptors()
	msgalert.Interceptors[0] = msgalertMixinInters1[0]
	msgalertFields := schema.MsgAlert{}.Fields()
	_ = msgalertFields
	// msgalertDescTimeout is the schema descriptor for timeout field.
	msgalertDescTimeout := msgalertFields[5].Descriptor()
	// msgalert.DefaultTimeout holds the default value on creation for the timeout field.
	msgalert.DefaultTimeout = msgalertDescTimeout.Default.(bool)
	// msgalertDescCreatedAt is the schema descriptor for created_at field.
	msgalertDescCreatedAt := msgalertFields[8].Descriptor()
	// msgalert.DefaultCreatedAt holds the default value on creation for the created_at field.
	msgalert.DefaultCreatedAt = msgalertDescCreatedAt.Default.(func() time.Time)
	// msgalertDescDeleted is the schema descriptor for deleted field.
	msgalertDescDeleted := msgalertFields[10].Descriptor()
	// msgalert.DefaultDeleted holds the default value on creation for the deleted field.
	msgalert.DefaultDeleted = msgalertDescDeleted.Default.(bool)
	msgchannelMixin := schema.MsgChannel{}.Mixin()
	msgchannelMixinHooks1 := msgchannelMixin[1].Hooks()
	msgchannelMixinHooks2 := msgchannelMixin[2].Hooks()
	msgchannelHooks := schema.MsgChannel{}.Hooks()
	msgchannel.Hooks[0] = msgchannelMixinHooks1[0]
	msgchannel.Hooks[1] = msgchannelMixinHooks2[0]
	msgchannel.Hooks[2] = msgchannelHooks[0]
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
	msgeventMixinHooks2 := msgeventMixin[2].Hooks()
	msgeventHooks := schema.MsgEvent{}.Hooks()
	msgevent.Hooks[0] = msgeventMixinHooks1[0]
	msgevent.Hooks[1] = msgeventMixinHooks2[0]
	msgevent.Hooks[2] = msgeventHooks[0]
	msgevent.Hooks[3] = msgeventHooks[1]
	msgevent.Hooks[4] = msgeventHooks[2]
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
	msginternalMixin := schema.MsgInternal{}.Mixin()
	msginternalMixinHooks1 := msginternalMixin[1].Hooks()
	msginternalMixinHooks2 := msginternalMixin[2].Hooks()
	msginternalMixinHooks3 := msginternalMixin[3].Hooks()
	msginternal.Hooks[0] = msginternalMixinHooks1[0]
	msginternal.Hooks[1] = msginternalMixinHooks2[0]
	msginternal.Hooks[2] = msginternalMixinHooks3[0]
	msginternalMixinInters1 := msginternalMixin[1].Interceptors()
	msginternal.Interceptors[0] = msginternalMixinInters1[0]
	msginternalMixinFields2 := msginternalMixin[2].Fields()
	_ = msginternalMixinFields2
	msginternalFields := schema.MsgInternal{}.Fields()
	_ = msginternalFields
	// msginternalDescCreatedAt is the schema descriptor for created_at field.
	msginternalDescCreatedAt := msginternalMixinFields2[1].Descriptor()
	// msginternal.DefaultCreatedAt holds the default value on creation for the created_at field.
	msginternal.DefaultCreatedAt = msginternalDescCreatedAt.Default.(func() time.Time)
	// msginternalDescCategory is the schema descriptor for category field.
	msginternalDescCategory := msginternalFields[0].Descriptor()
	// msginternal.CategoryValidator is a validator for the "category" field. It is called by the builders before save.
	msginternal.CategoryValidator = msginternalDescCategory.Validators[0].(func(string) error)
	msginternaltoMixin := schema.MsgInternalTo{}.Mixin()
	msginternaltoMixinHooks1 := msginternaltoMixin[1].Hooks()
	msginternaltoMixinHooks2 := msginternaltoMixin[2].Hooks()
	msginternalto.Hooks[0] = msginternaltoMixinHooks1[0]
	msginternalto.Hooks[1] = msginternaltoMixinHooks2[0]
	msginternaltoMixinInters1 := msginternaltoMixin[1].Interceptors()
	msginternalto.Interceptors[0] = msginternaltoMixinInters1[0]
	msginternaltoFields := schema.MsgInternalTo{}.Fields()
	_ = msginternaltoFields
	// msginternaltoDescCreatedAt is the schema descriptor for created_at field.
	msginternaltoDescCreatedAt := msginternaltoFields[4].Descriptor()
	// msginternalto.DefaultCreatedAt holds the default value on creation for the created_at field.
	msginternalto.DefaultCreatedAt = msginternaltoDescCreatedAt.Default.(func() time.Time)
	msgsubscriberMixin := schema.MsgSubscriber{}.Mixin()
	msgsubscriberMixinHooks1 := msgsubscriberMixin[1].Hooks()
	msgsubscriberMixinHooks2 := msgsubscriberMixin[2].Hooks()
	msgsubscriberHooks := schema.MsgSubscriber{}.Hooks()
	msgsubscriber.Hooks[0] = msgsubscriberMixinHooks1[0]
	msgsubscriber.Hooks[1] = msgsubscriberMixinHooks2[0]
	msgsubscriber.Hooks[2] = msgsubscriberHooks[0]
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
	msgtemplateMixinHooks2 := msgtemplateMixin[2].Hooks()
	msgtemplate.Hooks[0] = msgtemplateMixinHooks1[0]
	msgtemplate.Hooks[1] = msgtemplateMixinHooks2[0]
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
	msgtypeMixinHooks2 := msgtypeMixin[2].Hooks()
	msgtypeHooks := schema.MsgType{}.Hooks()
	msgtype.Hooks[0] = msgtypeMixinHooks1[0]
	msgtype.Hooks[1] = msgtypeMixinHooks2[0]
	msgtype.Hooks[2] = msgtypeHooks[0]
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
	nlogMixin := schema.Nlog{}.Mixin()
	nlogMixinHooks1 := nlogMixin[1].Hooks()
	nlogMixinHooks2 := nlogMixin[2].Hooks()
	nlog.Hooks[0] = nlogMixinHooks1[0]
	nlog.Hooks[1] = nlogMixinHooks2[0]
	nlogMixinInters1 := nlogMixin[1].Interceptors()
	nlog.Interceptors[0] = nlogMixinInters1[0]
	nlogFields := schema.Nlog{}.Fields()
	_ = nlogFields
	// nlogDescCreatedAt is the schema descriptor for created_at field.
	nlogDescCreatedAt := nlogFields[5].Descriptor()
	// nlog.DefaultCreatedAt holds the default value on creation for the created_at field.
	nlog.DefaultCreatedAt = nlogDescCreatedAt.Default.(func() time.Time)
	nlogalertFields := schema.NlogAlert{}.Fields()
	_ = nlogalertFields
	// nlogalertDescCreatedAt is the schema descriptor for created_at field.
	nlogalertDescCreatedAt := nlogalertFields[2].Descriptor()
	// nlogalert.DefaultCreatedAt holds the default value on creation for the created_at field.
	nlogalert.DefaultCreatedAt = nlogalertDescCreatedAt.Default.(func() time.Time)
	orgroleuserHooks := schema.OrgRoleUser{}.Hooks()
	orgroleuser.Hooks[0] = orgroleuserHooks[0]
	silenceMixin := schema.Silence{}.Mixin()
	silenceMixinHooks1 := silenceMixin[1].Hooks()
	silenceMixinHooks2 := silenceMixin[2].Hooks()
	silenceMixinHooks3 := silenceMixin[3].Hooks()
	silence.Hooks[0] = silenceMixinHooks1[0]
	silence.Hooks[1] = silenceMixinHooks2[0]
	silence.Hooks[2] = silenceMixinHooks3[0]
	silenceMixinInters2 := silenceMixin[2].Interceptors()
	silence.Interceptors[0] = silenceMixinInters2[0]
	silenceMixinFields0 := silenceMixin[0].Fields()
	_ = silenceMixinFields0
	silenceMixinFields1 := silenceMixin[1].Fields()
	_ = silenceMixinFields1
	silenceFields := schema.Silence{}.Fields()
	_ = silenceFields
	// silenceDescCreatedAt is the schema descriptor for created_at field.
	silenceDescCreatedAt := silenceMixinFields1[1].Descriptor()
	// silence.DefaultCreatedAt holds the default value on creation for the created_at field.
	silence.DefaultCreatedAt = silenceDescCreatedAt.Default.(func() time.Time)
	// silenceDescID is the schema descriptor for id field.
	silenceDescID := silenceMixinFields0[0].Descriptor()
	// silence.DefaultID holds the default value on creation for the id field.
	silence.DefaultID = silenceDescID.Default.(func() int)
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
	Version = "v0.12.4"                                         // Version of ent codegen.
	Sum     = "h1:LddPnAyxls/O7DTXZvUGDj0NZIdGSu317+aoNLJWbD8=" // Sum of ent codegen.
)
