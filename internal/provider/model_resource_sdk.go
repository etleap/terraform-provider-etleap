// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *ModelResourceModel) ToSharedModelInput() *shared.ModelInput {
	name := r.Name.ValueString()
	var warehouse shared.WarehouseTypesInput
	var warehouseRedshiftInput *shared.WarehouseRedshiftInput
	if r.Warehouse.Redshift != nil {
		typeVar := shared.WarehouseRedshiftType(r.Warehouse.Redshift.Type.ValueString())
		connectionID := r.Warehouse.Redshift.ConnectionID.ValueString()
		schema := new(string)
		if !r.Warehouse.Redshift.Schema.IsUnknown() && !r.Warehouse.Redshift.Schema.IsNull() {
			*schema = r.Warehouse.Redshift.Schema.ValueString()
		} else {
			schema = nil
		}
		table := r.Warehouse.Redshift.Table.ValueString()
		materializedView := r.Warehouse.Redshift.MaterializedView.ValueBool()
		waitForUpdatePreparation := r.Warehouse.Redshift.WaitForUpdatePreparation.ValueBool()
		var sortColumns []string = nil
		for _, sortColumnsItem := range r.Warehouse.Redshift.SortColumns {
			sortColumns = append(sortColumns, sortColumnsItem.ValueString())
		}
		var distributionStyle shared.DistributionStyle
		one := new(shared.One)
		if !r.Warehouse.Redshift.DistributionStyle.One.IsUnknown() && !r.Warehouse.Redshift.DistributionStyle.One.IsNull() {
			*one = shared.One(r.Warehouse.Redshift.DistributionStyle.One.ValueString())
		} else {
			one = nil
		}
		if one != nil {
			distributionStyle = shared.DistributionStyle{
				One: one,
			}
		}
		var distributionStyleKey *shared.DistributionStyleKey
		if r.Warehouse.Redshift.DistributionStyle.DistributionStyleKey != nil {
			typeVar1 := shared.DistributionStyleKeyType(r.Warehouse.Redshift.DistributionStyle.DistributionStyleKey.Type.ValueString())
			column := r.Warehouse.Redshift.DistributionStyle.DistributionStyleKey.Column.ValueString()
			distributionStyleKey = &shared.DistributionStyleKey{
				Type:   typeVar1,
				Column: column,
			}
		}
		if distributionStyleKey != nil {
			distributionStyle = shared.DistributionStyle{
				DistributionStyleKey: distributionStyleKey,
			}
		}
		warehouseRedshiftInput = &shared.WarehouseRedshiftInput{
			Type:                     typeVar,
			ConnectionID:             connectionID,
			Schema:                   schema,
			Table:                    table,
			MaterializedView:         materializedView,
			WaitForUpdatePreparation: waitForUpdatePreparation,
			SortColumns:              sortColumns,
			DistributionStyle:        distributionStyle,
		}
	}
	if warehouseRedshiftInput != nil {
		warehouse = shared.WarehouseTypesInput{
			WarehouseRedshiftInput: warehouseRedshiftInput,
		}
	}
	var warehouseSnowflakeInput *shared.WarehouseSnowflakeInput
	if r.Warehouse.Snowflake != nil {
		typeVar2 := shared.WarehouseSnowflakeType(r.Warehouse.Snowflake.Type.ValueString())
		connectionId1 := r.Warehouse.Snowflake.ConnectionID.ValueString()
		schema1 := new(string)
		if !r.Warehouse.Snowflake.Schema.IsUnknown() && !r.Warehouse.Snowflake.Schema.IsNull() {
			*schema1 = r.Warehouse.Snowflake.Schema.ValueString()
		} else {
			schema1 = nil
		}
		table1 := r.Warehouse.Snowflake.Table.ValueString()
		materializedView1 := r.Warehouse.Snowflake.MaterializedView.ValueBool()
		waitForUpdatePreparation1 := r.Warehouse.Snowflake.WaitForUpdatePreparation.ValueBool()
		warehouseSnowflakeInput = &shared.WarehouseSnowflakeInput{
			Type:                     typeVar2,
			ConnectionID:             connectionId1,
			Schema:                   schema1,
			Table:                    table1,
			MaterializedView:         materializedView1,
			WaitForUpdatePreparation: waitForUpdatePreparation1,
		}
	}
	if warehouseSnowflakeInput != nil {
		warehouse = shared.WarehouseTypesInput{
			WarehouseSnowflakeInput: warehouseSnowflakeInput,
		}
	}
	query := r.QueryAndTriggers.Query.ValueString()
	var triggers []string = nil
	for _, triggersItem := range r.QueryAndTriggers.Triggers {
		triggers = append(triggers, triggersItem.ValueString())
	}
	queryAndTriggers := shared.QueryAndTriggers{
		Query:    query,
		Triggers: triggers,
	}
	var updateSchedule shared.RefreshScheduleTypes
	var refreshScheduleModeNever *shared.RefreshScheduleModeNever
	if r.UpdateSchedule.Never != nil {
		mode := shared.RefreshScheduleModeNeverMode(r.UpdateSchedule.Never.Mode.ValueString())
		refreshScheduleModeNever = &shared.RefreshScheduleModeNever{
			Mode: mode,
		}
	}
	if refreshScheduleModeNever != nil {
		updateSchedule = shared.RefreshScheduleTypes{
			RefreshScheduleModeNever: refreshScheduleModeNever,
		}
	}
	var refreshScheduleModeHourly *shared.RefreshScheduleModeHourly
	if r.UpdateSchedule.Hourly != nil {
		mode1 := shared.RefreshScheduleModeHourlyMode(r.UpdateSchedule.Hourly.Mode.ValueString())
		refreshScheduleModeHourly = &shared.RefreshScheduleModeHourly{
			Mode: mode1,
		}
	}
	if refreshScheduleModeHourly != nil {
		updateSchedule = shared.RefreshScheduleTypes{
			RefreshScheduleModeHourly: refreshScheduleModeHourly,
		}
	}
	var refreshScheduleModeDaily *shared.RefreshScheduleModeDaily
	if r.UpdateSchedule.Daily != nil {
		mode2 := shared.RefreshScheduleModeDailyMode(r.UpdateSchedule.Daily.Mode.ValueString())
		hourOfDay := r.UpdateSchedule.Daily.HourOfDay.ValueInt64()
		refreshScheduleModeDaily = &shared.RefreshScheduleModeDaily{
			Mode:      mode2,
			HourOfDay: hourOfDay,
		}
	}
	if refreshScheduleModeDaily != nil {
		updateSchedule = shared.RefreshScheduleTypes{
			RefreshScheduleModeDaily: refreshScheduleModeDaily,
		}
	}
	var refreshScheduleModeWeekly *shared.RefreshScheduleModeWeekly
	if r.UpdateSchedule.Weekly != nil {
		mode3 := shared.RefreshScheduleModeWeeklyMode(r.UpdateSchedule.Weekly.Mode.ValueString())
		dayOfWeek := r.UpdateSchedule.Weekly.DayOfWeek.ValueInt64()
		hourOfDay1 := r.UpdateSchedule.Weekly.HourOfDay.ValueInt64()
		refreshScheduleModeWeekly = &shared.RefreshScheduleModeWeekly{
			Mode:      mode3,
			DayOfWeek: dayOfWeek,
			HourOfDay: hourOfDay1,
		}
	}
	if refreshScheduleModeWeekly != nil {
		updateSchedule = shared.RefreshScheduleTypes{
			RefreshScheduleModeWeekly: refreshScheduleModeWeekly,
		}
	}
	var refreshScheduleModeMonthly *shared.RefreshScheduleModeMonthly
	if r.UpdateSchedule.Monthly != nil {
		mode4 := shared.RefreshScheduleModeMonthlyMode(r.UpdateSchedule.Monthly.Mode.ValueString())
		dayOfMonth := r.UpdateSchedule.Monthly.DayOfMonth.ValueInt64()
		hourOfDay2 := r.UpdateSchedule.Monthly.HourOfDay.ValueInt64()
		refreshScheduleModeMonthly = &shared.RefreshScheduleModeMonthly{
			Mode:       mode4,
			DayOfMonth: dayOfMonth,
			HourOfDay:  hourOfDay2,
		}
	}
	if refreshScheduleModeMonthly != nil {
		updateSchedule = shared.RefreshScheduleTypes{
			RefreshScheduleModeMonthly: refreshScheduleModeMonthly,
		}
	}
	var shares []string = nil
	for _, sharesItem := range r.Shares {
		shares = append(shares, sharesItem.ValueString())
	}
	out := shared.ModelInput{
		Name:             name,
		Warehouse:        warehouse,
		QueryAndTriggers: queryAndTriggers,
		UpdateSchedule:   updateSchedule,
		Shares:           shares,
	}
	return &out
}

