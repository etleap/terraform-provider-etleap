// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
)

// PipelineUpdateAction - Whether Etleap should STOP the pipeline or NOTIFY once the `threshold` is reached.
type PipelineUpdateAction string

const (
	PipelineUpdateActionStop   PipelineUpdateAction = "STOP"
	PipelineUpdateActionNotify PipelineUpdateAction = "NOTIFY"
)

func (e PipelineUpdateAction) ToPointer() *PipelineUpdateAction {
	return &e
}

func (e *PipelineUpdateAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "STOP":
		fallthrough
	case "NOTIFY":
		*e = PipelineUpdateAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PipelineUpdateAction: %v", v)
	}
}

type PipelineUpdateParsingErrorSettings struct {
	// The parsing error threshold, in percentage points, for the `action` to be triggered.
	Threshold float64 `json:"threshold"`
	// Whether Etleap should STOP the pipeline or NOTIFY once the `threshold` is reached.
	Action PipelineUpdateAction `json:"action"`
}

func (o *PipelineUpdateParsingErrorSettings) GetThreshold() float64 {
	if o == nil {
		return 0.0
	}
	return o.Threshold
}

func (o *PipelineUpdateParsingErrorSettings) GetAction() PipelineUpdateAction {
	if o == nil {
		return PipelineUpdateAction("")
	}
	return o.Action
}

type UpdateScheduleModeMonthlyUpdateScheduleTypesMode string

const (
	UpdateScheduleModeMonthlyUpdateScheduleTypesModeMonthly UpdateScheduleModeMonthlyUpdateScheduleTypesMode = "MONTHLY"
)

func (e UpdateScheduleModeMonthlyUpdateScheduleTypesMode) ToPointer() *UpdateScheduleModeMonthlyUpdateScheduleTypesMode {
	return &e
}

func (e *UpdateScheduleModeMonthlyUpdateScheduleTypesMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MONTHLY":
		*e = UpdateScheduleModeMonthlyUpdateScheduleTypesMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateScheduleModeMonthlyUpdateScheduleTypesMode: %v", v)
	}
}

// MonthlyUpdateScheduleMode - The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information.
type MonthlyUpdateScheduleMode struct {
	Mode       UpdateScheduleModeMonthlyUpdateScheduleTypesMode `json:"mode"`
	DayOfMonth int64                                            `json:"dayOfMonth"`
	// Hour of day the  pipeline update should be started at (in UTC).
	HourOfDay int64 `json:"hourOfDay"`
}

func (o *MonthlyUpdateScheduleMode) GetMode() UpdateScheduleModeMonthlyUpdateScheduleTypesMode {
	if o == nil {
		return UpdateScheduleModeMonthlyUpdateScheduleTypesMode("")
	}
	return o.Mode
}

func (o *MonthlyUpdateScheduleMode) GetDayOfMonth() int64 {
	if o == nil {
		return 0
	}
	return o.DayOfMonth
}

func (o *MonthlyUpdateScheduleMode) GetHourOfDay() int64 {
	if o == nil {
		return 0
	}
	return o.HourOfDay
}

type UpdateScheduleModeWeeklyUpdateScheduleTypesMode string

const (
	UpdateScheduleModeWeeklyUpdateScheduleTypesModeWeekly UpdateScheduleModeWeeklyUpdateScheduleTypesMode = "WEEKLY"
)

func (e UpdateScheduleModeWeeklyUpdateScheduleTypesMode) ToPointer() *UpdateScheduleModeWeeklyUpdateScheduleTypesMode {
	return &e
}

func (e *UpdateScheduleModeWeeklyUpdateScheduleTypesMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "WEEKLY":
		*e = UpdateScheduleModeWeeklyUpdateScheduleTypesMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateScheduleModeWeeklyUpdateScheduleTypesMode: %v", v)
	}
}

// WeeklyUpdateScheduleMode - The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information.
type WeeklyUpdateScheduleMode struct {
	Mode      UpdateScheduleModeWeeklyUpdateScheduleTypesMode `json:"mode"`
	DayOfWeek int64                                           `json:"dayOfWeek"`
	// Hour of day the  pipeline update should be started at (in UTC).
	HourOfDay int64 `json:"hourOfDay"`
}

func (o *WeeklyUpdateScheduleMode) GetMode() UpdateScheduleModeWeeklyUpdateScheduleTypesMode {
	if o == nil {
		return UpdateScheduleModeWeeklyUpdateScheduleTypesMode("")
	}
	return o.Mode
}

