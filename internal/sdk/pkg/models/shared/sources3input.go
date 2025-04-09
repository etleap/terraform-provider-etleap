// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/types"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
)

type SourceS3InputType string

const (
	SourceS3InputTypeS3Input SourceS3InputType = "S3_INPUT"
)

func (e SourceS3InputType) ToPointer() *SourceS3InputType {
	return &e
}

func (e *SourceS3InputType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "S3_INPUT":
		*e = SourceS3InputType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceS3InputType: %v", v)
	}
}

// SourceS3InputNewFileBehavior - Specifies whether new files update, add to or replace existing files. See <a target="_blank" href="https://docs.etleap.com/docs/documentation/ZG9jOjI0NTQwNzI2-create-a-file-based-pipeline#update-method">the documentation</a> for more details.
type SourceS3InputNewFileBehavior string

const (
	SourceS3InputNewFileBehaviorUpdate  SourceS3InputNewFileBehavior = "UPDATE"
	SourceS3InputNewFileBehaviorAppend  SourceS3InputNewFileBehavior = "APPEND"
	SourceS3InputNewFileBehaviorReplace SourceS3InputNewFileBehavior = "REPLACE"
)

func (e SourceS3InputNewFileBehavior) ToPointer() *SourceS3InputNewFileBehavior {
	return &e
}

func (e *SourceS3InputNewFileBehavior) UnmarshalJSON(data []byte) error {
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
		*e = SourceS3InputNewFileBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceS3InputNewFileBehavior: %v", v)
	}
}

type SourceS3Input struct {
	Type SourceS3InputType `json:"type"`
	// The universally unique identifier for the source.
	ConnectionID string `json:"connectionId"`
	// Notify if we can't extract for `x` hours. Setting it to `null` disables the notification. Defaults to `null`.
	LatencyThreshold *int64 `json:"latencyThreshold,omitempty"`
	// Regular expression matching the names of the files to be processed by this pipeline. A single value for `paths` is required when `fileNameFilter` is specified.
	FileNameFilter *string `json:"fileNameFilter,omitempty"`
	// Specifies whether new files update, add to or replace existing files. See <a target="_blank" href="https://docs.etleap.com/docs/documentation/ZG9jOjI0NTQwNzI2-create-a-file-based-pipeline#update-method">the documentation</a> for more details.
	NewFileBehavior SourceS3InputNewFileBehavior `json:"newFileBehavior"`
	// Timestamp of the earliest modified file that should be processed by the pipeline. Only the files modified after this timestamp will be processed. Format of the timestamp: 'yyyy-MM-dd'.
	LowWatermark *types.Date `json:"lowWatermark,omitempty"`
	// File or folder paths for the files to be extracted from the source. In the case when `fileNameFilter` is specified exactly one folder path must be given here.
	Paths []string `json:"paths"`
	// Etleap can check whether files that were already processed have changed. If the file has changed, then Etleap fetches the new file and removes the old file's data in the destination and adds the changed data. <br> This can only be enabled when `newFileBehavior` is set to `APPEND`. Defaults to `false`.
	FilesCanChange *bool `json:"filesCanChange,omitempty"`
	// Whether this source should be triggered by a `Batch Added` event (`true`) or Etleap should inspect the source to find new files to process (`false`). Defaults to `false`.
	TriggeredByEvent *bool `json:"triggeredByEvent,omitempty"`
}

func (s SourceS3Input) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceS3Input) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceS3Input) GetType() SourceS3InputType {
	if o == nil {
		return SourceS3InputType("")
	}
	return o.Type
}

func (o *SourceS3Input) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *SourceS3Input) GetLatencyThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LatencyThreshold
}

func (o *SourceS3Input) GetFileNameFilter() *string {
	if o == nil {
		return nil
	}
	return o.FileNameFilter
}

func (o *SourceS3Input) GetNewFileBehavior() SourceS3InputNewFileBehavior {
	if o == nil {
		return SourceS3InputNewFileBehavior("")
	}
	return o.NewFileBehavior
}

func (o *SourceS3Input) GetLowWatermark() *types.Date {
	if o == nil {
		return nil
	}
	return o.LowWatermark
}

func (o *SourceS3Input) GetPaths() []string {
	if o == nil {
		return []string{}
	}
	return o.Paths
}

func (o *SourceS3Input) GetFilesCanChange() *bool {
	if o == nil {
		return nil
	}
	return o.FilesCanChange
}

func (o *SourceS3Input) GetTriggeredByEvent() *bool {
	if o == nil {
		return nil
	}
	return o.TriggeredByEvent
}
