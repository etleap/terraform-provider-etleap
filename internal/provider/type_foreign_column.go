// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type ForeignColumn struct {
	ColumnPath           []types.String `tfsdk:"column_path"`
	ReferencedColumnName types.String   `tfsdk:"referenced_column_name"`
	ReferencedEntityID   types.String   `tfsdk:"referenced_entity_id"`
}
