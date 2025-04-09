// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceElasticSearchType string

const (
	SourceElasticSearchTypeElasticsearch SourceElasticSearchType = "ELASTICSEARCH"
)

func (e SourceElasticSearchType) ToPointer() *SourceElasticSearchType {
	return &e
}

func (e *SourceElasticSearchType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ELASTICSEARCH":
		*e = SourceElasticSearchType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceElasticSearchType: %v", v)
	}
}

type SourceElasticSearch struct {
	// The universally unique identifier for the source.
	ConnectionID string                  `json:"connectionId"`
	Type         SourceElasticSearchType `json:"type"`
	// Notify if we can't extract for `x` hours. Setting it to `null` disables the notification. Defaults to `null`.
	LatencyThreshold *int64 `json:"latencyThreshold,omitempty"`
	// The index name.
	Entity string `json:"entity"`
}

func (o *SourceElasticSearch) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *SourceElasticSearch) GetType() SourceElasticSearchType {
	if o == nil {
		return SourceElasticSearchType("")
	}
	return o.Type
}

func (o *SourceElasticSearch) GetLatencyThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LatencyThreshold
}

func (o *SourceElasticSearch) GetEntity() string {
	if o == nil {
		return ""
	}
	return o.Entity
}
