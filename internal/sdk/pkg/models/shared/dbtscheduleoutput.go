// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
	"time"
)

// CurrentActivity - This field is deprecated and will be removed and replaced by the properties in `latestRun` when that field is implemented.
//
// Deprecated type: This will be removed in a future release, please migrate away from it as soon as possible.
type CurrentActivity string

const (
	CurrentActivityLoading  CurrentActivity = "LOADING"
	CurrentActivityBuilding CurrentActivity = "BUILDING"
)

func (e CurrentActivity) ToPointer() *CurrentActivity {
	return &e
}

func (e *CurrentActivity) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "LOADING":
		fallthrough
	case "BUILDING":
		*e = CurrentActivity(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CurrentActivity: %v", v)
	}
}

// DbtScheduleOutput - Response body for GET /dbtSchedules/{id}
type DbtScheduleOutput struct {
	// The id of the dbt schedule.
	ID string `json:"id"`
	// The name of the dbt schedule.
	Name string `json:"name"`
	// `true` if the schedule is paused.
	Paused bool `json:"paused"`
	// The selector this schedule runs.
	Selector string `json:"selector"`
	// The target schema for the dbt build. See [here](https://docs.getdbt.com/docs/build/custom-schemas) for details on how it's used.
	TargetSchema string `json:"targetSchema"`
	Owner        User   `json:"owner"`
	// The cron expression that defines triggers for this schedule. The maximum supported cron schedule precision is 1 minute.
	Cron string `json:"cron"`
	// The [connection](https://docs.etleap.com/docs/api-v2/edbec13814bbc-connection) where the dbt build runs. The only supported connections are Redshift, Snowflake or Databricks Delta Lake destinations.
	ConnectionID string `json:"connectionId"`
	// Whether the dbt build is skipped if no new data has been ingested for any of the pipelines in the table above.
	SkipBuildIfNoNewData bool `json:"skipBuildIfNoNewData"`
	// This field is deprecated and will be removed and replaced by the properties in `latestRun` when that field is implemented.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	CurrentActivity *CurrentActivity `json:"currentActivity,omitempty"`
	// The last time that a successful dbt build started. This field is deprecated and will be removed and replaced by the properties in `latestRun` when that field is implemented.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	LastDbtBuildDate *time.Time `json:"lastDbtBuildDate,omitempty"`
	// The duration of the last successful dbt build. This field is deprecated and will be removed and replaced by the properties in `latestRun` when that field is implemented.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	LastDbtRunTime *int64               `json:"lastDbtRunTime,omitempty"`
	CreateDate     time.Time            `json:"createDate"`
	LatestRun      *DbtScheduleRunTypes `json:"latestRun,omitempty"`
}

func (d DbtScheduleOutput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DbtScheduleOutput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DbtScheduleOutput) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DbtScheduleOutput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DbtScheduleOutput) GetPaused() bool {
	if o == nil {
		return false
	}
	return o.Paused
}

func (o *DbtScheduleOutput) GetSelector() string {
	if o == nil {
		return ""
	}
	return o.Selector
}

func (o *DbtScheduleOutput) GetTargetSchema() string {
	if o == nil {
		return ""
	}
	return o.TargetSchema
}

func (o *DbtScheduleOutput) GetOwner() User {
	if o == nil {
		return User{}
	}
	return o.Owner
}

func (o *DbtScheduleOutput) GetCron() string {
	if o == nil {
		return ""
	}
	return o.Cron
}

func (o *DbtScheduleOutput) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *DbtScheduleOutput) GetSkipBuildIfNoNewData() bool {
	if o == nil {
		return false
	}
	return o.SkipBuildIfNoNewData
}

func (o *DbtScheduleOutput) GetCurrentActivity() *CurrentActivity {
	if o == nil {
		return nil
	}
	return o.CurrentActivity
}

func (o *DbtScheduleOutput) GetLastDbtBuildDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.LastDbtBuildDate
}

func (o *DbtScheduleOutput) GetLastDbtRunTime() *int64 {
	if o == nil {
		return nil
	}
	return o.LastDbtRunTime
}

func (o *DbtScheduleOutput) GetCreateDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreateDate
}

func (o *DbtScheduleOutput) GetLatestRun() *DbtScheduleRunTypes {
	if o == nil {
		return nil
	}
	return o.LatestRun
}

func (o *DbtScheduleOutput) GetLatestRunInProgress() *DbtScheduleRunInProgress {
	if v := o.GetLatestRun(); v != nil {
		return v.DbtScheduleRunInProgress
	}
	return nil
}

func (o *DbtScheduleOutput) GetLatestRunIngestCouldNotComplete() *DbtScheduleRunFailure {
	if v := o.GetLatestRun(); v != nil {
		return v.DbtScheduleRunFailure
	}
	return nil
}

func (o *DbtScheduleOutput) GetLatestRunDbtError() *DbtScheduleRunFailure {
	if v := o.GetLatestRun(); v != nil {
		return v.DbtScheduleRunFailure
	}
	return nil
}

func (o *DbtScheduleOutput) GetLatestRunSuccessWithDbtWarnings() *DbtScheduleRunSuccess {
	if v := o.GetLatestRun(); v != nil {
		return v.DbtScheduleRunSuccess
	}
	return nil
}

func (o *DbtScheduleOutput) GetLatestRunSuccess() *DbtScheduleRunSuccess {
	if v := o.GetLatestRun(); v != nil {
		return v.DbtScheduleRunSuccess
	}
	return nil
}
