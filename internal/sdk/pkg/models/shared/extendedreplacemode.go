// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
)

type ExtendedReplaceModeType string

const (
	ExtendedReplaceModeTypeReplace ExtendedReplaceModeType = "REPLACE"
)

func (e ExtendedReplaceModeType) ToPointer() *ExtendedReplaceModeType {
	return &e
}

func (e *ExtendedReplaceModeType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "REPLACE":
		*e = ExtendedReplaceModeType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ExtendedReplaceModeType: %v", v)
	}
}

type ExtendedReplaceMode struct {
	Type              *ExtendedReplaceModeType `default:"REPLACE" json:"type"`
	PrimaryKeyColumns []string                 `json:"primaryKeyColumns"`
	// The foreign columns of the entity.
	ForeignKeyColumns []ForeignKeyColumn `json:"foreignKeyColumns,omitempty"`
}

func (e ExtendedReplaceMode) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(e, "", false)
}

func (e *ExtendedReplaceMode) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &e, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *ExtendedReplaceMode) GetType() *ExtendedReplaceModeType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ExtendedReplaceMode) GetPrimaryKeyColumns() []string {
	if o == nil {
		return []string{}
	}
	return o.PrimaryKeyColumns
}

func (o *ExtendedReplaceMode) GetForeignKeyColumns() []ForeignKeyColumn {
	if o == nil {
		return nil
	}
	return o.ForeignKeyColumns
}
