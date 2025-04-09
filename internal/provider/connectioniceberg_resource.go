// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	speakeasy_boolplanmodifier "github.com/etleap/terraform-provider-etleap/internal/planmodifiers/boolplanmodifier"
	speakeasy_int64planmodifier "github.com/etleap/terraform-provider-etleap/internal/planmodifiers/int64planmodifier"
	speakeasy_listplanmodifier "github.com/etleap/terraform-provider-etleap/internal/planmodifiers/listplanmodifier"
	speakeasy_objectplanmodifier "github.com/etleap/terraform-provider-etleap/internal/planmodifiers/objectplanmodifier"
	speakeasy_stringplanmodifier "github.com/etleap/terraform-provider-etleap/internal/planmodifiers/stringplanmodifier"
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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ConnectionICEBERGResource{}
var _ resource.ResourceWithImportState = &ConnectionICEBERGResource{}

func NewConnectionICEBERGResource() resource.Resource {
	return &ConnectionICEBERGResource{}
}

// ConnectionICEBERGResource defines the resource implementation.
type ConnectionICEBERGResource struct {
	client *sdk.SDK
}

// ConnectionICEBERGResourceModel describes the resource data model.
type ConnectionICEBERGResourceModel struct {
	Active                   types.Bool                                      `tfsdk:"active"`
	BaseDirectory            types.String                                    `tfsdk:"base_directory"`
	CreateDate               types.String                                    `tfsdk:"create_date"`
	DataBucket               types.String                                    `tfsdk:"data_bucket"`
	DefaultUpdateSchedule    []ConnectionActiveCampaignDefaultUpdateSchedule `tfsdk:"default_update_schedule"`
	DeletionOfExportProducts types.Bool                                      `tfsdk:"deletion_of_export_products"`
	GlueDatabase             types.String                                    `tfsdk:"glue_database"`
	GlueRegion               types.String                                    `tfsdk:"glue_region"`
	IamRole                  types.String                                    `tfsdk:"iam_role"`
	ID                       types.String                                    `tfsdk:"id"`
	Name                     types.String                                    `tfsdk:"name"`
	Status                   types.String                                    `tfsdk:"status"`
	Type                     types.String                                    `tfsdk:"type"`
	UpdateSchedule           *UpdateScheduleTypes                            `tfsdk:"update_schedule"`
	WarehouseConnection      types.String                                    `tfsdk:"warehouse_connection"`
}

func (r *ConnectionICEBERGResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_connection_iceberg"
}

