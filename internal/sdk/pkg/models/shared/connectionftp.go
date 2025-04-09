// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
	"time"
)

type ConnectionFtpType string

const (
	ConnectionFtpTypeFtp ConnectionFtpType = "FTP"
)

func (e ConnectionFtpType) ToPointer() *ConnectionFtpType {
	return &e
}

func (e *ConnectionFtpType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "FTP":
		*e = ConnectionFtpType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionFtpType: %v", v)
	}
}

// ConnectionFtpStatus - The current status of the connection.
type ConnectionFtpStatus string

const (
	ConnectionFtpStatusUnknown     ConnectionFtpStatus = "UNKNOWN"
	ConnectionFtpStatusUp          ConnectionFtpStatus = "UP"
	ConnectionFtpStatusDown        ConnectionFtpStatus = "DOWN"
	ConnectionFtpStatusResize      ConnectionFtpStatus = "RESIZE"
	ConnectionFtpStatusMaintenance ConnectionFtpStatus = "MAINTENANCE"
	ConnectionFtpStatusQuota       ConnectionFtpStatus = "QUOTA"
	ConnectionFtpStatusCreating    ConnectionFtpStatus = "CREATING"
)

func (e ConnectionFtpStatus) ToPointer() *ConnectionFtpStatus {
	return &e
}

func (e *ConnectionFtpStatus) UnmarshalJSON(data []byte) error {
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
		*e = ConnectionFtpStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionFtpStatus: %v", v)
	}
}

type ConnectionFtpDefaultUpdateSchedule struct {
	// The pipeline mode refers to how the pipeline fetches data changes from the source and how those changes are applied to the destination table. See <a target="_blank" href="https://docs.etleap.com/docs/documentation/ZG9jOjIyMjE3ODA2-introduction">the documentation</a> for more details.
	PipelineMode *PipelineUpdateModes `json:"pipelineMode,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
}

func (o *ConnectionFtpDefaultUpdateSchedule) GetPipelineMode() *PipelineUpdateModes {
	if o == nil {
		return nil
	}
	return o.PipelineMode
}

func (o *ConnectionFtpDefaultUpdateSchedule) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionFtpDefaultUpdateSchedule) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionFtpDefaultUpdateSchedule) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionFtpDefaultUpdateSchedule) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionFtpDefaultUpdateSchedule) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionFtpDefaultUpdateSchedule) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

type ConnectionFtp struct {
	// The unique identifier of the connection.
	ID string `json:"id"`
	// The unique name of this connection.
	Name string            `json:"name"`
	Type ConnectionFtpType `json:"type"`
	// Whether this connection has been marked as active.
	Active bool `json:"active"`
	// The current status of the connection.
	Status ConnectionFtpStatus `json:"status"`
	// The date and time when then the connection was created.
	CreateDate time.Time `json:"createDate"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// When an update schedule is not defined for a connection, the default schedule is used. The default defined individually per `pipelineMode` and may be subject to change.
	DefaultUpdateSchedule []ConnectionFtpDefaultUpdateSchedule `json:"defaultUpdateSchedule"`
	// E.g. 'ftp.etleap.com' or '10.0.0.2'.
	Hostname    string `json:"hostname"`
	Port        int64  `json:"port"`
	Username    string `json:"username"`
	PassiveMode bool   `json:"passiveMode"`
}

func (c ConnectionFtp) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionFtp) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionFtp) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ConnectionFtp) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionFtp) GetType() ConnectionFtpType {
	if o == nil {
		return ConnectionFtpType("")
	}
	return o.Type
}

func (o *ConnectionFtp) GetActive() bool {
	if o == nil {
		return false
	}
	return o.Active
}

func (o *ConnectionFtp) GetStatus() ConnectionFtpStatus {
	if o == nil {
		return ConnectionFtpStatus("")
	}
	return o.Status
}

func (o *ConnectionFtp) GetCreateDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreateDate
}

func (o *ConnectionFtp) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionFtp) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionFtp) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionFtp) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionFtp) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionFtp) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionFtp) GetDefaultUpdateSchedule() []ConnectionFtpDefaultUpdateSchedule {
	if o == nil {
		return []ConnectionFtpDefaultUpdateSchedule{}
	}
	return o.DefaultUpdateSchedule
}

func (o *ConnectionFtp) GetHostname() string {
	if o == nil {
		return ""
	}
	return o.Hostname
}

func (o *ConnectionFtp) GetPort() int64 {
	if o == nil {
		return 0
	}
	return o.Port
}

func (o *ConnectionFtp) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *ConnectionFtp) GetPassiveMode() bool {
	if o == nil {
		return false
	}
	return o.PassiveMode
}

type ConnectionFtpInput struct {
	// The unique name of this connection.
	Name string            `json:"name"`
	Type ConnectionFtpType `json:"type"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// E.g. 'ftp.etleap.com' or '10.0.0.2'.
	Hostname    string `json:"hostname"`
	Port        int64  `json:"port"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	PassiveMode bool   `json:"passiveMode"`
}

func (o *ConnectionFtpInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionFtpInput) GetType() ConnectionFtpType {
	if o == nil {
		return ConnectionFtpType("")
	}
	return o.Type
}

func (o *ConnectionFtpInput) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionFtpInput) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionFtpInput) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionFtpInput) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionFtpInput) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionFtpInput) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionFtpInput) GetHostname() string {
	if o == nil {
		return ""
	}
	return o.Hostname
}

func (o *ConnectionFtpInput) GetPort() int64 {
	if o == nil {
		return 0
	}
	return o.Port
}

func (o *ConnectionFtpInput) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *ConnectionFtpInput) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *ConnectionFtpInput) GetPassiveMode() bool {
	if o == nil {
		return false
	}
	return o.PassiveMode
}
