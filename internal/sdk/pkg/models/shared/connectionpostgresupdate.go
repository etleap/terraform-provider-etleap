// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionPostgresUpdateType string

const (
	ConnectionPostgresUpdateTypePostgres ConnectionPostgresUpdateType = "POSTGRES"
)

func (e ConnectionPostgresUpdateType) ToPointer() *ConnectionPostgresUpdateType {
	return &e
}

func (e *ConnectionPostgresUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "POSTGRES":
		*e = ConnectionPostgresUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionPostgresUpdateType: %v", v)
	}
}

type ConnectionPostgresUpdateSSHConfigurationUpdate struct {
	// The server address for the SSH connection.
	Address *string `json:"address,omitempty"`
	// The username for the SSH connection.
	Username *string `json:"username,omitempty"`
}

func (o *ConnectionPostgresUpdateSSHConfigurationUpdate) GetAddress() *string {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *ConnectionPostgresUpdateSSHConfigurationUpdate) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

// ConnectionPostgresUpdate - Specifies the location of a database.
type ConnectionPostgresUpdate struct {
	// The unique name of this connection.
	Name *string                      `json:"name,omitempty"`
	Type ConnectionPostgresUpdateType `json:"type"`
	// Whether this connection should be marked as active.
	Active *bool `json:"active,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// If not specified, the default schema will be used.
	Schema *string `json:"schema,omitempty"`
	// If you want Etleap to create pipelines for each source table automatically, specify the id of an Etleap destination connection here. If you want to create pipelines manually, omit this property. Note that only the connection owner can change this setting.
	AutoReplicate *string                                         `json:"autoReplicate,omitempty"`
	Address       *string                                         `json:"address,omitempty"`
	Port          *int64                                          `json:"port,omitempty"`
	Database      *string                                         `json:"database,omitempty"`
	Username      *string                                         `json:"username,omitempty"`
	Password      *string                                         `json:"password,omitempty"`
	SSHConfig     *ConnectionPostgresUpdateSSHConfigurationUpdate `json:"sshConfig,omitempty"`
}

func (o *ConnectionPostgresUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionPostgresUpdate) GetType() ConnectionPostgresUpdateType {
	if o == nil {
		return ConnectionPostgresUpdateType("")
	}
	return o.Type
}

func (o *ConnectionPostgresUpdate) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ConnectionPostgresUpdate) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionPostgresUpdate) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionPostgresUpdate) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionPostgresUpdate) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionPostgresUpdate) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionPostgresUpdate) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionPostgresUpdate) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *ConnectionPostgresUpdate) GetAutoReplicate() *string {
	if o == nil {
		return nil
	}
	return o.AutoReplicate
}

func (o *ConnectionPostgresUpdate) GetAddress() *string {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *ConnectionPostgresUpdate) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *ConnectionPostgresUpdate) GetDatabase() *string {
	if o == nil {
		return nil
	}
	return o.Database
}

func (o *ConnectionPostgresUpdate) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *ConnectionPostgresUpdate) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *ConnectionPostgresUpdate) GetSSHConfig() *ConnectionPostgresUpdateSSHConfigurationUpdate {
	if o == nil {
		return nil
	}
	return o.SSHConfig
}
