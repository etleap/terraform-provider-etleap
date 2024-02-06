// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
	"time"
)

type ConnectionSquareType string

const (
	ConnectionSquareTypeSquare ConnectionSquareType = "SQUARE"
)

func (e ConnectionSquareType) ToPointer() *ConnectionSquareType {
	return &e
}

func (e *ConnectionSquareType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SQUARE":
		*e = ConnectionSquareType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionSquareType: %v", v)
	}
}

// ConnectionSquareStatus - The current status of the connection.
type ConnectionSquareStatus string

const (
	ConnectionSquareStatusUnknown     ConnectionSquareStatus = "UNKNOWN"
	ConnectionSquareStatusUp          ConnectionSquareStatus = "UP"
	ConnectionSquareStatusDown        ConnectionSquareStatus = "DOWN"
	ConnectionSquareStatusResize      ConnectionSquareStatus = "RESIZE"
	ConnectionSquareStatusMaintenance ConnectionSquareStatus = "MAINTENANCE"
	ConnectionSquareStatusQuota       ConnectionSquareStatus = "QUOTA"
	ConnectionSquareStatusCreating    ConnectionSquareStatus = "CREATING"
)

func (e ConnectionSquareStatus) ToPointer() *ConnectionSquareStatus {
	return &e
}

func (e *ConnectionSquareStatus) UnmarshalJSON(data []byte) error {
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
		*e = ConnectionSquareStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionSquareStatus: %v", v)
	}
}

type ConnectionSquareDefaultUpdateSchedule struct {
	// The pipeline mode refers to how the pipeline fetches data changes from the source and how those changes are applied to the destination table. See <a target="_blank" href="https://docs.etleap.com/docs/documentation/ZG9jOjIyMjE3ODA2-introduction">the documentation</a> for more details.
	PipelineMode *PipelineUpdateModes `json:"pipelineMode,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
}

func (o *ConnectionSquareDefaultUpdateSchedule) GetPipelineMode() *PipelineUpdateModes {
	if o == nil {
		return nil
	}
	return o.PipelineMode
}

func (o *ConnectionSquareDefaultUpdateSchedule) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionSquareDefaultUpdateSchedule) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionSquareDefaultUpdateSchedule) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionSquareDefaultUpdateSchedule) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionSquareDefaultUpdateSchedule) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionSquareDefaultUpdateSchedule) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

type ConnectionSquare struct {
	// The unique identifier of the connection.
	ID string `json:"id"`
	// The unique name of this connection.
	Name string               `json:"name"`
	Type ConnectionSquareType `json:"type"`
	// Whether this connection has been marked as active.
	Active bool `json:"active"`
	// The current status of the connection.
	Status ConnectionSquareStatus `json:"status"`
	// The date and time when then the connection was created.
	CreateDate time.Time `json:"createDate"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// When an update schedule is not defined for a connection, the default schedule is used. The default defined individually per `pipelineMode` and may be subject to change.
	DefaultUpdateSchedule []ConnectionSquareDefaultUpdateSchedule `json:"defaultUpdateSchedule"`
	// Under Developer Dashboard > OAuth
	ApplicationID  string `json:"applicationId"`
	SandboxAccount bool   `json:"sandboxAccount"`
}

func (c ConnectionSquare) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionSquare) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionSquare) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ConnectionSquare) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionSquare) GetType() ConnectionSquareType {
	if o == nil {
		return ConnectionSquareType("")
	}
	return o.Type
}

func (o *ConnectionSquare) GetActive() bool {
	if o == nil {
		return false
	}
	return o.Active
}

func (o *ConnectionSquare) GetStatus() ConnectionSquareStatus {
	if o == nil {
		return ConnectionSquareStatus("")
	}
	return o.Status
}

func (o *ConnectionSquare) GetCreateDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreateDate
}

func (o *ConnectionSquare) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionSquare) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionSquare) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionSquare) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionSquare) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionSquare) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionSquare) GetDefaultUpdateSchedule() []ConnectionSquareDefaultUpdateSchedule {
	if o == nil {
		return []ConnectionSquareDefaultUpdateSchedule{}
	}
	return o.DefaultUpdateSchedule
}

func (o *ConnectionSquare) GetApplicationID() string {
	if o == nil {
		return ""
	}
	return o.ApplicationID
}

func (o *ConnectionSquare) GetSandboxAccount() bool {
	if o == nil {
		return false
	}
	return o.SandboxAccount
}

type ConnectionSquareInput struct {
	// The unique name of this connection.
	Name string               `json:"name"`
	Type ConnectionSquareType `json:"type"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// Code retrieved from `/connections/oauth2-initiation`. **Note:** it is short-lived, therefore the connection creation should be done as soon as code is returned.
	Code string `json:"code"`
	// Under Developer Dashboard > OAuth
	ApplicationID string `json:"applicationId"`
	// Under Developer Dashboard > OAuth
	ApplicationSecret string `json:"applicationSecret"`
	SandboxAccount    bool   `json:"sandboxAccount"`
}

func (o *ConnectionSquareInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionSquareInput) GetType() ConnectionSquareType {
	if o == nil {
		return ConnectionSquareType("")
	}
	return o.Type
}

func (o *ConnectionSquareInput) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionSquareInput) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionSquareInput) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionSquareInput) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionSquareInput) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionSquareInput) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionSquareInput) GetCode() string {
	if o == nil {
		return ""
	}
	return o.Code
}

func (o *ConnectionSquareInput) GetApplicationID() string {
	if o == nil {
		return ""
	}
	return o.ApplicationID
}

func (o *ConnectionSquareInput) GetApplicationSecret() string {
	if o == nil {
		return ""
	}
	return o.ApplicationSecret
}

func (o *ConnectionSquareInput) GetSandboxAccount() bool {
	if o == nil {
		return false
	}
	return o.SandboxAccount
}
