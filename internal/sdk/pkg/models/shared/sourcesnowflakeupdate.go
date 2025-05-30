// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceSnowflakeUpdateType string

const (
	SourceSnowflakeUpdateTypeSnowflake SourceSnowflakeUpdateType = "SNOWFLAKE"
)

func (e SourceSnowflakeUpdateType) ToPointer() *SourceSnowflakeUpdateType {
	return &e
}

func (e *SourceSnowflakeUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SNOWFLAKE":
		*e = SourceSnowflakeUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSnowflakeUpdateType: %v", v)
	}
}

type SourceSnowflakeUpdate struct {
	// Notify if we can't extract for `x` hours. Setting it to `null` disables the notification. Defaults to `null`.
	LatencyThreshold *int64                     `json:"latencyThreshold,omitempty"`
	Type             *SourceSnowflakeUpdateType `json:"type,omitempty"`
}

func (o *SourceSnowflakeUpdate) GetLatencyThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LatencyThreshold
}

func (o *SourceSnowflakeUpdate) GetType() *SourceSnowflakeUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}
