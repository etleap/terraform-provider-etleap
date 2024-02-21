// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/operations"
	"github.com/etleap/terraform-provider-etleap/internal/validators"
	speakeasy_int64validators "github.com/etleap/terraform-provider-etleap/internal/validators/int64validators"
	speakeasy_stringvalidators "github.com/etleap/terraform-provider-etleap/internal/validators/stringvalidators"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ConnectionBINGADSResource{}
var _ resource.ResourceWithImportState = &ConnectionBINGADSResource{}

func NewConnectionBINGADSResource() resource.Resource {
	return &ConnectionBINGADSResource{}
}

// ConnectionBINGADSResource defines the resource implementation.
type ConnectionBINGADSResource struct {
	client *sdk.SDK
}

// ConnectionBINGADSResourceModel describes the resource data model.
type ConnectionBINGADSResourceModel struct {
	Active                   types.Bool              `tfsdk:"active"`
	Code                     types.String            `tfsdk:"code"`
	CreateDate               types.String            `tfsdk:"create_date"`
	DefaultUpdateSchedule    []DefaultUpdateSchedule `tfsdk:"default_update_schedule"`
	DeletionOfExportProducts types.Bool              `tfsdk:"deletion_of_export_products"`
	ID                       types.String            `tfsdk:"id"`
	Name                     types.String            `tfsdk:"name"`
	Status                   types.String            `tfsdk:"status"`
	Type                     types.String            `tfsdk:"type"`
	UpdateSchedule           *UpdateScheduleTypes    `tfsdk:"update_schedule"`
	UserID                   types.Int64             `tfsdk:"user_id"`
	Username                 types.String            `tfsdk:"username"`
}

func (r *ConnectionBINGADSResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_connection_bing_ads"
}

