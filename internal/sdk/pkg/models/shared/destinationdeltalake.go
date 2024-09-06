// Code generated by Speakeasy (https://speakeasyapi.dev). Patched by Etleap.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
)

// DestinationDeltaLakeType - <!-- theme: warning -->
// > Delta Lake connections are currently in Beta which means that they are subject to non-backwards-compatible and breaking changes.
type DestinationDeltaLakeType string

const (
	DestinationDeltaLakeTypeDeltaLake DestinationDeltaLakeType = "DELTA_LAKE"
)

func (e DestinationDeltaLakeType) ToPointer() *DestinationDeltaLakeType {
	return &e
}

func (e *DestinationDeltaLakeType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DELTA_LAKE":
		*e = DestinationDeltaLakeType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationDeltaLakeType: %v", v)
	}
}

type DestinationDeltaLake struct {
	// <!-- theme: warning -->
	// > Delta Lake connections are currently in Beta which means that they are subject to non-backwards-compatible and breaking changes.
	Type DestinationDeltaLakeType `json:"type"`
	// The universally unique identifier of the destination connection.
	ConnectionID string `json:"connectionId"`
	// If set to `true`, a `Transformation Complete` event is published once a transformation completes, and the pipeline waits for a `Quality Check Complete` event before loading to the destination. Defaults to `false`.
	WaitForQualityCheck *bool `default:"false" json:"waitForQualityCheck"`
	// The destination column names that constitute the primary key. <br> If the pipline has a sharded source include a column that specifies the shard identifier.
	PrimaryKey []string `json:"primaryKey,omitempty"`
	// Whether schema changes detected during transformation should be handled automatically or not. Defaults to `true`.
	AutomaticSchemaChanges *bool `default:"true" json:"automaticSchemaChanges"`
	// The schema in the destination that the tables will be created in.
	Schema string `json:"schema"`
	Table  string `json:"table"`
	// Name of a column that indicates the time the record was updated at the destination.
	LastUpdatedColumn *string `json:"lastUpdatedColumn,omitempty"`
	// If the destination table should retain the history of the source. More information here: https://docs.etleap.com/docs/documentation/56a1503dc499e-update-with-history-retention-mode. Defaults to `false`.
	RetainHistory *bool `default:"false" json:"retainHistory"`
	// This setting disables column mapping on the tables created by this pipeline.
	//
	// When enabled, this pipeline will create Delta Lake tables that can be read by Databricks clusters with runtime versions before 10.2.
	//
	// However, without column mapping, native schema changes are not supported and will cause the table's underlying Parquet files to be rewritten, which can be slow. Schema changes will also not preserve column constraints such as `NOT NULL` on the destination tables.
	Pre10Dot2RuntimeSupport *bool `default:"false" json:"pre10Dot2RuntimeSupport"`
}

func (d DestinationDeltaLake) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationDeltaLake) UnmarshalJSON(data []byte) error {
	// Etleap monkey-patch: the unmarshal function allows unknown fields for ignoring schemaChangingTo and tableChangingTo
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DestinationDeltaLake) GetType() DestinationDeltaLakeType {
	if o == nil {
		return DestinationDeltaLakeType("")
	}
	return o.Type
}

func (o *DestinationDeltaLake) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *DestinationDeltaLake) GetWaitForQualityCheck() *bool {
	if o == nil {
		return nil
	}
	return o.WaitForQualityCheck
}

func (o *DestinationDeltaLake) GetPrimaryKey() []string {
	if o == nil {
		return nil
	}
	return o.PrimaryKey
}

func (o *DestinationDeltaLake) GetAutomaticSchemaChanges() *bool {
	if o == nil {
		return nil
	}
	return o.AutomaticSchemaChanges
}

func (o *DestinationDeltaLake) GetSchema() string {
	if o == nil {
		return ""
	}
	return o.Schema
}

func (o *DestinationDeltaLake) GetTable() string {
	if o == nil {
		return ""
	}
	return o.Table
}

func (o *DestinationDeltaLake) GetLastUpdatedColumn() *string {
	if o == nil {
		return nil
	}
	return o.LastUpdatedColumn
}

func (o *DestinationDeltaLake) GetRetainHistory() *bool {
	if o == nil {
		return nil
	}
	return o.RetainHistory
}

func (o *DestinationDeltaLake) GetPre10Dot2RuntimeSupport() *bool {
	if o == nil {
		return nil
	}
	return o.Pre10Dot2RuntimeSupport
}
