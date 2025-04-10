// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceActiveCampaignType string

const (
	SourceActiveCampaignTypeActiveCampaign SourceActiveCampaignType = "ACTIVE_CAMPAIGN"
)

func (e SourceActiveCampaignType) ToPointer() *SourceActiveCampaignType {
	return &e
}

func (e *SourceActiveCampaignType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ACTIVE_CAMPAIGN":
		*e = SourceActiveCampaignType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceActiveCampaignType: %v", v)
	}
}

type SourceActiveCampaign struct {
	// The universally unique identifier for the source.
	ConnectionID string                   `json:"connectionId"`
	Type         SourceActiveCampaignType `json:"type"`
	// Notify if we can't extract for `x` hours. Setting it to `null` disables the notification. Defaults to `null`.
	LatencyThreshold *int64 `json:"latencyThreshold,omitempty"`
	// The ActiveCampaign resource. Example: Contacts, Custom Fields and Custom Values
	Entity string `json:"entity"`
}

func (o *SourceActiveCampaign) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *SourceActiveCampaign) GetType() SourceActiveCampaignType {
	if o == nil {
		return SourceActiveCampaignType("")
	}
	return o.Type
}

func (o *SourceActiveCampaign) GetLatencyThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LatencyThreshold
}

func (o *SourceActiveCampaign) GetEntity() string {
	if o == nil {
		return ""
	}
	return o.Entity
}
