// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionCoupaUpdateType string

const (
	ConnectionCoupaUpdateTypeCoupa ConnectionCoupaUpdateType = "COUPA"
)

func (e ConnectionCoupaUpdateType) ToPointer() *ConnectionCoupaUpdateType {
	return &e
}

func (e *ConnectionCoupaUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "COUPA":
		*e = ConnectionCoupaUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionCoupaUpdateType: %v", v)
	}
}

type ConnectionCoupaUpdate struct {
	// The unique name of this connection.
	Name *string                    `json:"name,omitempty"`
	Type *ConnectionCoupaUpdateType `json:"type"`
	// Whether this connection should be marked as active.
	Active *bool `json:"active,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// Your Coupa client id.
	ClientID *string `json:"clientId,omitempty"`
	// Your Coupa client secret.
	ClientSecret *string `json:"clientSecret,omitempty"`
	// The subdomain for your Coupa instance.
	Subdomain *string `json:"subdomain,omitempty"`
}

func (o *ConnectionCoupaUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionCoupaUpdate) GetType() *ConnectionCoupaUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ConnectionCoupaUpdate) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ConnectionCoupaUpdate) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionCoupaUpdate) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionCoupaUpdate) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionCoupaUpdate) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionCoupaUpdate) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionCoupaUpdate) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionCoupaUpdate) GetClientID() *string {
	if o == nil {
		return nil
	}
	return o.ClientID
}

func (o *ConnectionCoupaUpdate) GetClientSecret() *string {
	if o == nil {
		return nil
	}
	return o.ClientSecret
}

func (o *ConnectionCoupaUpdate) GetSubdomain() *string {
	if o == nil {
		return nil
	}
	return o.Subdomain
}