func (r *ConnectionICEBERGResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "ConnectionICEBERG Resource",

		Attributes: map[string]schema.Attribute{
			"active": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					speakeasy_boolplanmodifier.SuppressDiff(speakeasy_boolplanmodifier.ExplicitSuppress),
				},
				Description: `Whether this connection has been marked as active.`,
			},
			"base_directory": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Required: true,
			},
			"create_date": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Description: `The date and time when then the connection was created.`,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"data_bucket": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Required: true,
			},
			"default_update_schedule": schema.ListNestedAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.List{
					speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"pipeline_mode": schema.StringAttribute{
							Computed: true,
							PlanModifiers: []planmodifier.String{
								speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
							},
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
							PlanModifiers: []planmodifier.Object{
								speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
							},
							Attributes: map[string]schema.Attribute{
								"daily": schema.SingleNestedAttribute{
									Computed: true,
									PlanModifiers: []planmodifier.Object{
										speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
									},
									Attributes: map[string]schema.Attribute{
										"hour_of_day": schema.Int64Attribute{
											Computed: true,
											PlanModifiers: []planmodifier.Int64{
												speakeasy_int64planmodifier.SuppressDiff(speakeasy_int64planmodifier.ExplicitSuppress),
											},
											Description: `Hour of day the  pipeline update should be started at (in UTC).`,
										},
										"mode": schema.StringAttribute{
											Computed: true,
											PlanModifiers: []planmodifier.String{
												speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
											},
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
									PlanModifiers: []planmodifier.Object{
										speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
									},
									Attributes: map[string]schema.Attribute{
										"mode": schema.StringAttribute{
											Computed: true,
											PlanModifiers: []planmodifier.String{
												speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
											},
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
									PlanModifiers: []planmodifier.Object{
										speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
									},
									Attributes: map[string]schema.Attribute{
										"interval_minutes": schema.Int64Attribute{
											Computed: true,
											PlanModifiers: []planmodifier.Int64{
												speakeasy_int64planmodifier.SuppressDiff(speakeasy_int64planmodifier.ExplicitSuppress),
											},
											Description: `Time to wait before new data is pulled (in minutes).`,
										},
										"mode": schema.StringAttribute{
											Computed: true,
											PlanModifiers: []planmodifier.String{
												speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
											},
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
									PlanModifiers: []planmodifier.Object{
										speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
									},
									Attributes: map[string]schema.Attribute{
										"day_of_month": schema.Int64Attribute{
											Computed: true,
											PlanModifiers: []planmodifier.Int64{
												speakeasy_int64planmodifier.SuppressDiff(speakeasy_int64planmodifier.ExplicitSuppress),
											},
										},
										"hour_of_day": schema.Int64Attribute{
											Computed: true,
											PlanModifiers: []planmodifier.Int64{
												speakeasy_int64planmodifier.SuppressDiff(speakeasy_int64planmodifier.ExplicitSuppress),
											},
											Description: `Hour of day the  pipeline update should be started at (in UTC).`,
										},
										"mode": schema.StringAttribute{
											Computed: true,
											PlanModifiers: []planmodifier.String{
												speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
											},
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
									PlanModifiers: []planmodifier.Object{
										speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
									},
									Attributes: map[string]schema.Attribute{
										"day_of_week": schema.Int64Attribute{
											Computed: true,
											PlanModifiers: []planmodifier.Int64{
												speakeasy_int64planmodifier.SuppressDiff(speakeasy_int64planmodifier.ExplicitSuppress),
											},
										},
										"hour_of_day": schema.Int64Attribute{
											Computed: true,
											PlanModifiers: []planmodifier.Int64{
												speakeasy_int64planmodifier.SuppressDiff(speakeasy_int64planmodifier.ExplicitSuppress),
											},
											Description: `Hour of day the  pipeline update should be started at (in UTC).`,
										},
										"mode": schema.StringAttribute{
											Computed: true,
											PlanModifiers: []planmodifier.String{
												speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
											},
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
			"glue_database": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Required: true,
			},
			"glue_region": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Required: true,
			},
			"iam_role": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Required: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Description: `The unique identifier of the connection.`,
			},
			"name": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Required:    true,
				Description: `The unique name of this connection.`,
			},
			"status": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
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
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Required:    true,
				Description: `must be one of ["ICEBERG"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"ICEBERG",
					),
				},
			},
			"update_schedule": schema.SingleNestedAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Object{
					speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
				},
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"daily": schema.SingleNestedAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.Object{
							speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
						},
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"hour_of_day": schema.Int64Attribute{
								Computed: true,
								PlanModifiers: []planmodifier.Int64{
									speakeasy_int64planmodifier.SuppressDiff(speakeasy_int64planmodifier.ExplicitSuppress),
								},
								Optional:    true,
								Description: `Hour of day the  pipeline update should be started at (in UTC). Not Null`,
								Validators: []validator.Int64{
									speakeasy_int64validators.NotNull(),
								},
							},
							"mode": schema.StringAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.String{
									speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
								},
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
						PlanModifiers: []planmodifier.Object{
							speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
						},
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"mode": schema.StringAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.String{
									speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
								},
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
						PlanModifiers: []planmodifier.Object{
							speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
						},
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"interval_minutes": schema.Int64Attribute{
								Computed: true,
								PlanModifiers: []planmodifier.Int64{
									speakeasy_int64planmodifier.SuppressDiff(speakeasy_int64planmodifier.ExplicitSuppress),
								},
								Optional:    true,
								Description: `Time to wait before new data is pulled (in minutes). Not Null`,
								Validators: []validator.Int64{
									speakeasy_int64validators.NotNull(),
								},
							},
							"mode": schema.StringAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.String{
									speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
								},
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
						PlanModifiers: []planmodifier.Object{
							speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
						},
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"day_of_month": schema.Int64Attribute{
								Computed: true,
								PlanModifiers: []planmodifier.Int64{
									speakeasy_int64planmodifier.SuppressDiff(speakeasy_int64planmodifier.ExplicitSuppress),
								},
								Optional:    true,
								Description: `Not Null`,
								Validators: []validator.Int64{
									speakeasy_int64validators.NotNull(),
								},
							},
							"hour_of_day": schema.Int64Attribute{
								Computed: true,
								PlanModifiers: []planmodifier.Int64{
									speakeasy_int64planmodifier.SuppressDiff(speakeasy_int64planmodifier.ExplicitSuppress),
								},
								Optional:    true,
								Description: `Hour of day the  pipeline update should be started at (in UTC). Not Null`,
								Validators: []validator.Int64{
									speakeasy_int64validators.NotNull(),
								},
							},
							"mode": schema.StringAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.String{
									speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
								},
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
						PlanModifiers: []planmodifier.Object{
							speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
						},
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"day_of_week": schema.Int64Attribute{
								Computed: true,
								PlanModifiers: []planmodifier.Int64{
									speakeasy_int64planmodifier.SuppressDiff(speakeasy_int64planmodifier.ExplicitSuppress),
								},
								Optional:    true,
								Description: `Not Null`,
								Validators: []validator.Int64{
									speakeasy_int64validators.NotNull(),
								},
							},
							"hour_of_day": schema.Int64Attribute{
								Computed: true,
								PlanModifiers: []planmodifier.Int64{
									speakeasy_int64planmodifier.SuppressDiff(speakeasy_int64planmodifier.ExplicitSuppress),
								},
								Optional:    true,
								Description: `Hour of day the  pipeline update should be started at (in UTC). Not Null`,
								Validators: []validator.Int64{
									speakeasy_int64validators.NotNull(),
								},
							},
							"mode": schema.StringAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.String{
									speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
								},
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
			"warehouse_connection": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Optional: true,
			},
		},
	}
}

func (r *ConnectionICEBERGResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ConnectionICEBERGResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ConnectionICEBERGResourceModel
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

	request := *data.ToSharedConnectionIcebergInput()
	res, err := r.client.Connection.CreateICEBERGConnection(ctx, request)
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
	if res.ConnectionIceberg == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedConnectionIceberg(res.ConnectionIceberg)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ConnectionICEBERGResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ConnectionICEBERGResourceModel
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
	request := operations.GetICEBERGConnectionRequest{
		ID: id,
	}
	res, err := r.client.Connection.GetICEBERGConnection(ctx, request)
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
	if res.ConnectionIceberg == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedConnectionIceberg(res.ConnectionIceberg)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ConnectionICEBERGResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ConnectionICEBERGResourceModel
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
	connectionIcebergUpdate := data.ToSharedConnectionIcebergUpdate()
	request := operations.UpdateICEBERGConnectionRequest{
		ID:                      id,
		ConnectionIcebergUpdate: connectionIcebergUpdate,
	}
	res, err := r.client.Connection.UpdateICEBERGConnection(ctx, request)
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
	if res.ConnectionIceberg == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedConnectionIceberg(res.ConnectionIceberg)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ConnectionICEBERGResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ConnectionICEBERGResourceModel
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
	request := operations.DeleteICEBERGConnectionRequest{
		ID:               id,
		ConnectionDelete: connectionDelete,
	}
	res, err := r.client.Connection.DeleteICEBERGConnection(ctx, request)
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

func (r *ConnectionICEBERGResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
