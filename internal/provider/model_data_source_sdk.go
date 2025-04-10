// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *ModelDataSourceModel) RefreshFromSharedModelOutput(resp *shared.ModelOutput) {
	if resp.CreateDate != nil {
		r.CreateDate = types.StringValue(resp.CreateDate.Format(time.RFC3339Nano))
	} else {
		r.CreateDate = types.StringNull()
	}
	if len(r.Dependencies) > len(resp.Dependencies) {
		r.Dependencies = r.Dependencies[:len(resp.Dependencies)]
	}
	for dependenciesCount, dependenciesItem := range resp.Dependencies {
		var dependencies1 ModelDependency
		dependencies1.ID = types.StringValue(dependenciesItem.ID)
		dependencies1.Name = types.StringValue(dependenciesItem.Name)
		dependencies1.Type = types.StringValue(string(dependenciesItem.Type))
		if dependenciesCount+1 > len(r.Dependencies) {
			r.Dependencies = append(r.Dependencies, dependencies1)
		} else {
			r.Dependencies[dependenciesCount].ID = dependencies1.ID
			r.Dependencies[dependenciesCount].Name = dependencies1.Name
			r.Dependencies[dependenciesCount].Type = dependencies1.Type
		}
	}
	r.ID = types.StringValue(resp.ID)
	r.LastUpdateDuration = types.Int64PointerValue(resp.LastUpdateDuration)
	if resp.LastUpdateTime != nil {
		r.LastUpdateTime = types.StringValue(resp.LastUpdateTime.Format(time.RFC3339Nano))
	} else {
		r.LastUpdateTime = types.StringNull()
	}
	r.Name = types.StringValue(resp.Name)
	r.Owner.EmailAddress = types.StringValue(resp.Owner.EmailAddress)
	r.Owner.FirstName = types.StringValue(resp.Owner.FirstName)
	r.Owner.ID = types.StringValue(resp.Owner.ID)
	r.Owner.LastName = types.StringValue(resp.Owner.LastName)
	r.Paused = types.BoolValue(resp.Paused)
	r.QueryAndTriggers.Query = types.StringValue(resp.QueryAndTriggers.Query)
	r.QueryAndTriggers.Triggers = nil
	for _, v := range resp.QueryAndTriggers.Triggers {
		r.QueryAndTriggers.Triggers = append(r.QueryAndTriggers.Triggers, types.StringValue(v))
	}
	r.Shares = nil
	for _, v := range resp.Shares {
		r.Shares = append(r.Shares, types.StringValue(v))
	}
	if resp.UpdateSchedule.RefreshScheduleModeDaily != nil {
		r.UpdateSchedule.Daily = &UpdateScheduleModeDaily{}
		r.UpdateSchedule.Daily.HourOfDay = types.Int64Value(resp.UpdateSchedule.RefreshScheduleModeDaily.HourOfDay)
		r.UpdateSchedule.Daily.Mode = types.StringValue(string(resp.UpdateSchedule.RefreshScheduleModeDaily.Mode))
	}
	if resp.UpdateSchedule.RefreshScheduleModeHourly != nil {
		r.UpdateSchedule.Hourly = &UpdateScheduleModeHourly{}
		r.UpdateSchedule.Hourly.Mode = types.StringValue(string(resp.UpdateSchedule.RefreshScheduleModeHourly.Mode))
	}
	if resp.UpdateSchedule.RefreshScheduleModeMonthly != nil {
		r.UpdateSchedule.Monthly = &UpdateScheduleModeMonthly{}
		r.UpdateSchedule.Monthly.DayOfMonth = types.Int64Value(resp.UpdateSchedule.RefreshScheduleModeMonthly.DayOfMonth)
		r.UpdateSchedule.Monthly.HourOfDay = types.Int64Value(resp.UpdateSchedule.RefreshScheduleModeMonthly.HourOfDay)
		r.UpdateSchedule.Monthly.Mode = types.StringValue(string(resp.UpdateSchedule.RefreshScheduleModeMonthly.Mode))
	}
	if resp.UpdateSchedule.RefreshScheduleModeNever != nil {
		r.UpdateSchedule.Never = &RefreshScheduleModeNever{}
		r.UpdateSchedule.Never.Mode = types.StringValue(string(resp.UpdateSchedule.RefreshScheduleModeNever.Mode))
	}
	if resp.UpdateSchedule.RefreshScheduleModeWeekly != nil {
		r.UpdateSchedule.Weekly = &UpdateScheduleModeWeekly{}
		r.UpdateSchedule.Weekly.DayOfWeek = types.Int64Value(resp.UpdateSchedule.RefreshScheduleModeWeekly.DayOfWeek)
		r.UpdateSchedule.Weekly.HourOfDay = types.Int64Value(resp.UpdateSchedule.RefreshScheduleModeWeekly.HourOfDay)
		r.UpdateSchedule.Weekly.Mode = types.StringValue(string(resp.UpdateSchedule.RefreshScheduleModeWeekly.Mode))
	}
	if resp.Warehouse.WarehouseRedshift != nil {
		r.Warehouse.Redshift = &WarehouseRedshift{}
		r.Warehouse.Redshift.ConnectionID = types.StringValue(resp.Warehouse.WarehouseRedshift.ConnectionID)
		if resp.Warehouse.WarehouseRedshift.DistributionStyle.DistributionStyle1 != nil {
			if resp.Warehouse.WarehouseRedshift.DistributionStyle.DistributionStyle1 != nil {
				r.Warehouse.Redshift.DistributionStyle.One = types.StringValue(string(*resp.Warehouse.WarehouseRedshift.DistributionStyle.DistributionStyle1))
			} else {
				r.Warehouse.Redshift.DistributionStyle.One = types.StringNull()
			}
		}
		if resp.Warehouse.WarehouseRedshift.DistributionStyle.DistributionStyleKey != nil {
			r.Warehouse.Redshift.DistributionStyle.DistributionStyleKey = &DistributionStyleKey{}
			r.Warehouse.Redshift.DistributionStyle.DistributionStyleKey.Column = types.StringValue(resp.Warehouse.WarehouseRedshift.DistributionStyle.DistributionStyleKey.Column)
			r.Warehouse.Redshift.DistributionStyle.DistributionStyleKey.Type = types.StringValue(string(resp.Warehouse.WarehouseRedshift.DistributionStyle.DistributionStyleKey.Type))
		}
		r.Warehouse.Redshift.MaterializedView = types.BoolValue(resp.Warehouse.WarehouseRedshift.MaterializedView)
		r.Warehouse.Redshift.PendingRenamedTable = types.StringPointerValue(resp.Warehouse.WarehouseRedshift.PendingRenamedTable)
		r.Warehouse.Redshift.Schema = types.StringPointerValue(resp.Warehouse.WarehouseRedshift.Schema)
		r.Warehouse.Redshift.SortColumns = nil
		for _, v := range resp.Warehouse.WarehouseRedshift.SortColumns {
			r.Warehouse.Redshift.SortColumns = append(r.Warehouse.Redshift.SortColumns, types.StringValue(v))
		}
		r.Warehouse.Redshift.Table = types.StringValue(resp.Warehouse.WarehouseRedshift.Table)
		r.Warehouse.Redshift.Type = types.StringValue(string(resp.Warehouse.WarehouseRedshift.Type))
		r.Warehouse.Redshift.WaitForUpdatePreparation = types.BoolValue(resp.Warehouse.WarehouseRedshift.WaitForUpdatePreparation)
	}
	if resp.Warehouse.WarehouseSnowflake != nil {
		r.Warehouse.Snowflake = &WarehouseSnowflake{}
		r.Warehouse.Snowflake.ConnectionID = types.StringValue(resp.Warehouse.WarehouseSnowflake.ConnectionID)
		r.Warehouse.Snowflake.MaterializedView = types.BoolValue(resp.Warehouse.WarehouseSnowflake.MaterializedView)
		r.Warehouse.Snowflake.PendingRenamedTable = types.StringPointerValue(resp.Warehouse.WarehouseSnowflake.PendingRenamedTable)
		r.Warehouse.Snowflake.Schema = types.StringPointerValue(resp.Warehouse.WarehouseSnowflake.Schema)
		r.Warehouse.Snowflake.Table = types.StringValue(resp.Warehouse.WarehouseSnowflake.Table)
		r.Warehouse.Snowflake.Type = types.StringValue(string(resp.Warehouse.WarehouseSnowflake.Type))
		r.Warehouse.Snowflake.WaitForUpdatePreparation = types.BoolValue(resp.Warehouse.WarehouseSnowflake.WaitForUpdatePreparation)
	}
}
