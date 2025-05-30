// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceErpxUpdateType string

const (
	SourceErpxUpdateTypeErpx SourceErpxUpdateType = "ERPX"
)

func (e SourceErpxUpdateType) ToPointer() *SourceErpxUpdateType {
	return &e
}

func (e *SourceErpxUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ERPX":
		*e = SourceErpxUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceErpxUpdateType: %v", v)
	}
}

type SourceErpxUpdate struct {
	// Notify if we can't extract for `x` hours. Setting it to `null` disables the notification. Defaults to `null`.
	LatencyThreshold *int64                `json:"latencyThreshold,omitempty"`
	Type             *SourceErpxUpdateType `json:"type,omitempty"`
}

func (o *SourceErpxUpdate) GetLatencyThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LatencyThreshold
}

func (o *SourceErpxUpdate) GetType() *SourceErpxUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}
