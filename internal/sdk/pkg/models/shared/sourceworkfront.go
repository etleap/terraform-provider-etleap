// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceWorkfrontType string

const (
	SourceWorkfrontTypeWorkfront SourceWorkfrontType = "WORKFRONT"
)

func (e SourceWorkfrontType) ToPointer() *SourceWorkfrontType {
	return &e
}

func (e *SourceWorkfrontType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "WORKFRONT":
		*e = SourceWorkfrontType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceWorkfrontType: %v", v)
	}
}

type SourceWorkfront struct {
	// The universally unique identifier for the source.
	ConnectionID string              `json:"connectionId"`
	Type         SourceWorkfrontType `json:"type"`
	// Notify if we can't extract for `x` hours. Setting it to `null` disables the notification. Defaults to `null`.
	LatencyThreshold *int64 `json:"latencyThreshold,omitempty"`
	// The Workfront entity. Spelled capitalized without spaces. For the full list, start creating a pipeline in the Etleap UI.
	Entity string `json:"entity"`
}

func (o *SourceWorkfront) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *SourceWorkfront) GetType() SourceWorkfrontType {
	if o == nil {
		return SourceWorkfrontType("")
	}
	return o.Type
}

func (o *SourceWorkfront) GetLatencyThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LatencyThreshold
}

func (o *SourceWorkfront) GetEntity() string {
	if o == nil {
		return ""
	}
	return o.Entity
}
