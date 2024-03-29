// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationRedshift struct {
	AutomaticSchemaChanges types.Bool         `tfsdk:"automatic_schema_changes"`
	CompressColumns        types.Bool         `tfsdk:"compress_columns"`
	ConnectionID           types.String       `tfsdk:"connection_id"`
	DistributionStyle      *DistributionStyle `tfsdk:"distribution_style"`
	LastUpdatedColumn      types.String       `tfsdk:"last_updated_column"`
	PrimaryKey             []types.String     `tfsdk:"primary_key"`
	RetainHistory          types.Bool         `tfsdk:"retain_history"`
	Schema                 types.String       `tfsdk:"schema"`
	SortColumns            []types.String     `tfsdk:"sort_columns"`
	Table                  types.String       `tfsdk:"table"`
	TruncateStrings        types.Bool         `tfsdk:"truncate_strings"`
	Type                   types.String       `tfsdk:"type"`
	WaitForQualityCheck    types.Bool         `tfsdk:"wait_for_quality_check"`
}
