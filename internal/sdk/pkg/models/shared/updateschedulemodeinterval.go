// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type Mode string

const (
	ModeInterval Mode = "INTERVAL"
)

func (e Mode) ToPointer() *Mode {
	return &e
}

func (e *Mode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "INTERVAL":
		*e = Mode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Mode: %v", v)
	}
}

// UpdateScheduleModeInterval - The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information.
type UpdateScheduleModeInterval struct {
	Mode Mode `json:"mode"`
	// Time to wait before new data is pulled (in minutes).
	IntervalMinutes int64 `json:"intervalMinutes"`
}

func (o *UpdateScheduleModeInterval) GetMode() Mode {
	if o == nil {
		return Mode("")
	}
	return o.Mode
}

func (o *UpdateScheduleModeInterval) GetIntervalMinutes() int64 {
	if o == nil {
		return 0
	}
	return o.IntervalMinutes
}
