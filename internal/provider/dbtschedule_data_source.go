// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/operations"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &DbtScheduleDataSource{}
var _ datasource.DataSourceWithConfigure = &DbtScheduleDataSource{}

func NewDbtScheduleDataSource() datasource.DataSource {
	return &DbtScheduleDataSource{}
}

// DbtScheduleDataSource is the data source implementation.
type DbtScheduleDataSource struct {
	client *sdk.SDK
}

// DbtScheduleDataSourceModel describes the data model.
type DbtScheduleDataSourceModel struct {
	ConnectionID         types.String        `tfsdk:"connection_id"`
	CreateDate           types.String        `tfsdk:"create_date"`
	Cron                 types.String        `tfsdk:"cron"`
	CurrentActivity      types.String        `tfsdk:"current_activity"`
	ID                   types.String        `tfsdk:"id"`
	LastDbtBuildDate     types.String        `tfsdk:"last_dbt_build_date"`
	LastDbtRunTime       types.Int64         `tfsdk:"last_dbt_run_time"`
	LatestRun            DbtScheduleRunTypes `tfsdk:"latest_run"`
	Name                 types.String        `tfsdk:"name"`
	Owner                User                `tfsdk:"owner"`
	Paused               types.Bool          `tfsdk:"paused"`
	Selector             types.String        `tfsdk:"selector"`
	SkipBuildIfNoNewData types.Bool          `tfsdk:"skip_build_if_no_new_data"`
	TargetSchema         types.String        `tfsdk:"target_schema"`
}

// Metadata returns the data source type name.
func (r *DbtScheduleDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dbt_schedule"
}

