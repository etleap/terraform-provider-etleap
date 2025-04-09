// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionSQLServerUpdateType string

const (
	ConnectionSQLServerUpdateTypeSQLServer ConnectionSQLServerUpdateType = "SQL_SERVER"
)

func (e ConnectionSQLServerUpdateType) ToPointer() *ConnectionSQLServerUpdateType {
	return &e
}

func (e *ConnectionSQLServerUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SQL_SERVER":
		*e = ConnectionSQLServerUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionSQLServerUpdateType: %v", v)
	}
}

type SSHConfigurationUpdate struct {
	// The username for the SSH connection.
	Username *string `json:"username,omitempty"`
	// The server address for the SSH connection.
	Address *string `json:"address,omitempty"`
}

func (o *SSHConfigurationUpdate) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *SSHConfigurationUpdate) GetAddress() *string {
	if o == nil {
		return nil
	}
	return o.Address
}

// ConnectionSQLServerUpdate - Specifies the location of a database.
type ConnectionSQLServerUpdate struct {
	// Whether this connection should be marked as active.
	Active *bool                         `json:"active,omitempty"`
	Type   ConnectionSQLServerUpdateType `json:"type"`
	// The unique name of this connection.
	Name *string `json:"name,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// If not specified, the default schema will be used.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	Schema    *string                 `json:"schema,omitempty"`
	Username  *string                 `json:"username,omitempty"`
	SSHConfig *SSHConfigurationUpdate `json:"sshConfig,omitempty"`
	Password  *string                 `json:"password,omitempty"`
	Port      *int64                  `json:"port,omitempty"`
	Address   *string                 `json:"address,omitempty"`
	Database  *string                 `json:"database,omitempty"`
}

func (o *ConnectionSQLServerUpdate) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ConnectionSQLServerUpdate) GetType() ConnectionSQLServerUpdateType {
	if o == nil {
		return ConnectionSQLServerUpdateType("")
	}
	return o.Type
}

func (o *ConnectionSQLServerUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionSQLServerUpdate) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionSQLServerUpdate) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionSQLServerUpdate) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionSQLServerUpdate) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionSQLServerUpdate) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionSQLServerUpdate) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionSQLServerUpdate) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *ConnectionSQLServerUpdate) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *ConnectionSQLServerUpdate) GetSSHConfig() *SSHConfigurationUpdate {
	if o == nil {
		return nil
	}
	return o.SSHConfig
}

func (o *ConnectionSQLServerUpdate) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *ConnectionSQLServerUpdate) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *ConnectionSQLServerUpdate) GetAddress() *string {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *ConnectionSQLServerUpdate) GetDatabase() *string {
	if o == nil {
		return nil
	}
	return o.Database
}
