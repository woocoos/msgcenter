// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/woocoos/knockout-go/ent/schemax/typex"
	"github.com/woocoos/msgcenter/ent/msgtemplate"
	"github.com/woocoos/msgcenter/pkg/alert"
	"github.com/woocoos/msgcenter/pkg/label"
	"github.com/woocoos/msgcenter/pkg/profile"
)

// CreateMsgChannelInput represents a mutation input for creating msgchannels.
type CreateMsgChannelInput struct {
	Name         string
	TenantID     int
	ReceiverType profile.ReceiverType
	Receiver     *profile.Receiver
	Comments     *string
}

// Mutate applies the CreateMsgChannelInput on the MsgChannelMutation builder.
func (i *CreateMsgChannelInput) Mutate(m *MsgChannelMutation) {
	m.SetName(i.Name)
	m.SetTenantID(i.TenantID)
	m.SetReceiverType(i.ReceiverType)
	if v := i.Receiver; v != nil {
		m.SetReceiver(v)
	}
	if v := i.Comments; v != nil {
		m.SetComments(*v)
	}
}

// SetInput applies the change-set in the CreateMsgChannelInput on the MsgChannelCreate builder.
func (c *MsgChannelCreate) SetInput(i CreateMsgChannelInput) *MsgChannelCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateMsgChannelInput represents a mutation input for updating msgchannels.
type UpdateMsgChannelInput struct {
	Name          *string
	TenantID      *int
	ReceiverType  *profile.ReceiverType
	ClearReceiver bool
	Receiver      *profile.Receiver
	ClearComments bool
	Comments      *string
}

// Mutate applies the UpdateMsgChannelInput on the MsgChannelMutation builder.
func (i *UpdateMsgChannelInput) Mutate(m *MsgChannelMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.TenantID; v != nil {
		m.SetTenantID(*v)
	}
	if v := i.ReceiverType; v != nil {
		m.SetReceiverType(*v)
	}
	if i.ClearReceiver {
		m.ClearReceiver()
	}
	if v := i.Receiver; v != nil {
		m.SetReceiver(v)
	}
	if i.ClearComments {
		m.ClearComments()
	}
	if v := i.Comments; v != nil {
		m.SetComments(*v)
	}
}

// SetInput applies the change-set in the UpdateMsgChannelInput on the MsgChannelUpdate builder.
func (c *MsgChannelUpdate) SetInput(i UpdateMsgChannelInput) *MsgChannelUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateMsgChannelInput on the MsgChannelUpdateOne builder.
func (c *MsgChannelUpdateOne) SetInput(i UpdateMsgChannelInput) *MsgChannelUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateMsgEventInput represents a mutation input for creating msgevents.
type CreateMsgEventInput struct {
	Name      string
	Comments  *string
	Route     *profile.Route
	Modes     string
	MsgTypeID int
}

// Mutate applies the CreateMsgEventInput on the MsgEventMutation builder.
func (i *CreateMsgEventInput) Mutate(m *MsgEventMutation) {
	m.SetName(i.Name)
	if v := i.Comments; v != nil {
		m.SetComments(*v)
	}
	if v := i.Route; v != nil {
		m.SetRoute(v)
	}
	m.SetModes(i.Modes)
	m.SetMsgTypeID(i.MsgTypeID)
}

// SetInput applies the change-set in the CreateMsgEventInput on the MsgEventCreate builder.
func (c *MsgEventCreate) SetInput(i CreateMsgEventInput) *MsgEventCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateMsgEventInput represents a mutation input for updating msgevents.
type UpdateMsgEventInput struct {
	Name          *string
	ClearComments bool
	Comments      *string
	ClearRoute    bool
	Route         *profile.Route
	Modes         *string
	MsgTypeID     *int
}

// Mutate applies the UpdateMsgEventInput on the MsgEventMutation builder.
func (i *UpdateMsgEventInput) Mutate(m *MsgEventMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if i.ClearComments {
		m.ClearComments()
	}
	if v := i.Comments; v != nil {
		m.SetComments(*v)
	}
	if i.ClearRoute {
		m.ClearRoute()
	}
	if v := i.Route; v != nil {
		m.SetRoute(v)
	}
	if v := i.Modes; v != nil {
		m.SetModes(*v)
	}
	if v := i.MsgTypeID; v != nil {
		m.SetMsgTypeID(*v)
	}
}

