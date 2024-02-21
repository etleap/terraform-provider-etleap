// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	speakeasy_boolplanmodifier "github.com/etleap/terraform-provider-etleap/internal/planmodifiers/boolplanmodifier"
	"github.com/etleap/terraform-provider-etleap/internal/sdk"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/operations"
	"github.com/etleap/terraform-provider-etleap/internal/validators"
	speakeasy_int64validators "github.com/etleap/terraform-provider-etleap/internal/validators/int64validators"
	speakeasy_stringvalidators "github.com/etleap/terraform-provider-etleap/internal/validators/stringvalidators"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ConnectionSAPHANASHARDEDResource{}
var _ resource.ResourceWithImportState = &ConnectionSAPHANASHARDEDResource{}

func NewConnectionSAPHANASHARDEDResource() resource.Resource {
	return &ConnectionSAPHANASHARDEDResource{}
}

// ConnectionSAPHANASHARDEDResource defines the resource implementation.
type ConnectionSAPHANASHARDEDResource struct {
	client *sdk.SDK
}

// ConnectionSAPHANASHARDEDResourceModel describes the resource data model.
type ConnectionSAPHANASHARDEDResourceModel struct {
	Active                   types.Bool              `tfsdk:"active"`
	CdcEnabled               types.Bool              `tfsdk:"cdc_enabled"`
	CreateDate               types.String            `tfsdk:"create_date"`
	DefaultUpdateSchedule    []DefaultUpdateSchedule `tfsdk:"default_update_schedule"`
	DeletionOfExportProducts types.Bool              `tfsdk:"deletion_of_export_products"`
	ID                       types.String            `tfsdk:"id"`
	Name                     types.String            `tfsdk:"name"`
	Schema                   types.String            `tfsdk:"schema"`
	Shards                   []DatabaseShard         `tfsdk:"shards"`
	Status                   types.String            `tfsdk:"status"`
	Type                     types.String            `tfsdk:"type"`
	UpdateSchedule           *UpdateScheduleTypes    `tfsdk:"update_schedule"`
}

func (r *ConnectionSAPHANASHARDEDResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_connection_sap_hana_sharded"
}

func (r *ConnectionSAPHANASHARDEDResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "ConnectionSAPHANASHARDED Resource",

		Attributes: map[string]schema.Attribute{
			"active": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether this connection has been marked as active.`,
			},
			"cdc_enabled": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_boolplanmodifier.SuppressDiff(speakeasy_boolplanmodifier.ExplicitSuppress),
				},
				Optional: true,
				Default:  booldefault.StaticBool(false),
				MarkdownDescription: `<i>Please note, this feature is currently in alpha and may be unstable.</i>` + "\n" +
					`` + "\n" +
					`Should Etleap use a change-tracking table and triggers defined on the source tables to capture changes from this database?` + "\n" +
					`` + "\n" +
					`For this setting to be enabled, the ` + "`" + `ETLEAP_CTT` + "`" + ` schema must be present in the source database, with the following privileges granted to the authenticating user: ` + "`" + `SELECT` + "`" + `, ` + "`" + `TRIGGER` + "`" + `, ` + "`" + `CREATE ANY` + "`" + `, and ` + "`" + `EXECUTE` + "`" + `.` + "\n" +
					`` + "\n" +
					`When enabled, Etleap will create a table and procedure in the ` + "`" + `ETLEAP_CTT` + "`" + ` schema and define triggers on any source table that has an Etleap pipeline. This setting cannot be changed later.` + "\n" +
					`Requires replacement if changed. ; Default: false`,
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
			"schema": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `If not specified, the default schema will be used.`,
			},
			"shards": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							Computed:    true,
							Optional:    true,
							Description: `Not Null`,
							Validators: []validator.String{
								speakeasy_stringvalidators.NotNull(),
							},
						},
						"database": schema.StringAttribute{
							Computed:    true,
							Optional:    true,
							Description: `Not Null`,
							Validators: []validator.String{
								speakeasy_stringvalidators.NotNull(),
							},
						},
						"password": schema.StringAttribute{
							Computed:    true,
							Optional:    true,
							Description: `Not Null`,
							Validators: []validator.String{
								speakeasy_stringvalidators.NotNull(),
							},
						},
						"port": schema.Int64Attribute{
							Computed:    true,
							Optional:    true,
							Description: `Not Null`,
							Validators: []validator.Int64{
								speakeasy_int64validators.NotNull(),
							},
						},
						"shard_id": schema.StringAttribute{
							Computed:    true,
							Optional:    true,
							Description: `Not Null`,
							Validators: []validator.String{
								speakeasy_stringvalidators.NotNull(),
							},
						},
						"ssh_config": schema.SingleNestedAttribute{
							Computed: true,
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"address": schema.StringAttribute{
									Computed:    true,
									Optional:    true,
									Description: `The server address for the SSH connection. Not Null`,
									Validators: []validator.String{
										speakeasy_stringvalidators.NotNull(),
									},
								},
								"username": schema.StringAttribute{
									Computed:    true,
									Optional:    true,
									Description: `The username for the SSH connection. Not Null`,
									Validators: []validator.String{
										speakeasy_stringvalidators.NotNull(),
									},
								},
							},
						},
						"username": schema.StringAttribute{
							Computed:    true,
							Optional:    true,
							Description: `Not Null`,
							Validators: []validator.String{
								speakeasy_stringvalidators.NotNull(),
							},
						},
					},
				},
				Validators: []validator.List{
					listvalidator.SizeAtLeast(2),
				},
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
				Description: `must be one of ["SAP_HANA_SHARDED"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"SAP_HANA_SHARDED",
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
		},
	}
}

func (r *ConnectionSAPHANASHARDEDResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ConnectionSAPHANASHARDEDResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ConnectionSAPHANASHARDEDResourceModel
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

	request := *data.ToSharedConnectionSapHanaShardedInput()
	res, err := r.client.Connection.CreateSAPHANASHARDEDConnection(ctx, request)
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
	if res.ConnectionSapHanaSharded == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedConnectionSapHanaSharded(res.ConnectionSapHanaSharded)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	id := data.ID.ValueString()
	request1 := operations.GetSAPHANASHARDEDConnectionRequest{
		ID: id,
	}
	res1, err := r.client.Connection.GetSAPHANASHARDEDConnection(ctx, request1)
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
	if res1.ConnectionSapHanaSharded == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedConnectionSapHanaSharded(res1.ConnectionSapHanaSharded)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ConnectionSAPHANASHARDEDResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ConnectionSAPHANASHARDEDResourceModel
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
	request := operations.GetSAPHANASHARDEDConnectionRequest{
		ID: id,
	}
	res, err := r.client.Connection.GetSAPHANASHARDEDConnection(ctx, request)
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
	if res.ConnectionSapHanaSharded == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedConnectionSapHanaSharded(res.ConnectionSapHanaSharded)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ConnectionSAPHANASHARDEDResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ConnectionSAPHANASHARDEDResourceModel
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
	connectionSapHanaShardedUpdate := data.ToSharedConnectionSapHanaShardedUpdate()
	request := operations.UpdateSAPHANASHARDEDConnectionRequest{
		ID:                             id,
		ConnectionSapHanaShardedUpdate: connectionSapHanaShardedUpdate,
	}
	res, err := r.client.Connection.UpdateSAPHANASHARDEDConnection(ctx, request)
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
	if res.ConnectionSapHanaSharded == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedConnectionSapHanaSharded(res.ConnectionSapHanaSharded)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)
	id1 := data.ID.ValueString()
	request1 := operations.GetSAPHANASHARDEDConnectionRequest{
		ID: id1,
	}
	res1, err := r.client.Connection.GetSAPHANASHARDEDConnection(ctx, request1)
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
	if res1.ConnectionSapHanaSharded == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res1.RawResponse))
		return
	}
	data.RefreshFromSharedConnectionSapHanaSharded(res1.ConnectionSapHanaSharded)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ConnectionSAPHANASHARDEDResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ConnectionSAPHANASHARDEDResourceModel
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
	request := operations.DeleteSAPHANASHARDEDConnectionRequest{
		ID:               id,
		ConnectionDelete: connectionDelete,
	}
	res, err := r.client.Connection.DeleteSAPHANASHARDEDConnection(ctx, request)
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

func (r *ConnectionSAPHANASHARDEDResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