func (r *ConnectionBINGADSResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "ConnectionBINGADS Resource",

		Attributes: map[string]schema.Attribute{
			"active": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether this connection has been marked as active.`,
			},
			"code": schema.StringAttribute{
				Required:    true,
				Description: `Code retrieved from ` + "`" + `/connections/oauth2-initiation` + "`" + `. **Note:** it is short-lived, therefore the connection creation should be done as soon as code is returned.`,
			},
			"create_date": schema.StringAttribute{
				Computed:    true,
				Description: `The date and time when then the connection was created.`,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"default_update_schedule": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"pipeline_mode": schema.StringAttribute{
							Computed:    true,
							Description: `The pipeline mode refers to how the pipeline fetches data changes from the source and how those changes are applied to the destination table. See <a target="_blank" href="https://docs.etleap.com/docs/documentation/ZG9jOjIyMjE3ODA2-introduction">the documentation</a> for more details. must be one of ["APPEND", "REPLACE", "UPDATE", "QUERY"]`,
							Validators: []validator.String{
								stringvalidator.OneOf(
									"APPEND",
									"REPLACE",
									"UPDATE",
									"QUERY",
								),
							},
						},
						"update_schedule": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"daily": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"hour_of_day": schema.Int64Attribute{
											Computed:    true,
											Description: `Hour of day the  pipeline update should be started at (in UTC).`,
										},
										"mode": schema.StringAttribute{
											Computed:    true,
											Description: `must be one of ["DAILY"]`,
											Validators: []validator.String{
												stringvalidator.OneOf(
													"DAILY",
												),
											},
										},
									},
									Description: `The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information.`,
								},
								"hourly": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"mode": schema.StringAttribute{
											Computed:    true,
											Description: `must be one of ["HOURLY"]`,
											Validators: []validator.String{
												stringvalidator.OneOf(
													"HOURLY",
												),
											},
										},
									},
									Description: `The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information.`,
								},
								"interval": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"interval_minutes": schema.Int64Attribute{
											Computed:    true,
											Description: `Time to wait before new data is pulled (in minutes).`,
										},
										"mode": schema.StringAttribute{
											Computed:    true,
											Description: `must be one of ["INTERVAL"]`,
											Validators: []validator.String{
												stringvalidator.OneOf(
													"INTERVAL",
												),
											},
										},
									},
									Description: `The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information.`,
								},
								"monthly": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"day_of_month": schema.Int64Attribute{
											Computed: true,
										},
										"hour_of_day": schema.Int64Attribute{
											Computed:    true,
											Description: `Hour of day the  pipeline update should be started at (in UTC).`,
										},
										"mode": schema.StringAttribute{
											Computed:    true,
											Description: `must be one of ["MONTHLY"]`,
											Validators: []validator.String{
												stringvalidator.OneOf(
													"MONTHLY",
												),
											},
										},
									},
									Description: `The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information.`,
								},
								"weekly": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"day_of_week": schema.Int64Attribute{
											Computed: true,
										},
										"hour_of_day": schema.Int64Attribute{
											Computed:    true,
											Description: `Hour of day the  pipeline update should be started at (in UTC).`,
										},
										"mode": schema.StringAttribute{
											Computed:    true,
											Description: `must be one of ["WEEKLY"]`,
											Validators: []validator.String{
												stringvalidator.OneOf(
													"WEEKLY",
												),
											},
										},
									},
									Description: `The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information.`,
								},
							},
							Description: `The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.`,
							Validators: []validator.Object{
								validators.ExactlyOneChild(),
							},
						},
					},
				},
				Description: `When an update schedule is not defined for a connection, the default schedule is used. The default defined individually per ` + "`" + `pipelineMode` + "`" + ` and may be subject to change.`,
			},
			"deletion_of_export_products": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
				Description: `Applicable for REDSHIFT and SNOWFLAKE connections only in the case when there are pipelines that use this connection as a destination, and these pipelines have been migrated to use a different destination. Specifies whether any tables created by these pipelines in this destination should be deleted. Defaults to ` + "`" + `false` + "`" + `. Default: false`,
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `The unique identifier of the connection.`,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: `The unique name of this connection.`,
			},
			"status": schema.StringAttribute{
				Computed:    true,
				Description: `The current status of the connection. must be one of ["UNKNOWN", "UP", "DOWN", "RESIZE", "MAINTENANCE", "QUOTA", "CREATING"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"UNKNOWN",
						"UP",
						"DOWN",
						"RESIZE",
						"MAINTENANCE",
						"QUOTA",
						"CREATING",
					),
				},
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: `must be one of ["BING_ADS"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"BING_ADS",
					),
				},
			},
			"update_schedule": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"daily": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"hour_of_day": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `Hour of day the  pipeline update should be started at (in UTC). Not Null`,
								Validators: []validator.Int64{
									speakeasy_int64validators.NotNull(),
								},
							},
							"mode": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Not Null; must be one of ["DAILY"]`,
								Validators: []validator.String{
									speakeasy_stringvalidators.NotNull(),
									stringvalidator.OneOf(
										"DAILY",
									),
								},
							},
						},
						Description: `The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information.`,
					},
					"hourly": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"mode": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Not Null; must be one of ["HOURLY"]`,
								Validators: []validator.String{
									speakeasy_stringvalidators.NotNull(),
									stringvalidator.OneOf(
										"HOURLY",
									),
								},
							},
						},
						Description: `The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information.`,
					},
					"interval": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"interval_minutes": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `Time to wait before new data is pulled (in minutes). Not Null`,
								Validators: []validator.Int64{
									speakeasy_int64validators.NotNull(),
								},
							},
							"mode": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Not Null; must be one of ["INTERVAL"]`,
								Validators: []validator.String{
									speakeasy_stringvalidators.NotNull(),
									stringvalidator.OneOf(
										"INTERVAL",
									),
								},
							},
						},
						Description: `The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information.`,
					},
					"monthly": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"day_of_month": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `Not Null`,
								Validators: []validator.Int64{
									speakeasy_int64validators.NotNull(),
								},
							},
							"hour_of_day": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `Hour of day the  pipeline update should be started at (in UTC). Not Null`,
								Validators: []validator.Int64{
									speakeasy_int64validators.NotNull(),
								},
							},
							"mode": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Not Null; must be one of ["MONTHLY"]`,
								Validators: []validator.String{
									speakeasy_stringvalidators.NotNull(),
									stringvalidator.OneOf(
										"MONTHLY",
									),
								},
							},
						},
						Description: `The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information.`,
					},
					"weekly": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"day_of_week": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `Not Null`,
								Validators: []validator.Int64{
									speakeasy_int64validators.NotNull(),
								},
							},
							"hour_of_day": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `Hour of day the  pipeline update should be started at (in UTC). Not Null`,
								Validators: []validator.Int64{
									speakeasy_int64validators.NotNull(),
								},
							},
							"mode": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Not Null; must be one of ["WEEKLY"]`,
								Validators: []validator.String{
									speakeasy_stringvalidators.NotNull(),
									stringvalidator.OneOf(
										"WEEKLY",
									),
								},
							},
						},
						Description: `The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information.`,
					},
				},
				Description: `The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.`,
				Validators: []validator.Object{
					validators.ExactlyOneChild(),
				},
			},
			"user_id": schema.Int64Attribute{
				Computed: true,
			},
			"username": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *ConnectionBINGADSResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *ConnectionBINGADSResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ConnectionBINGADSResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	request := *data.ToSharedConnectionBingInput()
	res, err := r.client.Connection.CreateBINGADSConnection(ctx, request)
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
	if res.ConnectionBing == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedConnectionBing(res.ConnectionBing)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	id := data.ID.ValueString()
	request1 := operations.GetBINGADSConnectionRequest{
		ID: id,
	}
	res1, err := r.client.Connection.GetBINGADSConnection(ctx, request1)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res1 != nil && res1.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res1.RawResponse))
		}
		return
	}
	if res1 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res1))
		return
	}
	if res1.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res1.StatusCode), debugResponse(res1.RawResponse))
		return
	}
	if res1.ConnectionBing == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedConnectionBing(res1.ConnectionBing)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ConnectionBINGADSResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ConnectionBINGADSResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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
	request := operations.GetBINGADSConnectionRequest{
		ID: id,
	}
	res, err := r.client.Connection.GetBINGADSConnection(ctx, request)
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
	if res.ConnectionBing == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedConnectionBing(res.ConnectionBing)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ConnectionBINGADSResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ConnectionBINGADSResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	id := data.ID.ValueString()
	connectionBingUpdate := data.ToSharedConnectionBingUpdate()
	request := operations.UpdateBINGADSConnectionRequest{
		ID:                   id,
		ConnectionBingUpdate: connectionBingUpdate,
	}
	res, err := r.client.Connection.UpdateBINGADSConnection(ctx, request)
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
	if res.ConnectionBing == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedConnectionBing(res.ConnectionBing)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	id1 := data.ID.ValueString()
	request1 := operations.GetBINGADSConnectionRequest{
		ID: id1,
	}
	res1, err := r.client.Connection.GetBINGADSConnection(ctx, request1)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res1 != nil && res1.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res1.RawResponse))
		}
		return
	}
	if res1 == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res1))
		return
	}
	if res1.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res1.StatusCode), debugResponse(res1.RawResponse))
		return
	}
	if res1.ConnectionBing == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedConnectionBing(res1.ConnectionBing)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ConnectionBINGADSResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ConnectionBINGADSResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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
	connectionDelete := data.ToSharedConnectionDelete()
	request := operations.DeleteBINGADSConnectionRequest{
		ID:               id,
		ConnectionDelete: connectionDelete,
	}
	res, err := r.client.Connection.DeleteBINGADSConnection(ctx, request)
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
	if res.StatusCode != 204 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *ConnectionBINGADSResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