// Schema defines the schema for the data source.
func (r *DbtScheduleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "DbtSchedule DataSource",

		Attributes: map[string]schema.Attribute{
			"connection_id": schema.StringAttribute{
				Computed:    true,
				Description: `The [connection](https://docs.etleap.com/docs/api-v2/edbec13814bbc-connection) where the dbt build runs. The only supported connections are Redshift, Snowflake or Databricks Delta Lake destinations.`,
			},
			"create_date": schema.StringAttribute{
				Computed: true,
			},
			"cron": schema.StringAttribute{
				Computed:    true,
				Description: `The cron expression that defines triggers for this schedule. The maximum supported cron schedule precision is 1 minute.`,
			},
			"current_activity": schema.StringAttribute{
				Computed:    true,
				Description: `This field is deprecated and will be removed and replaced by the properties in ` + "`" + `latestRun` + "`" + ` when that field is implemented. must be one of ["LOADING", "BUILDING"]`,
			},
			"id": schema.StringAttribute{
				Required:    true,
				Description: `The id of the dbt schedule`,
			},
			"last_dbt_build_date": schema.StringAttribute{
				Computed:    true,
				Description: `The last time that a successful dbt build started. This field is deprecated and will be removed and replaced by the properties in ` + "`" + `latestRun` + "`" + ` when that field is implemented.`,
			},
			"last_dbt_run_time": schema.Int64Attribute{
				Computed:    true,
				Description: `The duration of the last successful dbt build. This field is deprecated and will be removed and replaced by the properties in ` + "`" + `latestRun` + "`" + ` when that field is implemented.`,
			},
			"latest_run": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"etleap_error": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"duration": schema.Int64Attribute{
								Computed:    true,
								Description: `The duration, in seconds, between the time this dbt run was triggered and the time the dbt build for this run completed.`,
							},
							"last_successful_dbt_build_date": schema.StringAttribute{
								Computed:    true,
								Description: `The last time that a successful dbt build finished.`,
							},
							"next_trigger_date": schema.StringAttribute{
								Computed:    true,
								Description: `Timestamp for the next dbt schedule trigger.`,
							},
							"start_date": schema.StringAttribute{
								Computed:    true,
								Description: `The time that this dbt run was triggered.`,
							},
							"status": schema.StringAttribute{
								Computed:    true,
								Description: `must be one of ["ETLEAP_ERROR", "DBT_ERROR"]`,
							},
						},
					},
					"in_progress": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"build_is_taking_too_long": schema.BoolAttribute{
								Computed:    true,
								Description: `Whether the dbt build phase is taking too long.`,
							},
							"last_successful_dbt_build_date": schema.StringAttribute{
								Computed:    true,
								Description: `The last time that a successful dbt build finished.`,
							},
							"phase": schema.StringAttribute{
								Computed:    true,
								Description: `The phase that this dbt run is currently in. dbt runs consist of an ` + "`" + `INGEST` + "`" + ` phase where source pipelines ingest data into the warehouse, followed by a ` + "`" + `BUILD` + "`" + ` phase where the dbt models are built. must be one of ["INGEST", "BUILD"]`,
							},
							"previous_run_duration": schema.Int64Attribute{
								Computed:    true,
								Description: `The duration, in seconds, between the time the previous run was triggered and the time it completed. This will be ` + "`" + `null` + "`" + ` if this is the first time this schedule has run.`,
							},
							"previous_run_status": schema.StringAttribute{
								Computed:    true,
								Description: `must be one of ["IN_PROGRESS", "ETLEAP_ERROR", "DBT_ERROR", "SUCCESS_WITH_DBT_WARNINGS", "SUCCESS"]`,
							},
							"start_date": schema.StringAttribute{
								Computed:    true,
								Description: `The time that this dbt run was triggered.`,
							},
							"status": schema.StringAttribute{
								Computed:    true,
								Description: `must be one of ["IN_PROGRESS"]`,
							},
						},
					},
					"not_yet_run": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"next_trigger_date": schema.StringAttribute{
								Computed:    true,
								Description: `Timestamp for the next dbt schedule trigger.`,
							},
							"status": schema.StringAttribute{
								Computed:    true,
								Description: `must be one of ["NOT_YET_RUN"]`,
							},
						},
					},
					"success_with_dbt_warnings": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"duration": schema.Int64Attribute{
								Computed:    true,
								Description: `The duration, in seconds, between the time this dbt run was triggered and the time the dbt build for this run completed.`,
							},
							"last_successful_dbt_build_date": schema.StringAttribute{
								Computed:    true,
								Description: `The last time that a successful dbt build finished.`,
							},
							"next_trigger_date": schema.StringAttribute{
								Computed:    true,
								Description: `Timestamp for the next dbt schedule trigger.`,
							},
							"start_date": schema.StringAttribute{
								Computed:    true,
								Description: `The time that this dbt run was triggered.`,
							},
							"status": schema.StringAttribute{
								Computed:    true,
								Description: `must be one of ["SUCCESS", "SUCCESS_WITH_DBT_WARNINGS"]`,
							},
						},
					},
				},
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: `The name of the dbt schedule.`,
			},
			"owner": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"email_address": schema.StringAttribute{
						Computed: true,
					},
					"first_name": schema.StringAttribute{
						Computed: true,
					},
					"id": schema.StringAttribute{
						Computed: true,
					},
					"last_name": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"paused": schema.BoolAttribute{
				Computed:    true,
				Description: `` + "`" + `true` + "`" + ` if the schedule is paused.`,
			},
			"selector": schema.StringAttribute{
				Computed:    true,
				Description: `The selector this schedule runs.`,
			},
			"skip_build_if_no_new_data": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether the dbt build is skipped if no new data has been ingested for any of the pipelines in the table above.`,
			},
			"target_schema": schema.StringAttribute{
				Computed:    true,
				Description: `The target schema for the dbt build. See [here](https://docs.getdbt.com/docs/build/custom-schemas) for details on how it's used.`,
			},
		},
	}
}

func (r *DbtScheduleDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *DbtScheduleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *DbtScheduleDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	id := data.ID.ValueString()
	request := operations.GetDbtScheduleRequest{
		ID: id,
	}
	res, err := r.client.DbtSchedule.Get(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.DbtScheduleOutput == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedDbtScheduleOutput(res.DbtScheduleOutput)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
