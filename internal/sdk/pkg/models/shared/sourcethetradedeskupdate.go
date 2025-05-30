// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceTheTradeDeskUpdateType string

const (
	SourceTheTradeDeskUpdateTypeTheTradeDesk SourceTheTradeDeskUpdateType = "THE_TRADE_DESK"
)

func (e SourceTheTradeDeskUpdateType) ToPointer() *SourceTheTradeDeskUpdateType {
	return &e
}

func (e *SourceTheTradeDeskUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "THE_TRADE_DESK":
		*e = SourceTheTradeDeskUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceTheTradeDeskUpdateType: %v", v)
	}
}

type SourceTheTradeDeskUpdate struct {
	// Notify if we can't extract for `x` hours. Setting it to `null` disables the notification. Defaults to `null`.
	LatencyThreshold *int64                        `json:"latencyThreshold,omitempty"`
	Type             *SourceTheTradeDeskUpdateType `json:"type,omitempty"`
}

func (o *SourceTheTradeDeskUpdate) GetLatencyThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LatencyThreshold
}

func (o *SourceTheTradeDeskUpdate) GetType() *SourceTheTradeDeskUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}
