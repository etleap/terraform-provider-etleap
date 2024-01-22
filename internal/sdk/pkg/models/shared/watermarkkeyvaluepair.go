// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
)

type Format string

const (
	FormatYyyyMmDdThhMmSsX Format = "yyyy-MM-ddTHH:mm:ssX"
	FormatYyyyMmDdThhMmSsZ Format = "yyyy-MM-ddTHH:mm:ssZ"
	FormatYyyyMmDd         Format = "yyyy-MM-dd"
)

func (e Format) ToPointer() *Format {
	return &e
}

func (e *Format) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "yyyy-MM-ddTHH:mm:ssX":
		fallthrough
	case "yyyy-MM-ddTHH:mm:ssZ":
		fallthrough
	case "yyyy-MM-dd":
		*e = Format(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Format: %v", v)
	}
}

type WatermarkKeyValuePair struct {
	Key    string  `json:"key"`
	Value  string  `json:"value"`
	Format *Format `default:"yyyy-MM-ddTHH:mm:ssX" json:"format"`
}

func (w WatermarkKeyValuePair) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WatermarkKeyValuePair) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *WatermarkKeyValuePair) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

func (o *WatermarkKeyValuePair) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

func (o *WatermarkKeyValuePair) GetFormat() *Format {
	if o == nil {
		return nil
	}
	return o.Format
}
