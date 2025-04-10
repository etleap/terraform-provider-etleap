// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
	"time"
)

// ConnectionMixpanelStatus - The current status of the connection.
type ConnectionMixpanelStatus string

const (
	ConnectionMixpanelStatusUnknown     ConnectionMixpanelStatus = "UNKNOWN"
	ConnectionMixpanelStatusUp          ConnectionMixpanelStatus = "UP"
	ConnectionMixpanelStatusDown        ConnectionMixpanelStatus = "DOWN"
	ConnectionMixpanelStatusResize      ConnectionMixpanelStatus = "RESIZE"
	ConnectionMixpanelStatusMaintenance ConnectionMixpanelStatus = "MAINTENANCE"
	ConnectionMixpanelStatusQuota       ConnectionMixpanelStatus = "QUOTA"
	ConnectionMixpanelStatusCreating    ConnectionMixpanelStatus = "CREATING"
)

func (e ConnectionMixpanelStatus) ToPointer() *ConnectionMixpanelStatus {
	return &e
}

func (e *ConnectionMixpanelStatus) UnmarshalJSON(data []byte) error {
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
		*e = ConnectionMixpanelStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionMixpanelStatus: %v", v)
	}
}

type ConnectionMixpanelDefaultUpdateSchedule struct {
	// The pipeline mode refers to how the pipeline fetches data changes from the source and how those changes are applied to the destination table. See <a target="_blank" href="https://docs.etleap.com/docs/documentation/ZG9jOjIyMjE3ODA2-introduction">the documentation</a> for more details.
	PipelineMode *PipelineUpdateModes `json:"pipelineMode,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
}

func (o *ConnectionMixpanelDefaultUpdateSchedule) GetPipelineMode() *PipelineUpdateModes {
	if o == nil {
		return nil
	}
	return o.PipelineMode
}

func (o *ConnectionMixpanelDefaultUpdateSchedule) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionMixpanelDefaultUpdateSchedule) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionMixpanelDefaultUpdateSchedule) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionMixpanelDefaultUpdateSchedule) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionMixpanelDefaultUpdateSchedule) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionMixpanelDefaultUpdateSchedule) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

type ConnectionMixpanelType string

const (
	ConnectionMixpanelTypeMixpanel ConnectionMixpanelType = "MIXPANEL"
)

func (e ConnectionMixpanelType) ToPointer() *ConnectionMixpanelType {
	return &e
}

func (e *ConnectionMixpanelType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MIXPANEL":
		*e = ConnectionMixpanelType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionMixpanelType: %v", v)
	}
}

type ConnectionMixpanel struct {
	// The current status of the connection.
	Status ConnectionMixpanelStatus `json:"status"`
	// The unique name of this connection.
	Name string `json:"name"`
	// The date and time when then the connection was created.
	CreateDate time.Time `json:"createDate"`
	// When an update schedule is not defined for a connection, the default schedule is used. The default defined individually per `pipelineMode` and may be subject to change.
	DefaultUpdateSchedule []ConnectionMixpanelDefaultUpdateSchedule `json:"defaultUpdateSchedule"`
	// Whether this connection has been marked as active.
	Active bool                   `json:"active"`
	Type   ConnectionMixpanelType `json:"type"`
	// The unique identifier of the connection.
	ID string `json:"id"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// Mixpanel project timezone. Default: 'US/Pacific'
	Timezone string `json:"timezone"`
}

func (c ConnectionMixpanel) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionMixpanel) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionMixpanel) GetStatus() ConnectionMixpanelStatus {
	if o == nil {
		return ConnectionMixpanelStatus("")
	}
	return o.Status
}

func (o *ConnectionMixpanel) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionMixpanel) GetCreateDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreateDate
}

func (o *ConnectionMixpanel) GetDefaultUpdateSchedule() []ConnectionMixpanelDefaultUpdateSchedule {
	if o == nil {
		return []ConnectionMixpanelDefaultUpdateSchedule{}
	}
	return o.DefaultUpdateSchedule
}

func (o *ConnectionMixpanel) GetActive() bool {
	if o == nil {
		return false
	}
	return o.Active
}

func (o *ConnectionMixpanel) GetType() ConnectionMixpanelType {
	if o == nil {
		return ConnectionMixpanelType("")
	}
	return o.Type
}

func (o *ConnectionMixpanel) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ConnectionMixpanel) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionMixpanel) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionMixpanel) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionMixpanel) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionMixpanel) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionMixpanel) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionMixpanel) GetTimezone() string {
	if o == nil {
		return ""
	}
	return o.Timezone
}

type ConnectionMixpanelInput struct {
	// The unique name of this connection.
	Name string                 `json:"name"`
	Type ConnectionMixpanelType `json:"type"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// Mixpanel project timezone. Default: 'US/Pacific'
	Timezone string `json:"timezone"`
	// 'API Secret' under Mixpanel 'Project Settings'
	APISecret string `json:"apiSecret"`
}

func (o *ConnectionMixpanelInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionMixpanelInput) GetType() ConnectionMixpanelType {
	if o == nil {
		return ConnectionMixpanelType("")
	}
	return o.Type
}

func (o *ConnectionMixpanelInput) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionMixpanelInput) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionMixpanelInput) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionMixpanelInput) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionMixpanelInput) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionMixpanelInput) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionMixpanelInput) GetTimezone() string {
	if o == nil {
		return ""
	}
	return o.Timezone
}

func (o *ConnectionMixpanelInput) GetAPISecret() string {
	if o == nil {
		return ""
	}
	return o.APISecret
}
