// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceWorkfrontUpdateType string

const (
	SourceWorkfrontUpdateTypeWorkfront SourceWorkfrontUpdateType = "WORKFRONT"
)

func (e SourceWorkfrontUpdateType) ToPointer() *SourceWorkfrontUpdateType {
	return &e
}

func (e *SourceWorkfrontUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "WORKFRONT":
		*e = SourceWorkfrontUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceWorkfrontUpdateType: %v", v)
	}
}

type SourceWorkfrontUpdate struct {
	Type *SourceWorkfrontUpdateType `json:"type,omitempty"`
	// Notify if we can't extract for `x` hours. Setting it to `null` disables the notification. Defaults to `null`.
	LatencyThreshold *int64 `json:"latencyThreshold,omitempty"`
}

func (o *SourceWorkfrontUpdate) GetType() *SourceWorkfrontUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *SourceWorkfrontUpdate) GetLatencyThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LatencyThreshold
}
