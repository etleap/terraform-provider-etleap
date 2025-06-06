// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceBigQueryUpdateType string

const (
	SourceBigQueryUpdateTypeBigquery SourceBigQueryUpdateType = "BIGQUERY"
)

func (e SourceBigQueryUpdateType) ToPointer() *SourceBigQueryUpdateType {
	return &e
}

func (e *SourceBigQueryUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "BIGQUERY":
		*e = SourceBigQueryUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceBigQueryUpdateType: %v", v)
	}
}

type SourceBigQueryUpdate struct {
	// Notify if we can't extract for `x` hours. Setting it to `null` disables the notification. Defaults to `null`.
	LatencyThreshold *int64                    `json:"latencyThreshold,omitempty"`
	Type             *SourceBigQueryUpdateType `json:"type,omitempty"`
}

func (o *SourceBigQueryUpdate) GetLatencyThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LatencyThreshold
}

func (o *SourceBigQueryUpdate) GetType() *SourceBigQueryUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}
