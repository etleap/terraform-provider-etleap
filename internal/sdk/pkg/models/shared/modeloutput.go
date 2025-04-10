// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
	"time"
)

type ModelOutput struct {
	QueryAndTriggers QueryAndTriggers `json:"queryAndTriggers"`
	Name             string           `json:"name"`
	// The date and time of the latest successful update for this model.
	LastUpdateTime *time.Time `json:"lastUpdateTime,omitempty"`
	// The date and time when then the model was created.
	CreateDate *time.Time `json:"createDate,omitempty"`
	// An array of users' emails that the model is shared with.  Once shared, a model cannot be unshared, and future calls to `PATCH` can only add to this list.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	Shares       []string          `json:"shares"`
	Paused       bool              `json:"paused"`
	Owner        User              `json:"owner"`
	Dependencies []ModelDependency `json:"dependencies"`
	Warehouse    WarehouseTypes    `json:"warehouse"`
	// How long the latest update took to complete, in milliseconds, or the duration of the current update if one is in progress.
	LastUpdateDuration *int64 `json:"lastUpdateDuration,omitempty"`
	// The unique identifier of the model.
	ID string `json:"id"`
	// How often this model should update. Etleap will periodically update the model table in your warehouse according to this schedule. See [the Model Updates documentation](https://docs.etleap.com/docs/documentation/ZG9jOjI0MzU2NDY3-introduction-to-models#model-updates) for more information.
	UpdateSchedule RefreshScheduleTypes `json:"updateSchedule"`
}

func (m ModelOutput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *ModelOutput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ModelOutput) GetQueryAndTriggers() QueryAndTriggers {
	if o == nil {
		return QueryAndTriggers{}
	}
	return o.QueryAndTriggers
}

func (o *ModelOutput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ModelOutput) GetLastUpdateTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.LastUpdateTime
}

func (o *ModelOutput) GetCreateDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreateDate
}

func (o *ModelOutput) GetShares() []string {
	if o == nil {
		return []string{}
	}
	return o.Shares
}

func (o *ModelOutput) GetPaused() bool {
	if o == nil {
		return false
	}
	return o.Paused
}

func (o *ModelOutput) GetOwner() User {
	if o == nil {
		return User{}
	}
	return o.Owner
}

func (o *ModelOutput) GetDependencies() []ModelDependency {
	if o == nil {
		return []ModelDependency{}
	}
	return o.Dependencies
}

func (o *ModelOutput) GetWarehouse() WarehouseTypes {
	if o == nil {
		return WarehouseTypes{}
	}
	return o.Warehouse
}

func (o *ModelOutput) GetWarehouseSnowflake() *WarehouseSnowflake {
	return o.GetWarehouse().WarehouseSnowflake
}

func (o *ModelOutput) GetWarehouseRedshift() *WarehouseRedshift {
	return o.GetWarehouse().WarehouseRedshift
}

func (o *ModelOutput) GetLastUpdateDuration() *int64 {
	if o == nil {
		return nil
	}
	return o.LastUpdateDuration
}

func (o *ModelOutput) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ModelOutput) GetUpdateSchedule() RefreshScheduleTypes {
	if o == nil {
		return RefreshScheduleTypes{}
	}
	return o.UpdateSchedule
}

func (o *ModelOutput) GetUpdateScheduleMonthly() *RefreshScheduleModeMonthly {
	return o.GetUpdateSchedule().RefreshScheduleModeMonthly
}

func (o *ModelOutput) GetUpdateScheduleHourly() *RefreshScheduleModeHourly {
	return o.GetUpdateSchedule().RefreshScheduleModeHourly
}

func (o *ModelOutput) GetUpdateScheduleNever() *RefreshScheduleModeNever {
	return o.GetUpdateSchedule().RefreshScheduleModeNever
}

func (o *ModelOutput) GetUpdateScheduleDaily() *RefreshScheduleModeDaily {
	return o.GetUpdateSchedule().RefreshScheduleModeDaily
}

func (o *ModelOutput) GetUpdateScheduleWeekly() *RefreshScheduleModeWeekly {
	return o.GetUpdateSchedule().RefreshScheduleModeWeekly
}
