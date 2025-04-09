// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceCoupaUpdateType string

const (
	SourceCoupaUpdateTypeCoupa SourceCoupaUpdateType = "COUPA"
)

func (e SourceCoupaUpdateType) ToPointer() *SourceCoupaUpdateType {
	return &e
}

func (e *SourceCoupaUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "COUPA":
		*e = SourceCoupaUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceCoupaUpdateType: %v", v)
	}
}

type SourceCoupaUpdate struct {
	Type *SourceCoupaUpdateType `json:"type,omitempty"`
	// Notify if we can't extract for `x` hours. Setting it to `null` disables the notification. Defaults to `null`.
	LatencyThreshold *int64 `json:"latencyThreshold,omitempty"`
}

func (o *SourceCoupaUpdate) GetType() *SourceCoupaUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *SourceCoupaUpdate) GetLatencyThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LatencyThreshold
}
