// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type WarehouseSnowflake struct {
	ConnectionID             types.String `tfsdk:"connection_id"`
	MaterializedView         types.Bool   `tfsdk:"materialized_view"`
	PendingRenamedTable      types.String `tfsdk:"pending_renamed_table"`
	Schema                   types.String `tfsdk:"schema"`
	Table                    types.String `tfsdk:"table"`
	Type                     types.String `tfsdk:"type"`
	WaitForUpdatePreparation types.Bool   `tfsdk:"wait_for_update_preparation"`
}
