// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
)

type UpdateScheduleTypesType string

const (
	UpdateScheduleTypesTypeInterval UpdateScheduleTypesType = "INTERVAL"
	UpdateScheduleTypesTypeHourly   UpdateScheduleTypesType = "HOURLY"
	UpdateScheduleTypesTypeDaily    UpdateScheduleTypesType = "DAILY"
	UpdateScheduleTypesTypeWeekly   UpdateScheduleTypesType = "WEEKLY"
	UpdateScheduleTypesTypeMonthly  UpdateScheduleTypesType = "MONTHLY"
)

// UpdateScheduleTypes - The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
type UpdateScheduleTypes struct {
	UpdateScheduleModeInterval *UpdateScheduleModeInterval
	UpdateScheduleModeHourly   *UpdateScheduleModeHourly
	UpdateScheduleModeDaily    *UpdateScheduleModeDaily
	UpdateScheduleModeWeekly   *UpdateScheduleModeWeekly
	UpdateScheduleModeMonthly  *UpdateScheduleModeMonthly

	Type UpdateScheduleTypesType
}

func CreateUpdateScheduleTypesInterval(interval UpdateScheduleModeInterval) UpdateScheduleTypes {
	typ := UpdateScheduleTypesTypeInterval

	typStr := Mode(typ)
	interval.Mode = typStr

	return UpdateScheduleTypes{
		UpdateScheduleModeInterval: &interval,
		Type:                       typ,
	}
}

func CreateUpdateScheduleTypesHourly(hourly UpdateScheduleModeHourly) UpdateScheduleTypes {
	typ := UpdateScheduleTypesTypeHourly

	typStr := UpdateScheduleModeHourlyMode(typ)
	hourly.Mode = typStr

	return UpdateScheduleTypes{
		UpdateScheduleModeHourly: &hourly,
		Type:                     typ,
	}
}

func CreateUpdateScheduleTypesDaily(daily UpdateScheduleModeDaily) UpdateScheduleTypes {
	typ := UpdateScheduleTypesTypeDaily

	typStr := UpdateScheduleModeDailyMode(typ)
	daily.Mode = typStr

	return UpdateScheduleTypes{
		UpdateScheduleModeDaily: &daily,
		Type:                    typ,
	}
}

func CreateUpdateScheduleTypesWeekly(weekly UpdateScheduleModeWeekly) UpdateScheduleTypes {
	typ := UpdateScheduleTypesTypeWeekly

	typStr := UpdateScheduleModeWeeklyMode(typ)
	weekly.Mode = typStr

	return UpdateScheduleTypes{
		UpdateScheduleModeWeekly: &weekly,
		Type:                     typ,
	}
}

func CreateUpdateScheduleTypesMonthly(monthly UpdateScheduleModeMonthly) UpdateScheduleTypes {
	typ := UpdateScheduleTypesTypeMonthly

	typStr := UpdateScheduleModeMonthlyMode(typ)
	monthly.Mode = typStr

	return UpdateScheduleTypes{
		UpdateScheduleModeMonthly: &monthly,
		Type:                      typ,
	}
}

func (u *UpdateScheduleTypes) UnmarshalJSON(data []byte) error {

	type discriminator struct {
		Mode string
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.Mode {
	case "INTERVAL":
		updateScheduleModeInterval := new(UpdateScheduleModeInterval)
		if err := utils.UnmarshalJSON(data, &updateScheduleModeInterval, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.UpdateScheduleModeInterval = updateScheduleModeInterval
		u.Type = UpdateScheduleTypesTypeInterval
		return nil
	case "HOURLY":
		updateScheduleModeHourly := new(UpdateScheduleModeHourly)
		if err := utils.UnmarshalJSON(data, &updateScheduleModeHourly, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.UpdateScheduleModeHourly = updateScheduleModeHourly
		u.Type = UpdateScheduleTypesTypeHourly
		return nil
	case "DAILY":
		updateScheduleModeDaily := new(UpdateScheduleModeDaily)
		if err := utils.UnmarshalJSON(data, &updateScheduleModeDaily, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.UpdateScheduleModeDaily = updateScheduleModeDaily
		u.Type = UpdateScheduleTypesTypeDaily
		return nil
	case "WEEKLY":
		updateScheduleModeWeekly := new(UpdateScheduleModeWeekly)
		if err := utils.UnmarshalJSON(data, &updateScheduleModeWeekly, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.UpdateScheduleModeWeekly = updateScheduleModeWeekly
		u.Type = UpdateScheduleTypesTypeWeekly
		return nil
	case "MONTHLY":
		updateScheduleModeMonthly := new(UpdateScheduleModeMonthly)
		if err := utils.UnmarshalJSON(data, &updateScheduleModeMonthly, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.UpdateScheduleModeMonthly = updateScheduleModeMonthly
		u.Type = UpdateScheduleTypesTypeMonthly
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u UpdateScheduleTypes) MarshalJSON() ([]byte, error) {
	if u.UpdateScheduleModeInterval != nil {
		return utils.MarshalJSON(u.UpdateScheduleModeInterval, "", true)
	}

	if u.UpdateScheduleModeHourly != nil {
		return utils.MarshalJSON(u.UpdateScheduleModeHourly, "", true)
	}

	if u.UpdateScheduleModeDaily != nil {
		return utils.MarshalJSON(u.UpdateScheduleModeDaily, "", true)
	}

	if u.UpdateScheduleModeWeekly != nil {
		return utils.MarshalJSON(u.UpdateScheduleModeWeekly, "", true)
	}

	if u.UpdateScheduleModeMonthly != nil {
		return utils.MarshalJSON(u.UpdateScheduleModeMonthly, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}
