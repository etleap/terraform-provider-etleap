// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceQuoraAdsType string

const (
	SourceQuoraAdsTypeQuoraAds SourceQuoraAdsType = "QUORA_ADS"
)

func (e SourceQuoraAdsType) ToPointer() *SourceQuoraAdsType {
	return &e
}

func (e *SourceQuoraAdsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "QUORA_ADS":
		*e = SourceQuoraAdsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceQuoraAdsType: %v", v)
	}
}

type SourceQuoraAds struct {
	Type SourceQuoraAdsType `json:"type"`
	// The universally unique identifier for the source.
	ConnectionID string `json:"connectionId"`
	// Notify if we can't extract for `x` hours. Setting it to `null` disables the notification. Defaults to `null`.
	LatencyThreshold *int64 `json:"latencyThreshold,omitempty"`
	// The level of aggregation for your Quora Ads data. Example values: [Account, Campaign, Ad Set, Ad]
	Entity string `json:"entity"`
}

func (o *SourceQuoraAds) GetType() SourceQuoraAdsType {
	if o == nil {
		return SourceQuoraAdsType("")
	}
	return o.Type
}

func (o *SourceQuoraAds) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *SourceQuoraAds) GetLatencyThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LatencyThreshold
}

func (o *SourceQuoraAds) GetEntity() string {
	if o == nil {
		return ""
	}
	return o.Entity
}
