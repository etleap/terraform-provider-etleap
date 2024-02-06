// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionOracleUpdateType string

const (
	ConnectionOracleUpdateTypeOracle ConnectionOracleUpdateType = "ORACLE"
)

func (e ConnectionOracleUpdateType) ToPointer() *ConnectionOracleUpdateType {
	return &e
}

func (e *ConnectionOracleUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ORACLE":
		*e = ConnectionOracleUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionOracleUpdateType: %v", v)
	}
}

type ConnectionOracleUpdateSSHConfigurationUpdate struct {
	// The server address for the SSH connection.
	Address *string `json:"address,omitempty"`
	// The username for the SSH connection.
	Username *string `json:"username,omitempty"`
}

func (o *ConnectionOracleUpdateSSHConfigurationUpdate) GetAddress() *string {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *ConnectionOracleUpdateSSHConfigurationUpdate) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

// ConnectionOracleUpdate - Specifies the location of a database.
type ConnectionOracleUpdate struct {
	// The unique name of this connection.
	Name *string                    `json:"name,omitempty"`
	Type ConnectionOracleUpdateType `json:"type"`
	// Whether this connection should be marked as active.
	Active *bool `json:"active,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// If not specified, the default schema will be used.
	Schema    *string                                       `json:"schema,omitempty"`
	Address   *string                                       `json:"address,omitempty"`
	Port      *int64                                        `json:"port,omitempty"`
	Database  *string                                       `json:"database,omitempty"`
	Username  *string                                       `json:"username,omitempty"`
	Password  *string                                       `json:"password,omitempty"`
	SSHConfig *ConnectionOracleUpdateSSHConfigurationUpdate `json:"sshConfig,omitempty"`
}

func (o *ConnectionOracleUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionOracleUpdate) GetType() ConnectionOracleUpdateType {
	if o == nil {
		return ConnectionOracleUpdateType("")
	}
	return o.Type
}

func (o *ConnectionOracleUpdate) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ConnectionOracleUpdate) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionOracleUpdate) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionOracleUpdate) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionOracleUpdate) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionOracleUpdate) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionOracleUpdate) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionOracleUpdate) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *ConnectionOracleUpdate) GetAddress() *string {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *ConnectionOracleUpdate) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *ConnectionOracleUpdate) GetDatabase() *string {
	if o == nil {
		return nil
	}
	return o.Database
}

func (o *ConnectionOracleUpdate) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *ConnectionOracleUpdate) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *ConnectionOracleUpdate) GetSSHConfig() *ConnectionOracleUpdateSSHConfigurationUpdate {
	if o == nil {
		return nil
	}
	return o.SSHConfig
}
