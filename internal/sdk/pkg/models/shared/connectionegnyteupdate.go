// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionEgnyteUpdateType string

const (
	ConnectionEgnyteUpdateTypeEgnyte ConnectionEgnyteUpdateType = "EGNYTE"
)

func (e ConnectionEgnyteUpdateType) ToPointer() *ConnectionEgnyteUpdateType {
	return &e
}

func (e *ConnectionEgnyteUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "EGNYTE":
		*e = ConnectionEgnyteUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionEgnyteUpdateType: %v", v)
	}
}

type ConnectionEgnyteUpdate struct {
	// Whether this connection should be marked as active.
	Active *bool                       `json:"active,omitempty"`
	Type   *ConnectionEgnyteUpdateType `json:"type"`
	// The unique name of this connection.
	Name *string `json:"name,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// Code retrieved from `/connections/oauth2-initiation`. **Note:** it is short-lived, therefore the connection creation should be done as soon as code is returned.
	Code *string `json:"code,omitempty"`
	// The path for your base directory. Use "/" for the whole file system.
	BaseDirectory *string `json:"baseDirectory,omitempty"`
	// The name of your Egnyte domain.
	DomainName *string `json:"domainName,omitempty"`
}

func (o *ConnectionEgnyteUpdate) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ConnectionEgnyteUpdate) GetType() *ConnectionEgnyteUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ConnectionEgnyteUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionEgnyteUpdate) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionEgnyteUpdate) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionEgnyteUpdate) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionEgnyteUpdate) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionEgnyteUpdate) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionEgnyteUpdate) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionEgnyteUpdate) GetCode() *string {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *ConnectionEgnyteUpdate) GetBaseDirectory() *string {
	if o == nil {
		return nil
	}
	return o.BaseDirectory
}

func (o *ConnectionEgnyteUpdate) GetDomainName() *string {
	if o == nil {
		return nil
	}
	return o.DomainName
}
