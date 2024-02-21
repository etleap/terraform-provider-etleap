// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionImpactRadiusUpdateType string

const (
	ConnectionImpactRadiusUpdateTypeImpactRadius ConnectionImpactRadiusUpdateType = "IMPACT_RADIUS"
)

func (e ConnectionImpactRadiusUpdateType) ToPointer() *ConnectionImpactRadiusUpdateType {
	return &e
}

func (e *ConnectionImpactRadiusUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "IMPACT_RADIUS":
		*e = ConnectionImpactRadiusUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionImpactRadiusUpdateType: %v", v)
	}
}

type ConnectionImpactRadiusUpdate struct {
	// The unique name of this connection.
	Name *string                           `json:"name,omitempty"`
	Type *ConnectionImpactRadiusUpdateType `json:"type"`
	// Whether this connection should be marked as active.
	Active *bool `json:"active,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// To find the Account SID to Impact Radius, click on the Cloud Icon in the bottom left > Click API > Find Rest API information
	AccountSid *string `json:"accountSid,omitempty"`
	// To find the Auth Token to Impact Radius, click on the Cloud Icon in the bottom left > Click API > Find Rest API information
	AuthToken *string `json:"authToken,omitempty"`
}

func (o *ConnectionImpactRadiusUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionImpactRadiusUpdate) GetType() *ConnectionImpactRadiusUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ConnectionImpactRadiusUpdate) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ConnectionImpactRadiusUpdate) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionImpactRadiusUpdate) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionImpactRadiusUpdate) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionImpactRadiusUpdate) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionImpactRadiusUpdate) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionImpactRadiusUpdate) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionImpactRadiusUpdate) GetAccountSid() *string {
	if o == nil {
		return nil
	}
	return o.AccountSid
}

func (o *ConnectionImpactRadiusUpdate) GetAuthToken() *string {
	if o == nil {
		return nil
	}
	return o.AuthToken
}