// SetInput applies the change-set in the UpdateMsgEventInput on the MsgEventUpdate builder.
func (c *MsgEventUpdate) SetInput(i UpdateMsgEventInput) *MsgEventUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateMsgEventInput on the MsgEventUpdateOne builder.
func (c *MsgEventUpdateOne) SetInput(i UpdateMsgEventInput) *MsgEventUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateMsgSubscriberInput represents a mutation input for creating msgsubscribers.
type CreateMsgSubscriberInput struct {
	TenantID  int
	OrgRoleID *int
	Exclude   *bool
	MsgTypeID int
	UserID    *int
}

// Mutate applies the CreateMsgSubscriberInput on the MsgSubscriberMutation builder.
func (i *CreateMsgSubscriberInput) Mutate(m *MsgSubscriberMutation) {
	m.SetTenantID(i.TenantID)
	if v := i.OrgRoleID; v != nil {
		m.SetOrgRoleID(*v)
	}
	if v := i.Exclude; v != nil {
		m.SetExclude(*v)
	}
	m.SetMsgTypeID(i.MsgTypeID)
	if v := i.UserID; v != nil {
		m.SetUserID(*v)
	}
}

// SetInput applies the change-set in the CreateMsgSubscriberInput on the MsgSubscriberCreate builder.
func (c *MsgSubscriberCreate) SetInput(i CreateMsgSubscriberInput) *MsgSubscriberCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateMsgSubscriberInput represents a mutation input for updating msgsubscribers.
type UpdateMsgSubscriberInput struct {
	TenantID       *int
	ClearOrgRoleID bool
	OrgRoleID      *int
	ClearExclude   bool
	Exclude        *bool
	MsgTypeID      *int
	ClearUser      bool
	UserID         *int
}

// Mutate applies the UpdateMsgSubscriberInput on the MsgSubscriberMutation builder.
func (i *UpdateMsgSubscriberInput) Mutate(m *MsgSubscriberMutation) {
	if v := i.TenantID; v != nil {
		m.SetTenantID(*v)
	}
	if i.ClearOrgRoleID {
		m.ClearOrgRoleID()
	}
	if v := i.OrgRoleID; v != nil {
		m.SetOrgRoleID(*v)
	}
	if i.ClearExclude {
		m.ClearExclude()
	}
	if v := i.Exclude; v != nil {
		m.SetExclude(*v)
	}
	if v := i.MsgTypeID; v != nil {
		m.SetMsgTypeID(*v)
	}
	if i.ClearUser {
		m.ClearUser()
	}
	if v := i.UserID; v != nil {
		m.SetUserID(*v)
	}
}

// SetInput applies the change-set in the UpdateMsgSubscriberInput on the MsgSubscriberUpdate builder.
func (c *MsgSubscriberUpdate) SetInput(i UpdateMsgSubscriberInput) *MsgSubscriberUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateMsgSubscriberInput on the MsgSubscriberUpdateOne builder.
func (c *MsgSubscriberUpdateOne) SetInput(i UpdateMsgSubscriberInput) *MsgSubscriberUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateMsgTemplateInput represents a mutation input for creating msgtemplates.
type CreateMsgTemplateInput struct {
	MsgTypeID    int
	TenantID     int
	Name         string
	ReceiverType profile.ReceiverType
	Format       msgtemplate.Format
	Subject      *string
	From         *string
	To           *string
	Cc           *string
	Bcc          *string
	Body         *string
	Tpl          *string
	Attachments  []string
	Comments     *string
	EventID      int
}

// Mutate applies the CreateMsgTemplateInput on the MsgTemplateMutation builder.
func (i *CreateMsgTemplateInput) Mutate(m *MsgTemplateMutation) {
	m.SetMsgTypeID(i.MsgTypeID)
	m.SetTenantID(i.TenantID)
	m.SetName(i.Name)
	m.SetReceiverType(i.ReceiverType)
	m.SetFormat(i.Format)
	if v := i.Subject; v != nil {
		m.SetSubject(*v)
	}
	if v := i.From; v != nil {
		m.SetFrom(*v)
	}
	if v := i.To; v != nil {
		m.SetTo(*v)
	}
	if v := i.Cc; v != nil {
		m.SetCc(*v)
	}
	if v := i.Bcc; v != nil {
		m.SetBcc(*v)
	}
	if v := i.Body; v != nil {
		m.SetBody(*v)
	}
	if v := i.Tpl; v != nil {
		m.SetTpl(*v)
	}
	if v := i.Attachments; v != nil {
		m.SetAttachments(v)
	}
	if v := i.Comments; v != nil {
		m.SetComments(*v)
	}
	m.SetEventID(i.EventID)
}

