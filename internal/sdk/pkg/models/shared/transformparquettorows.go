// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type TransformParquetToRowsType string

const (
	TransformParquetToRowsTypeParquetToRows TransformParquetToRowsType = "PARQUET_TO_ROWS"
)

func (e TransformParquetToRowsType) ToPointer() *TransformParquetToRowsType {
	return &e
}

func (e *TransformParquetToRowsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "PARQUET_TO_ROWS":
		*e = TransformParquetToRowsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TransformParquetToRowsType: %v", v)
	}
}

// TransformParquetToRows - Parses Parquet files to rows of JSON objects.
type TransformParquetToRows struct {
	Type TransformParquetToRowsType `json:"type"`
}

func (o *TransformParquetToRows) GetType() TransformParquetToRowsType {
	if o == nil {
		return TransformParquetToRowsType("")
	}
	return o.Type
}
