// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
)

type ExtendedNumberTypeName string

const (
	ExtendedNumberTypeNameNumber ExtendedNumberTypeName = "NUMBER"
)

func (e ExtendedNumberTypeName) ToPointer() *ExtendedNumberTypeName {
	return &e
}

func (e *ExtendedNumberTypeName) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NUMBER":
		*e = ExtendedNumberTypeName(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ExtendedNumberTypeName: %v", v)
	}
}

type ExtendedNumberType struct {
	Name      *ExtendedNumberTypeName `default:"NUMBER" json:"name"`
	Scale     float64                 `json:"scale"`
	Precision float64                 `json:"precision"`
}

func (e ExtendedNumberType) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(e, "", false)
}

func (e *ExtendedNumberType) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &e, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *ExtendedNumberType) GetName() *ExtendedNumberTypeName {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ExtendedNumberType) GetScale() float64 {
	if o == nil {
		return 0.0
	}
	return o.Scale
}

func (o *ExtendedNumberType) GetPrecision() float64 {
	if o == nil {
		return 0.0
	}
	return o.Precision
}
