// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceIntercomType string

const (
	SourceIntercomTypeIntercom SourceIntercomType = "INTERCOM"
)

func (e SourceIntercomType) ToPointer() *SourceIntercomType {
	return &e
}

func (e *SourceIntercomType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "INTERCOM":
		*e = SourceIntercomType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceIntercomType: %v", v)
	}
}

type SourceIntercom struct {
	Type SourceIntercomType `json:"type"`
	// The universally unique identifier for the source.
	ConnectionID string `json:"connectionId"`
	// Notify if we can't extract for `x` hours. Setting it to `null` disables the notification. Defaults to `null`.
	LatencyThreshold *int64 `json:"latencyThreshold,omitempty"`
	// The Intercom entity. Example values: [User, Lead, Contact, Company, Admin, Tag, Segment, Note, Event, Counts, Conversation Counts, Admin Conversation Counts, User Tags Counts, User Segments Counts, Company Tags Counts, Company Segments Counts, Conversation, Conversation Parts, Conversation Tags, Subscription]
	Entity string `json:"entity"`
}

func (o *SourceIntercom) GetType() SourceIntercomType {
	if o == nil {
		return SourceIntercomType("")
	}
	return o.Type
}

func (o *SourceIntercom) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *SourceIntercom) GetLatencyThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LatencyThreshold
}

func (o *SourceIntercom) GetEntity() string {
	if o == nil {
		return ""
	}
	return o.Entity
}
