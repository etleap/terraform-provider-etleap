// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceSapHanaUpdateType string

const (
	SourceSapHanaUpdateTypeSapHana SourceSapHanaUpdateType = "SAP_HANA"
)

func (e SourceSapHanaUpdateType) ToPointer() *SourceSapHanaUpdateType {
	return &e
}

func (e *SourceSapHanaUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SAP_HANA":
		*e = SourceSapHanaUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSapHanaUpdateType: %v", v)
	}
}

type SourceSapHanaUpdate struct {
	Type *SourceSapHanaUpdateType `json:"type,omitempty"`
	// Notify if we can't extract for `x` hours. Setting it to `null` disables the notification. Defaults to `null`.
	LatencyThreshold *int64 `json:"latencyThreshold,omitempty"`
}

func (o *SourceSapHanaUpdate) GetType() *SourceSapHanaUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *SourceSapHanaUpdate) GetLatencyThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LatencyThreshold
}
