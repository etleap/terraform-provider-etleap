// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
)

// DbtScheduleInput - Request body for POST /dbtSchedules
type DbtScheduleInput struct {
	// The cron expression that defines triggers for this schedule. The maximum supported cron schedule precision is 1 minute.
	Cron string `json:"cron"`
	// The name of the dbt schedule.
	Name string `json:"name"`
	// Whether the dbt build is skipped if no new data has been ingested for any of the pipelines this schedule depends on.
	SkipBuildIfNoNewData *bool `default:"false" json:"skipBuildIfNoNewData"`
	// The [connection](https://docs.etleap.com/docs/api-v2/edbec13814bbc-connection) where the dbt build runs. The only supported connections are Redshift, Snowflake or Databricks Delta Lake destinations.
	ConnectionID string `json:"connectionId"`
	// The selector (from `selectors.yaml`) to run the build for.
	Selector string `json:"selector"`
	// `true` if the schedule should start as paused; defaults to `false`
	Paused *bool `default:"false" json:"paused"`
	// The target schema for the dbt build. See [here](https://docs.getdbt.com/docs/build/custom-schemas) for details on how it's used.
	TargetSchema string `json:"targetSchema"`
}

func (d DbtScheduleInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DbtScheduleInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DbtScheduleInput) GetCron() string {
	if o == nil {
		return ""
	}
	return o.Cron
}

func (o *DbtScheduleInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DbtScheduleInput) GetSkipBuildIfNoNewData() *bool {
	if o == nil {
		return nil
	}
	return o.SkipBuildIfNoNewData
}

func (o *DbtScheduleInput) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *DbtScheduleInput) GetSelector() string {
	if o == nil {
		return ""
	}
	return o.Selector
}

func (o *DbtScheduleInput) GetPaused() *bool {
	if o == nil {
		return nil
	}
	return o.Paused
}

func (o *DbtScheduleInput) GetTargetSchema() string {
	if o == nil {
		return ""
	}
	return o.TargetSchema
}
