// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceDb2UpdateType string

const (
	SourceDb2UpdateTypeDb2 SourceDb2UpdateType = "DB2"
)

func (e SourceDb2UpdateType) ToPointer() *SourceDb2UpdateType {
	return &e
}

func (e *SourceDb2UpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DB2":
		*e = SourceDb2UpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceDb2UpdateType: %v", v)
	}
}

type SourceDb2Update struct {
	// Notify if we can't extract for `x` hours. Setting it to `null` disables the notification. Defaults to `null`.
	LatencyThreshold *int64               `json:"latencyThreshold,omitempty"`
	Type             *SourceDb2UpdateType `json:"type,omitempty"`
}

func (o *SourceDb2Update) GetLatencyThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LatencyThreshold
}

func (o *SourceDb2Update) GetType() *SourceDb2UpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}
