// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type TopLevelForeignKeyColumn struct {
	ColumnName           types.String `tfsdk:"column_name"`
	ReferencedColumnName types.String `tfsdk:"referenced_column_name"`
	ReferencedEntityID   types.String `tfsdk:"referenced_entity_id"`
}