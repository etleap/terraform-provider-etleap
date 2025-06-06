// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceSapConcurUpdateType string

const (
	SourceSapConcurUpdateTypeSapConcur SourceSapConcurUpdateType = "SAP_CONCUR"
)

func (e SourceSapConcurUpdateType) ToPointer() *SourceSapConcurUpdateType {
	return &e
}

func (e *SourceSapConcurUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SAP_CONCUR":
		*e = SourceSapConcurUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSapConcurUpdateType: %v", v)
	}
}

type SourceSapConcurUpdate struct {
	// Notify if we can't extract for `x` hours. Setting it to `null` disables the notification. Defaults to `null`.
	LatencyThreshold *int64                     `json:"latencyThreshold,omitempty"`
	Type             *SourceSapConcurUpdateType `json:"type,omitempty"`
}

func (o *SourceSapConcurUpdate) GetLatencyThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LatencyThreshold
}

func (o *SourceSapConcurUpdate) GetType() *SourceSapConcurUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}
