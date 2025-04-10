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
var _ datasource.DataSource = &ConnectionUSERDEFINEDAPIDataSource{}
var _ datasource.DataSourceWithConfigure = &ConnectionUSERDEFINEDAPIDataSource{}

func NewConnectionUSERDEFINEDAPIDataSource() datasource.DataSource {
	return &ConnectionUSERDEFINEDAPIDataSource{}
}

// ConnectionUSERDEFINEDAPIDataSource is the data source implementation.
type ConnectionUSERDEFINEDAPIDataSource struct {
	client *sdk.SDK
}

// ConnectionUSERDEFINEDAPIDataSourceModel describes the data model.
type ConnectionUSERDEFINEDAPIDataSourceModel struct {
	Active                types.Bool                                      `tfsdk:"active"`
	Authentication        Authentication                                  `tfsdk:"authentication"`
	CreateDate            types.String                                    `tfsdk:"create_date"`
	DefaultUpdateSchedule []ConnectionActiveCampaignDefaultUpdateSchedule `tfsdk:"default_update_schedule"`
	Entities              []UserDefinedAPIEntity                          `tfsdk:"entities"`
	ID                    types.String                                    `tfsdk:"id"`
	Name                  types.String                                    `tfsdk:"name"`
	Status                types.String                                    `tfsdk:"status"`
	Type                  types.String                                    `tfsdk:"type"`
	UpdateSchedule        *UpdateScheduleTypes                            `tfsdk:"update_schedule"`
}

// Metadata returns the data source type name.
func (r *ConnectionUSERDEFINEDAPIDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_connection_user_defined_api"
}

