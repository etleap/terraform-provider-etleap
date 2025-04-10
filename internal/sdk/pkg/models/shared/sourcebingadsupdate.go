// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceBingAdsUpdateType string

const (
	SourceBingAdsUpdateTypeBingAds SourceBingAdsUpdateType = "BING_ADS"
)

func (e SourceBingAdsUpdateType) ToPointer() *SourceBingAdsUpdateType {
	return &e
}

func (e *SourceBingAdsUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "BING_ADS":
		*e = SourceBingAdsUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceBingAdsUpdateType: %v", v)
	}
}

type SourceBingAdsUpdate struct {
	Type *SourceBingAdsUpdateType `json:"type,omitempty"`
	// Notify if we can't extract for `x` hours. Setting it to `null` disables the notification. Defaults to `null`.
	LatencyThreshold *int64 `json:"latencyThreshold,omitempty"`
}

func (o *SourceBingAdsUpdate) GetType() *SourceBingAdsUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *SourceBingAdsUpdate) GetLatencyThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LatencyThreshold
}
