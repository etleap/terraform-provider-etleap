// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionSnapchatAdsUpdateType string

const (
	ConnectionSnapchatAdsUpdateTypeSnapchatAds ConnectionSnapchatAdsUpdateType = "SNAPCHAT_ADS"
)

func (e ConnectionSnapchatAdsUpdateType) ToPointer() *ConnectionSnapchatAdsUpdateType {
	return &e
}

func (e *ConnectionSnapchatAdsUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SNAPCHAT_ADS":
		*e = ConnectionSnapchatAdsUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionSnapchatAdsUpdateType: %v", v)
	}
}

type ConnectionSnapchatAdsUpdate struct {
	// Whether this connection should be marked as active.
	Active *bool                           `json:"active,omitempty"`
	Type   ConnectionSnapchatAdsUpdateType `json:"type"`
	// The unique name of this connection.
	Name *string `json:"name,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// Code retrieved from `/connections/oauth2-initiation`. **Note:** it is short-lived, therefore the connection creation should be done as soon as code is returned.
	Code *string `json:"code,omitempty"`
}

func (o *ConnectionSnapchatAdsUpdate) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ConnectionSnapchatAdsUpdate) GetType() ConnectionSnapchatAdsUpdateType {
	if o == nil {
		return ConnectionSnapchatAdsUpdateType("")
	}
	return o.Type
}

func (o *ConnectionSnapchatAdsUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionSnapchatAdsUpdate) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionSnapchatAdsUpdate) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionSnapchatAdsUpdate) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionSnapchatAdsUpdate) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionSnapchatAdsUpdate) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionSnapchatAdsUpdate) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionSnapchatAdsUpdate) GetCode() *string {
	if o == nil {
		return nil
	}
	return o.Code
}
