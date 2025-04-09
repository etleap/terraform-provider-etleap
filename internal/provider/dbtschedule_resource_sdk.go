// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *DbtScheduleResourceModel) ToSharedDbtScheduleInput() *shared.DbtScheduleInput {
	cron := r.Cron.ValueString()
	name := r.Name.ValueString()
	skipBuildIfNoNewData := new(bool)
	if !r.SkipBuildIfNoNewData.IsUnknown() && !r.SkipBuildIfNoNewData.IsNull() {
		*skipBuildIfNoNewData = r.SkipBuildIfNoNewData.ValueBool()
	} else {
		skipBuildIfNoNewData = nil
	}
	connectionID := r.ConnectionID.ValueString()
	selector := r.Selector.ValueString()
	paused := new(bool)
	if !r.Paused.IsUnknown() && !r.Paused.IsNull() {
		*paused = r.Paused.ValueBool()
	} else {
		paused = nil
	}
	targetSchema := r.TargetSchema.ValueString()
	out := shared.DbtScheduleInput{
		Cron:                 cron,
		Name:                 name,
		SkipBuildIfNoNewData: skipBuildIfNoNewData,
		ConnectionID:         connectionID,
		Selector:             selector,
		Paused:               paused,
		TargetSchema:         targetSchema,
	}
	return &out
}

func (r *DbtScheduleResourceModel) RefreshFromSharedDbtScheduleOutput(resp *shared.DbtScheduleOutput) {
	r.ConnectionID = types.StringValue(resp.ConnectionID)
	r.CreateDate = types.StringValue(resp.CreateDate.Format(time.RFC3339Nano))
	r.Cron = types.StringValue(resp.Cron)
	r.ID = types.StringValue(resp.ID)
	if resp.LatestRun == nil {
		r.LatestRun = nil
	} else {
		r.LatestRun = &DbtScheduleRunTypes{}
		if resp.LatestRun.DbtScheduleRunFailure != nil {
			r.LatestRun.IngestCouldNotComplete = &DbtScheduleRunFailure{}
			r.LatestRun.IngestCouldNotComplete.Duration = types.Int64Value(resp.LatestRun.DbtScheduleRunFailure.Duration)
			if resp.LatestRun.DbtScheduleRunFailure.LastSuccessfulDbtBuildDate != nil {
				r.LatestRun.IngestCouldNotComplete.LastSuccessfulDbtBuildDate = types.StringValue(resp.LatestRun.DbtScheduleRunFailure.LastSuccessfulDbtBuildDate.Format(time.RFC3339Nano))
			} else {
				r.LatestRun.IngestCouldNotComplete.LastSuccessfulDbtBuildDate = types.StringNull()
			}
			r.LatestRun.IngestCouldNotComplete.NextTriggerDate = types.StringValue(resp.LatestRun.DbtScheduleRunFailure.NextTriggerDate.Format(time.RFC3339Nano))
			r.LatestRun.IngestCouldNotComplete.StartDate = types.StringValue(resp.LatestRun.DbtScheduleRunFailure.StartDate.Format(time.RFC3339Nano))
			r.LatestRun.IngestCouldNotComplete.Status = types.StringValue(string(resp.LatestRun.DbtScheduleRunFailure.Status))
		}
		if resp.LatestRun.DbtScheduleRunInProgress != nil {
			r.LatestRun.InProgress = &DbtScheduleRunInProgress{}
			r.LatestRun.InProgress.BuildIsTakingTooLong = types.BoolValue(resp.LatestRun.DbtScheduleRunInProgress.BuildIsTakingTooLong)
			if resp.LatestRun.DbtScheduleRunInProgress.LastSuccessfulDbtBuildDate != nil {
				r.LatestRun.InProgress.LastSuccessfulDbtBuildDate = types.StringValue(resp.LatestRun.DbtScheduleRunInProgress.LastSuccessfulDbtBuildDate.Format(time.RFC3339Nano))
			} else {
				r.LatestRun.InProgress.LastSuccessfulDbtBuildDate = types.StringNull()
			}
			r.LatestRun.InProgress.Phase = types.StringValue(string(resp.LatestRun.DbtScheduleRunInProgress.Phase))
			r.LatestRun.InProgress.PreviousRunDuration = types.Int64PointerValue(resp.LatestRun.DbtScheduleRunInProgress.PreviousRunDuration)
			if resp.LatestRun.DbtScheduleRunInProgress.PreviousRunStatus != nil {
				r.LatestRun.InProgress.PreviousRunStatus = types.StringValue(string(*resp.LatestRun.DbtScheduleRunInProgress.PreviousRunStatus))
			} else {
				r.LatestRun.InProgress.PreviousRunStatus = types.StringNull()
			}
			r.LatestRun.InProgress.StartDate = types.StringValue(resp.LatestRun.DbtScheduleRunInProgress.StartDate.Format(time.RFC3339Nano))
			r.LatestRun.InProgress.Status = types.StringValue(string(resp.LatestRun.DbtScheduleRunInProgress.Status))
		}
		if resp.LatestRun.DbtScheduleRunSuccess != nil {
			r.LatestRun.SuccessWithDbtWarnings = &DbtScheduleRunSuccess{}
			r.LatestRun.SuccessWithDbtWarnings.Duration = types.Int64Value(resp.LatestRun.DbtScheduleRunSuccess.Duration)
			r.LatestRun.SuccessWithDbtWarnings.LastSuccessfulDbtBuildDate = types.StringValue(resp.LatestRun.DbtScheduleRunSuccess.LastSuccessfulDbtBuildDate.Format(time.RFC3339Nano))
			r.LatestRun.SuccessWithDbtWarnings.NextTriggerDate = types.StringValue(resp.LatestRun.DbtScheduleRunSuccess.NextTriggerDate.Format(time.RFC3339Nano))
			r.LatestRun.SuccessWithDbtWarnings.StartDate = types.StringValue(resp.LatestRun.DbtScheduleRunSuccess.StartDate.Format(time.RFC3339Nano))
			r.LatestRun.SuccessWithDbtWarnings.Status = types.StringValue(string(resp.LatestRun.DbtScheduleRunSuccess.Status))
		}
	}
	r.Name = types.StringValue(resp.Name)
	r.Owner.EmailAddress = types.StringValue(resp.Owner.EmailAddress)
	r.Owner.FirstName = types.StringValue(resp.Owner.FirstName)
	r.Owner.ID = types.StringValue(resp.Owner.ID)
	r.Owner.LastName = types.StringValue(resp.Owner.LastName)
	r.Paused = types.BoolValue(resp.Paused)
	r.Selector = types.StringValue(resp.Selector)
	r.SkipBuildIfNoNewData = types.BoolValue(resp.SkipBuildIfNoNewData)
	r.TargetSchema = types.StringValue(resp.TargetSchema)
}

