// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceUserDefinedAPIType string

const (
	SourceUserDefinedAPITypeUserDefinedAPI SourceUserDefinedAPIType = "USER_DEFINED_API"
)

func (e SourceUserDefinedAPIType) ToPointer() *SourceUserDefinedAPIType {
	return &e
}

func (e *SourceUserDefinedAPIType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "USER_DEFINED_API":
		*e = SourceUserDefinedAPIType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceUserDefinedAPIType: %v", v)
	}
}

type SourceUserDefinedAPI struct {
	// The universally unique identifier for the source.
	ConnectionID string                   `json:"connectionId"`
	Type         SourceUserDefinedAPIType `json:"type"`
	// Notify if we can't extract for `x` hours. Setting it to `null` disables the notification. Defaults to `null`.
	LatencyThreshold *int64 `json:"latencyThreshold,omitempty"`
	// The User-Defined API entity.
	Entity string `json:"entity"`
}

func (o *SourceUserDefinedAPI) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *SourceUserDefinedAPI) GetType() SourceUserDefinedAPIType {
	if o == nil {
		return SourceUserDefinedAPIType("")
	}
	return o.Type
}

func (o *SourceUserDefinedAPI) GetLatencyThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LatencyThreshold
}

func (o *SourceUserDefinedAPI) GetEntity() string {
	if o == nil {
		return ""
	}
	return o.Entity
}
