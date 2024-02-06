// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SnowflakeShard struct {
	Address   types.String `tfsdk:"address"`
	Database  types.String `tfsdk:"database"`
	Password  types.String `tfsdk:"password"`
	Role      types.String `tfsdk:"role"`
	ShardID   types.String `tfsdk:"shard_id"`
	Username  types.String `tfsdk:"username"`
	Warehouse types.String `tfsdk:"warehouse"`
}
