// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type MysqlShard struct {
	Address   types.String `tfsdk:"address"`
	Database  types.String `tfsdk:"database"`
	Password  types.String `tfsdk:"password"`
	Port      types.Int64  `tfsdk:"port"`
	ShardID   types.String `tfsdk:"shard_id"`
	SSHConfig *SSHConfig   `tfsdk:"ssh_config"`
	Username  types.String `tfsdk:"username"`
}
