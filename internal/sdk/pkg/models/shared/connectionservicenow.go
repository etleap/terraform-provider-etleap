// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
	"time"
)

type ConnectionServiceNowType string

const (
	ConnectionServiceNowTypeServiceNow ConnectionServiceNowType = "SERVICE_NOW"
)

func (e ConnectionServiceNowType) ToPointer() *ConnectionServiceNowType {
	return &e
}

func (e *ConnectionServiceNowType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SERVICE_NOW":
		*e = ConnectionServiceNowType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionServiceNowType: %v", v)
	}
}

// ConnectionServiceNowStatus - The current status of the connection.
type ConnectionServiceNowStatus string

const (
	ConnectionServiceNowStatusUnknown     ConnectionServiceNowStatus = "UNKNOWN"
	ConnectionServiceNowStatusUp          ConnectionServiceNowStatus = "UP"
	ConnectionServiceNowStatusDown        ConnectionServiceNowStatus = "DOWN"
	ConnectionServiceNowStatusResize      ConnectionServiceNowStatus = "RESIZE"
	ConnectionServiceNowStatusMaintenance ConnectionServiceNowStatus = "MAINTENANCE"
	ConnectionServiceNowStatusQuota       ConnectionServiceNowStatus = "QUOTA"
	ConnectionServiceNowStatusCreating    ConnectionServiceNowStatus = "CREATING"
)

func (e ConnectionServiceNowStatus) ToPointer() *ConnectionServiceNowStatus {
	return &e
}

func (e *ConnectionServiceNowStatus) UnmarshalJSON(data []byte) error {
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
		*e = ConnectionServiceNowStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionServiceNowStatus: %v", v)
	}
}

type ConnectionServiceNowDefaultUpdateSchedule struct {
	// The pipeline mode refers to how the pipeline fetches data changes from the source and how those changes are applied to the destination table. See <a target="_blank" href="https://docs.etleap.com/docs/documentation/ZG9jOjIyMjE3ODA2-introduction">the documentation</a> for more details.
	PipelineMode *PipelineUpdateModes `json:"pipelineMode,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
}

func (o *ConnectionServiceNowDefaultUpdateSchedule) GetPipelineMode() *PipelineUpdateModes {
	if o == nil {
		return nil
	}
	return o.PipelineMode
}

func (o *ConnectionServiceNowDefaultUpdateSchedule) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionServiceNowDefaultUpdateSchedule) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionServiceNowDefaultUpdateSchedule) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionServiceNowDefaultUpdateSchedule) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionServiceNowDefaultUpdateSchedule) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionServiceNowDefaultUpdateSchedule) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

type ConnectionServiceNow struct {
	// The unique identifier of the connection.
	ID string `json:"id"`
	// The unique name of this connection.
	Name string                   `json:"name"`
	Type ConnectionServiceNowType `json:"type"`
	// Whether this connection has been marked as active.
	Active bool `json:"active"`
	// The current status of the connection.
	Status ConnectionServiceNowStatus `json:"status"`
	// The date and time when then the connection was created.
	CreateDate time.Time `json:"createDate"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// When an update schedule is not defined for a connection, the default schedule is used. The default defined individually per `pipelineMode` and may be subject to change.
	DefaultUpdateSchedule []ConnectionServiceNowDefaultUpdateSchedule `json:"defaultUpdateSchedule"`
}

func (c ConnectionServiceNow) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionServiceNow) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionServiceNow) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ConnectionServiceNow) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionServiceNow) GetType() ConnectionServiceNowType {
	if o == nil {
		return ConnectionServiceNowType("")
	}
	return o.Type
}

func (o *ConnectionServiceNow) GetActive() bool {
	if o == nil {
		return false
	}
	return o.Active
}

func (o *ConnectionServiceNow) GetStatus() ConnectionServiceNowStatus {
	if o == nil {
		return ConnectionServiceNowStatus("")
	}
	return o.Status
}

func (o *ConnectionServiceNow) GetCreateDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreateDate
}

func (o *ConnectionServiceNow) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionServiceNow) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionServiceNow) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionServiceNow) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionServiceNow) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionServiceNow) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionServiceNow) GetDefaultUpdateSchedule() []ConnectionServiceNowDefaultUpdateSchedule {
	if o == nil {
		return []ConnectionServiceNowDefaultUpdateSchedule{}
	}
	return o.DefaultUpdateSchedule
}

type ConnectionServiceNowInput struct {
	// The unique name of this connection.
	Name string                   `json:"name"`
	Type ConnectionServiceNowType `json:"type"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// E.g., 12345.service-now.com
	SvnInstanceURL string `json:"svnInstanceUrl"`
	// ServiceNow Instance Login
	Username string `json:"username"`
	// ServiceNow Instance Password
	Password string `json:"password"`
}

func (o *ConnectionServiceNowInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionServiceNowInput) GetType() ConnectionServiceNowType {
	if o == nil {
		return ConnectionServiceNowType("")
	}
	return o.Type
}

func (o *ConnectionServiceNowInput) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionServiceNowInput) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionServiceNowInput) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionServiceNowInput) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionServiceNowInput) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionServiceNowInput) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionServiceNowInput) GetSvnInstanceURL() string {
	if o == nil {
		return ""
	}
	return o.SvnInstanceURL
}

func (o *ConnectionServiceNowInput) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *ConnectionServiceNowInput) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}
