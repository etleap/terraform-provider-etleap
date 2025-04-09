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
var _ datasource.DataSource = &ConnectionLDAPDataSource{}
var _ datasource.DataSourceWithConfigure = &ConnectionLDAPDataSource{}

func NewConnectionLDAPDataSource() datasource.DataSource {
	return &ConnectionLDAPDataSource{}
}

// ConnectionLDAPDataSource is the data source implementation.
type ConnectionLDAPDataSource struct {
	client *sdk.SDK
}

// ConnectionLDAPDataSourceModel describes the data model.
type ConnectionLDAPDataSourceModel struct {
	Active                types.Bool              `tfsdk:"active"`
	BaseDn                types.String            `tfsdk:"base_dn"`
	CreateDate            types.String            `tfsdk:"create_date"`
	DefaultUpdateSchedule []DefaultUpdateSchedule `tfsdk:"default_update_schedule"`
	Hostname              types.String            `tfsdk:"hostname"`
	ID                    types.String            `tfsdk:"id"`
	Name                  types.String            `tfsdk:"name"`
	Pen                   types.Int64             `tfsdk:"pen"`
	Port                  types.Int64             `tfsdk:"port"`
	Status                types.String            `tfsdk:"status"`
	Type                  types.String            `tfsdk:"type"`
	UpdateSchedule        *UpdateScheduleTypes    `tfsdk:"update_schedule"`
	User                  types.String            `tfsdk:"user"`
	UseSsl                types.Bool              `tfsdk:"use_ssl"`
}

// Metadata returns the data source type name.
func (r *ConnectionLDAPDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_connection_ldap"
}

// Schema defines the schema for the data source.
func (r *ConnectionLDAPDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "ConnectionLDAP DataSource",

		Attributes: map[string]schema.Attribute{
			"active": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether this connection has been marked as active.`,
			},
			"base_dn": schema.StringAttribute{
				Computed:    true,
				Description: `The path of your DIT, typically a DC, OU, or O entry`,
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
			"hostname": schema.StringAttribute{
				Computed:    true,
				Description: `LDAP server name or ip address`,
			},
			"id": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Computed:    true,
				Description: `The unique name of this connection.`,
			},
			"pen": schema.Int64Attribute{
				Computed:    true,
				Description: `The IANA-assigned number for your custom schema entities`,
			},
			"port": schema.Int64Attribute{
				Computed: true,
			},
			"status": schema.StringAttribute{
				Computed:    true,
				Description: `The current status of the connection. must be one of ["UNKNOWN", "UP", "DOWN", "RESIZE", "MAINTENANCE", "QUOTA", "CREATING"]`,
			},
			"type": schema.StringAttribute{
				Computed:    true,
				Description: `must be one of ["LDAP"]`,
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
			"user": schema.StringAttribute{
				Computed:    true,
				Description: `The login string to use, similar to 'uid=admin,ou=system'`,
			},
			"use_ssl": schema.BoolAttribute{
				Computed:    true,
				Description: `Enable this if you are using a secure port to connect to LDAP. Usually, this should be enabled if you are connecting via port 636.`,
			},
		},
	}
}

func (r *ConnectionLDAPDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *ConnectionLDAPDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *ConnectionLDAPDataSourceModel
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
	request := operations.GetLDAPConnectionRequest{
		ID: id,
	}
	res, err := r.client.Connection.GetLDAPConnection(ctx, request)
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
	if res.ConnectionLdap == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedConnectionLdap(res.ConnectionLdap)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
