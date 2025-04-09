// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceBigQuery struct {
	ConnectionID      types.String   `tfsdk:"connection_id"`
	Dataset           types.String   `tfsdk:"dataset"`
	LastUpdatedColumn types.String   `tfsdk:"last_updated_column"`
	LatencyThreshold  types.Int64    `tfsdk:"latency_threshold"`
	PrimaryKeyColumns []types.String `tfsdk:"primary_key_columns"`
	Table             types.String   `tfsdk:"table"`
	TableNameFilter   types.String   `tfsdk:"table_name_filter"`
}
