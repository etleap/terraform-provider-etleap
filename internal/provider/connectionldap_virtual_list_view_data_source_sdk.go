// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *ConnectionLDAPVIRTUALLISTVIEWDataSourceModel) RefreshFromSharedConnectionLdapVirtualListView(resp *shared.ConnectionLdapVirtualListView) {
	r.Active = types.BoolValue(resp.Active)
	r.BaseDn = types.StringValue(resp.BaseDn)
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
	r.Filter = types.StringValue(resp.Filter)
	r.Hostname = types.StringValue(resp.Hostname)
	r.ID = types.StringValue(resp.ID)
	r.Name = types.StringValue(resp.Name)
	r.Port = types.Int64Value(resp.Port)
	r.Scope = types.StringValue(string(resp.Scope))
	r.SortOrder = types.StringValue(resp.SortOrder)
	r.Status = types.StringValue(string(resp.Status))
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
	r.User = types.StringValue(resp.User)
	r.UseSsl = types.BoolValue(resp.UseSsl)
}
