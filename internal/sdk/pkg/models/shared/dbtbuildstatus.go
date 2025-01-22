// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type DbtBuildStatus string

const (
	DbtBuildStatusInProgress             DbtBuildStatus = "IN_PROGRESS"
	DbtBuildStatusEtleapError            DbtBuildStatus = "ETLEAP_ERROR"
	DbtBuildStatusDbtError               DbtBuildStatus = "DBT_ERROR"
	DbtBuildStatusSuccessWithDbtWarnings DbtBuildStatus = "SUCCESS_WITH_DBT_WARNINGS"
	DbtBuildStatusSuccess                DbtBuildStatus = "SUCCESS"
)

func (e DbtBuildStatus) ToPointer() *DbtBuildStatus {
	return &e
}

func (e *DbtBuildStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "IN_PROGRESS":
		fallthrough
	case "ETLEAP_ERROR":
		fallthrough
	case "DBT_ERROR":
		fallthrough
	case "SUCCESS_WITH_DBT_WARNINGS":
		fallthrough
	case "SUCCESS":
		*e = DbtBuildStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DbtBuildStatus: %v", v)
	}
}
