// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceZuoraType string

const (
	SourceZuoraTypeZuora SourceZuoraType = "ZUORA"
)

func (e SourceZuoraType) ToPointer() *SourceZuoraType {
	return &e
}

func (e *SourceZuoraType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ZUORA":
		*e = SourceZuoraType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceZuoraType: %v", v)
	}
}

type SourceZuora struct {
	// The universally unique identifier for the source.
	ConnectionID string          `json:"connectionId"`
	Type         SourceZuoraType `json:"type"`
	// Notify if we can't extract for `x` hours. Setting it to `null` disables the notification. Defaults to `null`.
	LatencyThreshold *int64 `json:"latencyThreshold,omitempty"`
	// The Zuora entity. Spelled capitalized with spaces. For the full list, start creating a pipeline in the Etleap UI.
	Entity string `json:"entity"`
}

func (o *SourceZuora) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *SourceZuora) GetType() SourceZuoraType {
	if o == nil {
		return SourceZuoraType("")
	}
	return o.Type
}

func (o *SourceZuora) GetLatencyThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LatencyThreshold
}

func (o *SourceZuora) GetEntity() string {
	if o == nil {
		return ""
	}
	return o.Entity
}