func (o *WeeklyUpdateScheduleMode) GetDayOfWeek() int64 {
	if o == nil {
		return 0
	}
	return o.DayOfWeek
}

func (o *WeeklyUpdateScheduleMode) GetHourOfDay() int64 {
	if o == nil {
		return 0
	}
	return o.HourOfDay
}

type UpdateScheduleModeDailyUpdateScheduleTypesMode string

const (
	UpdateScheduleModeDailyUpdateScheduleTypesModeDaily UpdateScheduleModeDailyUpdateScheduleTypesMode = "DAILY"
)

func (e UpdateScheduleModeDailyUpdateScheduleTypesMode) ToPointer() *UpdateScheduleModeDailyUpdateScheduleTypesMode {
	return &e
}

func (e *UpdateScheduleModeDailyUpdateScheduleTypesMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DAILY":
		*e = UpdateScheduleModeDailyUpdateScheduleTypesMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateScheduleModeDailyUpdateScheduleTypesMode: %v", v)
	}
}

// DailyUpdateScheduleMode - The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information.
type DailyUpdateScheduleMode struct {
	Mode UpdateScheduleModeDailyUpdateScheduleTypesMode `json:"mode"`
	// Hour of day the  pipeline update should be started at (in UTC).
	HourOfDay int64 `json:"hourOfDay"`
}

func (o *DailyUpdateScheduleMode) GetMode() UpdateScheduleModeDailyUpdateScheduleTypesMode {
	if o == nil {
		return UpdateScheduleModeDailyUpdateScheduleTypesMode("")
	}
	return o.Mode
}

func (o *DailyUpdateScheduleMode) GetHourOfDay() int64 {
	if o == nil {
		return 0
	}
	return o.HourOfDay
}

type UpdateScheduleModeHourlyUpdateScheduleTypesMode string

const (
	UpdateScheduleModeHourlyUpdateScheduleTypesModeHourly UpdateScheduleModeHourlyUpdateScheduleTypesMode = "HOURLY"
)

func (e UpdateScheduleModeHourlyUpdateScheduleTypesMode) ToPointer() *UpdateScheduleModeHourlyUpdateScheduleTypesMode {
	return &e
}

func (e *UpdateScheduleModeHourlyUpdateScheduleTypesMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "HOURLY":
		*e = UpdateScheduleModeHourlyUpdateScheduleTypesMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateScheduleModeHourlyUpdateScheduleTypesMode: %v", v)
	}
}

// HourlyUpdateScheduleMode - The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information.
type HourlyUpdateScheduleMode struct {
	Mode UpdateScheduleModeHourlyUpdateScheduleTypesMode `json:"mode"`
}

func (o *HourlyUpdateScheduleMode) GetMode() UpdateScheduleModeHourlyUpdateScheduleTypesMode {
	if o == nil {
		return UpdateScheduleModeHourlyUpdateScheduleTypesMode("")
	}
	return o.Mode
}

type UpdateScheduleTypesMode string

const (
	UpdateScheduleTypesModeInterval UpdateScheduleTypesMode = "INTERVAL"
)

func (e UpdateScheduleTypesMode) ToPointer() *UpdateScheduleTypesMode {
	return &e
}

func (e *UpdateScheduleTypesMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "INTERVAL":
		*e = UpdateScheduleTypesMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateScheduleTypesMode: %v", v)
	}
}

// IntervalUpdateScheduleMode - The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information.
type IntervalUpdateScheduleMode struct {
	Mode UpdateScheduleTypesMode `json:"mode"`
	// Time to wait before new data is pulled (in minutes).
	IntervalMinutes int64 `json:"intervalMinutes"`
}

func (o *IntervalUpdateScheduleMode) GetMode() UpdateScheduleTypesMode {
	if o == nil {
		return UpdateScheduleTypesMode("")
	}
	return o.Mode
}

func (o *IntervalUpdateScheduleMode) GetIntervalMinutes() int64 {
	if o == nil {
		return 0
	}
	return o.IntervalMinutes
}

type PipelineUpdateUpdateScheduleTypesType string

