// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionDb2ShardedUpdateType string

const (
	ConnectionDb2ShardedUpdateTypeDb2Sharded ConnectionDb2ShardedUpdateType = "DB2_SHARDED"
)

func (e ConnectionDb2ShardedUpdateType) ToPointer() *ConnectionDb2ShardedUpdateType {
	return &e
}

func (e *ConnectionDb2ShardedUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DB2_SHARDED":
		*e = ConnectionDb2ShardedUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionDb2ShardedUpdateType: %v", v)
	}
}

type ConnectionDb2ShardedUpdate struct {
	// Whether this connection should be marked as active.
	Active *bool                          `json:"active,omitempty"`
	Type   ConnectionDb2ShardedUpdateType `json:"type"`
	// The unique name of this connection.
	Name *string `json:"name,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// Etleap secures all connections with TLS encryption. You can provide your own TLS certificate, or if none is provided, the AWS RDS global certificate bundle will be used by default.
	Certificate *string `json:"certificate,omitempty"`
	// If not specified, the default schema will be used.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	Schema *string         `json:"schema,omitempty"`
	Shards []DatabaseShard `json:"shards,omitempty"`
}

func (o *ConnectionDb2ShardedUpdate) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ConnectionDb2ShardedUpdate) GetType() ConnectionDb2ShardedUpdateType {
	if o == nil {
		return ConnectionDb2ShardedUpdateType("")
	}
	return o.Type
}

func (o *ConnectionDb2ShardedUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionDb2ShardedUpdate) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionDb2ShardedUpdate) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionDb2ShardedUpdate) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionDb2ShardedUpdate) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionDb2ShardedUpdate) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionDb2ShardedUpdate) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionDb2ShardedUpdate) GetCertificate() *string {
	if o == nil {
		return nil
	}
	return o.Certificate
}

func (o *ConnectionDb2ShardedUpdate) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *ConnectionDb2ShardedUpdate) GetShards() []DatabaseShard {
	if o == nil {
		return nil
	}
	return o.Shards
}
