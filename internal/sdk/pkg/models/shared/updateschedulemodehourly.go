// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type UpdateScheduleModeHourlyMode string

const (
	UpdateScheduleModeHourlyModeHourly UpdateScheduleModeHourlyMode = "HOURLY"
)

func (e UpdateScheduleModeHourlyMode) ToPointer() *UpdateScheduleModeHourlyMode {
	return &e
}

func (e *UpdateScheduleModeHourlyMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "HOURLY":
		*e = UpdateScheduleModeHourlyMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateScheduleModeHourlyMode: %v", v)
	}
}

// UpdateScheduleModeHourly - The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information.
type UpdateScheduleModeHourly struct {
	Mode UpdateScheduleModeHourlyMode `json:"mode"`
}

func (o *UpdateScheduleModeHourly) GetMode() UpdateScheduleModeHourlyMode {
	if o == nil {
		return UpdateScheduleModeHourlyMode("")
	}
	return o.Mode
}
