// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceSalesforceMarketingCloudUpdateType string

const (
	SourceSalesforceMarketingCloudUpdateTypeSalesforceMarketingCloud SourceSalesforceMarketingCloudUpdateType = "SALESFORCE_MARKETING_CLOUD"
)

func (e SourceSalesforceMarketingCloudUpdateType) ToPointer() *SourceSalesforceMarketingCloudUpdateType {
	return &e
}

func (e *SourceSalesforceMarketingCloudUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SALESFORCE_MARKETING_CLOUD":
		*e = SourceSalesforceMarketingCloudUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSalesforceMarketingCloudUpdateType: %v", v)
	}
}

type SourceSalesforceMarketingCloudUpdate struct {
	Type *SourceSalesforceMarketingCloudUpdateType `json:"type,omitempty"`
	// Notify if we can't extract for `x` hours. Setting it to `null` disables the notification. Defaults to `null`.
	LatencyThreshold *int64 `json:"latencyThreshold,omitempty"`
}

func (o *SourceSalesforceMarketingCloudUpdate) GetType() *SourceSalesforceMarketingCloudUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *SourceSalesforceMarketingCloudUpdate) GetLatencyThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LatencyThreshold
}
