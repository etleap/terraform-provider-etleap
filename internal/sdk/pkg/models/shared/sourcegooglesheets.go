// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceGoogleSheetsType string

const (
	SourceGoogleSheetsTypeGoogleSheets SourceGoogleSheetsType = "GOOGLE_SHEETS"
)

func (e SourceGoogleSheetsType) ToPointer() *SourceGoogleSheetsType {
	return &e
}

func (e *SourceGoogleSheetsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "GOOGLE_SHEETS":
		*e = SourceGoogleSheetsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGoogleSheetsType: %v", v)
	}
}

type SourceGoogleSheets struct {
	// The universally unique identifier for the source.
	ConnectionID string                 `json:"connectionId"`
	Type         SourceGoogleSheetsType `json:"type"`
	// Notify if we can't extract for `x` hours. Setting it to `null` disables the notification. Defaults to `null`.
	LatencyThreshold *int64 `json:"latencyThreshold,omitempty"`
	// Google Sheets entities are in the form `SpreadsheetID/SheetID`. You can find both values by clicking on the sheet (tab) you want and looking at the URL: docs.google.com/spreadsheets/d/`1pRAGMSRpEEG31kbtG2qcpr-HDeDfvafp_v00`/edit#gid=`642381756`
	Entity string `json:"entity"`
}

func (o *SourceGoogleSheets) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *SourceGoogleSheets) GetType() SourceGoogleSheetsType {
	if o == nil {
		return SourceGoogleSheetsType("")
	}
	return o.Type
}

func (o *SourceGoogleSheets) GetLatencyThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LatencyThreshold
}

func (o *SourceGoogleSheets) GetEntity() string {
	if o == nil {
		return ""
	}
	return o.Entity
}
