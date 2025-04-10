// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionTikTokAdsUpdateType string

const (
	ConnectionTikTokAdsUpdateTypeTikTokAds ConnectionTikTokAdsUpdateType = "TIK_TOK_ADS"
)

func (e ConnectionTikTokAdsUpdateType) ToPointer() *ConnectionTikTokAdsUpdateType {
	return &e
}

func (e *ConnectionTikTokAdsUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TIK_TOK_ADS":
		*e = ConnectionTikTokAdsUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionTikTokAdsUpdateType: %v", v)
	}
}

type ConnectionTikTokAdsUpdate struct {
	// Whether this connection should be marked as active.
	Active *bool                         `json:"active,omitempty"`
	Type   ConnectionTikTokAdsUpdateType `json:"type"`
	// The unique name of this connection.
	Name *string `json:"name,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// Code retrieved from `/connections/oauth2-initiation`. **Note:** it is short-lived, therefore the connection creation should be done as soon as code is returned.
	Code *string `json:"code,omitempty"`
}

func (o *ConnectionTikTokAdsUpdate) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ConnectionTikTokAdsUpdate) GetType() ConnectionTikTokAdsUpdateType {
	if o == nil {
		return ConnectionTikTokAdsUpdateType("")
	}
	return o.Type
}

func (o *ConnectionTikTokAdsUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionTikTokAdsUpdate) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionTikTokAdsUpdate) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionTikTokAdsUpdate) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionTikTokAdsUpdate) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionTikTokAdsUpdate) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionTikTokAdsUpdate) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionTikTokAdsUpdate) GetCode() *string {
	if o == nil {
		return nil
	}
	return o.Code
}
