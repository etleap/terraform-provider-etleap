// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type Type struct {
	One                     types.String             `tfsdk:"one"`
	TypeDecimal             *TypeDecimal             `tfsdk:"type_decimal"`
	TypeStringWithMaxLength *TypeStringWithMaxLength `tfsdk:"type_string_with_max_length"`
}
