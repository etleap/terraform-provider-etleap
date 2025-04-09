// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type WallClockStrategyType string

const (
	WallClockStrategyTypeWallClock WallClockStrategyType = "WALL_CLOCK"
)

func (e WallClockStrategyType) ToPointer() *WallClockStrategyType {
	return &e
}

func (e *WallClockStrategyType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "WALL_CLOCK":
		*e = WallClockStrategyType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for WallClockStrategyType: %v", v)
	}
}

type WallClockStrategy struct {
	Type             WallClockStrategyType                              `json:"type"`
	EndDateParameter *InclusiveOrExclusiveWatermarkDateonlyKeyValuePair `json:"endDateParameter,omitempty"`
	// The number of days we look back during each delta extraction.
	LookbackWindow     int64                         `json:"lookbackWindow"`
	BeginDateParameter WatermarkDateonlyKeyValuePair `json:"beginDateParameter"`
}

func (o *WallClockStrategy) GetType() WallClockStrategyType {
	if o == nil {
		return WallClockStrategyType("")
	}
	return o.Type
}

func (o *WallClockStrategy) GetEndDateParameter() *InclusiveOrExclusiveWatermarkDateonlyKeyValuePair {
	if o == nil {
		return nil
	}
	return o.EndDateParameter
}

func (o *WallClockStrategy) GetLookbackWindow() int64 {
	if o == nil {
		return 0
	}
	return o.LookbackWindow
}

func (o *WallClockStrategy) GetBeginDateParameter() WatermarkDateonlyKeyValuePair {
	if o == nil {
		return WatermarkDateonlyKeyValuePair{}
	}
	return o.BeginDateParameter
}
