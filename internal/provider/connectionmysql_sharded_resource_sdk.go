// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *ConnectionMYSQLSHARDEDResourceModel) ToSharedConnectionMysqlShardedInput() *shared.ConnectionMysqlShardedInput {
	name := r.Name.ValueString()
	typeVar := shared.ConnectionMysqlShardedType(r.Type.ValueString())
	var updateSchedule *shared.UpdateScheduleTypes
	if r.UpdateSchedule != nil {
		var updateScheduleModeInterval *shared.UpdateScheduleModeInterval
		if r.UpdateSchedule.Interval != nil {
			mode := shared.Mode(r.UpdateSchedule.Interval.Mode.ValueString())
			intervalMinutes := r.UpdateSchedule.Interval.IntervalMinutes.ValueInt64()
			updateScheduleModeInterval = &shared.UpdateScheduleModeInterval{
				Mode:            mode,
				IntervalMinutes: intervalMinutes,
			}
		}
		if updateScheduleModeInterval != nil {
			updateSchedule = &shared.UpdateScheduleTypes{
				UpdateScheduleModeInterval: updateScheduleModeInterval,
			}
		}
		var updateScheduleModeHourly *shared.UpdateScheduleModeHourly
		if r.UpdateSchedule.Hourly != nil {
			mode1 := shared.UpdateScheduleModeHourlyMode(r.UpdateSchedule.Hourly.Mode.ValueString())
			updateScheduleModeHourly = &shared.UpdateScheduleModeHourly{
				Mode: mode1,
			}
		}
		if updateScheduleModeHourly != nil {
			updateSchedule = &shared.UpdateScheduleTypes{
				UpdateScheduleModeHourly: updateScheduleModeHourly,
			}
		}
		var updateScheduleModeDaily *shared.UpdateScheduleModeDaily
		if r.UpdateSchedule.Daily != nil {
			mode2 := shared.UpdateScheduleModeDailyMode(r.UpdateSchedule.Daily.Mode.ValueString())
			hourOfDay := r.UpdateSchedule.Daily.HourOfDay.ValueInt64()
			updateScheduleModeDaily = &shared.UpdateScheduleModeDaily{
				Mode:      mode2,
				HourOfDay: hourOfDay,
			}
		}
		if updateScheduleModeDaily != nil {
			updateSchedule = &shared.UpdateScheduleTypes{
				UpdateScheduleModeDaily: updateScheduleModeDaily,
			}
		}
		var updateScheduleModeWeekly *shared.UpdateScheduleModeWeekly
		if r.UpdateSchedule.Weekly != nil {
			mode3 := shared.UpdateScheduleModeWeeklyMode(r.UpdateSchedule.Weekly.Mode.ValueString())
			dayOfWeek := r.UpdateSchedule.Weekly.DayOfWeek.ValueInt64()
			hourOfDay1 := r.UpdateSchedule.Weekly.HourOfDay.ValueInt64()
			updateScheduleModeWeekly = &shared.UpdateScheduleModeWeekly{
				Mode:      mode3,
				DayOfWeek: dayOfWeek,
				HourOfDay: hourOfDay1,
			}
		}
		if updateScheduleModeWeekly != nil {
			updateSchedule = &shared.UpdateScheduleTypes{
				UpdateScheduleModeWeekly: updateScheduleModeWeekly,
			}
		}
		var updateScheduleModeMonthly *shared.UpdateScheduleModeMonthly
		if r.UpdateSchedule.Monthly != nil {
			mode4 := shared.UpdateScheduleModeMonthlyMode(r.UpdateSchedule.Monthly.Mode.ValueString())
			dayOfMonth := r.UpdateSchedule.Monthly.DayOfMonth.ValueInt64()
			hourOfDay2 := r.UpdateSchedule.Monthly.HourOfDay.ValueInt64()
			updateScheduleModeMonthly = &shared.UpdateScheduleModeMonthly{
				Mode:       mode4,
				DayOfMonth: dayOfMonth,
				HourOfDay:  hourOfDay2,
			}
		}
		if updateScheduleModeMonthly != nil {
			updateSchedule = &shared.UpdateScheduleTypes{
				UpdateScheduleModeMonthly: updateScheduleModeMonthly,
			}
		}
	}
	validateSslCert := new(bool)
	if !r.ValidateSslCert.IsUnknown() && !r.ValidateSslCert.IsNull() {
		*validateSslCert = r.ValidateSslCert.ValueBool()
	} else {
		validateSslCert = nil
	}
	cdcEnabled := new(bool)
	if !r.CdcEnabled.IsUnknown() && !r.CdcEnabled.IsNull() {
		*cdcEnabled = r.CdcEnabled.ValueBool()
	} else {
		cdcEnabled = nil
	}
	autoReplicate := new(string)
	if !r.AutoReplicate.IsUnknown() && !r.AutoReplicate.IsNull() {
		*autoReplicate = r.AutoReplicate.ValueString()
	} else {
		autoReplicate = nil
	}
	tinyInt1IsBoolean := new(bool)
	if !r.TinyInt1IsBoolean.IsUnknown() && !r.TinyInt1IsBoolean.IsNull() {
		*tinyInt1IsBoolean = r.TinyInt1IsBoolean.ValueBool()
	} else {
		tinyInt1IsBoolean = nil
	}
	var shards []shared.DatabaseShard = nil
	for _, shardsItem := range r.Shards {
		shardID := shardsItem.ShardID.ValueString()
		address := shardsItem.Address.ValueString()
		port := shardsItem.Port.ValueInt64()
		database := shardsItem.Database.ValueString()
		username := shardsItem.Username.ValueString()
		password := shardsItem.Password.ValueString()
		var sshConfig *shared.SSHConfig
		if shardsItem.SSHConfig != nil {
			address1 := shardsItem.SSHConfig.Address.ValueString()
			username1 := shardsItem.SSHConfig.Username.ValueString()
			sshConfig = &shared.SSHConfig{
				Address:  address1,
				Username: username1,
			}
		}
		shards = append(shards, shared.DatabaseShard{
			ShardID:   shardID,
			Address:   address,
			Port:      port,
			Database:  database,
			Username:  username,
			Password:  password,
			SSHConfig: sshConfig,
		})
	}
	out := shared.ConnectionMysqlShardedInput{
		Name:              name,
		Type:              typeVar,
		UpdateSchedule:    updateSchedule,
		ValidateSslCert:   validateSslCert,
		CdcEnabled:        cdcEnabled,
		AutoReplicate:     autoReplicate,
		TinyInt1IsBoolean: tinyInt1IsBoolean,
		Shards:            shards,
	}
	return &out
}

