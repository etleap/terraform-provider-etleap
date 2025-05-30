// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceStripeUpdateType string

const (
	SourceStripeUpdateTypeStripe SourceStripeUpdateType = "STRIPE"
)

func (e SourceStripeUpdateType) ToPointer() *SourceStripeUpdateType {
	return &e
}

func (e *SourceStripeUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "STRIPE":
		*e = SourceStripeUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceStripeUpdateType: %v", v)
	}
}

type SourceStripeUpdate struct {
	// Notify if we can't extract for `x` hours. Setting it to `null` disables the notification. Defaults to `null`.
	LatencyThreshold *int64                  `json:"latencyThreshold,omitempty"`
	Type             *SourceStripeUpdateType `json:"type,omitempty"`
}

func (o *SourceStripeUpdate) GetLatencyThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LatencyThreshold
}

func (o *SourceStripeUpdate) GetType() *SourceStripeUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}
