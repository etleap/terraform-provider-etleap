// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceHubspotType string

const (
	SourceHubspotTypeHubspot SourceHubspotType = "HUBSPOT"
)

func (e SourceHubspotType) ToPointer() *SourceHubspotType {
	return &e
}

func (e *SourceHubspotType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "HUBSPOT":
		*e = SourceHubspotType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceHubspotType: %v", v)
	}
}

type SourceHubspot struct {
	// The universally unique identifier for the source.
	ConnectionID string            `json:"connectionId"`
	Type         SourceHubspotType `json:"type"`
	// Notify if we can't extract for `x` hours. Setting it to `null` disables the notification. Defaults to `null`.
	LatencyThreshold *int64 `json:"latencyThreshold,omitempty"`
	// The Hubspot entity. Example values: [Campaigns, Contacts, Email Events, Engagements, Deals, Owners, Deal Pipelines, Companies, Marketing Emails, Pages, Landing Pages Analytics]
	Entity string `json:"entity"`
}

func (o *SourceHubspot) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *SourceHubspot) GetType() SourceHubspotType {
	if o == nil {
		return SourceHubspotType("")
	}
	return o.Type
}

func (o *SourceHubspot) GetLatencyThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LatencyThreshold
}

func (o *SourceHubspot) GetEntity() string {
	if o == nil {
		return ""
	}
	return o.Entity
}
