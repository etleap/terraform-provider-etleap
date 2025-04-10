// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionSQLServerShardedUpdateType string

const (
	ConnectionSQLServerShardedUpdateTypeSQLServerSharded ConnectionSQLServerShardedUpdateType = "SQL_SERVER_SHARDED"
)

func (e ConnectionSQLServerShardedUpdateType) ToPointer() *ConnectionSQLServerShardedUpdateType {
	return &e
}

func (e *ConnectionSQLServerShardedUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SQL_SERVER_SHARDED":
		*e = ConnectionSQLServerShardedUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionSQLServerShardedUpdateType: %v", v)
	}
}

type ConnectionSQLServerShardedUpdate struct {
	// Whether this connection should be marked as active.
	Active *bool                                `json:"active,omitempty"`
	Type   ConnectionSQLServerShardedUpdateType `json:"type"`
	// The unique name of this connection.
	Name *string `json:"name,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// If not specified, the default schema will be used.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	Schema *string         `json:"schema,omitempty"`
	Shards []DatabaseShard `json:"shards,omitempty"`
}

func (o *ConnectionSQLServerShardedUpdate) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ConnectionSQLServerShardedUpdate) GetType() ConnectionSQLServerShardedUpdateType {
	if o == nil {
		return ConnectionSQLServerShardedUpdateType("")
	}
	return o.Type
}

func (o *ConnectionSQLServerShardedUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionSQLServerShardedUpdate) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionSQLServerShardedUpdate) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionSQLServerShardedUpdate) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionSQLServerShardedUpdate) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionSQLServerShardedUpdate) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionSQLServerShardedUpdate) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionSQLServerShardedUpdate) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *ConnectionSQLServerShardedUpdate) GetShards() []DatabaseShard {
	if o == nil {
		return nil
	}
	return o.Shards
}
