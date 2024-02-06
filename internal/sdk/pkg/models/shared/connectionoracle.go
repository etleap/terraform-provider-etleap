// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
	"time"
)

type ConnectionOracleType string

const (
	ConnectionOracleTypeOracle ConnectionOracleType = "ORACLE"
)

func (e ConnectionOracleType) ToPointer() *ConnectionOracleType {
	return &e
}

func (e *ConnectionOracleType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ORACLE":
		*e = ConnectionOracleType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionOracleType: %v", v)
	}
}

// ConnectionOracleStatus - The current status of the connection.
type ConnectionOracleStatus string

const (
	ConnectionOracleStatusUnknown     ConnectionOracleStatus = "UNKNOWN"
	ConnectionOracleStatusUp          ConnectionOracleStatus = "UP"
	ConnectionOracleStatusDown        ConnectionOracleStatus = "DOWN"
	ConnectionOracleStatusResize      ConnectionOracleStatus = "RESIZE"
	ConnectionOracleStatusMaintenance ConnectionOracleStatus = "MAINTENANCE"
	ConnectionOracleStatusQuota       ConnectionOracleStatus = "QUOTA"
	ConnectionOracleStatusCreating    ConnectionOracleStatus = "CREATING"
)

func (e ConnectionOracleStatus) ToPointer() *ConnectionOracleStatus {
	return &e
}

func (e *ConnectionOracleStatus) UnmarshalJSON(data []byte) error {
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
		*e = ConnectionOracleStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionOracleStatus: %v", v)
	}
}

type ConnectionOracleDefaultUpdateSchedule struct {
	// The pipeline mode refers to how the pipeline fetches data changes from the source and how those changes are applied to the destination table. See <a target="_blank" href="https://docs.etleap.com/docs/documentation/ZG9jOjIyMjE3ODA2-introduction">the documentation</a> for more details.
	PipelineMode *PipelineUpdateModes `json:"pipelineMode,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
}

func (o *ConnectionOracleDefaultUpdateSchedule) GetPipelineMode() *PipelineUpdateModes {
	if o == nil {
		return nil
	}
	return o.PipelineMode
}

func (o *ConnectionOracleDefaultUpdateSchedule) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionOracleDefaultUpdateSchedule) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionOracleDefaultUpdateSchedule) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionOracleDefaultUpdateSchedule) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionOracleDefaultUpdateSchedule) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionOracleDefaultUpdateSchedule) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

// ConnectionOracle - Specifies the location of a database.
type ConnectionOracle struct {
	// The unique identifier of the connection.
	ID string `json:"id"`
	// The unique name of this connection.
	Name string               `json:"name"`
	Type ConnectionOracleType `json:"type"`
	// Whether this connection has been marked as active.
	Active bool `json:"active"`
	// The current status of the connection.
	Status ConnectionOracleStatus `json:"status"`
	// The date and time when then the connection was created.
	CreateDate time.Time `json:"createDate"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// When an update schedule is not defined for a connection, the default schedule is used. The default defined individually per `pipelineMode` and may be subject to change.
	DefaultUpdateSchedule []ConnectionOracleDefaultUpdateSchedule `json:"defaultUpdateSchedule"`
	// If not specified, the default schema will be used.
	Schema *string `json:"schema,omitempty"`
	// Should Etleap use replication logs to capture changes from this database? This setting cannot be changed later.
	CdcEnabled *bool      `default:"false" json:"cdcEnabled"`
	Address    string     `json:"address"`
	Port       int64      `json:"port"`
	Database   string     `json:"database"`
	Username   string     `json:"username"`
	SSHConfig  *SSHConfig `json:"sshConfig,omitempty"`
}

func (c ConnectionOracle) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionOracle) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionOracle) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ConnectionOracle) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionOracle) GetType() ConnectionOracleType {
	if o == nil {
		return ConnectionOracleType("")
	}
	return o.Type
}

func (o *ConnectionOracle) GetActive() bool {
	if o == nil {
		return false
	}
	return o.Active
}

func (o *ConnectionOracle) GetStatus() ConnectionOracleStatus {
	if o == nil {
		return ConnectionOracleStatus("")
	}
	return o.Status
}

func (o *ConnectionOracle) GetCreateDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreateDate
}

func (o *ConnectionOracle) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionOracle) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionOracle) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionOracle) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionOracle) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionOracle) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionOracle) GetDefaultUpdateSchedule() []ConnectionOracleDefaultUpdateSchedule {
	if o == nil {
		return []ConnectionOracleDefaultUpdateSchedule{}
	}
	return o.DefaultUpdateSchedule
}

func (o *ConnectionOracle) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *ConnectionOracle) GetCdcEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.CdcEnabled
}

func (o *ConnectionOracle) GetAddress() string {
	if o == nil {
		return ""
	}
	return o.Address
}

func (o *ConnectionOracle) GetPort() int64 {
	if o == nil {
		return 0
	}
	return o.Port
}

func (o *ConnectionOracle) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *ConnectionOracle) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *ConnectionOracle) GetSSHConfig() *SSHConfig {
	if o == nil {
		return nil
	}
	return o.SSHConfig
}

// ConnectionOracleInput - Specifies the location of a database.
type ConnectionOracleInput struct {
	// The unique name of this connection.
	Name string               `json:"name"`
	Type ConnectionOracleType `json:"type"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// If not specified, the default schema will be used.
	Schema *string `json:"schema,omitempty"`
	// Should Etleap use replication logs to capture changes from this database? This setting cannot be changed later.
	CdcEnabled *bool      `default:"false" json:"cdcEnabled"`
	Address    string     `json:"address"`
	Port       int64      `json:"port"`
	Database   string     `json:"database"`
	Username   string     `json:"username"`
	Password   string     `json:"password"`
	SSHConfig  *SSHConfig `json:"sshConfig,omitempty"`
}

func (c ConnectionOracleInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionOracleInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionOracleInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionOracleInput) GetType() ConnectionOracleType {
	if o == nil {
		return ConnectionOracleType("")
	}
	return o.Type
}

func (o *ConnectionOracleInput) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionOracleInput) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionOracleInput) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionOracleInput) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionOracleInput) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionOracleInput) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionOracleInput) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *ConnectionOracleInput) GetCdcEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.CdcEnabled
}

func (o *ConnectionOracleInput) GetAddress() string {
	if o == nil {
		return ""
	}
	return o.Address
}

func (o *ConnectionOracleInput) GetPort() int64 {
	if o == nil {
		return 0
	}
	return o.Port
}

func (o *ConnectionOracleInput) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *ConnectionOracleInput) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *ConnectionOracleInput) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *ConnectionOracleInput) GetSSHConfig() *SSHConfig {
	if o == nil {
		return nil
	}
	return o.SSHConfig
}