func (r *ModelResourceModel) RefreshFromSharedModelOutput(resp *shared.ModelOutput) {
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
		if resp.Warehouse.WarehouseRedshift.DistributionStyle.One != nil {
			if resp.Warehouse.WarehouseRedshift.DistributionStyle.One != nil {
				r.Warehouse.Redshift.DistributionStyle.One = types.StringValue(string(*resp.Warehouse.WarehouseRedshift.DistributionStyle.One))
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

func (r *ModelResourceModel) ToSharedModelUpdate() *shared.ModelUpdate {
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = nil
	}
	paused := new(bool)
	if !r.Paused.IsUnknown() && !r.Paused.IsNull() {
		*paused = r.Paused.ValueBool()
	} else {
		paused = nil
	}
	var warehouse *shared.WarehouseUpdateTypes
	var warehouseRedshiftUpdate *shared.WarehouseRedshiftUpdate
	if r.Warehouse.Redshift != nil {
		typeVar := shared.WarehouseRedshiftUpdateType(r.Warehouse.Redshift.Type.ValueString())
		table := new(string)
		if !r.Warehouse.Redshift.Table.IsUnknown() && !r.Warehouse.Redshift.Table.IsNull() {
			*table = r.Warehouse.Redshift.Table.ValueString()
		} else {
			table = nil
		}
		var sortColumns []string = nil
		for _, sortColumnsItem := range r.Warehouse.Redshift.SortColumns {
			sortColumns = append(sortColumns, sortColumnsItem.ValueString())
		}
		var distributionStyle *shared.DistributionStyle
		one := new(shared.One)
		if !r.Warehouse.Redshift.DistributionStyle.One.IsUnknown() && !r.Warehouse.Redshift.DistributionStyle.One.IsNull() {
			*one = shared.One(r.Warehouse.Redshift.DistributionStyle.One.ValueString())
		} else {
			one = nil
		}
		if one != nil {
			distributionStyle = &shared.DistributionStyle{
				One: one,
			}
		}
		var distributionStyleKey *shared.DistributionStyleKey
		if r.Warehouse.Redshift.DistributionStyle.DistributionStyleKey != nil {
			typeVar1 := shared.DistributionStyleKeyType(r.Warehouse.Redshift.DistributionStyle.DistributionStyleKey.Type.ValueString())
			column := r.Warehouse.Redshift.DistributionStyle.DistributionStyleKey.Column.ValueString()
			distributionStyleKey = &shared.DistributionStyleKey{
				Type:   typeVar1,
				Column: column,
			}
		}
		if distributionStyleKey != nil {
			distributionStyle = &shared.DistributionStyle{
				DistributionStyleKey: distributionStyleKey,
			}
		}
		warehouseRedshiftUpdate = &shared.WarehouseRedshiftUpdate{
			Type:              typeVar,
			Table:             table,
			SortColumns:       sortColumns,
			DistributionStyle: distributionStyle,
		}
	}
	if warehouseRedshiftUpdate != nil {
		warehouse = &shared.WarehouseUpdateTypes{
			WarehouseRedshiftUpdate: warehouseRedshiftUpdate,
		}
	}
	var warehouseSnowflakeUpdate *shared.WarehouseSnowflakeUpdate
	if r.Warehouse.Snowflake != nil {
		typeVar2 := shared.WarehouseSnowflakeUpdateType(r.Warehouse.Snowflake.Type.ValueString())
		table1 := new(string)
		if !r.Warehouse.Snowflake.Table.IsUnknown() && !r.Warehouse.Snowflake.Table.IsNull() {
			*table1 = r.Warehouse.Snowflake.Table.ValueString()
		} else {
			table1 = nil
		}
		warehouseSnowflakeUpdate = &shared.WarehouseSnowflakeUpdate{
			Type:  typeVar2,
			Table: table1,
		}
	}
	if warehouseSnowflakeUpdate != nil {
		warehouse = &shared.WarehouseUpdateTypes{
			WarehouseSnowflakeUpdate: warehouseSnowflakeUpdate,
		}
	}
	var queryAndTriggers *shared.ModelQueryAndTriggers
	query := r.QueryAndTriggers.Query.ValueString()
	var triggers []string = nil
	for _, triggersItem := range r.QueryAndTriggers.Triggers {
		triggers = append(triggers, triggersItem.ValueString())
	}
	queryAndTriggers = &shared.ModelQueryAndTriggers{
		Query:    query,
		Triggers: triggers,
	}
	var updateSchedule *shared.ModelUpdateScheduleTypes
	var scheduleTypesNeverScheduleMode *shared.ScheduleTypesNeverScheduleMode
	if r.UpdateSchedule.Never != nil {
		mode := shared.RefreshScheduleModeNeverScheduleTypesMode(r.UpdateSchedule.Never.Mode.ValueString())
		scheduleTypesNeverScheduleMode = &shared.ScheduleTypesNeverScheduleMode{
			Mode: mode,
		}
	}
	if scheduleTypesNeverScheduleMode != nil {
		updateSchedule = &shared.ModelUpdateScheduleTypes{
			ScheduleTypesNeverScheduleMode: scheduleTypesNeverScheduleMode,
		}
	}
	var scheduleTypesHourlyScheduleMode *shared.ScheduleTypesHourlyScheduleMode
	if r.UpdateSchedule.Hourly != nil {
		mode1 := shared.RefreshScheduleModeHourlyScheduleTypesModelUpdateMode(r.UpdateSchedule.Hourly.Mode.ValueString())
		scheduleTypesHourlyScheduleMode = &shared.ScheduleTypesHourlyScheduleMode{
			Mode: mode1,
		}
	}
	if scheduleTypesHourlyScheduleMode != nil {
		updateSchedule = &shared.ModelUpdateScheduleTypes{
			ScheduleTypesHourlyScheduleMode: scheduleTypesHourlyScheduleMode,
		}
	}
	var scheduleTypesDailyScheduleMode *shared.ScheduleTypesDailyScheduleMode
	if r.UpdateSchedule.Daily != nil {
		mode2 := shared.RefreshScheduleModeDailyScheduleTypesModelUpdateMode(r.UpdateSchedule.Daily.Mode.ValueString())
		hourOfDay := r.UpdateSchedule.Daily.HourOfDay.ValueInt64()
		scheduleTypesDailyScheduleMode = &shared.ScheduleTypesDailyScheduleMode{
			Mode:      mode2,
			HourOfDay: hourOfDay,
		}
	}
	if scheduleTypesDailyScheduleMode != nil {
		updateSchedule = &shared.ModelUpdateScheduleTypes{
			ScheduleTypesDailyScheduleMode: scheduleTypesDailyScheduleMode,
		}
	}
	var scheduleTypesWeeklyScheduleMode *shared.ScheduleTypesWeeklyScheduleMode
	if r.UpdateSchedule.Weekly != nil {
		mode3 := shared.RefreshScheduleModeWeeklyScheduleTypesModelUpdateMode(r.UpdateSchedule.Weekly.Mode.ValueString())
		dayOfWeek := r.UpdateSchedule.Weekly.DayOfWeek.ValueInt64()
		hourOfDay1 := r.UpdateSchedule.Weekly.HourOfDay.ValueInt64()
		scheduleTypesWeeklyScheduleMode = &shared.ScheduleTypesWeeklyScheduleMode{
			Mode:      mode3,
			DayOfWeek: dayOfWeek,
			HourOfDay: hourOfDay1,
		}
	}
	if scheduleTypesWeeklyScheduleMode != nil {
		updateSchedule = &shared.ModelUpdateScheduleTypes{
			ScheduleTypesWeeklyScheduleMode: scheduleTypesWeeklyScheduleMode,
		}
	}
	var scheduleTypesMonthlyScheduleMode *shared.ScheduleTypesMonthlyScheduleMode
	if r.UpdateSchedule.Monthly != nil {
		mode4 := shared.RefreshScheduleModeMonthlyScheduleTypesModelUpdateMode(r.UpdateSchedule.Monthly.Mode.ValueString())
		dayOfMonth := r.UpdateSchedule.Monthly.DayOfMonth.ValueInt64()
		hourOfDay2 := r.UpdateSchedule.Monthly.HourOfDay.ValueInt64()
		scheduleTypesMonthlyScheduleMode = &shared.ScheduleTypesMonthlyScheduleMode{
			Mode:       mode4,
			DayOfMonth: dayOfMonth,
			HourOfDay:  hourOfDay2,
		}
	}
	if scheduleTypesMonthlyScheduleMode != nil {
		updateSchedule = &shared.ModelUpdateScheduleTypes{
			ScheduleTypesMonthlyScheduleMode: scheduleTypesMonthlyScheduleMode,
		}
	}
	var shares []string = nil
	for _, sharesItem := range r.Shares {
		shares = append(shares, sharesItem.ValueString())
	}
	out := shared.ModelUpdate{
		Name:             name,
		Paused:           paused,
		Warehouse:        warehouse,
		QueryAndTriggers: queryAndTriggers,
		UpdateSchedule:   updateSchedule,
		Shares:           shares,
	}
	return &out
}

func (r *ModelResourceModel) ToSharedModelDelete() *shared.ModelDelete {
	deletionOfExportProducts := new(bool)
	if !r.DeletionOfExportProducts.IsUnknown() && !r.DeletionOfExportProducts.IsNull() {
		*deletionOfExportProducts = r.DeletionOfExportProducts.ValueBool()
	} else {
		deletionOfExportProducts = nil
	}
	out := shared.ModelDelete{
		DeletionOfExportProducts: deletionOfExportProducts,
	}
	return &out
}
