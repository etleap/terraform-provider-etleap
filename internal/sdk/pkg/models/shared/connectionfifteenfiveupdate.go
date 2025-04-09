// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionFifteenFiveUpdateType string

const (
	ConnectionFifteenFiveUpdateTypeFifteenFive ConnectionFifteenFiveUpdateType = "FIFTEEN_FIVE"
)

func (e ConnectionFifteenFiveUpdateType) ToPointer() *ConnectionFifteenFiveUpdateType {
	return &e
}

func (e *ConnectionFifteenFiveUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "FIFTEEN_FIVE":
		*e = ConnectionFifteenFiveUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionFifteenFiveUpdateType: %v", v)
	}
}

type ConnectionFifteenFiveUpdate struct {
	// Whether this connection should be marked as active.
	Active *bool                            `json:"active,omitempty"`
	Type   *ConnectionFifteenFiveUpdateType `json:"type"`
	// The unique name of this connection.
	Name *string `json:"name,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// You company 15Five's subdomain, only required if your 15Five instance has a subdomain that's not https://(my).15five.com. Example: https://(subdomain).15Five.com
	Subdomain   *string `json:"subdomain,omitempty"`
	AccessToken *string `json:"accessToken,omitempty"`
}

func (o *ConnectionFifteenFiveUpdate) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ConnectionFifteenFiveUpdate) GetType() *ConnectionFifteenFiveUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ConnectionFifteenFiveUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionFifteenFiveUpdate) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionFifteenFiveUpdate) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionFifteenFiveUpdate) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionFifteenFiveUpdate) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionFifteenFiveUpdate) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionFifteenFiveUpdate) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionFifteenFiveUpdate) GetSubdomain() *string {
	if o == nil {
		return nil
	}
	return o.Subdomain
}

func (o *ConnectionFifteenFiveUpdate) GetAccessToken() *string {
	if o == nil {
		return nil
	}
	return o.AccessToken
}
