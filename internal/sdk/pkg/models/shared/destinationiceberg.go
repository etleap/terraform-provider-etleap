// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
)

type DestinationIcebergType string

const (
	DestinationIcebergTypeIceberg DestinationIcebergType = "ICEBERG"
)

func (e DestinationIcebergType) ToPointer() *DestinationIcebergType {
	return &e
}

func (e *DestinationIcebergType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ICEBERG":
		*e = DestinationIcebergType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationIcebergType: %v", v)
	}
}

type DestinationIceberg struct {
	Type DestinationIcebergType `json:"type"`
	// The universally unique identifier of the destination connection.
	ConnectionID string `json:"connectionId"`
	// If set to `true`, a `Transformation Complete` event is published once a transformation completes, and the pipeline waits for a `Quality Check Complete` event before loading to the destination. Defaults to `false`.
	WaitForQualityCheck *bool `default:"false" json:"waitForQualityCheck"`
	// The destination column names that constitute the primary key. <br> If the pipline has a sharded source include a column that specifies the shard identifier.
	PrimaryKey []string `json:"primaryKey,omitempty"`
	// Whether schema changes detected during transformation should be handled automatically or not. Defaults to `true`.
	AutomaticSchemaChanges *bool `default:"true" json:"automaticSchemaChanges"`
	// The schema in the destination that the tables will be created in. If this is not specified or set to `null` then the schema specified on the connection is used.
	Schema *string `json:"schema,omitempty"`
	Table  string  `json:"table"`
}

func (d DestinationIceberg) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DestinationIceberg) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *DestinationIceberg) GetType() DestinationIcebergType {
	if o == nil {
		return DestinationIcebergType("")
	}
	return o.Type
}

func (o *DestinationIceberg) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *DestinationIceberg) GetWaitForQualityCheck() *bool {
	if o == nil {
		return nil
	}
	return o.WaitForQualityCheck
}

func (o *DestinationIceberg) GetPrimaryKey() []string {
	if o == nil {
		return nil
	}
	return o.PrimaryKey
}

func (o *DestinationIceberg) GetAutomaticSchemaChanges() *bool {
	if o == nil {
		return nil
	}
	return o.AutomaticSchemaChanges
}

func (o *DestinationIceberg) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *DestinationIceberg) GetTable() string {
	if o == nil {
		return ""
	}
	return o.Table
}