const (
	PipelineUpdateUpdateScheduleTypesTypeInterval PipelineUpdateUpdateScheduleTypesType = "INTERVAL"
	PipelineUpdateUpdateScheduleTypesTypeHourly   PipelineUpdateUpdateScheduleTypesType = "HOURLY"
	PipelineUpdateUpdateScheduleTypesTypeDaily    PipelineUpdateUpdateScheduleTypesType = "DAILY"
	PipelineUpdateUpdateScheduleTypesTypeWeekly   PipelineUpdateUpdateScheduleTypesType = "WEEKLY"
	PipelineUpdateUpdateScheduleTypesTypeMonthly  PipelineUpdateUpdateScheduleTypesType = "MONTHLY"
)

// PipelineUpdateUpdateScheduleTypes - Setting the `updateSchedule` to `null` will remove the Pipeline Update Schedule and revert to using the Connection Update Schedule.
type PipelineUpdateUpdateScheduleTypes struct {
	IntervalUpdateScheduleMode *IntervalUpdateScheduleMode
	HourlyUpdateScheduleMode   *HourlyUpdateScheduleMode
	DailyUpdateScheduleMode    *DailyUpdateScheduleMode
	WeeklyUpdateScheduleMode   *WeeklyUpdateScheduleMode
	MonthlyUpdateScheduleMode  *MonthlyUpdateScheduleMode

	Type PipelineUpdateUpdateScheduleTypesType
}

func CreatePipelineUpdateUpdateScheduleTypesInterval(interval IntervalUpdateScheduleMode) PipelineUpdateUpdateScheduleTypes {
	typ := PipelineUpdateUpdateScheduleTypesTypeInterval

	typStr := UpdateScheduleTypesMode(typ)
	interval.Mode = typStr

	return PipelineUpdateUpdateScheduleTypes{
		IntervalUpdateScheduleMode: &interval,
		Type:                       typ,
	}
}

func CreatePipelineUpdateUpdateScheduleTypesHourly(hourly HourlyUpdateScheduleMode) PipelineUpdateUpdateScheduleTypes {
	typ := PipelineUpdateUpdateScheduleTypesTypeHourly

	typStr := UpdateScheduleModeHourlyUpdateScheduleTypesMode(typ)
	hourly.Mode = typStr

	return PipelineUpdateUpdateScheduleTypes{
		HourlyUpdateScheduleMode: &hourly,
		Type:                     typ,
	}
}

func CreatePipelineUpdateUpdateScheduleTypesDaily(daily DailyUpdateScheduleMode) PipelineUpdateUpdateScheduleTypes {
	typ := PipelineUpdateUpdateScheduleTypesTypeDaily

	typStr := UpdateScheduleModeDailyUpdateScheduleTypesMode(typ)
	daily.Mode = typStr

	return PipelineUpdateUpdateScheduleTypes{
		DailyUpdateScheduleMode: &daily,
		Type:                    typ,
	}
}

func CreatePipelineUpdateUpdateScheduleTypesWeekly(weekly WeeklyUpdateScheduleMode) PipelineUpdateUpdateScheduleTypes {
	typ := PipelineUpdateUpdateScheduleTypesTypeWeekly

	typStr := UpdateScheduleModeWeeklyUpdateScheduleTypesMode(typ)
	weekly.Mode = typStr

	return PipelineUpdateUpdateScheduleTypes{
		WeeklyUpdateScheduleMode: &weekly,
		Type:                     typ,
	}
}

func CreatePipelineUpdateUpdateScheduleTypesMonthly(monthly MonthlyUpdateScheduleMode) PipelineUpdateUpdateScheduleTypes {
	typ := PipelineUpdateUpdateScheduleTypesTypeMonthly

	typStr := UpdateScheduleModeMonthlyUpdateScheduleTypesMode(typ)
	monthly.Mode = typStr

	return PipelineUpdateUpdateScheduleTypes{
		MonthlyUpdateScheduleMode: &monthly,
		Type:                      typ,
	}
}

