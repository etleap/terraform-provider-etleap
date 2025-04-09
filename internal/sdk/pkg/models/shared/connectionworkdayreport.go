// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
	"time"
)

// ConnectionWorkdayReportStatus - The current status of the connection.
type ConnectionWorkdayReportStatus string

const (
	ConnectionWorkdayReportStatusUnknown     ConnectionWorkdayReportStatus = "UNKNOWN"
	ConnectionWorkdayReportStatusUp          ConnectionWorkdayReportStatus = "UP"
	ConnectionWorkdayReportStatusDown        ConnectionWorkdayReportStatus = "DOWN"
	ConnectionWorkdayReportStatusResize      ConnectionWorkdayReportStatus = "RESIZE"
	ConnectionWorkdayReportStatusMaintenance ConnectionWorkdayReportStatus = "MAINTENANCE"
	ConnectionWorkdayReportStatusQuota       ConnectionWorkdayReportStatus = "QUOTA"
	ConnectionWorkdayReportStatusCreating    ConnectionWorkdayReportStatus = "CREATING"
)

func (e ConnectionWorkdayReportStatus) ToPointer() *ConnectionWorkdayReportStatus {
	return &e
}

func (e *ConnectionWorkdayReportStatus) UnmarshalJSON(data []byte) error {
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
		*e = ConnectionWorkdayReportStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionWorkdayReportStatus: %v", v)
	}
}

type ConnectionWorkdayReportDefaultUpdateSchedule struct {
	// The pipeline mode refers to how the pipeline fetches data changes from the source and how those changes are applied to the destination table. See <a target="_blank" href="https://docs.etleap.com/docs/documentation/ZG9jOjIyMjE3ODA2-introduction">the documentation</a> for more details.
	PipelineMode *PipelineUpdateModes `json:"pipelineMode,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
}

func (o *ConnectionWorkdayReportDefaultUpdateSchedule) GetPipelineMode() *PipelineUpdateModes {
	if o == nil {
		return nil
	}
	return o.PipelineMode
}

func (o *ConnectionWorkdayReportDefaultUpdateSchedule) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionWorkdayReportDefaultUpdateSchedule) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionWorkdayReportDefaultUpdateSchedule) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionWorkdayReportDefaultUpdateSchedule) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionWorkdayReportDefaultUpdateSchedule) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionWorkdayReportDefaultUpdateSchedule) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

type ConnectionWorkdayReportType string

const (
	ConnectionWorkdayReportTypeWorkdayReport ConnectionWorkdayReportType = "WORKDAY_REPORT"
)

func (e ConnectionWorkdayReportType) ToPointer() *ConnectionWorkdayReportType {
	return &e
}

func (e *ConnectionWorkdayReportType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "WORKDAY_REPORT":
		*e = ConnectionWorkdayReportType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionWorkdayReportType: %v", v)
	}
}

type ConnectionWorkdayReport struct {
	// The current status of the connection.
	Status ConnectionWorkdayReportStatus `json:"status"`
	// The unique name of this connection.
	Name string `json:"name"`
	// The date and time when then the connection was created.
	CreateDate time.Time `json:"createDate"`
	// When an update schedule is not defined for a connection, the default schedule is used. The default defined individually per `pipelineMode` and may be subject to change.
	DefaultUpdateSchedule []ConnectionWorkdayReportDefaultUpdateSchedule `json:"defaultUpdateSchedule"`
	// Whether this connection has been marked as active.
	Active bool                        `json:"active"`
	Type   ConnectionWorkdayReportType `json:"type"`
	// The unique identifier of the connection.
	ID string `json:"id"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	ReportURL      string               `json:"reportUrl"`
	Username       string               `json:"username"`
}

func (c ConnectionWorkdayReport) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionWorkdayReport) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionWorkdayReport) GetStatus() ConnectionWorkdayReportStatus {
	if o == nil {
		return ConnectionWorkdayReportStatus("")
	}
	return o.Status
}

func (o *ConnectionWorkdayReport) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionWorkdayReport) GetCreateDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreateDate
}

func (o *ConnectionWorkdayReport) GetDefaultUpdateSchedule() []ConnectionWorkdayReportDefaultUpdateSchedule {
	if o == nil {
		return []ConnectionWorkdayReportDefaultUpdateSchedule{}
	}
	return o.DefaultUpdateSchedule
}

func (o *ConnectionWorkdayReport) GetActive() bool {
	if o == nil {
		return false
	}
	return o.Active
}

func (o *ConnectionWorkdayReport) GetType() ConnectionWorkdayReportType {
	if o == nil {
		return ConnectionWorkdayReportType("")
	}
	return o.Type
}

func (o *ConnectionWorkdayReport) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ConnectionWorkdayReport) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionWorkdayReport) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionWorkdayReport) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionWorkdayReport) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionWorkdayReport) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionWorkdayReport) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionWorkdayReport) GetReportURL() string {
	if o == nil {
		return ""
	}
	return o.ReportURL
}

func (o *ConnectionWorkdayReport) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

type ConnectionWorkdayReportInput struct {
	// The unique name of this connection.
	Name string                      `json:"name"`
	Type ConnectionWorkdayReportType `json:"type"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	ReportURL      string               `json:"reportUrl"`
	Username       string               `json:"username"`
	Password       string               `json:"password"`
}

func (o *ConnectionWorkdayReportInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionWorkdayReportInput) GetType() ConnectionWorkdayReportType {
	if o == nil {
		return ConnectionWorkdayReportType("")
	}
	return o.Type
}

func (o *ConnectionWorkdayReportInput) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionWorkdayReportInput) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionWorkdayReportInput) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionWorkdayReportInput) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionWorkdayReportInput) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionWorkdayReportInput) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionWorkdayReportInput) GetReportURL() string {
	if o == nil {
		return ""
	}
	return o.ReportURL
}

func (o *ConnectionWorkdayReportInput) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *ConnectionWorkdayReportInput) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}
