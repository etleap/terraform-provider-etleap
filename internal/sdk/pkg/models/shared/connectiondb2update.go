// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionDb2UpdateType string

const (
	ConnectionDb2UpdateTypeDb2 ConnectionDb2UpdateType = "DB2"
)

func (e ConnectionDb2UpdateType) ToPointer() *ConnectionDb2UpdateType {
	return &e
}

func (e *ConnectionDb2UpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DB2":
		*e = ConnectionDb2UpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionDb2UpdateType: %v", v)
	}
}

type ConnectionDb2UpdateSSHConfigurationUpdate struct {
	// The username for the SSH connection.
	Username *string `json:"username,omitempty"`
	// The server address for the SSH connection.
	Address *string `json:"address,omitempty"`
}

func (o *ConnectionDb2UpdateSSHConfigurationUpdate) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *ConnectionDb2UpdateSSHConfigurationUpdate) GetAddress() *string {
	if o == nil {
		return nil
	}
	return o.Address
}

// ConnectionDb2Update - Specifies the location of a database.
type ConnectionDb2Update struct {
	// Whether this connection should be marked as active.
	Active *bool                   `json:"active,omitempty"`
	Type   ConnectionDb2UpdateType `json:"type"`
	// The unique name of this connection.
	Name *string `json:"name,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// Etleap secures all connections with TLS encryption. You can provide your own TLS certificate, or if none is provided, the AWS RDS global certificate bundle will be used by default.
	Certificate *string `json:"certificate,omitempty"`
	// If not specified, the default schema will be used.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	Schema    *string                                    `json:"schema,omitempty"`
	Username  *string                                    `json:"username,omitempty"`
	SSHConfig *ConnectionDb2UpdateSSHConfigurationUpdate `json:"sshConfig,omitempty"`
	Password  *string                                    `json:"password,omitempty"`
	Port      *int64                                     `json:"port,omitempty"`
	Address   *string                                    `json:"address,omitempty"`
	Database  *string                                    `json:"database,omitempty"`
}

func (o *ConnectionDb2Update) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ConnectionDb2Update) GetType() ConnectionDb2UpdateType {
	if o == nil {
		return ConnectionDb2UpdateType("")
	}
	return o.Type
}

func (o *ConnectionDb2Update) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionDb2Update) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionDb2Update) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionDb2Update) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionDb2Update) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionDb2Update) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionDb2Update) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionDb2Update) GetCertificate() *string {
	if o == nil {
		return nil
	}
	return o.Certificate
}

func (o *ConnectionDb2Update) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *ConnectionDb2Update) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *ConnectionDb2Update) GetSSHConfig() *ConnectionDb2UpdateSSHConfigurationUpdate {
	if o == nil {
		return nil
	}
	return o.SSHConfig
}

func (o *ConnectionDb2Update) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *ConnectionDb2Update) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *ConnectionDb2Update) GetAddress() *string {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *ConnectionDb2Update) GetDatabase() *string {
	if o == nil {
		return nil
	}
	return o.Database
}
