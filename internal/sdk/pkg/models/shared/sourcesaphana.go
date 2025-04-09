// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceSapHanaType string

const (
	SourceSapHanaTypeSapHana SourceSapHanaType = "SAP_HANA"
)

func (e SourceSapHanaType) ToPointer() *SourceSapHanaType {
	return &e
}

func (e *SourceSapHanaType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SAP_HANA":
		*e = SourceSapHanaType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSapHanaType: %v", v)
	}
}

type SourceSapHana struct {
	// The universally unique identifier for the source.
	ConnectionID string            `json:"connectionId"`
	Type         SourceSapHanaType `json:"type"`
	// Notify if we can't extract for `x` hours. Setting it to `null` disables the notification. Defaults to `null`.
	LatencyThreshold *int64 `json:"latencyThreshold,omitempty"`
	// Name of the table to be extracted from the source. Either `table` or `tableNameFilter` must be specified, but not both.
	Table *string `json:"table,omitempty"`
	// Regular expression matching all partitions of a table. Partitions must have the same table schema. Either `tableNameFilter` or `table` must be specified, but not both.
	TableNameFilter *string `json:"tableNameFilter,omitempty"`
	// Columns that make up the primary key of the source. The specified order of columns matters for composite primary keys. <br> For source tables that do not have primary keys please specify an empty array. <br> For sharded sources include `shard_id` as first primary key column.<br><br>The **default value** is an empty array.
	PrimaryKeyColumns []string `json:"primaryKeyColumns"`
	// Name of a column that indicates the time the record was updated at the source.
	LastUpdatedColumn *string `json:"lastUpdatedColumn,omitempty"`
	// Name of the schema in the source from which the data is to be extracted. If not specified, the source connection schema or the default schema for connection type will be used.
	Schema *string `json:"schema,omitempty"`
}

func (o *SourceSapHana) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *SourceSapHana) GetType() SourceSapHanaType {
	if o == nil {
		return SourceSapHanaType("")
	}
	return o.Type
}

func (o *SourceSapHana) GetLatencyThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LatencyThreshold
}

func (o *SourceSapHana) GetTable() *string {
	if o == nil {
		return nil
	}
	return o.Table
}

func (o *SourceSapHana) GetTableNameFilter() *string {
	if o == nil {
		return nil
	}
	return o.TableNameFilter
}

func (o *SourceSapHana) GetPrimaryKeyColumns() []string {
	if o == nil {
		return []string{}
	}
	return o.PrimaryKeyColumns
}

func (o *SourceSapHana) GetLastUpdatedColumn() *string {
	if o == nil {
		return nil
	}
	return o.LastUpdatedColumn
}

func (o *SourceSapHana) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}
