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
	ConnectionID         types.String `tfsdk:"connection_id"`
	CreateDate           types.String `tfsdk:"create_date"`
	Cron                 types.String `tfsdk:"cron"`
	CurrentActivity      types.String `tfsdk:"current_activity"`
	ID                   types.String `tfsdk:"id"`
	LastDbtBuildDate     types.String `tfsdk:"last_dbt_build_date"`
	LastDbtRunTime       types.Int64  `tfsdk:"last_dbt_run_time"`
	Name                 types.String `tfsdk:"name"`
	Owner                User         `tfsdk:"owner"`
	Paused               types.Bool   `tfsdk:"paused"`
	Selector             types.String `tfsdk:"selector"`
	SkipBuildIfNoNewData types.Bool   `tfsdk:"skip_build_if_no_new_data"`
	TargetSchema         types.String `tfsdk:"target_schema"`
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
				Description: `must be one of ["LOADING", "BUILDING"]`,
			},
			"id": schema.StringAttribute{
				Required:    true,
				Description: `The id of the dbt schedule`,
			},
			"last_dbt_build_date": schema.StringAttribute{
				Computed:    true,
				Description: `The last time that a successful dbt build started.`,
			},
			"last_dbt_run_time": schema.Int64Attribute{
				Computed:    true,
				Description: `The duration of the last successful dbt build.`,
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
