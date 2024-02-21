// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionPostgresShardedUpdateType string

const (
	ConnectionPostgresShardedUpdateTypePostgresSharded ConnectionPostgresShardedUpdateType = "POSTGRES_SHARDED"
)

func (e ConnectionPostgresShardedUpdateType) ToPointer() *ConnectionPostgresShardedUpdateType {
	return &e
}

func (e *ConnectionPostgresShardedUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "POSTGRES_SHARDED":
		*e = ConnectionPostgresShardedUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionPostgresShardedUpdateType: %v", v)
	}
}

type ConnectionPostgresShardedUpdate struct {
	// The unique name of this connection.
	Name *string                             `json:"name,omitempty"`
	Type ConnectionPostgresShardedUpdateType `json:"type"`
	// Whether this connection should be marked as active.
	Active *bool `json:"active,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// If not specified, the default schema will be used.
	Schema *string `json:"schema,omitempty"`
	// If you want Etleap to create pipelines for each source table automatically, specify the id of an Etleap destination connection here. If you want to create pipelines manually, omit this property.<br/><br/>If a schema is not specified on this connection, then all schemas will be replicated to the selected destination. Any schemas not present in the destination will be created as needed.<br/><br/>If a schema is specified on this connection, then only tables in that schema will be replicated to the selected destination. Tables will be created in the schema specified on the destination connection.
	AutoReplicate *string         `json:"autoReplicate,omitempty"`
	Shards        []DatabaseShard `json:"shards,omitempty"`
}

func (o *ConnectionPostgresShardedUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionPostgresShardedUpdate) GetType() ConnectionPostgresShardedUpdateType {
	if o == nil {
		return ConnectionPostgresShardedUpdateType("")
	}
	return o.Type
}

func (o *ConnectionPostgresShardedUpdate) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ConnectionPostgresShardedUpdate) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionPostgresShardedUpdate) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionPostgresShardedUpdate) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionPostgresShardedUpdate) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionPostgresShardedUpdate) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionPostgresShardedUpdate) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionPostgresShardedUpdate) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *ConnectionPostgresShardedUpdate) GetAutoReplicate() *string {
	if o == nil {
		return nil
	}
	return o.AutoReplicate
}

func (o *ConnectionPostgresShardedUpdate) GetShards() []DatabaseShard {
	if o == nil {
		return nil
	}
	return o.Shards
}
