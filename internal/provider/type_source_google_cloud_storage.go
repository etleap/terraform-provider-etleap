// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceGoogleCloudStorage struct {
	ConnectionID     types.String   `tfsdk:"connection_id"`
	FileNameFilter   types.String   `tfsdk:"file_name_filter"`
	GlobPattern      types.String   `tfsdk:"glob_pattern"`
	LatencyThreshold types.Int64    `tfsdk:"latency_threshold"`
	LowWatermark     types.String   `tfsdk:"low_watermark"`
	NewFileBehavior  types.String   `tfsdk:"new_file_behavior"`
	Paths            []types.String `tfsdk:"paths"`
	Type             types.String   `tfsdk:"type"`
}