func (r *DbtScheduleResourceModel) ToSharedDbtScheduleUpdate() *shared.DbtScheduleUpdate {
	paused := new(bool)
	if !r.Paused.IsUnknown() && !r.Paused.IsNull() {
		*paused = r.Paused.ValueBool()
	} else {
		paused = nil
	}
	cron := new(string)
	if !r.Cron.IsUnknown() && !r.Cron.IsNull() {
		*cron = r.Cron.ValueString()
	} else {
		cron = nil
	}
	targetSchema := new(string)
	if !r.TargetSchema.IsUnknown() && !r.TargetSchema.IsNull() {
		*targetSchema = r.TargetSchema.ValueString()
	} else {
		targetSchema = nil
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = nil
	}
	skipBuildIfNoNewData := new(bool)
	if !r.SkipBuildIfNoNewData.IsUnknown() && !r.SkipBuildIfNoNewData.IsNull() {
		*skipBuildIfNoNewData = r.SkipBuildIfNoNewData.ValueBool()
	} else {
		skipBuildIfNoNewData = nil
	}
	out := shared.DbtScheduleUpdate{
		Paused:               paused,
		Cron:                 cron,
		TargetSchema:         targetSchema,
		Name:                 name,
		SkipBuildIfNoNewData: skipBuildIfNoNewData,
	}
	return &out
}
