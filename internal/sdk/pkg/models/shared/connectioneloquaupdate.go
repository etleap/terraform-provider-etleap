// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionEloquaUpdateType string

const (
	ConnectionEloquaUpdateTypeEloqua ConnectionEloquaUpdateType = "ELOQUA"
)

func (e ConnectionEloquaUpdateType) ToPointer() *ConnectionEloquaUpdateType {
	return &e
}

func (e *ConnectionEloquaUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ELOQUA":
		*e = ConnectionEloquaUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionEloquaUpdateType: %v", v)
	}
}

type ConnectionEloquaUpdate struct {
	// The unique name of this connection.
	Name *string                     `json:"name,omitempty"`
	Type *ConnectionEloquaUpdateType `json:"type"`
	// Whether this connection should be marked as active.
	Active *bool `json:"active,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	Company        *string              `json:"company,omitempty"`
	Username       *string              `json:"username,omitempty"`
	Password       *string              `json:"password,omitempty"`
}

func (o *ConnectionEloquaUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionEloquaUpdate) GetType() *ConnectionEloquaUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ConnectionEloquaUpdate) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ConnectionEloquaUpdate) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionEloquaUpdate) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionEloquaUpdate) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionEloquaUpdate) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionEloquaUpdate) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionEloquaUpdate) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionEloquaUpdate) GetCompany() *string {
	if o == nil {
		return nil
	}
	return o.Company
}

func (o *ConnectionEloquaUpdate) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *ConnectionEloquaUpdate) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}