func (r *ConnectionMYSQLSHARDEDResourceModel) RefreshFromSharedConnectionMysqlSharded(resp *shared.ConnectionMysqlSharded) {
	r.Active = types.BoolValue(resp.Active)
	r.AutoReplicate = types.StringPointerValue(resp.AutoReplicate)
	r.CdcEnabled = types.BoolPointerValue(resp.CdcEnabled)
	r.CreateDate = types.StringValue(resp.CreateDate.Format(time.RFC3339Nano))
	if len(r.DefaultUpdateSchedule) > len(resp.DefaultUpdateSchedule) {
		r.DefaultUpdateSchedule = r.DefaultUpdateSchedule[:len(resp.DefaultUpdateSchedule)]
	}
	for defaultUpdateScheduleCount, defaultUpdateScheduleItem := range resp.DefaultUpdateSchedule {
		var defaultUpdateSchedule1 DefaultUpdateSchedule
		if defaultUpdateScheduleItem.PipelineMode != nil {
			defaultUpdateSchedule1.PipelineMode = types.StringValue(string(*defaultUpdateScheduleItem.PipelineMode))
		} else {
			defaultUpdateSchedule1.PipelineMode = types.StringNull()
		}
		if defaultUpdateScheduleItem.UpdateSchedule == nil {
			defaultUpdateSchedule1.UpdateSchedule = nil
		} else {
			defaultUpdateSchedule1.UpdateSchedule = &UpdateScheduleTypes{}
			if defaultUpdateScheduleItem.UpdateSchedule.UpdateScheduleModeDaily != nil {
				defaultUpdateSchedule1.UpdateSchedule.Daily = &UpdateScheduleModeDaily{}
				defaultUpdateSchedule1.UpdateSchedule.Daily.HourOfDay = types.Int64Value(defaultUpdateScheduleItem.UpdateSchedule.UpdateScheduleModeDaily.HourOfDay)
				defaultUpdateSchedule1.UpdateSchedule.Daily.Mode = types.StringValue(string(defaultUpdateScheduleItem.UpdateSchedule.UpdateScheduleModeDaily.Mode))
			}
			if defaultUpdateScheduleItem.UpdateSchedule.UpdateScheduleModeHourly != nil {
				defaultUpdateSchedule1.UpdateSchedule.Hourly = &UpdateScheduleModeHourly{}
				defaultUpdateSchedule1.UpdateSchedule.Hourly.Mode = types.StringValue(string(defaultUpdateScheduleItem.UpdateSchedule.UpdateScheduleModeHourly.Mode))
			}
			if defaultUpdateScheduleItem.UpdateSchedule.UpdateScheduleModeInterval != nil {
				defaultUpdateSchedule1.UpdateSchedule.Interval = &UpdateScheduleModeInterval{}
				defaultUpdateSchedule1.UpdateSchedule.Interval.IntervalMinutes = types.Int64Value(defaultUpdateScheduleItem.UpdateSchedule.UpdateScheduleModeInterval.IntervalMinutes)
				defaultUpdateSchedule1.UpdateSchedule.Interval.Mode = types.StringValue(string(defaultUpdateScheduleItem.UpdateSchedule.UpdateScheduleModeInterval.Mode))
			}
			if defaultUpdateScheduleItem.UpdateSchedule.UpdateScheduleModeMonthly != nil {
				defaultUpdateSchedule1.UpdateSchedule.Monthly = &UpdateScheduleModeMonthly{}
				defaultUpdateSchedule1.UpdateSchedule.Monthly.DayOfMonth = types.Int64Value(defaultUpdateScheduleItem.UpdateSchedule.UpdateScheduleModeMonthly.DayOfMonth)
				defaultUpdateSchedule1.UpdateSchedule.Monthly.HourOfDay = types.Int64Value(defaultUpdateScheduleItem.UpdateSchedule.UpdateScheduleModeMonthly.HourOfDay)
				defaultUpdateSchedule1.UpdateSchedule.Monthly.Mode = types.StringValue(string(defaultUpdateScheduleItem.UpdateSchedule.UpdateScheduleModeMonthly.Mode))
			}
			if defaultUpdateScheduleItem.UpdateSchedule.UpdateScheduleModeWeekly != nil {
				defaultUpdateSchedule1.UpdateSchedule.Weekly = &UpdateScheduleModeWeekly{}
				defaultUpdateSchedule1.UpdateSchedule.Weekly.DayOfWeek = types.Int64Value(defaultUpdateScheduleItem.UpdateSchedule.UpdateScheduleModeWeekly.DayOfWeek)
				defaultUpdateSchedule1.UpdateSchedule.Weekly.HourOfDay = types.Int64Value(defaultUpdateScheduleItem.UpdateSchedule.UpdateScheduleModeWeekly.HourOfDay)
				defaultUpdateSchedule1.UpdateSchedule.Weekly.Mode = types.StringValue(string(defaultUpdateScheduleItem.UpdateSchedule.UpdateScheduleModeWeekly.Mode))
			}
		}
		if defaultUpdateScheduleCount+1 > len(r.DefaultUpdateSchedule) {
			r.DefaultUpdateSchedule = append(r.DefaultUpdateSchedule, defaultUpdateSchedule1)
		} else {
			r.DefaultUpdateSchedule[defaultUpdateScheduleCount].PipelineMode = defaultUpdateSchedule1.PipelineMode
			r.DefaultUpdateSchedule[defaultUpdateScheduleCount].UpdateSchedule = defaultUpdateSchedule1.UpdateSchedule
		}
	}
	r.ID = types.StringValue(resp.ID)
	r.Name = types.StringValue(resp.Name)
	if len(r.Shards) > len(resp.Shards) {
		r.Shards = r.Shards[:len(resp.Shards)]
	}
	for shardsCount, shardsItem := range resp.Shards {
		var shards1 DatabaseShard
		shards1.Address = types.StringValue(shardsItem.Address)
		shards1.Database = types.StringValue(shardsItem.Database)
		shards1.Port = types.Int64Value(shardsItem.Port)
		shards1.ShardID = types.StringValue(shardsItem.ShardID)
		if shardsItem.SSHConfig == nil {
			shards1.SSHConfig = nil
		} else {
			shards1.SSHConfig = &SSHConfig{}
			shards1.SSHConfig.Address = types.StringValue(shardsItem.SSHConfig.Address)
			shards1.SSHConfig.Username = types.StringValue(shardsItem.SSHConfig.Username)
		}
		shards1.Username = types.StringValue(shardsItem.Username)
		if shardsCount+1 > len(r.Shards) {
			r.Shards = append(r.Shards, shards1)
		} else {
			r.Shards[shardsCount].Address = shards1.Address
			r.Shards[shardsCount].Database = shards1.Database
			r.Shards[shardsCount].Port = shards1.Port
			r.Shards[shardsCount].ShardID = shards1.ShardID
			r.Shards[shardsCount].SSHConfig = shards1.SSHConfig
			r.Shards[shardsCount].Username = shards1.Username
		}
	}
	r.Status = types.StringValue(string(resp.Status))
	r.TinyInt1IsBoolean = types.BoolPointerValue(resp.TinyInt1IsBoolean)
	r.Type = types.StringValue(string(resp.Type))
	if resp.UpdateSchedule == nil {
		r.UpdateSchedule = nil
	} else {
		r.UpdateSchedule = &UpdateScheduleTypes{}
		if resp.UpdateSchedule.UpdateScheduleModeDaily != nil {
			r.UpdateSchedule.Daily = &UpdateScheduleModeDaily{}
			r.UpdateSchedule.Daily.HourOfDay = types.Int64Value(resp.UpdateSchedule.UpdateScheduleModeDaily.HourOfDay)
			r.UpdateSchedule.Daily.Mode = types.StringValue(string(resp.UpdateSchedule.UpdateScheduleModeDaily.Mode))
		}
		if resp.UpdateSchedule.UpdateScheduleModeHourly != nil {
			r.UpdateSchedule.Hourly = &UpdateScheduleModeHourly{}
			r.UpdateSchedule.Hourly.Mode = types.StringValue(string(resp.UpdateSchedule.UpdateScheduleModeHourly.Mode))
		}
		if resp.UpdateSchedule.UpdateScheduleModeInterval != nil {
			r.UpdateSchedule.Interval = &UpdateScheduleModeInterval{}
			r.UpdateSchedule.Interval.IntervalMinutes = types.Int64Value(resp.UpdateSchedule.UpdateScheduleModeInterval.IntervalMinutes)
			r.UpdateSchedule.Interval.Mode = types.StringValue(string(resp.UpdateSchedule.UpdateScheduleModeInterval.Mode))
		}
		if resp.UpdateSchedule.UpdateScheduleModeMonthly != nil {
			r.UpdateSchedule.Monthly = &UpdateScheduleModeMonthly{}
			r.UpdateSchedule.Monthly.DayOfMonth = types.Int64Value(resp.UpdateSchedule.UpdateScheduleModeMonthly.DayOfMonth)
			r.UpdateSchedule.Monthly.HourOfDay = types.Int64Value(resp.UpdateSchedule.UpdateScheduleModeMonthly.HourOfDay)
			r.UpdateSchedule.Monthly.Mode = types.StringValue(string(resp.UpdateSchedule.UpdateScheduleModeMonthly.Mode))
		}
		if resp.UpdateSchedule.UpdateScheduleModeWeekly != nil {
			r.UpdateSchedule.Weekly = &UpdateScheduleModeWeekly{}
			r.UpdateSchedule.Weekly.DayOfWeek = types.Int64Value(resp.UpdateSchedule.UpdateScheduleModeWeekly.DayOfWeek)
			r.UpdateSchedule.Weekly.HourOfDay = types.Int64Value(resp.UpdateSchedule.UpdateScheduleModeWeekly.HourOfDay)
			r.UpdateSchedule.Weekly.Mode = types.StringValue(string(resp.UpdateSchedule.UpdateScheduleModeWeekly.Mode))
		}
	}
	r.ValidateSslCert = types.BoolPointerValue(resp.ValidateSslCert)
}

