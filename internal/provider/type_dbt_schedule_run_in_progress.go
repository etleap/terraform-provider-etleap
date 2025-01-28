// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type DbtScheduleRunInProgress struct {
	BuildIsTakingTooLong       types.Bool   `tfsdk:"build_is_taking_too_long"`
	LastSuccessfulDbtBuildDate types.String `tfsdk:"last_successful_dbt_build_date"`
	Phase                      types.String `tfsdk:"phase"`
	PreviousRunDuration        types.Int64  `tfsdk:"previous_run_duration"`
	PreviousRunStatus          types.String `tfsdk:"previous_run_status"`
	StartDate                  types.String `tfsdk:"start_date"`
	Status                     types.String `tfsdk:"status"`
}
