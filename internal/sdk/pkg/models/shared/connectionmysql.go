// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
	"time"
)

type ConnectionMysqlType string

const (
	ConnectionMysqlTypeMysql ConnectionMysqlType = "MYSQL"
)

func (e ConnectionMysqlType) ToPointer() *ConnectionMysqlType {
	return &e
}

func (e *ConnectionMysqlType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MYSQL":
		*e = ConnectionMysqlType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionMysqlType: %v", v)
	}
}

// ConnectionMysqlStatus - The current status of the connection.
type ConnectionMysqlStatus string

const (
	ConnectionMysqlStatusUnknown     ConnectionMysqlStatus = "UNKNOWN"
	ConnectionMysqlStatusUp          ConnectionMysqlStatus = "UP"
	ConnectionMysqlStatusDown        ConnectionMysqlStatus = "DOWN"
	ConnectionMysqlStatusResize      ConnectionMysqlStatus = "RESIZE"
	ConnectionMysqlStatusMaintenance ConnectionMysqlStatus = "MAINTENANCE"
	ConnectionMysqlStatusQuota       ConnectionMysqlStatus = "QUOTA"
	ConnectionMysqlStatusCreating    ConnectionMysqlStatus = "CREATING"
)

func (e ConnectionMysqlStatus) ToPointer() *ConnectionMysqlStatus {
	return &e
}

func (e *ConnectionMysqlStatus) UnmarshalJSON(data []byte) error {
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
		*e = ConnectionMysqlStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionMysqlStatus: %v", v)
	}
}

type ConnectionMysqlDefaultUpdateSchedule struct {
	// The pipeline mode refers to how the pipeline fetches data changes from the source and how those changes are applied to the destination table. See <a target="_blank" href="https://docs.etleap.com/docs/documentation/ZG9jOjIyMjE3ODA2-introduction">the documentation</a> for more details.
	PipelineMode *PipelineUpdateModes `json:"pipelineMode,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
}

func (o *ConnectionMysqlDefaultUpdateSchedule) GetPipelineMode() *PipelineUpdateModes {
	if o == nil {
		return nil
	}
	return o.PipelineMode
}

func (o *ConnectionMysqlDefaultUpdateSchedule) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionMysqlDefaultUpdateSchedule) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionMysqlDefaultUpdateSchedule) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionMysqlDefaultUpdateSchedule) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionMysqlDefaultUpdateSchedule) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionMysqlDefaultUpdateSchedule) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

// ConnectionMysql - Specifies the properties of a database connection.
type ConnectionMysql struct {
	// The unique identifier of the connection.
	ID string `json:"id"`
	// The unique name of this connection.
	Name string              `json:"name"`
	Type ConnectionMysqlType `json:"type"`
	// Whether this connection has been marked as active.
	Active bool `json:"active"`
	// The current status of the connection.
	Status ConnectionMysqlStatus `json:"status"`
	// The date and time when then the connection was created.
	CreateDate time.Time `json:"createDate"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// When an update schedule is not defined for a connection, the default schedule is used. The default defined individually per `pipelineMode` and may be subject to change.
	DefaultUpdateSchedule            []ConnectionMysqlDefaultUpdateSchedule `json:"defaultUpdateSchedule"`
	RequireSslAndValidateCertificate *bool                                  `default:"true" json:"requireSslAndValidateCertificate"`
	// Should Etleap use MySQL binlogs to capture changes from this database? This setting cannot be changed later.
	CdcEnabled *bool `default:"false" json:"cdcEnabled"`
	// If you want Etleap to create pipelines for each source table automatically, specify the id of an Etleap destination connection here. If you want to create pipelines manually, omit this property.<br/><br/>If a database is not specified on this connection, then all databases will be replicated to the selected destination. Any databases not present in the destination will be created as needed.<br/><br/>If a database is specified on this connection, then only tables in that database will be replicated to the selected destination. Tables will be created in the database specified on the destination connection.
	AutoReplicate *string `json:"autoReplicate,omitempty"`
	// Should Etleap interpret columns with type Tinyint(1) as Boolean (i.e. true/false)?
	TinyInt1IsBoolean *bool `default:"false" json:"tinyInt1IsBoolean"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	Database  *string    `json:"database,omitempty"`
	Address   string     `json:"address"`
	Port      int64      `json:"port"`
	Username  string     `json:"username"`
	SSHConfig *SSHConfig `json:"sshConfig,omitempty"`
}

func (c ConnectionMysql) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionMysql) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionMysql) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ConnectionMysql) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionMysql) GetType() ConnectionMysqlType {
	if o == nil {
		return ConnectionMysqlType("")
	}
	return o.Type
}

func (o *ConnectionMysql) GetActive() bool {
	if o == nil {
		return false
	}
	return o.Active
}

func (o *ConnectionMysql) GetStatus() ConnectionMysqlStatus {
	if o == nil {
		return ConnectionMysqlStatus("")
	}
	return o.Status
}

func (o *ConnectionMysql) GetCreateDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreateDate
}

func (o *ConnectionMysql) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionMysql) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionMysql) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionMysql) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionMysql) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionMysql) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionMysql) GetDefaultUpdateSchedule() []ConnectionMysqlDefaultUpdateSchedule {
	if o == nil {
		return []ConnectionMysqlDefaultUpdateSchedule{}
	}
	return o.DefaultUpdateSchedule
}

func (o *ConnectionMysql) GetRequireSslAndValidateCertificate() *bool {
	if o == nil {
		return nil
	}
	return o.RequireSslAndValidateCertificate
}

func (o *ConnectionMysql) GetCdcEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.CdcEnabled
}

func (o *ConnectionMysql) GetAutoReplicate() *string {
	if o == nil {
		return nil
	}
	return o.AutoReplicate
}

func (o *ConnectionMysql) GetTinyInt1IsBoolean() *bool {
	if o == nil {
		return nil
	}
	return o.TinyInt1IsBoolean
}

func (o *ConnectionMysql) GetDatabase() *string {
	if o == nil {
		return nil
	}
	return o.Database
}

func (o *ConnectionMysql) GetAddress() string {
	if o == nil {
		return ""
	}
	return o.Address
}

func (o *ConnectionMysql) GetPort() int64 {
	if o == nil {
		return 0
	}
	return o.Port
}

func (o *ConnectionMysql) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *ConnectionMysql) GetSSHConfig() *SSHConfig {
	if o == nil {
		return nil
	}
	return o.SSHConfig
}

// ConnectionMysqlInput - Specifies the properties of a database connection.
type ConnectionMysqlInput struct {
	// The unique name of this connection.
	Name string              `json:"name"`
	Type ConnectionMysqlType `json:"type"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule                   *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	RequireSslAndValidateCertificate *bool                `default:"true" json:"requireSslAndValidateCertificate"`
	// Should Etleap use MySQL binlogs to capture changes from this database? This setting cannot be changed later.
	CdcEnabled *bool `default:"false" json:"cdcEnabled"`
	// If you want Etleap to create pipelines for each source table automatically, specify the id of an Etleap destination connection here. If you want to create pipelines manually, omit this property.<br/><br/>If a database is not specified on this connection, then all databases will be replicated to the selected destination. Any databases not present in the destination will be created as needed.<br/><br/>If a database is specified on this connection, then only tables in that database will be replicated to the selected destination. Tables will be created in the database specified on the destination connection.
	AutoReplicate *string `json:"autoReplicate,omitempty"`
	// Should Etleap interpret columns with type Tinyint(1) as Boolean (i.e. true/false)?
	TinyInt1IsBoolean *bool `default:"false" json:"tinyInt1IsBoolean"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	Database  *string    `json:"database,omitempty"`
	Address   string     `json:"address"`
	Port      int64      `json:"port"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	SSHConfig *SSHConfig `json:"sshConfig,omitempty"`
}

