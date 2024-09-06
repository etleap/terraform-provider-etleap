// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	speakeasy_boolplanmodifier "github.com/etleap/terraform-provider-etleap/internal/planmodifiers/boolplanmodifier"
	speakeasy_stringplanmodifier "github.com/etleap/terraform-provider-etleap/internal/planmodifiers/stringplanmodifier"
	"github.com/etleap/terraform-provider-etleap/internal/sdk"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/operations"
	"github.com/etleap/terraform-provider-etleap/internal/validators"
	speakeasy_boolvalidators "github.com/etleap/terraform-provider-etleap/internal/validators/boolvalidators"
	speakeasy_int64validators "github.com/etleap/terraform-provider-etleap/internal/validators/int64validators"
	speakeasy_objectvalidators "github.com/etleap/terraform-provider-etleap/internal/validators/objectvalidators"
	speakeasy_stringvalidators "github.com/etleap/terraform-provider-etleap/internal/validators/stringvalidators"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ModelResource{}
var _ resource.ResourceWithImportState = &ModelResource{}

func NewModelResource() resource.Resource {
	return &ModelResource{}
}

// ModelResource defines the resource implementation.
type ModelResource struct {
	client *sdk.SDK
}

// ModelResourceModel describes the resource data model.
type ModelResourceModel struct {
	CreateDate               types.String         `tfsdk:"create_date"`
	DeletionOfExportProducts types.Bool           `tfsdk:"deletion_of_export_products"`
	Dependencies             []ModelDependency    `tfsdk:"dependencies"`
	ID                       types.String         `tfsdk:"id"`
	LastUpdateDuration       types.Int64          `tfsdk:"last_update_duration"`
	LastUpdateTime           types.String         `tfsdk:"last_update_time"`
	Name                     types.String         `tfsdk:"name"`
	Owner                    User                 `tfsdk:"owner"`
	Paused                   types.Bool           `tfsdk:"paused"`
	QueryAndTriggers         QueryAndTriggers     `tfsdk:"query_and_triggers"`
	Shares                   []types.String       `tfsdk:"shares"`
	UpdateSchedule           RefreshScheduleTypes `tfsdk:"update_schedule"`
	Warehouse                WarehouseTypesInput  `tfsdk:"warehouse"`
}

func (r *ModelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_model"
}

