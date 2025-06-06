// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceFreshworksUpdateType string

const (
	SourceFreshworksUpdateTypeFreshworks SourceFreshworksUpdateType = "FRESHWORKS"
)

func (e SourceFreshworksUpdateType) ToPointer() *SourceFreshworksUpdateType {
	return &e
}

func (e *SourceFreshworksUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "FRESHWORKS":
		*e = SourceFreshworksUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFreshworksUpdateType: %v", v)
	}
}

type SourceFreshworksUpdate struct {
	// Notify if we can't extract for `x` hours. Setting it to `null` disables the notification. Defaults to `null`.
	LatencyThreshold *int64                      `json:"latencyThreshold,omitempty"`
	Type             *SourceFreshworksUpdateType `json:"type,omitempty"`
}

func (o *SourceFreshworksUpdate) GetLatencyThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LatencyThreshold
}

func (o *SourceFreshworksUpdate) GetType() *SourceFreshworksUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}