func (u *PipelineUpdateUpdateScheduleTypes) UnmarshalJSON(data []byte) error {

	type discriminator struct {
		Mode string
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.Mode {
	case "INTERVAL":
		intervalUpdateScheduleMode := new(IntervalUpdateScheduleMode)
		if err := utils.UnmarshalJSON(data, &intervalUpdateScheduleMode, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.IntervalUpdateScheduleMode = intervalUpdateScheduleMode
		u.Type = PipelineUpdateUpdateScheduleTypesTypeInterval
		return nil
	case "HOURLY":
		hourlyUpdateScheduleMode := new(HourlyUpdateScheduleMode)
		if err := utils.UnmarshalJSON(data, &hourlyUpdateScheduleMode, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.HourlyUpdateScheduleMode = hourlyUpdateScheduleMode
		u.Type = PipelineUpdateUpdateScheduleTypesTypeHourly
		return nil
	case "DAILY":
		dailyUpdateScheduleMode := new(DailyUpdateScheduleMode)
		if err := utils.UnmarshalJSON(data, &dailyUpdateScheduleMode, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.DailyUpdateScheduleMode = dailyUpdateScheduleMode
		u.Type = PipelineUpdateUpdateScheduleTypesTypeDaily
		return nil
	case "WEEKLY":
		weeklyUpdateScheduleMode := new(WeeklyUpdateScheduleMode)
		if err := utils.UnmarshalJSON(data, &weeklyUpdateScheduleMode, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.WeeklyUpdateScheduleMode = weeklyUpdateScheduleMode
		u.Type = PipelineUpdateUpdateScheduleTypesTypeWeekly
		return nil
	case "MONTHLY":
		monthlyUpdateScheduleMode := new(MonthlyUpdateScheduleMode)
		if err := utils.UnmarshalJSON(data, &monthlyUpdateScheduleMode, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.MonthlyUpdateScheduleMode = monthlyUpdateScheduleMode
		u.Type = PipelineUpdateUpdateScheduleTypesTypeMonthly
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u PipelineUpdateUpdateScheduleTypes) MarshalJSON() ([]byte, error) {
	if u.IntervalUpdateScheduleMode != nil {
		return utils.MarshalJSON(u.IntervalUpdateScheduleMode, "", true)
	}

	if u.HourlyUpdateScheduleMode != nil {
		return utils.MarshalJSON(u.HourlyUpdateScheduleMode, "", true)
	}

	if u.DailyUpdateScheduleMode != nil {
		return utils.MarshalJSON(u.DailyUpdateScheduleMode, "", true)
	}

	if u.WeeklyUpdateScheduleMode != nil {
		return utils.MarshalJSON(u.WeeklyUpdateScheduleMode, "", true)
	}

	if u.MonthlyUpdateScheduleMode != nil {
		return utils.MarshalJSON(u.MonthlyUpdateScheduleMode, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type RefreshScheduleModeMonthlyScheduleTypesMode string

const (
	RefreshScheduleModeMonthlyScheduleTypesModeMonthly RefreshScheduleModeMonthlyScheduleTypesMode = "MONTHLY"
)

func (e RefreshScheduleModeMonthlyScheduleTypesMode) ToPointer() *RefreshScheduleModeMonthlyScheduleTypesMode {
	return &e
}

func (e *RefreshScheduleModeMonthlyScheduleTypesMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MONTHLY":
		*e = RefreshScheduleModeMonthlyScheduleTypesMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RefreshScheduleModeMonthlyScheduleTypesMode: %v", v)
	}
}

type MonthlyScheduleMode struct {
	Mode RefreshScheduleModeMonthlyScheduleTypesMode `json:"mode"`
	// Day of the month this schedule should trigger at (in UTC).
	DayOfMonth int64 `json:"dayOfMonth"`
	// Hour of day this schedule should trigger at (in UTC).
	HourOfDay int64 `json:"hourOfDay"`
}

func (o *MonthlyScheduleMode) GetMode() RefreshScheduleModeMonthlyScheduleTypesMode {
	if o == nil {
		return RefreshScheduleModeMonthlyScheduleTypesMode("")
	}
	return o.Mode
}

func (o *MonthlyScheduleMode) GetDayOfMonth() int64 {
	if o == nil {
		return 0
	}
	return o.DayOfMonth
}

func (o *MonthlyScheduleMode) GetHourOfDay() int64 {
	if o == nil {
		return 0
	}
	return o.HourOfDay
}

type RefreshScheduleModeWeeklyScheduleTypesMode string

const (
	RefreshScheduleModeWeeklyScheduleTypesModeWeekly RefreshScheduleModeWeeklyScheduleTypesMode = "WEEKLY"
)

func (e RefreshScheduleModeWeeklyScheduleTypesMode) ToPointer() *RefreshScheduleModeWeeklyScheduleTypesMode {
	return &e
}

func (e *RefreshScheduleModeWeeklyScheduleTypesMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "WEEKLY":
		*e = RefreshScheduleModeWeeklyScheduleTypesMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RefreshScheduleModeWeeklyScheduleTypesMode: %v", v)
	}
}

type WeeklyScheduleMode struct {
	Mode RefreshScheduleModeWeeklyScheduleTypesMode `json:"mode"`
	// The day of the week this schedule should trigger at (in UTC).
	DayOfWeek int64 `json:"dayOfWeek"`
	// Hour of day this schedule should trigger at (in UTC).
	HourOfDay int64 `json:"hourOfDay"`
}

func (o *WeeklyScheduleMode) GetMode() RefreshScheduleModeWeeklyScheduleTypesMode {
	if o == nil {
		return RefreshScheduleModeWeeklyScheduleTypesMode("")
	}
	return o.Mode
}

func (o *WeeklyScheduleMode) GetDayOfWeek() int64 {
	if o == nil {
		return 0
	}
	return o.DayOfWeek
}

func (o *WeeklyScheduleMode) GetHourOfDay() int64 {
	if o == nil {
		return 0
	}
	return o.HourOfDay
}

type RefreshScheduleModeDailyScheduleTypesMode string

const (
	RefreshScheduleModeDailyScheduleTypesModeDaily RefreshScheduleModeDailyScheduleTypesMode = "DAILY"
)

func (e RefreshScheduleModeDailyScheduleTypesMode) ToPointer() *RefreshScheduleModeDailyScheduleTypesMode {
	return &e
}

func (e *RefreshScheduleModeDailyScheduleTypesMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DAILY":
		*e = RefreshScheduleModeDailyScheduleTypesMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RefreshScheduleModeDailyScheduleTypesMode: %v", v)
	}
}

type DailyScheduleMode struct {
	Mode RefreshScheduleModeDailyScheduleTypesMode `json:"mode"`
	// Hour of day this schedule should trigger at (in UTC).
	HourOfDay int64 `json:"hourOfDay"`
}

func (o *DailyScheduleMode) GetMode() RefreshScheduleModeDailyScheduleTypesMode {
	if o == nil {
		return RefreshScheduleModeDailyScheduleTypesMode("")
	}
	return o.Mode
}

func (o *DailyScheduleMode) GetHourOfDay() int64 {
	if o == nil {
		return 0
	}
	return o.HourOfDay
}

type RefreshScheduleModeHourlyScheduleTypesMode string

const (
	RefreshScheduleModeHourlyScheduleTypesModeHourly RefreshScheduleModeHourlyScheduleTypesMode = "HOURLY"
)

func (e RefreshScheduleModeHourlyScheduleTypesMode) ToPointer() *RefreshScheduleModeHourlyScheduleTypesMode {
	return &e
}

func (e *RefreshScheduleModeHourlyScheduleTypesMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "HOURLY":
		*e = RefreshScheduleModeHourlyScheduleTypesMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RefreshScheduleModeHourlyScheduleTypesMode: %v", v)
	}
}

type HourlyScheduleMode struct {
	Mode RefreshScheduleModeHourlyScheduleTypesMode `json:"mode"`
}

func (o *HourlyScheduleMode) GetMode() RefreshScheduleModeHourlyScheduleTypesMode {
	if o == nil {
		return RefreshScheduleModeHourlyScheduleTypesMode("")
	}
	return o.Mode
}

type ScheduleTypesMode string

const (
	ScheduleTypesModeNever ScheduleTypesMode = "NEVER"
)

func (e ScheduleTypesMode) ToPointer() *ScheduleTypesMode {
	return &e
}

func (e *ScheduleTypesMode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NEVER":
		*e = ScheduleTypesMode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ScheduleTypesMode: %v", v)
	}
}

type NeverScheduleMode struct {
	Mode ScheduleTypesMode `json:"mode"`
}

func (o *NeverScheduleMode) GetMode() ScheduleTypesMode {
	if o == nil {
		return ScheduleTypesMode("")
	}
	return o.Mode
}

type ScheduleTypesType string

const (
	ScheduleTypesTypeNever   ScheduleTypesType = "NEVER"
	ScheduleTypesTypeHourly  ScheduleTypesType = "HOURLY"
	ScheduleTypesTypeDaily   ScheduleTypesType = "DAILY"
	ScheduleTypesTypeWeekly  ScheduleTypesType = "WEEKLY"
	ScheduleTypesTypeMonthly ScheduleTypesType = "MONTHLY"
)

// ScheduleTypes - A pipeline refresh processes all data in your source from the beginning to re-establish consistency with your destination. The pipeline refresh schedule defines when Etleap should automatically refresh the pipeline. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information.
//
// Setting this to `null` is equivalent to setting the Refresh Schedule to `NEVER`.
type ScheduleTypes struct {
	NeverScheduleMode   *NeverScheduleMode
	HourlyScheduleMode  *HourlyScheduleMode
	DailyScheduleMode   *DailyScheduleMode
	WeeklyScheduleMode  *WeeklyScheduleMode
	MonthlyScheduleMode *MonthlyScheduleMode

	Type ScheduleTypesType
}

func CreateScheduleTypesNever(never NeverScheduleMode) ScheduleTypes {
	typ := ScheduleTypesTypeNever

	typStr := ScheduleTypesMode(typ)
	never.Mode = typStr

	return ScheduleTypes{
		NeverScheduleMode: &never,
		Type:              typ,
	}
}

func CreateScheduleTypesHourly(hourly HourlyScheduleMode) ScheduleTypes {
	typ := ScheduleTypesTypeHourly

	typStr := RefreshScheduleModeHourlyScheduleTypesMode(typ)
	hourly.Mode = typStr

	return ScheduleTypes{
		HourlyScheduleMode: &hourly,
		Type:               typ,
	}
}

func CreateScheduleTypesDaily(daily DailyScheduleMode) ScheduleTypes {
	typ := ScheduleTypesTypeDaily

	typStr := RefreshScheduleModeDailyScheduleTypesMode(typ)
	daily.Mode = typStr

	return ScheduleTypes{
		DailyScheduleMode: &daily,
		Type:              typ,
	}
}

func CreateScheduleTypesWeekly(weekly WeeklyScheduleMode) ScheduleTypes {
	typ := ScheduleTypesTypeWeekly

	typStr := RefreshScheduleModeWeeklyScheduleTypesMode(typ)
	weekly.Mode = typStr

	return ScheduleTypes{
		WeeklyScheduleMode: &weekly,
		Type:               typ,
	}
}

func CreateScheduleTypesMonthly(monthly MonthlyScheduleMode) ScheduleTypes {
	typ := ScheduleTypesTypeMonthly

	typStr := RefreshScheduleModeMonthlyScheduleTypesMode(typ)
	monthly.Mode = typStr

	return ScheduleTypes{
		MonthlyScheduleMode: &monthly,
		Type:                typ,
	}
}

func (u *ScheduleTypes) UnmarshalJSON(data []byte) error {

	type discriminator struct {
		Mode string
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.Mode {
	case "NEVER":
		neverScheduleMode := new(NeverScheduleMode)
		if err := utils.UnmarshalJSON(data, &neverScheduleMode, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.NeverScheduleMode = neverScheduleMode
		u.Type = ScheduleTypesTypeNever
		return nil
	case "HOURLY":
		hourlyScheduleMode := new(HourlyScheduleMode)
		if err := utils.UnmarshalJSON(data, &hourlyScheduleMode, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.HourlyScheduleMode = hourlyScheduleMode
		u.Type = ScheduleTypesTypeHourly
		return nil
	case "DAILY":
		dailyScheduleMode := new(DailyScheduleMode)
		if err := utils.UnmarshalJSON(data, &dailyScheduleMode, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.DailyScheduleMode = dailyScheduleMode
		u.Type = ScheduleTypesTypeDaily
		return nil
	case "WEEKLY":
		weeklyScheduleMode := new(WeeklyScheduleMode)
		if err := utils.UnmarshalJSON(data, &weeklyScheduleMode, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.WeeklyScheduleMode = weeklyScheduleMode
		u.Type = ScheduleTypesTypeWeekly
		return nil
	case "MONTHLY":
		monthlyScheduleMode := new(MonthlyScheduleMode)
		if err := utils.UnmarshalJSON(data, &monthlyScheduleMode, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.MonthlyScheduleMode = monthlyScheduleMode
		u.Type = ScheduleTypesTypeMonthly
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u ScheduleTypes) MarshalJSON() ([]byte, error) {
	if u.NeverScheduleMode != nil {
		return utils.MarshalJSON(u.NeverScheduleMode, "", true)
	}

	if u.HourlyScheduleMode != nil {
		return utils.MarshalJSON(u.HourlyScheduleMode, "", true)
	}

	if u.DailyScheduleMode != nil {
		return utils.MarshalJSON(u.DailyScheduleMode, "", true)
	}

	if u.WeeklyScheduleMode != nil {
		return utils.MarshalJSON(u.WeeklyScheduleMode, "", true)
	}

	if u.MonthlyScheduleMode != nil {
		return utils.MarshalJSON(u.MonthlyScheduleMode, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type PipelineUpdate struct {
	Name              *string             `json:"name,omitempty"`
	SourceUpdate      *SourceUpdate       `json:"source,omitempty"`
	DestinationUpdate []DestinationUpdate `json:"destination,omitempty"`
	Paused            *bool               `json:"paused,omitempty"`
	// A list of users' email to share the pipeline with.
	//
	// A pipeline cannot be unshared; therefore future calls can only add to this list.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	Shares               []string                            `json:"shares,omitempty"`
	ParsingErrorSettings *PipelineUpdateParsingErrorSettings `json:"parsingErrorSettings,omitempty"`
	// Setting the `updateSchedule` to `null` will remove the Pipeline Update Schedule and revert to using the Connection Update Schedule.
	UpdateSchedule *PipelineUpdateUpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// A pipeline refresh processes all data in your source from the beginning to re-establish consistency with your destination. The pipeline refresh schedule defines when Etleap should automatically refresh the pipeline. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information.
	//
	// Setting this to `null` is equivalent to setting the Refresh Schedule to `NEVER`.
	RefreshSchedule *ScheduleTypes `json:"refreshSchedule,omitempty"`
}

func (o *PipelineUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *PipelineUpdate) GetSourceUpdate() *SourceUpdate {
	if o == nil {
		return nil
	}
	return o.SourceUpdate
}

func (o *PipelineUpdate) GetDestinationUpdate() []DestinationUpdate {
	if o == nil {
		return nil
	}
	return o.DestinationUpdate
}

func (o *PipelineUpdate) GetPaused() *bool {
	if o == nil {
		return nil
	}
	return o.Paused
}

func (o *PipelineUpdate) GetShares() []string {
	if o == nil {
		return nil
	}
	return o.Shares
}

func (o *PipelineUpdate) GetParsingErrorSettings() *PipelineUpdateParsingErrorSettings {
	if o == nil {
		return nil
	}
	return o.ParsingErrorSettings
}

func (o *PipelineUpdate) GetUpdateSchedule() *PipelineUpdateUpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *PipelineUpdate) GetUpdateScheduleInterval() *IntervalUpdateScheduleMode {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.IntervalUpdateScheduleMode
	}
	return nil
}

func (o *PipelineUpdate) GetUpdateScheduleHourly() *HourlyUpdateScheduleMode {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.HourlyUpdateScheduleMode
	}
	return nil
}

func (o *PipelineUpdate) GetUpdateScheduleDaily() *DailyUpdateScheduleMode {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.DailyUpdateScheduleMode
	}
	return nil
}

func (o *PipelineUpdate) GetUpdateScheduleWeekly() *WeeklyUpdateScheduleMode {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.WeeklyUpdateScheduleMode
	}
	return nil
}

func (o *PipelineUpdate) GetUpdateScheduleMonthly() *MonthlyUpdateScheduleMode {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.MonthlyUpdateScheduleMode
	}
	return nil
}

func (o *PipelineUpdate) GetRefreshSchedule() *ScheduleTypes {
	if o == nil {
		return nil
	}
	return o.RefreshSchedule
}

func (o *PipelineUpdate) GetRefreshScheduleNever() *NeverScheduleMode {
	if v := o.GetRefreshSchedule(); v != nil {
		return v.NeverScheduleMode
	}
	return nil
}

func (o *PipelineUpdate) GetRefreshScheduleHourly() *HourlyScheduleMode {
	if v := o.GetRefreshSchedule(); v != nil {
		return v.HourlyScheduleMode
	}
	return nil
}

func (o *PipelineUpdate) GetRefreshScheduleDaily() *DailyScheduleMode {
	if v := o.GetRefreshSchedule(); v != nil {
		return v.DailyScheduleMode
	}
	return nil
}

func (o *PipelineUpdate) GetRefreshScheduleWeekly() *WeeklyScheduleMode {
	if v := o.GetRefreshSchedule(); v != nil {
		return v.WeeklyScheduleMode
	}
	return nil
}

func (o *PipelineUpdate) GetRefreshScheduleMonthly() *MonthlyScheduleMode {
	if v := o.GetRefreshSchedule(); v != nil {
		return v.MonthlyScheduleMode
	}
	return nil
}
