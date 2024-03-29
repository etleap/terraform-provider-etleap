// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"errors"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
)

type ForeignColumn struct {
	// Represents the path in the entity schema where the column is located. If the column is at the top level of an entity, use `Top Level Foreign Key Column` instead.
	ColumnPath []string `json:"columnPath"`
	// The id of the referenced entity.
	ReferencedEntityID string `json:"referencedEntityId"`
	// The column name of the referenced entity.
	ReferencedColumnName string `json:"referencedColumnName"`
}

func (o *ForeignColumn) GetColumnPath() []string {
	if o == nil {
		return []string{}
	}
	return o.ColumnPath
}

func (o *ForeignColumn) GetReferencedEntityID() string {
	if o == nil {
		return ""
	}
	return o.ReferencedEntityID
}

func (o *ForeignColumn) GetReferencedColumnName() string {
	if o == nil {
		return ""
	}
	return o.ReferencedColumnName
}

type TopLevelForeignKeyColumn struct {
	// The entity's foreign key column. If the column is nested inside the entity's structure and not at the top level, use `NestedForeignKeyColumn` instead.
	ColumnName string `json:"columnName"`
	// The id of the referenced entity.
	ReferencedEntityID string `json:"referencedEntityId"`
	// The column name of the referenced entity.
	ReferencedColumnName string `json:"referencedColumnName"`
}

func (o *TopLevelForeignKeyColumn) GetColumnName() string {
	if o == nil {
		return ""
	}
	return o.ColumnName
}

func (o *TopLevelForeignKeyColumn) GetReferencedEntityID() string {
	if o == nil {
		return ""
	}
	return o.ReferencedEntityID
}

func (o *TopLevelForeignKeyColumn) GetReferencedColumnName() string {
	if o == nil {
		return ""
	}
	return o.ReferencedColumnName
}

type SchemaV1ForeignKeyColumnType string

const (
	SchemaV1ForeignKeyColumnTypeTopLevelForeignKeyColumn SchemaV1ForeignKeyColumnType = "Top Level Foreign Key Column"
	SchemaV1ForeignKeyColumnTypeForeignColumn            SchemaV1ForeignKeyColumnType = "Foreign Column"
)

type SchemaV1ForeignKeyColumn struct {
	TopLevelForeignKeyColumn *TopLevelForeignKeyColumn
	ForeignColumn            *ForeignColumn

	Type SchemaV1ForeignKeyColumnType
}

func CreateSchemaV1ForeignKeyColumnTopLevelForeignKeyColumn(topLevelForeignKeyColumn TopLevelForeignKeyColumn) SchemaV1ForeignKeyColumn {
	typ := SchemaV1ForeignKeyColumnTypeTopLevelForeignKeyColumn

	return SchemaV1ForeignKeyColumn{
		TopLevelForeignKeyColumn: &topLevelForeignKeyColumn,
		Type:                     typ,
	}
}

func CreateSchemaV1ForeignKeyColumnForeignColumn(foreignColumn ForeignColumn) SchemaV1ForeignKeyColumn {
	typ := SchemaV1ForeignKeyColumnTypeForeignColumn

	return SchemaV1ForeignKeyColumn{
		ForeignColumn: &foreignColumn,
		Type:          typ,
	}
}

func (u *SchemaV1ForeignKeyColumn) UnmarshalJSON(data []byte) error {

	topLevelForeignKeyColumn := new(TopLevelForeignKeyColumn)
	if err := utils.UnmarshalJSON(data, &topLevelForeignKeyColumn, "", true, true); err == nil {
		u.TopLevelForeignKeyColumn = topLevelForeignKeyColumn
		u.Type = SchemaV1ForeignKeyColumnTypeTopLevelForeignKeyColumn
		return nil
	}

	foreignColumn := new(ForeignColumn)
	if err := utils.UnmarshalJSON(data, &foreignColumn, "", true, true); err == nil {
		u.ForeignColumn = foreignColumn
		u.Type = SchemaV1ForeignKeyColumnTypeForeignColumn
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SchemaV1ForeignKeyColumn) MarshalJSON() ([]byte, error) {
	if u.TopLevelForeignKeyColumn != nil {
		return utils.MarshalJSON(u.TopLevelForeignKeyColumn, "", true)
	}

	if u.ForeignColumn != nil {
		return utils.MarshalJSON(u.ForeignColumn, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}