func (c ConnectionMysqlInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionMysqlInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionMysqlInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionMysqlInput) GetType() ConnectionMysqlType {
	if o == nil {
		return ConnectionMysqlType("")
	}
	return o.Type
}

func (o *ConnectionMysqlInput) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionMysqlInput) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionMysqlInput) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionMysqlInput) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionMysqlInput) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionMysqlInput) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionMysqlInput) GetRequireSslAndValidateCertificate() *bool {
	if o == nil {
		return nil
	}
	return o.RequireSslAndValidateCertificate
}

func (o *ConnectionMysqlInput) GetCdcEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.CdcEnabled
}

func (o *ConnectionMysqlInput) GetAutoReplicate() *string {
	if o == nil {
		return nil
	}
	return o.AutoReplicate
}

func (o *ConnectionMysqlInput) GetTinyInt1IsBoolean() *bool {
	if o == nil {
		return nil
	}
	return o.TinyInt1IsBoolean
}

func (o *ConnectionMysqlInput) GetDatabase() *string {
	if o == nil {
		return nil
	}
	return o.Database
}

func (o *ConnectionMysqlInput) GetAddress() string {
	if o == nil {
		return ""
	}
	return o.Address
}

func (o *ConnectionMysqlInput) GetPort() int64 {
	if o == nil {
		return 0
	}
	return o.Port
}

func (o *ConnectionMysqlInput) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *ConnectionMysqlInput) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *ConnectionMysqlInput) GetSSHConfig() *SSHConfig {
	if o == nil {
		return nil
	}
	return o.SSHConfig
}
