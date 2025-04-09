// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceErpxType string

const (
	SourceErpxTypeErpx SourceErpxType = "ERPX"
)

func (e SourceErpxType) ToPointer() *SourceErpxType {
	return &e
}

func (e *SourceErpxType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ERPX":
		*e = SourceErpxType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceErpxType: %v", v)
	}
}

type SourceErpx struct {
	// The universally unique identifier for the source.
	ConnectionID string         `json:"connectionId"`
	Type         SourceErpxType `json:"type"`
	// Notify if we can't extract for `x` hours. Setting it to `null` disables the notification. Defaults to `null`.
	LatencyThreshold *int64 `json:"latencyThreshold,omitempty"`
	// The ERPx resource.
	Entity string `json:"entity"`
}

func (o *SourceErpx) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *SourceErpx) GetType() SourceErpxType {
	if o == nil {
		return SourceErpxType("")
	}
	return o.Type
}

func (o *SourceErpx) GetLatencyThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LatencyThreshold
}

func (o *SourceErpx) GetEntity() string {
	if o == nil {
		return ""
	}
	return o.Entity
}
