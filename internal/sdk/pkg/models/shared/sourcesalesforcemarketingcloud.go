// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceSalesforceMarketingCloudType string

const (
	SourceSalesforceMarketingCloudTypeSalesforceMarketingCloud SourceSalesforceMarketingCloudType = "SALESFORCE_MARKETING_CLOUD"
)

func (e SourceSalesforceMarketingCloudType) ToPointer() *SourceSalesforceMarketingCloudType {
	return &e
}

func (e *SourceSalesforceMarketingCloudType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SALESFORCE_MARKETING_CLOUD":
		*e = SourceSalesforceMarketingCloudType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSalesforceMarketingCloudType: %v", v)
	}
}

type SourceSalesforceMarketingCloud struct {
	Type SourceSalesforceMarketingCloudType `json:"type"`
	// The universally unique identifier for the source.
	ConnectionID string `json:"connectionId"`
	// Notify if we can't extract for `x` hours. Setting it to `null` disables the notification. Defaults to `null`.
	LatencyThreshold *int64 `json:"latencyThreshold,omitempty"`
	// The Salesforce Marketing Cloud entity. Example Values: [Bounce Event, Campaign, Click Event, Content Area, Data Extension, Data Extension Object, Email, Folders, List Subscriber, Lists, Open Event, Send, Sent Event, Subscribers, Unsub Event]
	Entity string `json:"entity"`
}

func (o *SourceSalesforceMarketingCloud) GetType() SourceSalesforceMarketingCloudType {
	if o == nil {
		return SourceSalesforceMarketingCloudType("")
	}
	return o.Type
}

func (o *SourceSalesforceMarketingCloud) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *SourceSalesforceMarketingCloud) GetLatencyThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LatencyThreshold
}

func (o *SourceSalesforceMarketingCloud) GetEntity() string {
	if o == nil {
		return ""
	}
	return o.Entity
}
