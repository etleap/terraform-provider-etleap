// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceSapHanaShardedUpdateType string

const (
	SourceSapHanaShardedUpdateTypeSapHanaSharded SourceSapHanaShardedUpdateType = "SAP_HANA_SHARDED"
)

func (e SourceSapHanaShardedUpdateType) ToPointer() *SourceSapHanaShardedUpdateType {
	return &e
}

func (e *SourceSapHanaShardedUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SAP_HANA_SHARDED":
		*e = SourceSapHanaShardedUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSapHanaShardedUpdateType: %v", v)
	}
}

type SourceSapHanaShardedUpdate struct {
	// Notify if we can't extract for `x` hours. Setting it to `null` disables the notification. Defaults to `null`.
	LatencyThreshold *int64                          `json:"latencyThreshold,omitempty"`
	Type             *SourceSapHanaShardedUpdateType `json:"type,omitempty"`
}

func (o *SourceSapHanaShardedUpdate) GetLatencyThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LatencyThreshold
}

func (o *SourceSapHanaShardedUpdate) GetType() *SourceSapHanaShardedUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}
