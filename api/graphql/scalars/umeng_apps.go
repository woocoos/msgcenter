package scalars

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
	"github.com/woocoos/msgcenter/pkg/profile"
)

// UmengApps is a custom scalar for map[string]*profile.UmengAppConfig
type UmengApps map[string]*profile.UmengAppConfig

func (a *UmengApps) UnmarshalGQL(v any) error {
	data, ok := v.(map[string]any)
	if !ok {
		return fmt.Errorf("UmengApps must be a JSON object")
	}
	result := make(map[string]*profile.UmengAppConfig, len(data))
	for key, val := range data {
		app := &profile.UmengAppConfig{}
		appMap, ok := val.(map[string]any)
		if !ok {
			return fmt.Errorf("UmengAppConfig value must be an object")
		}
		if appKey, ok := appMap["appKey"].(string); ok {
			app.AppKey = appKey
		}
		if appMasterSecret, ok := appMap["appMasterSecret"].(string); ok {
			app.AppMasterSecret = appMasterSecret
		}
		if platform, ok := appMap["platform"].(string); ok {
			app.Platform = platform
		}
		if appSet, ok := appMap["appSet"].(string); ok {
			app.AppSet = appSet
		}
		if aliasType, ok := appMap["aliasType"].(string); ok {
			app.AliasType = aliasType
		}
		if afterOpen, ok := appMap["afterOpen"].(string); ok {
			app.AfterOpen = afterOpen
		}
		if activity, ok := appMap["activity"].(string); ok {
			app.Activity = activity
		}
		result[key] = app
	}
	*a = result
	return nil
}

func (a UmengApps) MarshalGQL(w io.Writer) {
	if a == nil {
		io.WriteString(w, "null")
		return
	}
	data := make(map[string]any, len(a))
	for k, v := range a {
		data[k] = map[string]any{
			"appKey":          v.AppKey,
			"appMasterSecret": v.AppMasterSecret,
			"platform":        v.Platform,
			"appSet":          v.AppSet,
			"aliasType":       v.AliasType,
			"afterOpen":       v.AfterOpen,
			"activity":        v.Activity,
		}
	}
	b, _ := json.Marshal(data)
	io.WriteString(w, string(b))
}

// Ensure UmengApps implements the required interfaces
var (
	_ graphql.Unmarshaler = (*UmengApps)(nil)
	_ graphql.Marshaler   = UmengApps(nil)
)
