// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type QueryAndTriggers struct {
	Query    types.String   `tfsdk:"query"`
	Triggers []types.String `tfsdk:"triggers"`
}
