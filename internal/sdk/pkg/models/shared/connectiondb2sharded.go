// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
	"time"
)

// ConnectionDb2ShardedStatus - The current status of the connection.
type ConnectionDb2ShardedStatus string

const (
	ConnectionDb2ShardedStatusUnknown     ConnectionDb2ShardedStatus = "UNKNOWN"
	ConnectionDb2ShardedStatusUp          ConnectionDb2ShardedStatus = "UP"
	ConnectionDb2ShardedStatusDown        ConnectionDb2ShardedStatus = "DOWN"
	ConnectionDb2ShardedStatusResize      ConnectionDb2ShardedStatus = "RESIZE"
	ConnectionDb2ShardedStatusMaintenance ConnectionDb2ShardedStatus = "MAINTENANCE"
	ConnectionDb2ShardedStatusQuota       ConnectionDb2ShardedStatus = "QUOTA"
	ConnectionDb2ShardedStatusCreating    ConnectionDb2ShardedStatus = "CREATING"
)

func (e ConnectionDb2ShardedStatus) ToPointer() *ConnectionDb2ShardedStatus {
	return &e
}

func (e *ConnectionDb2ShardedStatus) UnmarshalJSON(data []byte) error {
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
		*e = ConnectionDb2ShardedStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionDb2ShardedStatus: %v", v)
	}
}

type ConnectionDb2ShardedDefaultUpdateSchedule struct {
	// The pipeline mode refers to how the pipeline fetches data changes from the source and how those changes are applied to the destination table. See <a target="_blank" href="https://docs.etleap.com/docs/documentation/ZG9jOjIyMjE3ODA2-introduction">the documentation</a> for more details.
	PipelineMode *PipelineUpdateModes `json:"pipelineMode,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
}

func (o *ConnectionDb2ShardedDefaultUpdateSchedule) GetPipelineMode() *PipelineUpdateModes {
	if o == nil {
		return nil
	}
	return o.PipelineMode
}

func (o *ConnectionDb2ShardedDefaultUpdateSchedule) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionDb2ShardedDefaultUpdateSchedule) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionDb2ShardedDefaultUpdateSchedule) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionDb2ShardedDefaultUpdateSchedule) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionDb2ShardedDefaultUpdateSchedule) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionDb2ShardedDefaultUpdateSchedule) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

type ConnectionDb2ShardedType string

const (
	ConnectionDb2ShardedTypeDb2Sharded ConnectionDb2ShardedType = "DB2_SHARDED"
)

func (e ConnectionDb2ShardedType) ToPointer() *ConnectionDb2ShardedType {
	return &e
}

func (e *ConnectionDb2ShardedType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DB2_SHARDED":
		*e = ConnectionDb2ShardedType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionDb2ShardedType: %v", v)
	}
}

type ConnectionDb2Sharded struct {
	// The current status of the connection.
	Status ConnectionDb2ShardedStatus `json:"status"`
	// The unique name of this connection.
	Name string `json:"name"`
	// The date and time when then the connection was created.
	CreateDate time.Time `json:"createDate"`
	// When an update schedule is not defined for a connection, the default schedule is used. The default defined individually per `pipelineMode` and may be subject to change.
	DefaultUpdateSchedule []ConnectionDb2ShardedDefaultUpdateSchedule `json:"defaultUpdateSchedule"`
	// Whether this connection has been marked as active.
	Active bool                     `json:"active"`
	Type   ConnectionDb2ShardedType `json:"type"`
	// The unique identifier of the connection.
	ID string `json:"id"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// Etleap secures all connections with TLS encryption. You can provide your own TLS certificate, or if none is provided, the AWS RDS global certificate bundle will be used by default.
	Certificate *string `json:"certificate,omitempty"`
	// If not specified, the default schema will be used.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	Schema *string               `json:"schema,omitempty"`
	Shards []DatabaseShardOutput `json:"shards"`
}

func (c ConnectionDb2Sharded) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionDb2Sharded) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionDb2Sharded) GetStatus() ConnectionDb2ShardedStatus {
	if o == nil {
		return ConnectionDb2ShardedStatus("")
	}
	return o.Status
}

func (o *ConnectionDb2Sharded) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionDb2Sharded) GetCreateDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreateDate
}

func (o *ConnectionDb2Sharded) GetDefaultUpdateSchedule() []ConnectionDb2ShardedDefaultUpdateSchedule {
	if o == nil {
		return []ConnectionDb2ShardedDefaultUpdateSchedule{}
	}
	return o.DefaultUpdateSchedule
}

func (o *ConnectionDb2Sharded) GetActive() bool {
	if o == nil {
		return false
	}
	return o.Active
}

func (o *ConnectionDb2Sharded) GetType() ConnectionDb2ShardedType {
	if o == nil {
		return ConnectionDb2ShardedType("")
	}
	return o.Type
}

func (o *ConnectionDb2Sharded) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ConnectionDb2Sharded) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionDb2Sharded) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionDb2Sharded) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionDb2Sharded) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionDb2Sharded) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionDb2Sharded) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionDb2Sharded) GetCertificate() *string {
	if o == nil {
		return nil
	}
	return o.Certificate
}

func (o *ConnectionDb2Sharded) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *ConnectionDb2Sharded) GetShards() []DatabaseShardOutput {
	if o == nil {
		return []DatabaseShardOutput{}
	}
	return o.Shards
}

type ConnectionDb2ShardedInput struct {
	// The unique name of this connection.
	Name string                   `json:"name"`
	Type ConnectionDb2ShardedType `json:"type"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// Etleap secures all connections with TLS encryption. You can provide your own TLS certificate, or if none is provided, the AWS RDS global certificate bundle will be used by default.
	Certificate *string `json:"certificate,omitempty"`
	// If not specified, the default schema will be used.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	Schema *string         `json:"schema,omitempty"`
	Shards []DatabaseShard `json:"shards"`
}

func (o *ConnectionDb2ShardedInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionDb2ShardedInput) GetType() ConnectionDb2ShardedType {
	if o == nil {
		return ConnectionDb2ShardedType("")
	}
	return o.Type
}

func (o *ConnectionDb2ShardedInput) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionDb2ShardedInput) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionDb2ShardedInput) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionDb2ShardedInput) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionDb2ShardedInput) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionDb2ShardedInput) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionDb2ShardedInput) GetCertificate() *string {
	if o == nil {
		return nil
	}
	return o.Certificate
}

func (o *ConnectionDb2ShardedInput) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *ConnectionDb2ShardedInput) GetShards() []DatabaseShard {
	if o == nil {
		return []DatabaseShard{}
	}
	return o.Shards
}
