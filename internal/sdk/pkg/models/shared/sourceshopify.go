// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceShopifyType string

const (
	SourceShopifyTypeShopify SourceShopifyType = "SHOPIFY"
)

func (e SourceShopifyType) ToPointer() *SourceShopifyType {
	return &e
}

func (e *SourceShopifyType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SHOPIFY":
		*e = SourceShopifyType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceShopifyType: %v", v)
	}
}

type SourceShopify struct {
	// The universally unique identifier for the source.
	ConnectionID string            `json:"connectionId"`
	Type         SourceShopifyType `json:"type"`
	// Notify if we can't extract for `x` hours. Setting it to `null` disables the notification. Defaults to `null`.
	LatencyThreshold *int64 `json:"latencyThreshold,omitempty"`
	// The Shopify entity. Spelled with spaces and only first word capitalized. Nested JSON objects are selected by appending the field name. For example, `Orders fulfillments line items` has the lineItems field from the `Order fulfillments` entity. Start creating a pipeline in the Etleap UI for the full list of entities.
	Entity string `json:"entity"`
}

func (o *SourceShopify) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *SourceShopify) GetType() SourceShopifyType {
	if o == nil {
		return SourceShopifyType("")
	}
	return o.Type
}

func (o *SourceShopify) GetLatencyThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LatencyThreshold
}

func (o *SourceShopify) GetEntity() string {
	if o == nil {
		return ""
	}
	return o.Entity
}
