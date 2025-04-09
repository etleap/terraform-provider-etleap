// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ExtractedDataStrategyType string

const (
	ExtractedDataStrategyTypeExtractedData ExtractedDataStrategyType = "EXTRACTED_DATA"
)

func (e ExtractedDataStrategyType) ToPointer() *ExtractedDataStrategyType {
	return &e
}

func (e *ExtractedDataStrategyType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "EXTRACTED_DATA":
		*e = ExtractedDataStrategyType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ExtractedDataStrategyType: %v", v)
	}
}

type KeyValuePair struct {
	Value string `json:"value"`
	Key   string `json:"key"`
}

func (o *KeyValuePair) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

func (o *KeyValuePair) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

type ExtractedDataStrategy struct {
	Type             ExtractedDataStrategyType      `json:"type"`
	EndTimeParameter *WatermarkDatetimeKeyValuePair `json:"endTimeParameter,omitempty"`
	// Defines the query parameters to be included when fetching the most recently updated record. E.g., [{"key": "sort", "value": "updated_at"}, {"key": "order", "value": "desc"}]
	HighWatermarkQueryParameters []KeyValuePair                `json:"highWatermarkQueryParameters"`
	BeginTimeParameter           WatermarkDatetimeKeyValuePair `json:"beginTimeParameter"`
	LastUpdatedColumn            string                        `json:"lastUpdatedColumn"`
}

func (o *ExtractedDataStrategy) GetType() ExtractedDataStrategyType {
	if o == nil {
		return ExtractedDataStrategyType("")
	}
	return o.Type
}

func (o *ExtractedDataStrategy) GetEndTimeParameter() *WatermarkDatetimeKeyValuePair {
	if o == nil {
		return nil
	}
	return o.EndTimeParameter
}

func (o *ExtractedDataStrategy) GetHighWatermarkQueryParameters() []KeyValuePair {
	if o == nil {
		return []KeyValuePair{}
	}
	return o.HighWatermarkQueryParameters
}

func (o *ExtractedDataStrategy) GetBeginTimeParameter() WatermarkDatetimeKeyValuePair {
	if o == nil {
		return WatermarkDatetimeKeyValuePair{}
	}
	return o.BeginTimeParameter
}

func (o *ExtractedDataStrategy) GetLastUpdatedColumn() string {
	if o == nil {
		return ""
	}
	return o.LastUpdatedColumn
}