func (r *ModelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Model Resource",

		Attributes: map[string]schema.Attribute{
			"create_date": schema.StringAttribute{
				Computed:    true,
				Description: `The date and time when then the model was created.`,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"deletion_of_export_products": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
				Description: `Specifies whether the table in the destination created by this model should be deleted. Defaults to ` + "`" + `false` + "`" + `. Default: false`,
			},
			"dependencies": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed:    true,
							Description: `The unique identifier of the pipeline or model.`,
						},
						"name": schema.StringAttribute{
							Computed:    true,
							Description: `The name of the pipeline or model.`,
						},
						"type": schema.StringAttribute{
							Computed:    true,
							Description: `must be one of ["PIPELINE", "MODEL"]`,
							Validators: []validator.String{
								stringvalidator.OneOf(
									"PIPELINE",
									"MODEL",
								),
							},
						},
					},
				},
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `The unique identifier of the model.`,
			},
			"last_update_duration": schema.Int64Attribute{
				Computed:    true,
				Description: `How long the latest update took to complete, in milliseconds, or the duration of the current update if one is in progress.`,
			},
			"last_update_time": schema.StringAttribute{
				Computed:    true,
				Description: `The date and time of the latest successful update for this model.`,
				Validators: []validator.String{
					validators.IsRFC3339(),
				},
			},
			"name": schema.StringAttribute{
				Required: true,
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
				Computed: true,
			},
			"query_and_triggers": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"query": schema.StringAttribute{
						Required: true,
						MarkdownDescription: `The SQL query used to build this model. To specify dependencies on pipelines or other models, replace the schema and table name of the dependency with the id of the dependency enclosed in ` + "`" + `{{"{{"}}` + "`" + ` and ` + "`" + `{{"}}"}}` + "`" + `. The dependency must load data into the same Etleap connection as the one given in ` + "`" + `warehouse.connectionId` + "`" + ` for this model.` + "\n" +
							`` + "\n" +
							`**For Example**` + "\n" +
							`Say there is a pipeline with the id ` + "`" + `abcd1234` + "`" + ` which loads data to the table "schema"."my_table". To create a model in Etleap that has a dependency on this pipeline, the following query:` + "\n" +
							`` + "\n" +
							`` + "```" + `sql` + "\n" +
							`SELECT col1, col2 FROM "schema"."my_table";` + "\n" +
							`` + "```" + `` + "\n" +
							`` + "\n" +
							`becomes:` + "\n" +
							`` + "```" + `sql` + "\n" +
							`SELECT col1, col2 FROM {{"{{"}}abcd1234{{"}}"}};` + "\n" +
							`` + "```" + `` + "\n" +
							`` + "\n" +
							`[See the Model documentation](https://docs.etleap.com/docs/documentation/ZG9jOjI0MzU2NDY3-introduction-to-models#model-dependencies) for more information on Model dependencies.`,
					},
					"triggers": schema.ListAttribute{
						Required:    true,
						ElementType: types.StringType,
						Description: `A list of model dependency ids. An update will be automatically triggered in this model if any of the dependencies listed here get new data. Any ids given here must be present as dependencies in the ` + "`" + `query` + "`" + `.`,
						Validators: []validator.List{
							listvalidator.UniqueValues(),
						},
					},
				},
			},
			"shares": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
				Description: `An array of users' emails to share the model with. Once shared, a model cannot be unshared, and future calls to ` + "`" + `PATCH` + "`" + ` can only add to this list.`,
			},
			"update_schedule": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"daily": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"hour_of_day": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `Hour of day this schedule should trigger at (in UTC). Not Null`,
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
					},
					"monthly": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"day_of_month": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `Day of the month this schedule should trigger at (in UTC). Not Null`,
								Validators: []validator.Int64{
									speakeasy_int64validators.NotNull(),
								},
							},
							"hour_of_day": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `Hour of day this schedule should trigger at (in UTC). Not Null`,
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
					},
					"never": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"mode": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Not Null; must be one of ["NEVER"]`,
								Validators: []validator.String{
									speakeasy_stringvalidators.NotNull(),
									stringvalidator.OneOf(
										"NEVER",
									),
								},
							},
						},
					},
					"weekly": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"day_of_week": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `The day of the week this schedule should trigger at (in UTC). Not Null`,
								Validators: []validator.Int64{
									speakeasy_int64validators.NotNull(),
								},
							},
							"hour_of_day": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `Hour of day this schedule should trigger at (in UTC). Not Null`,
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
					},
				},
				Description: `How often this model should update. Etleap will periodically update the model table in your warehouse according to this schedule. See [the Model Updates documentation](https://docs.etleap.com/docs/documentation/ZG9jOjI0MzU2NDY3-introduction-to-models#model-updates) for more information.`,
				Validators: []validator.Object{
					validators.ExactlyOneChild(),
				},
			},
			"warehouse": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"redshift": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"connection_id": schema.StringAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplaceIfConfigured(),
									speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
								},
								Optional:    true,
								Description: `Requires replacement if changed. ; Not Null`,
								Validators: []validator.String{
									speakeasy_stringvalidators.NotNull(),
								},
							},
							"distribution_style": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"one": schema.StringAttribute{
										Computed:    true,
										Optional:    true,
										Description: `must be one of ["ALL", "AUTO", "EVEN"]`,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"ALL",
												"AUTO",
												"EVEN",
											),
										},
									},
									"distribution_style_key": schema.SingleNestedAttribute{
										Computed: true,
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"column": schema.StringAttribute{
												Computed:    true,
												Optional:    true,
												Description: `Not Null`,
												Validators: []validator.String{
													speakeasy_stringvalidators.NotNull(),
												},
											},
											"type": schema.StringAttribute{
												Computed:    true,
												Optional:    true,
												Description: `Not Null; must be one of ["KEY"]`,
												Validators: []validator.String{
													speakeasy_stringvalidators.NotNull(),
													stringvalidator.OneOf(
														"KEY",
													),
												},
											},
										},
									},
								},
								Description: `Can either be one the strings ` + "`" + `ALL` + "`" + `, ` + "`" + `AUTO` + "`" + ` or ` + "`" + `EVEN` + "`" + `, or an object for ` + "`" + `KEY` + "`" + ` distribution that specifies a column. Not Null`,
								Validators: []validator.Object{
									speakeasy_objectvalidators.NotNull(),
									validators.ExactlyOneChild(),
								},
							},
							"materialized_view": schema.BoolAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.RequiresReplaceIfConfigured(),
									speakeasy_boolplanmodifier.SuppressDiff(speakeasy_boolplanmodifier.ExplicitSuppress),
								},
								Optional:    true,
								Description: `Requires replacement if changed. ; Not Null`,
								Validators: []validator.Bool{
									speakeasy_boolvalidators.NotNull(),
								},
							},
							"pending_renamed_table": schema.StringAttribute{
								Computed:    true,
								Description: `Only set when a table rename was triggered but is not complete yet.`,
							},
							"schema": schema.StringAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplaceIfConfigured(),
									speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
								},
								Optional:    true,
								Description: `Requires replacement if changed. `,
							},
							"sort_columns": schema.ListAttribute{
								Computed:    true,
								Optional:    true,
								ElementType: types.StringType,
								Description: `The sort columns to use.`,
							},
							"table": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Not Null`,
								Validators: []validator.String{
									speakeasy_stringvalidators.NotNull(),
								},
							},
							"type": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Not Null; must be one of ["REDSHIFT"]`,
								Validators: []validator.String{
									speakeasy_stringvalidators.NotNull(),
									stringvalidator.OneOf(
										"REDSHIFT",
									),
								},
							},
							"wait_for_update_preparation": schema.BoolAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.RequiresReplaceIfConfigured(),
									speakeasy_boolplanmodifier.SuppressDiff(speakeasy_boolplanmodifier.ExplicitSuppress),
								},
								Optional:    true,
								Description: `Requires replacement if changed. ; Not Null`,
								Validators: []validator.Bool{
									speakeasy_boolvalidators.NotNull(),
								},
							},
						},
					},
					"snowflake": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"connection_id": schema.StringAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplaceIfConfigured(),
									speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
								},
								Optional:    true,
								Description: `Requires replacement if changed. ; Not Null`,
								Validators: []validator.String{
									speakeasy_stringvalidators.NotNull(),
								},
							},
							"materialized_view": schema.BoolAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.RequiresReplaceIfConfigured(),
									speakeasy_boolplanmodifier.SuppressDiff(speakeasy_boolplanmodifier.ExplicitSuppress),
								},
								Optional:    true,
								Description: `Requires replacement if changed. ; Not Null`,
								Validators: []validator.Bool{
									speakeasy_boolvalidators.NotNull(),
								},
							},
							"pending_renamed_table": schema.StringAttribute{
								Computed:    true,
								Description: `Only set when a table rename was triggered but is not complete yet.`,
							},
							"schema": schema.StringAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplaceIfConfigured(),
									speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
								},
								Optional:    true,
								Description: `Requires replacement if changed. `,
							},
							"table": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Not Null`,
								Validators: []validator.String{
									speakeasy_stringvalidators.NotNull(),
								},
							},
							"type": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Not Null; must be one of ["SNOWFLAKE"]`,
								Validators: []validator.String{
									speakeasy_stringvalidators.NotNull(),
									stringvalidator.OneOf(
										"SNOWFLAKE",
									),
								},
							},
							"wait_for_update_preparation": schema.BoolAttribute{
								Computed: true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.RequiresReplaceIfConfigured(),
									speakeasy_boolplanmodifier.SuppressDiff(speakeasy_boolplanmodifier.ExplicitSuppress),
								},
								Optional:    true,
								Description: `Requires replacement if changed. ; Not Null`,
								Validators: []validator.Bool{
									speakeasy_boolvalidators.NotNull(),
								},
							},
						},
					},
				},
				Validators: []validator.Object{
					validators.ExactlyOneChild(),
				},
			},
		},
	}
}

func (r *ModelResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *ModelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ModelResourceModel
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

	request := data.ToSharedModelInput()
	res, err := r.client.Model.Create(ctx, request)
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
	if res.ModelOutput == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedModelOutput(res.ModelOutput)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ModelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ModelResourceModel
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
	request := operations.GetModelRequest{
		ID: id,
	}
	res, err := r.client.Model.Get(ctx, request)
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
	if res.ModelOutput == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedModelOutput(res.ModelOutput)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ModelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ModelResourceModel
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
	modelUpdate := *data.ToSharedModelUpdate()
	request := operations.UpdateModelRequest{
		ID:          id,
		ModelUpdate: modelUpdate,
	}
	res, err := r.client.Model.Update(ctx, request)
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
	if res.ModelOutput == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedModelOutput(res.ModelOutput)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ModelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ModelResourceModel
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
	modelDelete := data.ToSharedModelDelete()
	request := operations.DeleteModelRequest{
		ID:          id,
		ModelDelete: modelDelete,
	}
	res, err := r.client.Model.Delete(ctx, request)
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

func (r *ModelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
