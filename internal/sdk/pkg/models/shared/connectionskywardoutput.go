// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
	"time"
)

// ConnectionSkywardStatus - The current status of the connection.
type ConnectionSkywardStatus string

const (
	ConnectionSkywardStatusUnknown     ConnectionSkywardStatus = "UNKNOWN"
	ConnectionSkywardStatusUp          ConnectionSkywardStatus = "UP"
	ConnectionSkywardStatusDown        ConnectionSkywardStatus = "DOWN"
	ConnectionSkywardStatusResize      ConnectionSkywardStatus = "RESIZE"
	ConnectionSkywardStatusMaintenance ConnectionSkywardStatus = "MAINTENANCE"
	ConnectionSkywardStatusQuota       ConnectionSkywardStatus = "QUOTA"
	ConnectionSkywardStatusCreating    ConnectionSkywardStatus = "CREATING"
)

func (e ConnectionSkywardStatus) ToPointer() *ConnectionSkywardStatus {
	return &e
}

func (e *ConnectionSkywardStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "UNKNOWN":
		fallthrough
	case "UP":
		fallthrough
	case "DOWN":
		fallthrough
	case "RESIZE":
		fallthrough
	case "MAINTENANCE":
		fallthrough
	case "QUOTA":
		fallthrough
	case "CREATING":
		*e = ConnectionSkywardStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionSkywardStatus: %v", v)
	}
}

type ConnectionSkywardDefaultUpdateSchedule struct {
	// The pipeline mode refers to how the pipeline fetches data changes from the source and how those changes are applied to the destination table. See <a target="_blank" href="https://docs.etleap.com/docs/documentation/ZG9jOjIyMjE3ODA2-introduction">the documentation</a> for more details.
	PipelineMode *PipelineUpdateModes `json:"pipelineMode,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
}

func (o *ConnectionSkywardDefaultUpdateSchedule) GetPipelineMode() *PipelineUpdateModes {
	if o == nil {
		return nil
	}
	return o.PipelineMode
}

func (o *ConnectionSkywardDefaultUpdateSchedule) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionSkywardDefaultUpdateSchedule) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionSkywardDefaultUpdateSchedule) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionSkywardDefaultUpdateSchedule) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionSkywardDefaultUpdateSchedule) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionSkywardDefaultUpdateSchedule) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

type ConnectionSkywardType string

const (
	ConnectionSkywardTypeSkyward ConnectionSkywardType = "SKYWARD"
)

func (e ConnectionSkywardType) ToPointer() *ConnectionSkywardType {
	return &e
}

func (e *ConnectionSkywardType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SKYWARD":
		*e = ConnectionSkywardType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionSkywardType: %v", v)
	}
}

type ConnectionSkywardOutput struct {
	// The current status of the connection.
	Status ConnectionSkywardStatus `json:"status"`
	// The unique name of this connection.
	Name string `json:"name"`
	// The date and time when then the connection was created.
	CreateDate time.Time `json:"createDate"`
	// When an update schedule is not defined for a connection, the default schedule is used. The default defined individually per `pipelineMode` and may be subject to change.
	DefaultUpdateSchedule []ConnectionSkywardDefaultUpdateSchedule `json:"defaultUpdateSchedule"`
	// Whether this connection has been marked as active.
	Active bool                  `json:"active"`
	Type   ConnectionSkywardType `json:"type"`
	// The unique identifier of the connection.
	ID string `json:"id"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	ClientID       string               `json:"clientId"`
}

func (c ConnectionSkywardOutput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionSkywardOutput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionSkywardOutput) GetStatus() ConnectionSkywardStatus {
	if o == nil {
		return ConnectionSkywardStatus("")
	}
	return o.Status
}

func (o *ConnectionSkywardOutput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionSkywardOutput) GetCreateDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreateDate
}

func (o *ConnectionSkywardOutput) GetDefaultUpdateSchedule() []ConnectionSkywardDefaultUpdateSchedule {
	if o == nil {
		return []ConnectionSkywardDefaultUpdateSchedule{}
	}
	return o.DefaultUpdateSchedule
}

func (o *ConnectionSkywardOutput) GetActive() bool {
	if o == nil {
		return false
	}
	return o.Active
}

func (o *ConnectionSkywardOutput) GetType() ConnectionSkywardType {
	if o == nil {
		return ConnectionSkywardType("")
	}
	return o.Type
}

func (o *ConnectionSkywardOutput) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ConnectionSkywardOutput) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionSkywardOutput) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionSkywardOutput) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionSkywardOutput) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionSkywardOutput) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionSkywardOutput) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionSkywardOutput) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

type ConnectionSkyward struct {
	// The unique name of this connection.
	Name string                `json:"name"`
	Type ConnectionSkywardType `json:"type"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	ClientSecret   string               `json:"clientSecret"`
	ClientID       string               `json:"clientId"`
}

func (o *ConnectionSkyward) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionSkyward) GetType() ConnectionSkywardType {
	if o == nil {
		return ConnectionSkywardType("")
	}
	return o.Type
}

func (o *ConnectionSkyward) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionSkyward) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionSkyward) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionSkyward) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionSkyward) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionSkyward) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionSkyward) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *ConnectionSkyward) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}
