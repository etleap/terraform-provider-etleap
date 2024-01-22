// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
)

type TimestampMillisecondsTypeType string

const (
	TimestampMillisecondsTypeTypeTimestampMilliseconds TimestampMillisecondsTypeType = "TIMESTAMP_MILLISECONDS"
)

func (e TimestampMillisecondsTypeType) ToPointer() *TimestampMillisecondsTypeType {
	return &e
}

func (e *TimestampMillisecondsTypeType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TIMESTAMP_MILLISECONDS":
		*e = TimestampMillisecondsTypeType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TimestampMillisecondsTypeType: %v", v)
	}
}

type TimestampMillisecondsType struct {
	Type *TimestampMillisecondsTypeType `default:"TIMESTAMP_MILLISECONDS" json:"type"`
}

func (t TimestampMillisecondsType) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TimestampMillisecondsType) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *TimestampMillisecondsType) GetType() *TimestampMillisecondsTypeType {
	if o == nil {
		return nil
	}
	return o.Type
}
