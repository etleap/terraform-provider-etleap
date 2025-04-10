// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
	"time"
)

// ConnectionRedshiftShardedStatus - The current status of the connection.
type ConnectionRedshiftShardedStatus string

const (
	ConnectionRedshiftShardedStatusUnknown     ConnectionRedshiftShardedStatus = "UNKNOWN"
	ConnectionRedshiftShardedStatusUp          ConnectionRedshiftShardedStatus = "UP"
	ConnectionRedshiftShardedStatusDown        ConnectionRedshiftShardedStatus = "DOWN"
	ConnectionRedshiftShardedStatusResize      ConnectionRedshiftShardedStatus = "RESIZE"
	ConnectionRedshiftShardedStatusMaintenance ConnectionRedshiftShardedStatus = "MAINTENANCE"
	ConnectionRedshiftShardedStatusQuota       ConnectionRedshiftShardedStatus = "QUOTA"
	ConnectionRedshiftShardedStatusCreating    ConnectionRedshiftShardedStatus = "CREATING"
)

func (e ConnectionRedshiftShardedStatus) ToPointer() *ConnectionRedshiftShardedStatus {
	return &e
}

func (e *ConnectionRedshiftShardedStatus) UnmarshalJSON(data []byte) error {
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
		*e = ConnectionRedshiftShardedStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionRedshiftShardedStatus: %v", v)
	}
}

type ConnectionRedshiftShardedDefaultUpdateSchedule struct {
	// The pipeline mode refers to how the pipeline fetches data changes from the source and how those changes are applied to the destination table. See <a target="_blank" href="https://docs.etleap.com/docs/documentation/ZG9jOjIyMjE3ODA2-introduction">the documentation</a> for more details.
	PipelineMode *PipelineUpdateModes `json:"pipelineMode,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
}

func (o *ConnectionRedshiftShardedDefaultUpdateSchedule) GetPipelineMode() *PipelineUpdateModes {
	if o == nil {
		return nil
	}
	return o.PipelineMode
}

func (o *ConnectionRedshiftShardedDefaultUpdateSchedule) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionRedshiftShardedDefaultUpdateSchedule) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionRedshiftShardedDefaultUpdateSchedule) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionRedshiftShardedDefaultUpdateSchedule) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionRedshiftShardedDefaultUpdateSchedule) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionRedshiftShardedDefaultUpdateSchedule) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

type ConnectionRedshiftShardedType string

const (
	ConnectionRedshiftShardedTypeRedshiftSharded ConnectionRedshiftShardedType = "REDSHIFT_SHARDED"
)

func (e ConnectionRedshiftShardedType) ToPointer() *ConnectionRedshiftShardedType {
	return &e
}

func (e *ConnectionRedshiftShardedType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "REDSHIFT_SHARDED":
		*e = ConnectionRedshiftShardedType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionRedshiftShardedType: %v", v)
	}
}

type ConnectionRedshiftSharded struct {
	// The current status of the connection.
	Status ConnectionRedshiftShardedStatus `json:"status"`
	// The unique name of this connection.
	Name string `json:"name"`
	// The date and time when then the connection was created.
	CreateDate time.Time `json:"createDate"`
	// When an update schedule is not defined for a connection, the default schedule is used. The default defined individually per `pipelineMode` and may be subject to change.
	DefaultUpdateSchedule []ConnectionRedshiftShardedDefaultUpdateSchedule `json:"defaultUpdateSchedule"`
	// Whether this connection has been marked as active.
	Active bool                          `json:"active"`
	Type   ConnectionRedshiftShardedType `json:"type"`
	// The unique identifier of the connection.
	ID string `json:"id"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// The id of another Etleap Redshift connection. If specified, Etleap will make the data loaded available to the other cluster via Redshift Data Sharing.
	DataSharingDestinations []string `json:"dataSharingDestinations,omitempty"`
	// Should Etleap prefix each load query with metadata? More info can be found <a href="https://docs.etleap.com/docs/documentation/ba7744fcf6114-redshift-optional-connection-settings#include-query-tags">here</a>.
	QueryTagsEnabled *bool `default:"false" json:"queryTagsEnabled"`
	// When Etleap creates Redshift tables, SELECT privileges will be granted to user groups specified here.
	UserGroups []string `json:"userGroups,omitempty"`
	// Are you going to use this connection only as a source for pipelines? When `true`, this connection will only be available as an ETL source only, and Etleap will skip the creation of an audit table in the database.
	SourceOnly *bool `default:"false" json:"sourceOnly"`
	// Etleap will create VARCHAR columns with the minimal required width based on the data it's loading, and expand the column width as required. This can improve performance but there are <a target="_blank" href="https://docs.etleap.com/docs/documentation/ba7744fcf6114-redshift-optional-connection-settings#enable-dynamic-varchar-widths">some limitations</a>. Note: if set to `true`, it can't later be updated to `false`.
	DynamicVarcharWidthEnabled *bool `default:"false" json:"dynamicVarcharWidthEnabled"`
	// If not specified, the default schema will be used.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	Schema *string               `json:"schema,omitempty"`
	Shards []DatabaseShardOutput `json:"shards"`
}

