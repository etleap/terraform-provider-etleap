// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceLinkedInAdsUpdateType string

const (
	SourceLinkedInAdsUpdateTypeLinkedInAds SourceLinkedInAdsUpdateType = "LINKED_IN_ADS"
)

func (e SourceLinkedInAdsUpdateType) ToPointer() *SourceLinkedInAdsUpdateType {
	return &e
}

func (e *SourceLinkedInAdsUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "LINKED_IN_ADS":
		*e = SourceLinkedInAdsUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceLinkedInAdsUpdateType: %v", v)
	}
}

type SourceLinkedInAdsUpdate struct {
	// Notify if we can't extract for `x` hours. Setting it to `null` disables the notification. Defaults to `null`.
	LatencyThreshold *int64                       `json:"latencyThreshold,omitempty"`
	Type             *SourceLinkedInAdsUpdateType `json:"type,omitempty"`
}

func (o *SourceLinkedInAdsUpdate) GetLatencyThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LatencyThreshold
}

func (o *SourceLinkedInAdsUpdate) GetType() *SourceLinkedInAdsUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}
