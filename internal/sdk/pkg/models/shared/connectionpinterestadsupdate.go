// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionPinterestAdsUpdateType string

const (
	ConnectionPinterestAdsUpdateTypePinterestAds ConnectionPinterestAdsUpdateType = "PINTEREST_ADS"
)

func (e ConnectionPinterestAdsUpdateType) ToPointer() *ConnectionPinterestAdsUpdateType {
	return &e
}

func (e *ConnectionPinterestAdsUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "PINTEREST_ADS":
		*e = ConnectionPinterestAdsUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionPinterestAdsUpdateType: %v", v)
	}
}

type ConnectionPinterestAdsUpdate struct {
	// Whether this connection should be marked as active.
	Active *bool                            `json:"active,omitempty"`
	Type   ConnectionPinterestAdsUpdateType `json:"type"`
	// The unique name of this connection.
	Name *string `json:"name,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// Code retrieved from `/connections/oauth2-initiation`. **Note:** it is short-lived, therefore the connection creation should be done as soon as code is returned.
	Code *string `json:"code,omitempty"`
}

func (o *ConnectionPinterestAdsUpdate) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ConnectionPinterestAdsUpdate) GetType() ConnectionPinterestAdsUpdateType {
	if o == nil {
		return ConnectionPinterestAdsUpdateType("")
	}
	return o.Type
}

func (o *ConnectionPinterestAdsUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionPinterestAdsUpdate) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionPinterestAdsUpdate) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionPinterestAdsUpdate) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionPinterestAdsUpdate) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionPinterestAdsUpdate) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionPinterestAdsUpdate) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionPinterestAdsUpdate) GetCode() *string {
	if o == nil {
		return nil
	}
	return o.Code
}
