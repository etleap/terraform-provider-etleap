// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
)

// One - Note: JSON_OBJECT is deprecated and is a synonym of JSON.
type One string

const (
	OneAuto       One = "AUTO"
	OneBigint     One = "BIGINT"
	OneBoolean    One = "BOOLEAN"
	OneDate       One = "DATE"
	OneDatetime   One = "DATETIME"
	OneDouble     One = "DOUBLE"
	OneString     One = "STRING"
	OneJSONObject One = "JSON_OBJECT"
	OneJSON       One = "JSON"
)

func (e One) ToPointer() *One {
	return &e
}

func (e *One) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "AUTO":
		fallthrough
	case "BIGINT":
		fallthrough
	case "BOOLEAN":
		fallthrough
	case "DATE":
		fallthrough
	case "DATETIME":
		fallthrough
	case "DOUBLE":
		fallthrough
	case "STRING":
		fallthrough
	case "JSON_OBJECT":
		fallthrough
	case "JSON":
		*e = One(v)
		return nil
	default:
		return fmt.Errorf("invalid value for One: %v", v)
	}
}

type TypeType string

const (
	TypeTypeOne                     TypeType = "1"
	TypeTypeTypeDecimal             TypeType = "type_decimal"
	TypeTypeTypeStringWithMaxLength TypeType = "type_string_with_max_length"
)

type Type struct {
	One                     *One
	TypeDecimal             *TypeDecimal
	TypeStringWithMaxLength *TypeStringWithMaxLength

	Type TypeType
}

func CreateTypeOne(one One) Type {
	typ := TypeTypeOne

	return Type{
		One:  &one,
		Type: typ,
	}
}

func CreateTypeTypeDecimal(typeDecimal TypeDecimal) Type {
	typ := TypeTypeTypeDecimal

	return Type{
		TypeDecimal: &typeDecimal,
		Type:        typ,
	}
}

func CreateTypeTypeStringWithMaxLength(typeStringWithMaxLength TypeStringWithMaxLength) Type {
	typ := TypeTypeTypeStringWithMaxLength

	return Type{
		TypeStringWithMaxLength: &typeStringWithMaxLength,
		Type:                    typ,
	}
}

func (u *Type) UnmarshalJSON(data []byte) error {

	typeStringWithMaxLength := new(TypeStringWithMaxLength)
	if err := utils.UnmarshalJSON(data, &typeStringWithMaxLength, "", true, true); err == nil {
		u.TypeStringWithMaxLength = typeStringWithMaxLength
		u.Type = TypeTypeTypeStringWithMaxLength
		return nil
	}

	typeDecimal := new(TypeDecimal)
	if err := utils.UnmarshalJSON(data, &typeDecimal, "", true, true); err == nil {
		u.TypeDecimal = typeDecimal
		u.Type = TypeTypeTypeDecimal
		return nil
	}

	one := new(One)
	if err := utils.UnmarshalJSON(data, &one, "", true, true); err == nil {
		u.One = one
		u.Type = TypeTypeOne
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u Type) MarshalJSON() ([]byte, error) {
	if u.One != nil {
		return utils.MarshalJSON(u.One, "", true)
	}

	if u.TypeDecimal != nil {
		return utils.MarshalJSON(u.TypeDecimal, "", true)
	}

	if u.TypeStringWithMaxLength != nil {
		return utils.MarshalJSON(u.TypeStringWithMaxLength, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}
