// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionSnowflakeUpdateType string

const (
	ConnectionSnowflakeUpdateTypeSnowflake ConnectionSnowflakeUpdateType = "SNOWFLAKE"
)

func (e ConnectionSnowflakeUpdateType) ToPointer() *ConnectionSnowflakeUpdateType {
	return &e
}

func (e *ConnectionSnowflakeUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SNOWFLAKE":
		*e = ConnectionSnowflakeUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionSnowflakeUpdateType: %v", v)
	}
}

type ConnectionSnowflakeUpdate struct {
	// Whether this connection should be marked as active.
	Active *bool                         `json:"active,omitempty"`
	Type   ConnectionSnowflakeUpdateType `json:"type"`
	// The unique name of this connection.
	Name *string `json:"name,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// Are you going to use this connection only as a source for pipelines? When `true`, this connection will only be available as an ETL source only, and Etleap will skip the creation of an audit table in the database.
	SourceOnly *bool `json:"sourceOnly,omitempty"`
	// When Etleap creates Snowflake tables, SELECT privileges will be granted to roles specified here. Take into account that the roles are case sensitive.
	Roles []string `json:"roles,omitempty"`
	// Take into account that the schema is case sensitive
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	Schema   *string `json:"schema,omitempty"`
	Username *string `json:"username,omitempty"`
	Database *string `json:"database,omitempty"`
	// Snowflake Authentication Types
	Authentication *SnowflakeAuthenticationTypesInput `json:"authentication,omitempty"`
	// The role the user will use to connect
	Role    *string `json:"role,omitempty"`
	Address *string `json:"address,omitempty"`
	// The virtual warehouse to use once connected.
	Warehouse *string `json:"warehouse,omitempty"`
	Password  *string `json:"password,omitempty"`
}

func (o *ConnectionSnowflakeUpdate) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ConnectionSnowflakeUpdate) GetType() ConnectionSnowflakeUpdateType {
	if o == nil {
		return ConnectionSnowflakeUpdateType("")
	}
	return o.Type
}

func (o *ConnectionSnowflakeUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionSnowflakeUpdate) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionSnowflakeUpdate) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionSnowflakeUpdate) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionSnowflakeUpdate) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionSnowflakeUpdate) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionSnowflakeUpdate) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionSnowflakeUpdate) GetSourceOnly() *bool {
	if o == nil {
		return nil
	}
	return o.SourceOnly
}

func (o *ConnectionSnowflakeUpdate) GetRoles() []string {
	if o == nil {
		return nil
	}
	return o.Roles
}

func (o *ConnectionSnowflakeUpdate) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *ConnectionSnowflakeUpdate) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *ConnectionSnowflakeUpdate) GetDatabase() *string {
	if o == nil {
		return nil
	}
	return o.Database
}

func (o *ConnectionSnowflakeUpdate) GetAuthentication() *SnowflakeAuthenticationTypesInput {
	if o == nil {
		return nil
	}
	return o.Authentication
}

func (o *ConnectionSnowflakeUpdate) GetAuthenticationPassword() *SnowflakeAuthenticationPassword {
	if v := o.GetAuthentication(); v != nil {
		return v.SnowflakeAuthenticationPassword
	}
	return nil
}

func (o *ConnectionSnowflakeUpdate) GetAuthenticationKeyPair() *SnowflakeAuthenticationKeyPairInput {
	if v := o.GetAuthentication(); v != nil {
		return v.SnowflakeAuthenticationKeyPairInput
	}
	return nil
}

func (o *ConnectionSnowflakeUpdate) GetRole() *string {
	if o == nil {
		return nil
	}
	return o.Role
}

func (o *ConnectionSnowflakeUpdate) GetAddress() *string {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *ConnectionSnowflakeUpdate) GetWarehouse() *string {
	if o == nil {
		return nil
	}
	return o.Warehouse
}

func (o *ConnectionSnowflakeUpdate) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}
