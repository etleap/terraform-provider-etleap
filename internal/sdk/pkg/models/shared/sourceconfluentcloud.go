// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceConfluentCloudType string

const (
	SourceConfluentCloudTypeConfluentCloud SourceConfluentCloudType = "CONFLUENT_CLOUD"
)

func (e SourceConfluentCloudType) ToPointer() *SourceConfluentCloudType {
	return &e
}

func (e *SourceConfluentCloudType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CONFLUENT_CLOUD":
		*e = SourceConfluentCloudType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceConfluentCloudType: %v", v)
	}
}

type SourceConfluentCloud struct {
	// The universally unique identifier for the source.
	ConnectionID string                   `json:"connectionId"`
	Type         SourceConfluentCloudType `json:"type"`
	// Notify if we can't extract for `x` hours. Setting it to `null` disables the notification. Defaults to `null`.
	LatencyThreshold *int64 `json:"latencyThreshold,omitempty"`
	// You can ingest data from Kafka topics in your Confluent Cloud cluster.
	Entity string `json:"entity"`
}

func (o *SourceConfluentCloud) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *SourceConfluentCloud) GetType() SourceConfluentCloudType {
	if o == nil {
		return SourceConfluentCloudType("")
	}
	return o.Type
}

func (o *SourceConfluentCloud) GetLatencyThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LatencyThreshold
}

func (o *SourceConfluentCloud) GetEntity() string {
	if o == nil {
		return ""
	}
	return o.Entity
}
