// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionZoomPhoneUpdateType string

const (
	ConnectionZoomPhoneUpdateTypeZoomPhone ConnectionZoomPhoneUpdateType = "ZOOM_PHONE"
)

func (e ConnectionZoomPhoneUpdateType) ToPointer() *ConnectionZoomPhoneUpdateType {
	return &e
}

func (e *ConnectionZoomPhoneUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ZOOM_PHONE":
		*e = ConnectionZoomPhoneUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionZoomPhoneUpdateType: %v", v)
	}
}

type ConnectionZoomPhoneUpdate struct {
	// The unique name of this connection.
	Name *string                       `json:"name,omitempty"`
	Type ConnectionZoomPhoneUpdateType `json:"type"`
	// Whether this connection should be marked as active.
	Active *bool `json:"active,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// Code retrieved from `/connections/oauth2-initiation`. **Note:** it is short-lived, therefore the connection creation should be done as soon as code is returned.
	Code *string `json:"code,omitempty"`
	// Under Admin Dashboard > Created App > App credentials
	ClientID *string `json:"clientId,omitempty"`
	// Under Admin Dashboard > Created App > App credentials
	ClientSecret *string `json:"clientSecret,omitempty"`
}

func (o *ConnectionZoomPhoneUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionZoomPhoneUpdate) GetType() ConnectionZoomPhoneUpdateType {
	if o == nil {
		return ConnectionZoomPhoneUpdateType("")
	}
	return o.Type
}

func (o *ConnectionZoomPhoneUpdate) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ConnectionZoomPhoneUpdate) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionZoomPhoneUpdate) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionZoomPhoneUpdate) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionZoomPhoneUpdate) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionZoomPhoneUpdate) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionZoomPhoneUpdate) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionZoomPhoneUpdate) GetCode() *string {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *ConnectionZoomPhoneUpdate) GetClientID() *string {
	if o == nil {
		return nil
	}
	return o.ClientID
}

func (o *ConnectionZoomPhoneUpdate) GetClientSecret() *string {
	if o == nil {
		return nil
	}
	return o.ClientSecret
}
