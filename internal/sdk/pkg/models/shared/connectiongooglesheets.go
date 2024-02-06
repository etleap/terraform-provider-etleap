// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
	"time"
)

type ConnectionGoogleSheetsType string

const (
	ConnectionGoogleSheetsTypeGoogleSheets ConnectionGoogleSheetsType = "GOOGLE_SHEETS"
)

func (e ConnectionGoogleSheetsType) ToPointer() *ConnectionGoogleSheetsType {
	return &e
}

func (e *ConnectionGoogleSheetsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "GOOGLE_SHEETS":
		*e = ConnectionGoogleSheetsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionGoogleSheetsType: %v", v)
	}
}

// ConnectionGoogleSheetsStatus - The current status of the connection.
type ConnectionGoogleSheetsStatus string

const (
	ConnectionGoogleSheetsStatusUnknown     ConnectionGoogleSheetsStatus = "UNKNOWN"
	ConnectionGoogleSheetsStatusUp          ConnectionGoogleSheetsStatus = "UP"
	ConnectionGoogleSheetsStatusDown        ConnectionGoogleSheetsStatus = "DOWN"
	ConnectionGoogleSheetsStatusResize      ConnectionGoogleSheetsStatus = "RESIZE"
	ConnectionGoogleSheetsStatusMaintenance ConnectionGoogleSheetsStatus = "MAINTENANCE"
	ConnectionGoogleSheetsStatusQuota       ConnectionGoogleSheetsStatus = "QUOTA"
	ConnectionGoogleSheetsStatusCreating    ConnectionGoogleSheetsStatus = "CREATING"
)

func (e ConnectionGoogleSheetsStatus) ToPointer() *ConnectionGoogleSheetsStatus {
	return &e
}

func (e *ConnectionGoogleSheetsStatus) UnmarshalJSON(data []byte) error {
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
		*e = ConnectionGoogleSheetsStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionGoogleSheetsStatus: %v", v)
	}
}

type ConnectionGoogleSheetsDefaultUpdateSchedule struct {
	// The pipeline mode refers to how the pipeline fetches data changes from the source and how those changes are applied to the destination table. See <a target="_blank" href="https://docs.etleap.com/docs/documentation/ZG9jOjIyMjE3ODA2-introduction">the documentation</a> for more details.
	PipelineMode *PipelineUpdateModes `json:"pipelineMode,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
}

func (o *ConnectionGoogleSheetsDefaultUpdateSchedule) GetPipelineMode() *PipelineUpdateModes {
	if o == nil {
		return nil
	}
	return o.PipelineMode
}

func (o *ConnectionGoogleSheetsDefaultUpdateSchedule) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionGoogleSheetsDefaultUpdateSchedule) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionGoogleSheetsDefaultUpdateSchedule) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionGoogleSheetsDefaultUpdateSchedule) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionGoogleSheetsDefaultUpdateSchedule) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionGoogleSheetsDefaultUpdateSchedule) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

type ConnectionGoogleSheets struct {
	// The unique identifier of the connection.
	ID string `json:"id"`
	// The unique name of this connection.
	Name string                     `json:"name"`
	Type ConnectionGoogleSheetsType `json:"type"`
	// Whether this connection has been marked as active.
	Active bool `json:"active"`
	// The current status of the connection.
	Status ConnectionGoogleSheetsStatus `json:"status"`
	// The date and time when then the connection was created.
	CreateDate time.Time `json:"createDate"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// When an update schedule is not defined for a connection, the default schedule is used. The default defined individually per `pipelineMode` and may be subject to change.
	DefaultUpdateSchedule []ConnectionGoogleSheetsDefaultUpdateSchedule `json:"defaultUpdateSchedule"`
	Username              string                                        `json:"username"`
}

func (c ConnectionGoogleSheets) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionGoogleSheets) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionGoogleSheets) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ConnectionGoogleSheets) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionGoogleSheets) GetType() ConnectionGoogleSheetsType {
	if o == nil {
		return ConnectionGoogleSheetsType("")
	}
	return o.Type
}

func (o *ConnectionGoogleSheets) GetActive() bool {
	if o == nil {
		return false
	}
	return o.Active
}

func (o *ConnectionGoogleSheets) GetStatus() ConnectionGoogleSheetsStatus {
	if o == nil {
		return ConnectionGoogleSheetsStatus("")
	}
	return o.Status
}

func (o *ConnectionGoogleSheets) GetCreateDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreateDate
}

func (o *ConnectionGoogleSheets) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionGoogleSheets) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionGoogleSheets) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionGoogleSheets) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionGoogleSheets) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionGoogleSheets) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionGoogleSheets) GetDefaultUpdateSchedule() []ConnectionGoogleSheetsDefaultUpdateSchedule {
	if o == nil {
		return []ConnectionGoogleSheetsDefaultUpdateSchedule{}
	}
	return o.DefaultUpdateSchedule
}

func (o *ConnectionGoogleSheets) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

type ConnectionGoogleSheetsInput struct {
	// The unique name of this connection.
	Name string                     `json:"name"`
	Type ConnectionGoogleSheetsType `json:"type"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// Code retrieved from `/connections/oauth2-initiation`. **Note:** it is short-lived, therefore the connection creation should be done as soon as code is returned.
	Code string `json:"code"`
}

func (o *ConnectionGoogleSheetsInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionGoogleSheetsInput) GetType() ConnectionGoogleSheetsType {
	if o == nil {
		return ConnectionGoogleSheetsType("")
	}
	return o.Type
}

func (o *ConnectionGoogleSheetsInput) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionGoogleSheetsInput) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionGoogleSheetsInput) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionGoogleSheetsInput) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionGoogleSheetsInput) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionGoogleSheetsInput) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionGoogleSheetsInput) GetCode() string {
	if o == nil {
		return ""
	}
	return o.Code
}