func (c ConnectionRedshiftSharded) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionRedshiftSharded) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionRedshiftSharded) GetStatus() ConnectionRedshiftShardedStatus {
	if o == nil {
		return ConnectionRedshiftShardedStatus("")
	}
	return o.Status
}

func (o *ConnectionRedshiftSharded) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionRedshiftSharded) GetCreateDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreateDate
}

func (o *ConnectionRedshiftSharded) GetDefaultUpdateSchedule() []ConnectionRedshiftShardedDefaultUpdateSchedule {
	if o == nil {
		return []ConnectionRedshiftShardedDefaultUpdateSchedule{}
	}
	return o.DefaultUpdateSchedule
}

func (o *ConnectionRedshiftSharded) GetActive() bool {
	if o == nil {
		return false
	}
	return o.Active
}

func (o *ConnectionRedshiftSharded) GetType() ConnectionRedshiftShardedType {
	if o == nil {
		return ConnectionRedshiftShardedType("")
	}
	return o.Type
}

func (o *ConnectionRedshiftSharded) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ConnectionRedshiftSharded) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionRedshiftSharded) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionRedshiftSharded) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionRedshiftSharded) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionRedshiftSharded) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionRedshiftSharded) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionRedshiftSharded) GetDataSharingDestinations() []string {
	if o == nil {
		return nil
	}
	return o.DataSharingDestinations
}

func (o *ConnectionRedshiftSharded) GetQueryTagsEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.QueryTagsEnabled
}

func (o *ConnectionRedshiftSharded) GetUserGroups() []string {
	if o == nil {
		return nil
	}
	return o.UserGroups
}

func (o *ConnectionRedshiftSharded) GetSourceOnly() *bool {
	if o == nil {
		return nil
	}
	return o.SourceOnly
}

func (o *ConnectionRedshiftSharded) GetDynamicVarcharWidthEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.DynamicVarcharWidthEnabled
}

func (o *ConnectionRedshiftSharded) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *ConnectionRedshiftSharded) GetShards() []DatabaseShardOutput {
	if o == nil {
		return []DatabaseShardOutput{}
	}
	return o.Shards
}

type ConnectionRedshiftShardedInput struct {
	// The unique name of this connection.
	Name string                        `json:"name"`
	Type ConnectionRedshiftShardedType `json:"type"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// The id of another Etleap Redshift connection. If specified, Etleap will make the data loaded available to the other cluster via Redshift Data Sharing.
	DataSharingDestinations []string `json:"dataSharingDestinations,omitempty"`
	// Should Etleap prefix each load query with metadata? More info can be found <a href="https://docs.etleap.com/docs/documentation/ba7744fcf6114-redshift-optional-connection-settings#include-query-tags">here</a>.
	QueryTagsEnabled *bool `default:"false" json:"queryTagsEnabled"`
	// When Etleap creates Redshift tables, SELECT privileges will be granted to user groups specified here.
	UserGroups []string `json:"userGroups,omitempty"`
	// Are you going to use this connection only as a source for pipelines? When `true`, this connection will only be available as an ETL source only, and Etleap will skip the creation of an audit table in the database.
	SourceOnly *bool `default:"false" json:"sourceOnly"`
	// Etleap will create VARCHAR columns with the minimal required width based on the data it's loading, and expand the column width as required. This can improve performance but there are <a target="_blank" href="https://docs.etleap.com/docs/documentation/ba7744fcf6114-redshift-optional-connection-settings#enable-dynamic-varchar-widths">some limitations</a>. Note: if set to `true`, it can't later be updated to `false`.
	DynamicVarcharWidthEnabled *bool `default:"false" json:"dynamicVarcharWidthEnabled"`
	// If not specified, the default schema will be used.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	Schema *string         `json:"schema,omitempty"`
	Shards []DatabaseShard `json:"shards"`
}

func (c ConnectionRedshiftShardedInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionRedshiftShardedInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionRedshiftShardedInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionRedshiftShardedInput) GetType() ConnectionRedshiftShardedType {
	if o == nil {
		return ConnectionRedshiftShardedType("")
	}
	return o.Type
}

func (o *ConnectionRedshiftShardedInput) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionRedshiftShardedInput) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionRedshiftShardedInput) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionRedshiftShardedInput) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionRedshiftShardedInput) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionRedshiftShardedInput) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionRedshiftShardedInput) GetDataSharingDestinations() []string {
	if o == nil {
		return nil
	}
	return o.DataSharingDestinations
}

func (o *ConnectionRedshiftShardedInput) GetQueryTagsEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.QueryTagsEnabled
}

func (o *ConnectionRedshiftShardedInput) GetUserGroups() []string {
	if o == nil {
		return nil
	}
	return o.UserGroups
}

func (o *ConnectionRedshiftShardedInput) GetSourceOnly() *bool {
	if o == nil {
		return nil
	}
	return o.SourceOnly
}

func (o *ConnectionRedshiftShardedInput) GetDynamicVarcharWidthEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.DynamicVarcharWidthEnabled
}

func (o *ConnectionRedshiftShardedInput) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *ConnectionRedshiftShardedInput) GetShards() []DatabaseShard {
	if o == nil {
		return []DatabaseShard{}
	}
	return o.Shards
}
