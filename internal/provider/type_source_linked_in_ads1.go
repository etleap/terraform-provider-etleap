// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceLinkedInAds1 struct {
	ConnectionID     types.String   `tfsdk:"connection_id"`
	Entity           types.String   `tfsdk:"entity"`
	LatencyThreshold types.Int64    `tfsdk:"latency_threshold"`
	Metrics          []types.String `tfsdk:"metrics"`
	Pivots           []types.String `tfsdk:"pivots"`
	Type             types.String   `tfsdk:"type"`
}
