// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceFreshchatType string

const (
	SourceFreshchatTypeFreshchat SourceFreshchatType = "FRESHCHAT"
)

func (e SourceFreshchatType) ToPointer() *SourceFreshchatType {
	return &e
}

func (e *SourceFreshchatType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "FRESHCHAT":
		*e = SourceFreshchatType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFreshchatType: %v", v)
	}
}

type SourceFreshchat struct {
	// The universally unique identifier for the source.
	ConnectionID string              `json:"connectionId"`
	Type         SourceFreshchatType `json:"type"`
	// Notify if we can't extract for `x` hours. Setting it to `null` disables the notification. Defaults to `null`.
	LatencyThreshold *int64 `json:"latencyThreshold,omitempty"`
	// The Freshchat resource. Example values: [Agents, Channels, Conversations, Conversation Messages]
	Entity string `json:"entity"`
	// Only when Entity is related to Deals. Select which views you want Etleap to pull data from.
	View []string `json:"view,omitempty"`
}

func (o *SourceFreshchat) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *SourceFreshchat) GetType() SourceFreshchatType {
	if o == nil {
		return SourceFreshchatType("")
	}
	return o.Type
}

func (o *SourceFreshchat) GetLatencyThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LatencyThreshold
}

func (o *SourceFreshchat) GetEntity() string {
	if o == nil {
		return ""
	}
	return o.Entity
}

func (o *SourceFreshchat) GetView() []string {
	if o == nil {
		return nil
	}
	return o.View
}
