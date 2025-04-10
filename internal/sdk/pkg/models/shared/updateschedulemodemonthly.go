// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type UpdateScheduleModeMonthlyMode string

const (
	UpdateScheduleModeMonthlyModeMonthly UpdateScheduleModeMonthlyMode = "MONTHLY"
)

func (e UpdateScheduleModeMonthlyMode) ToPointer() *UpdateScheduleModeMonthlyMode {
	return &e
}

func (e *UpdateScheduleModeMonthlyMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MONTHLY":
		*e = UpdateScheduleModeMonthlyMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateScheduleModeMonthlyMode: %v", v)
	}
}

// UpdateScheduleModeMonthly - The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information.
type UpdateScheduleModeMonthly struct {
	Mode UpdateScheduleModeMonthlyMode `json:"mode"`
	// Hour of day the  pipeline update should be started at (in UTC).
	HourOfDay  int64 `json:"hourOfDay"`
	DayOfMonth int64 `json:"dayOfMonth"`
}

func (o *UpdateScheduleModeMonthly) GetMode() UpdateScheduleModeMonthlyMode {
	if o == nil {
		return UpdateScheduleModeMonthlyMode("")
	}
	return o.Mode
}

func (o *UpdateScheduleModeMonthly) GetHourOfDay() int64 {
	if o == nil {
		return 0
	}
	return o.HourOfDay
}

func (o *UpdateScheduleModeMonthly) GetDayOfMonth() int64 {
	if o == nil {
		return 0
	}
	return o.DayOfMonth
}
