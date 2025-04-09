// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceBingAds1 struct {
	ConnectionID     types.String   `tfsdk:"connection_id"`
	Entity           types.String   `tfsdk:"entity"`
	Fields           []types.String `tfsdk:"fields"`
	LatencyThreshold types.Int64    `tfsdk:"latency_threshold"`
	Type             types.String   `tfsdk:"type"`
}
