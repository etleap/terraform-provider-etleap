// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceMongodb struct {
	ConnectionID     types.String `tfsdk:"connection_id"`
	LatencyThreshold types.Int64  `tfsdk:"latency_threshold"`
	Table            types.String `tfsdk:"table"`
	TableNameFilter  types.String `tfsdk:"table_name_filter"`
	Type             types.String `tfsdk:"type"`
}
