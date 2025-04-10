// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourcePostgresUpdateType string

const (
	SourcePostgresUpdateTypePostgres SourcePostgresUpdateType = "POSTGRES"
)

func (e SourcePostgresUpdateType) ToPointer() *SourcePostgresUpdateType {
	return &e
}

func (e *SourcePostgresUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "POSTGRES":
		*e = SourcePostgresUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePostgresUpdateType: %v", v)
	}
}

type SourcePostgresUpdate struct {
	Type *SourcePostgresUpdateType `json:"type,omitempty"`
	// Notify if we can't extract for `x` hours. Setting it to `null` disables the notification. Defaults to `null`.
	LatencyThreshold *int64 `json:"latencyThreshold,omitempty"`
}

func (o *SourcePostgresUpdate) GetType() *SourcePostgresUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *SourcePostgresUpdate) GetLatencyThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LatencyThreshold
}
