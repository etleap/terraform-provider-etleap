// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceRaveMedidataType string

const (
	SourceRaveMedidataTypeRaveMedidata SourceRaveMedidataType = "RAVE_MEDIDATA"
)

func (e SourceRaveMedidataType) ToPointer() *SourceRaveMedidataType {
	return &e
}

func (e *SourceRaveMedidataType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "RAVE_MEDIDATA":
		*e = SourceRaveMedidataType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceRaveMedidataType: %v", v)
	}
}

type SourceRaveMedidata struct {
	// The universally unique identifier for the source.
	ConnectionID string                 `json:"connectionId"`
	Type         SourceRaveMedidataType `json:"type"`
	// Notify if we can't extract for `x` hours. Setting it to `null` disables the notification. Defaults to `null`.
	LatencyThreshold *int64 `json:"latencyThreshold,omitempty"`
	// The Rave Medidata entity. Example values: [dataset, study, <study-oid>@-@<form-oid>]
	Entity string `json:"entity"`
}

func (o *SourceRaveMedidata) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *SourceRaveMedidata) GetType() SourceRaveMedidataType {
	if o == nil {
		return SourceRaveMedidataType("")
	}
	return o.Type
}

func (o *SourceRaveMedidata) GetLatencyThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LatencyThreshold
}

func (o *SourceRaveMedidata) GetEntity() string {
	if o == nil {
		return ""
	}
	return o.Entity
}
