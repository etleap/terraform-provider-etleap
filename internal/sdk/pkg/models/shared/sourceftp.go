// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/types"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
)

type SourceFtpType string

const (
	SourceFtpTypeFtp SourceFtpType = "FTP"
)

func (e SourceFtpType) ToPointer() *SourceFtpType {
	return &e
}

func (e *SourceFtpType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "FTP":
		*e = SourceFtpType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceFtpType: %v", v)
	}
}

// NewFileBehavior - Specifies whether new files update, add to or replace existing files. See <a target="_blank" href="https://docs.etleap.com/docs/documentation/ZG9jOjI0NTQwNzI2-create-a-file-based-pipeline#update-method">the documentation</a> for more details.
type NewFileBehavior string

const (
	NewFileBehaviorUpdate  NewFileBehavior = "UPDATE"
	NewFileBehaviorAppend  NewFileBehavior = "APPEND"
	NewFileBehaviorReplace NewFileBehavior = "REPLACE"
)

func (e NewFileBehavior) ToPointer() *NewFileBehavior {
	return &e
}

func (e *NewFileBehavior) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "UPDATE":
		fallthrough
	case "APPEND":
		fallthrough
	case "REPLACE":
		*e = NewFileBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for NewFileBehavior: %v", v)
	}
}

type SourceFtp struct {
	Type SourceFtpType `json:"type"`
	// The universally unique identifier for the source.
	ConnectionID string `json:"connectionId"`
	// Notify if we can't extract for `x` hours. Setting it to `null` disables the notification. Defaults to `null`.
	LatencyThreshold *int64 `json:"latencyThreshold,omitempty"`
	// File or folder paths for the files to be extracted from the source. In the case when `fileNameFilter` is specified exactly one folder path must be given here.
	Paths []string `json:"paths"`
	// Regular expression matching the names of the files to be processed by this pipeline. `fileNameFilter` or `paths` must be specified.
	FileNameFilter *string `json:"fileNameFilter,omitempty"`
	// Specifies whether new files update, add to or replace existing files. See <a target="_blank" href="https://docs.etleap.com/docs/documentation/ZG9jOjI0NTQwNzI2-create-a-file-based-pipeline#update-method">the documentation</a> for more details.
	NewFileBehavior NewFileBehavior `json:"newFileBehavior"`
	// Timestamp of the earliest modified file that should be processed by the pipeline. Only the files modified after this timestamp will be processed. Format of the timestamp: 'yyyy-MM-dd'.
	LowWatermark *types.Date `json:"lowWatermark,omitempty"`
}

func (s SourceFtp) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceFtp) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceFtp) GetType() SourceFtpType {
	if o == nil {
		return SourceFtpType("")
	}
	return o.Type
}

func (o *SourceFtp) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *SourceFtp) GetLatencyThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LatencyThreshold
}

func (o *SourceFtp) GetPaths() []string {
	if o == nil {
		return []string{}
	}
	return o.Paths
}

func (o *SourceFtp) GetFileNameFilter() *string {
	if o == nil {
		return nil
	}
	return o.FileNameFilter
}

func (o *SourceFtp) GetNewFileBehavior() NewFileBehavior {
	if o == nil {
		return NewFileBehavior("")
	}
	return o.NewFileBehavior
}

func (o *SourceFtp) GetLowWatermark() *types.Date {
	if o == nil {
		return nil
	}
	return o.LowWatermark
}