func (r *ConnectionMYSQLSHARDEDResourceModel) ToSharedConnectionMysqlShardedUpdate() *shared.ConnectionMysqlShardedUpdate {
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = nil
	}
	typeVar := shared.ConnectionMysqlShardedUpdateType(r.Type.ValueString())
	active := new(bool)
	if !r.Active.IsUnknown() && !r.Active.IsNull() {
		*active = r.Active.ValueBool()
	} else {
		active = nil
	}
	var updateSchedule *shared.UpdateScheduleTypes
	if r.UpdateSchedule != nil {
		var updateScheduleModeInterval *shared.UpdateScheduleModeInterval
		if r.UpdateSchedule.Interval != nil {
			mode := shared.Mode(r.UpdateSchedule.Interval.Mode.ValueString())
			intervalMinutes := r.UpdateSchedule.Interval.IntervalMinutes.ValueInt64()
			updateScheduleModeInterval = &shared.UpdateScheduleModeInterval{
				Mode:            mode,
				IntervalMinutes: intervalMinutes,
			}
		}
		if updateScheduleModeInterval != nil {
			updateSchedule = &shared.UpdateScheduleTypes{
				UpdateScheduleModeInterval: updateScheduleModeInterval,
			}
		}
		var updateScheduleModeHourly *shared.UpdateScheduleModeHourly
		if r.UpdateSchedule.Hourly != nil {
			mode1 := shared.UpdateScheduleModeHourlyMode(r.UpdateSchedule.Hourly.Mode.ValueString())
			updateScheduleModeHourly = &shared.UpdateScheduleModeHourly{
				Mode: mode1,
			}
		}
		if updateScheduleModeHourly != nil {
			updateSchedule = &shared.UpdateScheduleTypes{
				UpdateScheduleModeHourly: updateScheduleModeHourly,
			}
		}
		var updateScheduleModeDaily *shared.UpdateScheduleModeDaily
		if r.UpdateSchedule.Daily != nil {
			mode2 := shared.UpdateScheduleModeDailyMode(r.UpdateSchedule.Daily.Mode.ValueString())
			hourOfDay := r.UpdateSchedule.Daily.HourOfDay.ValueInt64()
			updateScheduleModeDaily = &shared.UpdateScheduleModeDaily{
				Mode:      mode2,
				HourOfDay: hourOfDay,
			}
		}
		if updateScheduleModeDaily != nil {
			updateSchedule = &shared.UpdateScheduleTypes{
				UpdateScheduleModeDaily: updateScheduleModeDaily,
			}
		}
		var updateScheduleModeWeekly *shared.UpdateScheduleModeWeekly
		if r.UpdateSchedule.Weekly != nil {
			mode3 := shared.UpdateScheduleModeWeeklyMode(r.UpdateSchedule.Weekly.Mode.ValueString())
			dayOfWeek := r.UpdateSchedule.Weekly.DayOfWeek.ValueInt64()
			hourOfDay1 := r.UpdateSchedule.Weekly.HourOfDay.ValueInt64()
			updateScheduleModeWeekly = &shared.UpdateScheduleModeWeekly{
				Mode:      mode3,
				DayOfWeek: dayOfWeek,
				HourOfDay: hourOfDay1,
			}
		}
		if updateScheduleModeWeekly != nil {
			updateSchedule = &shared.UpdateScheduleTypes{
				UpdateScheduleModeWeekly: updateScheduleModeWeekly,
			}
		}
		var updateScheduleModeMonthly *shared.UpdateScheduleModeMonthly
		if r.UpdateSchedule.Monthly != nil {
			mode4 := shared.UpdateScheduleModeMonthlyMode(r.UpdateSchedule.Monthly.Mode.ValueString())
			dayOfMonth := r.UpdateSchedule.Monthly.DayOfMonth.ValueInt64()
			hourOfDay2 := r.UpdateSchedule.Monthly.HourOfDay.ValueInt64()
			updateScheduleModeMonthly = &shared.UpdateScheduleModeMonthly{
				Mode:       mode4,
				DayOfMonth: dayOfMonth,
				HourOfDay:  hourOfDay2,
			}
		}
		if updateScheduleModeMonthly != nil {
			updateSchedule = &shared.UpdateScheduleTypes{
				UpdateScheduleModeMonthly: updateScheduleModeMonthly,
			}
		}
	}
	validateSslCert := new(bool)
	if !r.ValidateSslCert.IsUnknown() && !r.ValidateSslCert.IsNull() {
		*validateSslCert = r.ValidateSslCert.ValueBool()
	} else {
		validateSslCert = nil
	}
	autoReplicate := new(string)
	if !r.AutoReplicate.IsUnknown() && !r.AutoReplicate.IsNull() {
		*autoReplicate = r.AutoReplicate.ValueString()
	} else {
		autoReplicate = nil
	}
	tinyInt1IsBoolean := new(bool)
	if !r.TinyInt1IsBoolean.IsUnknown() && !r.TinyInt1IsBoolean.IsNull() {
		*tinyInt1IsBoolean = r.TinyInt1IsBoolean.ValueBool()
	} else {
		tinyInt1IsBoolean = nil
	}
	var shards []shared.DatabaseShard = nil
	for _, shardsItem := range r.Shards {
		shardID := shardsItem.ShardID.ValueString()
		address := shardsItem.Address.ValueString()
		port := shardsItem.Port.ValueInt64()
		database := shardsItem.Database.ValueString()
		username := shardsItem.Username.ValueString()
		password := shardsItem.Password.ValueString()
		var sshConfig *shared.SSHConfig
		if shardsItem.SSHConfig != nil {
			address1 := shardsItem.SSHConfig.Address.ValueString()
			username1 := shardsItem.SSHConfig.Username.ValueString()
			sshConfig = &shared.SSHConfig{
				Address:  address1,
				Username: username1,
			}
		}
		shards = append(shards, shared.DatabaseShard{
			ShardID:   shardID,
			Address:   address,
			Port:      port,
			Database:  database,
			Username:  username,
			Password:  password,
			SSHConfig: sshConfig,
		})
	}
	out := shared.ConnectionMysqlShardedUpdate{
		Name:              name,
		Type:              typeVar,
		Active:            active,
		UpdateSchedule:    updateSchedule,
		ValidateSslCert:   validateSslCert,
		AutoReplicate:     autoReplicate,
		TinyInt1IsBoolean: tinyInt1IsBoolean,
		Shards:            shards,
	}
	return &out
}

func (r *ConnectionMYSQLSHARDEDResourceModel) ToSharedConnectionDelete() *shared.ConnectionDelete {
	deletionOfExportProducts := new(bool)
	if !r.DeletionOfExportProducts.IsUnknown() && !r.DeletionOfExportProducts.IsNull() {
		*deletionOfExportProducts = r.DeletionOfExportProducts.ValueBool()
	} else {
		deletionOfExportProducts = nil
	}
	out := shared.ConnectionDelete{
		DeletionOfExportProducts: deletionOfExportProducts,
	}
	return &out
}
