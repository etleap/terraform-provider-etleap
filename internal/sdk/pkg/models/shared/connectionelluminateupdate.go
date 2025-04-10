// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionElluminateUpdateType string

const (
	ConnectionElluminateUpdateTypeElluminate ConnectionElluminateUpdateType = "ELLUMINATE"
)

func (e ConnectionElluminateUpdateType) ToPointer() *ConnectionElluminateUpdateType {
	return &e
}

func (e *ConnectionElluminateUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ELLUMINATE":
		*e = ConnectionElluminateUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionElluminateUpdateType: %v", v)
	}
}

type ConnectionElluminateUpdate struct {
	// Whether this connection should be marked as active.
	Active *bool                           `json:"active,omitempty"`
	Type   *ConnectionElluminateUpdateType `json:"type"`
	// The unique name of this connection.
	Name *string `json:"name,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// Your Elluminate API key
	APIKey *string `json:"apiKey,omitempty"`
	// Your Elluminate URL.
	BaseURL *string `json:"baseUrl,omitempty"`
	// Your Elluminate API secret
	APISecret *string `json:"apiSecret,omitempty"`
}

func (o *ConnectionElluminateUpdate) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ConnectionElluminateUpdate) GetType() *ConnectionElluminateUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ConnectionElluminateUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionElluminateUpdate) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionElluminateUpdate) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionElluminateUpdate) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionElluminateUpdate) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionElluminateUpdate) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionElluminateUpdate) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionElluminateUpdate) GetAPIKey() *string {
	if o == nil {
		return nil
	}
	return o.APIKey
}

func (o *ConnectionElluminateUpdate) GetBaseURL() *string {
	if o == nil {
		return nil
	}
	return o.BaseURL
}

func (o *ConnectionElluminateUpdate) GetAPISecret() *string {
	if o == nil {
		return nil
	}
	return o.APISecret
}
