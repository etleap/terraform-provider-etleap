// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionMysqlShardedUpdateType string

const (
	ConnectionMysqlShardedUpdateTypeMysqlSharded ConnectionMysqlShardedUpdateType = "MYSQL_SHARDED"
)

func (e ConnectionMysqlShardedUpdateType) ToPointer() *ConnectionMysqlShardedUpdateType {
	return &e
}

func (e *ConnectionMysqlShardedUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MYSQL_SHARDED":
		*e = ConnectionMysqlShardedUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionMysqlShardedUpdateType: %v", v)
	}
}

// ConnectionMysqlShardedUpdate - The properties available shared among both sharded and unsharded versions. Cannot be used directly; use the inheriting classes.
type ConnectionMysqlShardedUpdate struct {
	// The unique name of this connection.
	Name *string                          `json:"name,omitempty"`
	Type ConnectionMysqlShardedUpdateType `json:"type"`
	// Whether this connection should be marked as active.
	Active *bool `json:"active,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule                   *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	RequireSslAndValidateCertificate *bool                `json:"requireSslAndValidateCertificate,omitempty"`
	// If you want Etleap to create pipelines for each source table automatically, specify the id of an Etleap destination connection here. If you want to create pipelines manually, omit this property.<br/><br/>If a database is not specified on this connection, then all databases will be replicated to the selected destination. Any databases not present in the destination will be created as needed.<br/><br/>If a database is specified on this connection, then only tables in that database will be replicated to the selected destination. Tables will be created in the database specified on the destination connection.
	AutoReplicate *string `json:"autoReplicate,omitempty"`
	// Should Etleap interpret columns with type Tinyint(1) as Boolean (i.e. true/false)?
	TinyInt1IsBoolean *bool `json:"tinyInt1IsBoolean,omitempty"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	Database *string `json:"database,omitempty"`
	// All MySQL shards for a connection should specify the same database.
	Shards []MysqlShard `json:"shards,omitempty"`
}

func (o *ConnectionMysqlShardedUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionMysqlShardedUpdate) GetType() ConnectionMysqlShardedUpdateType {
	if o == nil {
		return ConnectionMysqlShardedUpdateType("")
	}
	return o.Type
}

func (o *ConnectionMysqlShardedUpdate) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ConnectionMysqlShardedUpdate) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionMysqlShardedUpdate) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionMysqlShardedUpdate) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionMysqlShardedUpdate) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionMysqlShardedUpdate) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionMysqlShardedUpdate) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionMysqlShardedUpdate) GetRequireSslAndValidateCertificate() *bool {
	if o == nil {
		return nil
	}
	return o.RequireSslAndValidateCertificate
}

func (o *ConnectionMysqlShardedUpdate) GetAutoReplicate() *string {
	if o == nil {
		return nil
	}
	return o.AutoReplicate
}

func (o *ConnectionMysqlShardedUpdate) GetTinyInt1IsBoolean() *bool {
	if o == nil {
		return nil
	}
	return o.TinyInt1IsBoolean
}

func (o *ConnectionMysqlShardedUpdate) GetDatabase() *string {
	if o == nil {
		return nil
	}
	return o.Database
}

func (o *ConnectionMysqlShardedUpdate) GetShards() []MysqlShard {
	if o == nil {
		return nil
	}
	return o.Shards
}