// SetInput applies the change-set in the CreateMsgTemplateInput on the MsgTemplateCreate builder.
func (c *MsgTemplateCreate) SetInput(i CreateMsgTemplateInput) *MsgTemplateCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateMsgTemplateInput represents a mutation input for updating msgtemplates.
type UpdateMsgTemplateInput struct {
	MsgTypeID         *int
	TenantID          *int
	Name              *string
	ReceiverType      *profile.ReceiverType
	Format            *msgtemplate.Format
	ClearSubject      bool
	Subject           *string
	ClearFrom         bool
	From              *string
	ClearTo           bool
	To                *string
	ClearCc           bool
	Cc                *string
	ClearBcc          bool
	Bcc               *string
	ClearBody         bool
	Body              *string
	ClearTpl          bool
	Tpl               *string
	ClearAttachments  bool
	Attachments       []string
	AppendAttachments []string
	ClearComments     bool
	Comments          *string
	EventID           *int
}

// Mutate applies the UpdateMsgTemplateInput on the MsgTemplateMutation builder.
func (i *UpdateMsgTemplateInput) Mutate(m *MsgTemplateMutation) {
	if v := i.MsgTypeID; v != nil {
		m.SetMsgTypeID(*v)
	}
	if v := i.TenantID; v != nil {
		m.SetTenantID(*v)
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.ReceiverType; v != nil {
		m.SetReceiverType(*v)
	}
	if v := i.Format; v != nil {
		m.SetFormat(*v)
	}
	if i.ClearSubject {
		m.ClearSubject()
	}
	if v := i.Subject; v != nil {
		m.SetSubject(*v)
	}
	if i.ClearFrom {
		m.ClearFrom()
	}
	if v := i.From; v != nil {
		m.SetFrom(*v)
	}
	if i.ClearTo {
		m.ClearTo()
	}
	if v := i.To; v != nil {
		m.SetTo(*v)
	}
	if i.ClearCc {
		m.ClearCc()
	}
	if v := i.Cc; v != nil {
		m.SetCc(*v)
	}
	if i.ClearBcc {
		m.ClearBcc()
	}
	if v := i.Bcc; v != nil {
		m.SetBcc(*v)
	}
	if i.ClearBody {
		m.ClearBody()
	}
	if v := i.Body; v != nil {
		m.SetBody(*v)
	}
	if i.ClearTpl {
		m.ClearTpl()
	}
	if v := i.Tpl; v != nil {
		m.SetTpl(*v)
	}
	if i.ClearAttachments {
		m.ClearAttachments()
	}
	if v := i.Attachments; v != nil {
		m.SetAttachments(v)
	}
	if i.AppendAttachments != nil {
		m.AppendAttachments(i.Attachments)
	}
	if i.ClearComments {
		m.ClearComments()
	}
	if v := i.Comments; v != nil {
		m.SetComments(*v)
	}
	if v := i.EventID; v != nil {
		m.SetEventID(*v)
	}
}

// SetInput applies the change-set in the UpdateMsgTemplateInput on the MsgTemplateUpdate builder.
func (c *MsgTemplateUpdate) SetInput(i UpdateMsgTemplateInput) *MsgTemplateUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateMsgTemplateInput on the MsgTemplateUpdateOne builder.
func (c *MsgTemplateUpdateOne) SetInput(i UpdateMsgTemplateInput) *MsgTemplateUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateMsgTypeInput represents a mutation input for creating msgtypes.
type CreateMsgTypeInput struct {
	AppID     *int
	Category  string
	Name      string
	Status    *typex.SimpleStatus
	Comments  *string
	CanSubs   *bool
	CanCustom *bool
}

// Mutate applies the CreateMsgTypeInput on the MsgTypeMutation builder.
func (i *CreateMsgTypeInput) Mutate(m *MsgTypeMutation) {
	if v := i.AppID; v != nil {
		m.SetAppID(*v)
	}
	m.SetCategory(i.Category)
	m.SetName(i.Name)
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if v := i.Comments; v != nil {
		m.SetComments(*v)
	}
	if v := i.CanSubs; v != nil {
		m.SetCanSubs(*v)
	}
	if v := i.CanCustom; v != nil {
		m.SetCanCustom(*v)
	}
}

// SetInput applies the change-set in the CreateMsgTypeInput on the MsgTypeCreate builder.
func (c *MsgTypeCreate) SetInput(i CreateMsgTypeInput) *MsgTypeCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateMsgTypeInput represents a mutation input for updating msgtypes.
type UpdateMsgTypeInput struct {
	ClearAppID     bool
	AppID          *int
	Category       *string
	Name           *string
	ClearStatus    bool
	Status         *typex.SimpleStatus
	ClearComments  bool
	Comments       *string
	ClearCanSubs   bool
	CanSubs        *bool
	ClearCanCustom bool
	CanCustom      *bool
}

// Mutate applies the UpdateMsgTypeInput on the MsgTypeMutation builder.
func (i *UpdateMsgTypeInput) Mutate(m *MsgTypeMutation) {
	if i.ClearAppID {
		m.ClearAppID()
	}
	if v := i.AppID; v != nil {
		m.SetAppID(*v)
	}
	if v := i.Category; v != nil {
		m.SetCategory(*v)
	}
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if i.ClearStatus {
		m.ClearStatus()
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if i.ClearComments {
		m.ClearComments()
	}
	if v := i.Comments; v != nil {
		m.SetComments(*v)
	}
	if i.ClearCanSubs {
		m.ClearCanSubs()
	}
	if v := i.CanSubs; v != nil {
		m.SetCanSubs(*v)
	}
	if i.ClearCanCustom {
		m.ClearCanCustom()
	}
	if v := i.CanCustom; v != nil {
		m.SetCanCustom(*v)
	}
}

// SetInput applies the change-set in the UpdateMsgTypeInput on the MsgTypeUpdate builder.
func (c *MsgTypeUpdate) SetInput(i UpdateMsgTypeInput) *MsgTypeUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateMsgTypeInput on the MsgTypeUpdateOne builder.
func (c *MsgTypeUpdateOne) SetInput(i UpdateMsgTypeInput) *MsgTypeUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateSilenceInput represents a mutation input for creating silences.
type CreateSilenceInput struct {
	TenantID int
	Matchers []*label.Matcher
	StartsAt time.Time
	EndsAt   time.Time
	Comments *string
	State    *alert.SilenceState
}

// Mutate applies the CreateSilenceInput on the SilenceMutation builder.
func (i *CreateSilenceInput) Mutate(m *SilenceMutation) {
	m.SetTenantID(i.TenantID)
	if v := i.Matchers; v != nil {
		m.SetMatchers(v)
	}
	m.SetStartsAt(i.StartsAt)
	m.SetEndsAt(i.EndsAt)
	if v := i.Comments; v != nil {
		m.SetComments(*v)
	}
	if v := i.State; v != nil {
		m.SetState(*v)
	}
}

// SetInput applies the change-set in the CreateSilenceInput on the SilenceCreate builder.
func (c *SilenceCreate) SetInput(i CreateSilenceInput) *SilenceCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateSilenceInput represents a mutation input for updating silences.
type UpdateSilenceInput struct {
	ClearMatchers  bool
	Matchers       []*label.Matcher
	AppendMatchers []*label.Matcher
	StartsAt       *time.Time
	EndsAt         *time.Time
	ClearComments  bool
	Comments       *string
	State          *alert.SilenceState
}

// Mutate applies the UpdateSilenceInput on the SilenceMutation builder.
func (i *UpdateSilenceInput) Mutate(m *SilenceMutation) {
	if i.ClearMatchers {
		m.ClearMatchers()
	}
	if v := i.Matchers; v != nil {
		m.SetMatchers(v)
	}
	if i.AppendMatchers != nil {
		m.AppendMatchers(i.Matchers)
	}
	if v := i.StartsAt; v != nil {
		m.SetStartsAt(*v)
	}
	if v := i.EndsAt; v != nil {
		m.SetEndsAt(*v)
	}
	if i.ClearComments {
		m.ClearComments()
	}
	if v := i.Comments; v != nil {
		m.SetComments(*v)
	}
	if v := i.State; v != nil {
		m.SetState(*v)
	}
}

// SetInput applies the change-set in the UpdateSilenceInput on the SilenceUpdate builder.
func (c *SilenceUpdate) SetInput(i UpdateSilenceInput) *SilenceUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateSilenceInput on the SilenceUpdateOne builder.
func (c *SilenceUpdateOne) SetInput(i UpdateSilenceInput) *SilenceUpdateOne {
	i.Mutate(c.Mutation())
	return c
}
