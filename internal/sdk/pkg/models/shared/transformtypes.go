// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
)

type TransformTypesType string

const (
	TransformTypesTypeAddFilePath       TransformTypesType = "ADD_FILE_PATH"
	TransformTypesTypeFlattenJSONObject TransformTypesType = "FLATTEN_JSON_OBJECT"
	TransformTypesTypeParquetToRows     TransformTypesType = "PARQUET_TO_ROWS"
	TransformTypesTypeParseByRegex      TransformTypesType = "PARSE_BY_REGEX"
	TransformTypesTypeRenameColumns     TransformTypesType = "RENAME_COLUMNS"
)

type TransformTypes struct {
	TransformAddFilePath       *TransformAddFilePath
	TransformExtractJSONFields *TransformExtractJSONFields
	TransformParquetToRows     *TransformParquetToRows
	TransformParseByRegex      *TransformParseByRegex
	TransformRenameColumns     *TransformRenameColumns

	Type TransformTypesType
}

func CreateTransformTypesAddFilePath(addFilePath TransformAddFilePath) TransformTypes {
	typ := TransformTypesTypeAddFilePath

	typStr := TransformAddFilePathType(typ)
	addFilePath.Type = typStr

	return TransformTypes{
		TransformAddFilePath: &addFilePath,
		Type:                 typ,
	}
}

func CreateTransformTypesFlattenJSONObject(flattenJSONObject TransformExtractJSONFields) TransformTypes {
	typ := TransformTypesTypeFlattenJSONObject

	typStr := TransformExtractJSONFieldsType(typ)
	flattenJSONObject.Type = typStr

	return TransformTypes{
		TransformExtractJSONFields: &flattenJSONObject,
		Type:                       typ,
	}
}

func CreateTransformTypesParquetToRows(parquetToRows TransformParquetToRows) TransformTypes {
	typ := TransformTypesTypeParquetToRows

	typStr := TransformParquetToRowsType(typ)
	parquetToRows.Type = typStr

	return TransformTypes{
		TransformParquetToRows: &parquetToRows,
		Type:                   typ,
	}
}

func CreateTransformTypesParseByRegex(parseByRegex TransformParseByRegex) TransformTypes {
	typ := TransformTypesTypeParseByRegex

	typStr := TransformParseByRegexType(typ)
	parseByRegex.Type = typStr

	return TransformTypes{
		TransformParseByRegex: &parseByRegex,
		Type:                  typ,
	}
}

func CreateTransformTypesRenameColumns(renameColumns TransformRenameColumns) TransformTypes {
	typ := TransformTypesTypeRenameColumns

	typStr := TransformRenameColumnsType(typ)
	renameColumns.Type = typStr

	return TransformTypes{
		TransformRenameColumns: &renameColumns,
		Type:                   typ,
	}
}

func (u *TransformTypes) UnmarshalJSON(data []byte) error {

	type discriminator struct {
		Type string
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.Type {
	case "ADD_FILE_PATH":
		transformAddFilePath := new(TransformAddFilePath)
		if err := utils.UnmarshalJSON(data, &transformAddFilePath, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.TransformAddFilePath = transformAddFilePath
		u.Type = TransformTypesTypeAddFilePath
		return nil
	case "FLATTEN_JSON_OBJECT":
		transformExtractJSONFields := new(TransformExtractJSONFields)
		if err := utils.UnmarshalJSON(data, &transformExtractJSONFields, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.TransformExtractJSONFields = transformExtractJSONFields
		u.Type = TransformTypesTypeFlattenJSONObject
		return nil
	case "PARQUET_TO_ROWS":
		transformParquetToRows := new(TransformParquetToRows)
		if err := utils.UnmarshalJSON(data, &transformParquetToRows, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.TransformParquetToRows = transformParquetToRows
		u.Type = TransformTypesTypeParquetToRows
		return nil
	case "PARSE_BY_REGEX":
		transformParseByRegex := new(TransformParseByRegex)
		if err := utils.UnmarshalJSON(data, &transformParseByRegex, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.TransformParseByRegex = transformParseByRegex
		u.Type = TransformTypesTypeParseByRegex
		return nil
	case "RENAME_COLUMNS":
		transformRenameColumns := new(TransformRenameColumns)
		if err := utils.UnmarshalJSON(data, &transformRenameColumns, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.TransformRenameColumns = transformRenameColumns
		u.Type = TransformTypesTypeRenameColumns
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u TransformTypes) MarshalJSON() ([]byte, error) {
	if u.TransformAddFilePath != nil {
		return utils.MarshalJSON(u.TransformAddFilePath, "", true)
	}

	if u.TransformExtractJSONFields != nil {
		return utils.MarshalJSON(u.TransformExtractJSONFields, "", true)
	}

	if u.TransformParquetToRows != nil {
		return utils.MarshalJSON(u.TransformParquetToRows, "", true)
	}

	if u.TransformParseByRegex != nil {
		return utils.MarshalJSON(u.TransformParseByRegex, "", true)
	}

	if u.TransformRenameColumns != nil {
		return utils.MarshalJSON(u.TransformRenameColumns, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}
