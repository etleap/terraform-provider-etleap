// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
	"time"
)

// Status - The current status of the connection.
type Status string

const (
	StatusUnknown     Status = "UNKNOWN"
	StatusUp          Status = "UP"
	StatusDown        Status = "DOWN"
	StatusResize      Status = "RESIZE"
	StatusMaintenance Status = "MAINTENANCE"
	StatusQuota       Status = "QUOTA"
	StatusCreating    Status = "CREATING"
)

func (e Status) ToPointer() *Status {
	return &e
}

func (e *Status) UnmarshalJSON(data []byte) error {
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
		*e = Status(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Status: %v", v)
	}
}

type DefaultUpdateSchedule struct {
	// The pipeline mode refers to how the pipeline fetches data changes from the source and how those changes are applied to the destination table. See <a target="_blank" href="https://docs.etleap.com/docs/documentation/ZG9jOjIyMjE3ODA2-introduction">the documentation</a> for more details.
	PipelineMode *PipelineUpdateModes `json:"pipelineMode,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
}

func (o *DefaultUpdateSchedule) GetPipelineMode() *PipelineUpdateModes {
	if o == nil {
		return nil
	}
	return o.PipelineMode
}

func (o *DefaultUpdateSchedule) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *DefaultUpdateSchedule) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *DefaultUpdateSchedule) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *DefaultUpdateSchedule) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *DefaultUpdateSchedule) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *DefaultUpdateSchedule) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

type ConnectionSQLServerType string

const (
	ConnectionSQLServerTypeSQLServer ConnectionSQLServerType = "SQL_SERVER"
)

func (e ConnectionSQLServerType) ToPointer() *ConnectionSQLServerType {
	return &e
}

func (e *ConnectionSQLServerType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SQL_SERVER":
		*e = ConnectionSQLServerType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionSQLServerType: %v", v)
	}
}

// ConnectionSQLServer - Specifies the location of a database.
type ConnectionSQLServer struct {
	// The current status of the connection.
	Status Status `json:"status"`
	// The unique name of this connection.
	Name string `json:"name"`
	// The date and time when then the connection was created.
	CreateDate time.Time `json:"createDate"`
	// When an update schedule is not defined for a connection, the default schedule is used. The default defined individually per `pipelineMode` and may be subject to change.
	DefaultUpdateSchedule []DefaultUpdateSchedule `json:"defaultUpdateSchedule"`
	// Whether this connection has been marked as active.
	Active bool                    `json:"active"`
	Type   ConnectionSQLServerType `json:"type"`
	// The unique identifier of the connection.
	ID string `json:"id"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// Should Etleap use the SQL Server transaction log to capture changes from this database? This setting cannot be changed later.
	CdcEnabled *bool `default:"false" json:"cdcEnabled"`
	// If not specified, the default schema will be used.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	Schema    *string    `json:"schema,omitempty"`
	Username  string     `json:"username"`
	SSHConfig *SSHConfig `json:"sshConfig,omitempty"`
	Port      int64      `json:"port"`
	Address   string     `json:"address"`
	Database  string     `json:"database"`
}

func (c ConnectionSQLServer) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionSQLServer) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionSQLServer) GetStatus() Status {
	if o == nil {
		return Status("")
	}
	return o.Status
}

func (o *ConnectionSQLServer) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionSQLServer) GetCreateDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreateDate
}

func (o *ConnectionSQLServer) GetDefaultUpdateSchedule() []DefaultUpdateSchedule {
	if o == nil {
		return []DefaultUpdateSchedule{}
	}
	return o.DefaultUpdateSchedule
}

func (o *ConnectionSQLServer) GetActive() bool {
	if o == nil {
		return false
	}
	return o.Active
}

func (o *ConnectionSQLServer) GetType() ConnectionSQLServerType {
	if o == nil {
		return ConnectionSQLServerType("")
	}
	return o.Type
}

func (o *ConnectionSQLServer) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ConnectionSQLServer) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionSQLServer) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionSQLServer) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionSQLServer) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionSQLServer) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionSQLServer) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionSQLServer) GetCdcEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.CdcEnabled
}

func (o *ConnectionSQLServer) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *ConnectionSQLServer) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *ConnectionSQLServer) GetSSHConfig() *SSHConfig {
	if o == nil {
		return nil
	}
	return o.SSHConfig
}

func (o *ConnectionSQLServer) GetPort() int64 {
	if o == nil {
		return 0
	}
	return o.Port
}

func (o *ConnectionSQLServer) GetAddress() string {
	if o == nil {
		return ""
	}
	return o.Address
}

func (o *ConnectionSQLServer) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

// ConnectionSQLServerInput - Specifies the location of a database.
type ConnectionSQLServerInput struct {
	// The unique name of this connection.
	Name string                  `json:"name"`
	Type ConnectionSQLServerType `json:"type"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// Should Etleap use the SQL Server transaction log to capture changes from this database? This setting cannot be changed later.
	CdcEnabled *bool `default:"false" json:"cdcEnabled"`
	// If not specified, the default schema will be used.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	Schema    *string    `json:"schema,omitempty"`
	Username  string     `json:"username"`
	SSHConfig *SSHConfig `json:"sshConfig,omitempty"`
	Password  string     `json:"password"`
	Port      int64      `json:"port"`
	Address   string     `json:"address"`
	Database  string     `json:"database"`
}

func (c ConnectionSQLServerInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionSQLServerInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionSQLServerInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionSQLServerInput) GetType() ConnectionSQLServerType {
	if o == nil {
		return ConnectionSQLServerType("")
	}
	return o.Type
}

func (o *ConnectionSQLServerInput) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionSQLServerInput) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionSQLServerInput) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionSQLServerInput) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionSQLServerInput) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionSQLServerInput) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionSQLServerInput) GetCdcEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.CdcEnabled
}

func (o *ConnectionSQLServerInput) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *ConnectionSQLServerInput) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *ConnectionSQLServerInput) GetSSHConfig() *SSHConfig {
	if o == nil {
		return nil
	}
	return o.SSHConfig
}

func (o *ConnectionSQLServerInput) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *ConnectionSQLServerInput) GetPort() int64 {
	if o == nil {
		return 0
	}
	return o.Port
}

func (o *ConnectionSQLServerInput) GetAddress() string {
	if o == nil {
		return ""
	}
	return o.Address
}

func (o *ConnectionSQLServerInput) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}