// Schema defines the schema for the data source.
func (r *ConnectionUSERDEFINEDAPIDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "ConnectionUSERDEFINEDAPI DataSource",

		Attributes: map[string]schema.Attribute{
			"active": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether this connection has been marked as active.`,
			},
			"authentication": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"basic": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"password": schema.StringAttribute{
								Computed:    true,
								Description: `The Basic password. This property is stored as a secret in Etleap.`,
							},
							"type": schema.StringAttribute{
								Computed:    true,
								Description: `must be one of ["BASIC"]`,
							},
							"username": schema.StringAttribute{
								Computed: true,
							},
						},
						Description: `Authentication method that uses a username/password pair.`,
					},
					"bearer": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"token": schema.StringAttribute{
								Computed:    true,
								Description: `The Bearer token used for authentication. This property is stored as a secret in Etleap.`,
							},
							"type": schema.StringAttribute{
								Computed:    true,
								Description: `must be one of ["BEARER"]`,
							},
						},
						Description: `Authentication method that uses a Bearer token.`,
					},
					"header": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"key": schema.StringAttribute{
								Computed: true,
							},
							"type": schema.StringAttribute{
								Computed:    true,
								Description: `must be one of ["HEADER"]`,
							},
							"value": schema.StringAttribute{
								Computed: true,
							},
						},
						Description: `Authentication method that uses a header.`,
					},
				},
			},
			"create_date": schema.StringAttribute{
				Computed:    true,
				Description: `The date and time when then the connection was created.`,
			},
			"default_update_schedule": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"pipeline_mode": schema.StringAttribute{
							Computed:    true,
							Description: `The pipeline mode refers to how the pipeline fetches data changes from the source and how those changes are applied to the destination table. See <a target="_blank" href="https://docs.etleap.com/docs/documentation/ZG9jOjIyMjE3ODA2-introduction">the documentation</a> for more details. must be one of ["APPEND", "REPLACE", "UPDATE", "QUERY"]`,
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
										},
									},
									Description: `The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information.`,
								},
							},
							Description: `The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.`,
						},
					},
				},
				Description: `When an update schedule is not defined for a connection, the default schedule is used. The default defined individually per ` + "`" + `pipelineMode` + "`" + ` and may be subject to change.`,
			},
			"entities": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"api_url": schema.StringAttribute{
							Computed:    true,
							Description: `The base URL to fetch data. Note that query parameters for things like pagination and sorting will be appended.`,
						},
						"columns": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
							Description: `The columns of the entity.`,
						},
						"display_name": schema.StringAttribute{
							Computed:    true,
							Description: `The name of the entity.`,
						},
						"header_parameters": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Computed: true,
									},
									"value": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							Description: `A list of headers to be passed with all the requests.`,
						},
						"id": schema.StringAttribute{
							Computed:    true,
							Description: `The unique identifier of the entity.`,
						},
						"paging_strategy": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"cursor_uri": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"max_page_size": schema.Int64Attribute{
											Computed:    true,
											Description: `The maximum page size supported by the API.`,
										},
										"page_size_field_name": schema.StringAttribute{
											Computed:    true,
											Description: `The name of the request parameter used to specify the page size.`,
										},
										"path_to_cursor": schema.StringAttribute{
											Computed:    true,
											Description: `The path to the paging cursor inside the response body.`,
										},
										"type": schema.StringAttribute{
											Computed:    true,
											Description: `must be one of ["CURSOR_URI"]`,
										},
										"url_prefix": schema.StringAttribute{
											Computed:    true,
											Description: `String prepended to the paging cursor string to turn it into a URL, e.g. because the cursor only contains the URL path.`,
										},
									},
									Description: `Paging strategy that uses a cursor to iterate through the results`,
								},
								"offset": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"max_page_size": schema.Int64Attribute{
											Computed:    true,
											Description: `The maximum page size supported by the API.`,
										},
										"offset_field_name": schema.StringAttribute{
											Computed:    true,
											Description: `The name of the request parameter used to specify the offset.`,
										},
										"page_size_field_name": schema.StringAttribute{
											Computed:    true,
											Description: `The name of the request parameter used to specify the page size.`,
										},
										"type": schema.StringAttribute{
											Computed:    true,
											Description: `must be one of ["OFFSET"]`,
										},
									},
									Description: `Paging strategy that is based on a result set offset.`,
								},
							},
							Description: `The paging strategy.`,
						},
						"path_to_results": schema.StringAttribute{
							Computed:    true,
							Description: `The [JMESPath](https://jmespath.org/) expression that converts the API response into an array containing one JSON object per record.`,
						},
						"pipeline_mode": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"one": schema.StringAttribute{
									Computed:    true,
									Description: `must be one of ["REPLACE"]`,
								},
								"user_defined_api_replace_mode": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"foreign_key_columns": schema.ListNestedAttribute{
											Computed: true,
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"foreign_column": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"column_path": schema.ListAttribute{
																Computed:    true,
																ElementType: types.StringType,
																Description: `Represents the path in the entity schema where the column is located. If the column is at the top level of an entity, use ` + "`" + `Top Level Foreign Key Column` + "`" + ` instead.`,
															},
															"referenced_column_name": schema.StringAttribute{
																Computed:    true,
																Description: `The column name of the referenced entity.`,
															},
															"referenced_entity_id": schema.StringAttribute{
																Computed:    true,
																Description: `The id of the referenced entity.`,
															},
														},
													},
													"top_level_foreign_key_column": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"column_name": schema.StringAttribute{
																Computed:    true,
																Description: `The entity's foreign key column. If the column is nested inside the entity's structure and not at the top level, use ` + "`" + `NestedForeignKeyColumn` + "`" + ` instead.`,
															},
															"referenced_column_name": schema.StringAttribute{
																Computed:    true,
																Description: `The column name of the referenced entity.`,
															},
															"referenced_entity_id": schema.StringAttribute{
																Computed:    true,
																Description: `The id of the referenced entity.`,
															},
														},
													},
												},
											},
											Description: `The foreign columns of the entity.`,
										},
										"primary_key_columns": schema.ListAttribute{
											Computed:    true,
											ElementType: types.StringType,
										},
										"type": schema.StringAttribute{
											Computed:    true,
											Description: `must be one of ["REPLACE"]`,
										},
									},
								},
								"user_defined_api_update_mode": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"foreign_key_columns": schema.ListNestedAttribute{
											Computed: true,
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"foreign_column": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"column_path": schema.ListAttribute{
																Computed:    true,
																ElementType: types.StringType,
																Description: `Represents the path in the entity schema where the column is located. If the column is at the top level of an entity, use ` + "`" + `Top Level Foreign Key Column` + "`" + ` instead.`,
															},
															"referenced_column_name": schema.StringAttribute{
																Computed:    true,
																Description: `The column name of the referenced entity.`,
															},
															"referenced_entity_id": schema.StringAttribute{
																Computed:    true,
																Description: `The id of the referenced entity.`,
															},
														},
													},
													"top_level_foreign_key_column": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"column_name": schema.StringAttribute{
																Computed:    true,
																Description: `The entity's foreign key column. If the column is nested inside the entity's structure and not at the top level, use ` + "`" + `NestedForeignKeyColumn` + "`" + ` instead.`,
															},
															"referenced_column_name": schema.StringAttribute{
																Computed:    true,
																Description: `The column name of the referenced entity.`,
															},
															"referenced_entity_id": schema.StringAttribute{
																Computed:    true,
																Description: `The id of the referenced entity.`,
															},
														},
													},
												},
											},
											Description: `The foreign columns of the entity.`,
										},
										"primary_key_columns": schema.ListAttribute{
											Computed:    true,
											ElementType: types.StringType,
										},
										"strategy": schema.SingleNestedAttribute{
											Computed: true,
											Attributes: map[string]schema.Attribute{
												"extracted_data": schema.SingleNestedAttribute{
													Computed: true,
													Attributes: map[string]schema.Attribute{
														"begin_time_parameter": schema.SingleNestedAttribute{
															Computed: true,
															Attributes: map[string]schema.Attribute{
																"format": schema.StringAttribute{
																	Computed:    true,
																	Description: `must be one of ["yyyy-MM-ddTHH:mm:ssX", "yyyy-MM-ddTHH:mm:ssZ", "yyyy-MM-dd"]`,
																},
																"key": schema.StringAttribute{
																	Computed: true,
																},
																"value": schema.StringAttribute{
																	Computed: true,
																},
															},
														},
														"end_time_parameter": schema.SingleNestedAttribute{
															Computed: true,
															Attributes: map[string]schema.Attribute{
																"format": schema.StringAttribute{
																	Computed:    true,
																	Description: `must be one of ["yyyy-MM-ddTHH:mm:ssX", "yyyy-MM-ddTHH:mm:ssZ", "yyyy-MM-dd"]`,
																},
																"key": schema.StringAttribute{
																	Computed: true,
																},
																"value": schema.StringAttribute{
																	Computed: true,
																},
															},
														},
														"high_watermark_query_parameters": schema.ListNestedAttribute{
															Computed: true,
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Computed: true,
																	},
																	"value": schema.StringAttribute{
																		Computed: true,
																	},
																},
															},
															Description: `Defines the query parameters to be included when fetching the most recently updated record. E.g., [{"key": "sort", "value": "updated_at"}, {"key": "order", "value": "desc"}]`,
														},
														"last_updated_column": schema.StringAttribute{
															Computed: true,
														},
														"type": schema.StringAttribute{
															Computed:    true,
															Description: `must be one of ["EXTRACTED_DATA"]`,
														},
													},
												},
												"wall_clock": schema.SingleNestedAttribute{
													Computed: true,
													Attributes: map[string]schema.Attribute{
														"begin_date_parameter": schema.SingleNestedAttribute{
															Computed: true,
															Attributes: map[string]schema.Attribute{
																"format": schema.StringAttribute{
																	Computed:    true,
																	Description: `must be one of ["yyyy-MM-dd"]`,
																},
																"key": schema.StringAttribute{
																	Computed: true,
																},
																"value": schema.StringAttribute{
																	Computed: true,
																},
															},
														},
														"end_date_parameter": schema.SingleNestedAttribute{
															Computed: true,
															Attributes: map[string]schema.Attribute{
																"format": schema.StringAttribute{
																	Computed:    true,
																	Description: `must be one of ["yyyy-MM-dd"]`,
																},
																"inclusive": schema.BoolAttribute{
																	Computed: true,
																},
																"key": schema.StringAttribute{
																	Computed: true,
																},
																"value": schema.StringAttribute{
																	Computed: true,
																},
															},
														},
														"lookback_window": schema.Int64Attribute{
															Computed:    true,
															Description: `The number of days we look back during each delta extraction.`,
														},
														"type": schema.StringAttribute{
															Computed:    true,
															Description: `must be one of ["WALL_CLOCK"]`,
														},
													},
												},
											},
										},
										"type": schema.StringAttribute{
											Computed:    true,
											Description: `must be one of ["UPDATE"]`,
										},
									},
								},
							},
							Description: `Can be either the string ` + "`" + `REPLACE` + "`" + ` or one of the supported objects.`,
						},
						"query_parameters": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Computed: true,
									},
									"value": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							Description: `A list of query parameters to be passed with all the requests.`,
						},
						"rest_method": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"one": schema.StringAttribute{
									Computed:    true,
									Description: `must be one of ["GET"]`,
								},
								"post_method": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"body_parameters": schema.ListNestedAttribute{
											Computed: true,
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Computed: true,
													},
													"value": schema.StringAttribute{
														Computed: true,
													},
												},
											},
											Description: `A list of body parameters to be passed with all the requests.`,
										},
										"type": schema.StringAttribute{
											Computed:    true,
											Description: `must be one of ["POST"]`,
										},
									},
								},
							},
							Description: `Can be either the string ` + "`" + `GET` + "`" + ` or ` + "`" + `POST` + "`" + `, which includes body parameters.`,
						},
					},
				},
				Description: `The list of entities.`,
			},
			"id": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: `The unique name of this connection.`,
			},
			"status": schema.StringAttribute{
				Computed:    true,
				Description: `The current status of the connection. must be one of ["UNKNOWN", "UP", "DOWN", "RESIZE", "MAINTENANCE", "QUOTA", "CREATING"]`,
			},
			"type": schema.StringAttribute{
				Computed:    true,
				Description: `must be one of ["USER_DEFINED_API"]`,
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
							},
						},
						Description: `The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information.`,
					},
				},
				Description: `The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.`,
			},
		},
	}
}

func (r *ConnectionUSERDEFINEDAPIDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *ConnectionUSERDEFINEDAPIDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *ConnectionUSERDEFINEDAPIDataSourceModel
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
	request := operations.GetUSERDEFINEDAPIConnectionRequest{
		ID: id,
	}
	res, err := r.client.Connection.GetUSERDEFINEDAPIConnection(ctx, request)
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
	if res.ConnectionUserDefinedAPI == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedConnectionUserDefinedAPI(res.ConnectionUserDefinedAPI)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
