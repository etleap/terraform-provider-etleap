// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceServiceNowType string

const (
	SourceServiceNowTypeServiceNow SourceServiceNowType = "SERVICE_NOW"
)

func (e SourceServiceNowType) ToPointer() *SourceServiceNowType {
	return &e
}

func (e *SourceServiceNowType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SERVICE_NOW":
		*e = SourceServiceNowType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceServiceNowType: %v", v)
	}
}

type SourceServiceNow struct {
	// The universally unique identifier for the source.
	ConnectionID string               `json:"connectionId"`
	Type         SourceServiceNowType `json:"type"`
	// Notify if we can't extract for `x` hours. Setting it to `null` disables the notification. Defaults to `null`.
	LatencyThreshold *int64 `json:"latencyThreshold,omitempty"`
	// The ServiceNow entity. Example values: [Task, Problem, Incident]
	Entity string `json:"entity"`
}

func (o *SourceServiceNow) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *SourceServiceNow) GetType() SourceServiceNowType {
	if o == nil {
		return SourceServiceNowType("")
	}
	return o.Type
}

func (o *SourceServiceNow) GetLatencyThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LatencyThreshold
}

func (o *SourceServiceNow) GetEntity() string {
	if o == nil {
		return ""
	}
	return o.Entity
}
