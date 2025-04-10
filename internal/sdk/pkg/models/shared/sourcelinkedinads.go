// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceLinkedInAdsType string

const (
	SourceLinkedInAdsTypeLinkedInAds SourceLinkedInAdsType = "LINKED_IN_ADS"
)

func (e SourceLinkedInAdsType) ToPointer() *SourceLinkedInAdsType {
	return &e
}

func (e *SourceLinkedInAdsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "LINKED_IN_ADS":
		*e = SourceLinkedInAdsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceLinkedInAdsType: %v", v)
	}
}

type SourceLinkedInAds struct {
	// The universally unique identifier for the source.
	ConnectionID string                `json:"connectionId"`
	Type         SourceLinkedInAdsType `json:"type"`
	// Notify if we can't extract for `x` hours. Setting it to `null` disables the notification. Defaults to `null`.
	LatencyThreshold *int64 `json:"latencyThreshold,omitempty"`
	// The LinkedIn resource. Example values: [ACCOUNTS, ACCOUNT_USERS, AD_ANALYTICS, CAMPAIGNS, CAMPAIGN_GROUPS, CONVERSIONS, INSIGHT_TAG_DOMAINS]
	Entity string `json:"entity"`
	// Specify the report `metrics` if and only if the entity is 'AD_ANALYTICS'. Example values: [dateRange, pivotValues, clicks]
	Metrics []string `json:"metrics,omitempty"`
	// Specify the report `pivots` groups if and only if the entity is 'AD_ANALYTICS'. Example values: [ACCOUNT, CAMPAIGN, COMPANY]
	Pivots []string `json:"pivots,omitempty"`
}

func (o *SourceLinkedInAds) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *SourceLinkedInAds) GetType() SourceLinkedInAdsType {
	if o == nil {
		return SourceLinkedInAdsType("")
	}
	return o.Type
}

func (o *SourceLinkedInAds) GetLatencyThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LatencyThreshold
}

func (o *SourceLinkedInAds) GetEntity() string {
	if o == nil {
		return ""
	}
	return o.Entity
}

func (o *SourceLinkedInAds) GetMetrics() []string {
	if o == nil {
		return nil
	}
	return o.Metrics
}

func (o *SourceLinkedInAds) GetPivots() []string {
	if o == nil {
		return nil
	}
	return o.Pivots
}
