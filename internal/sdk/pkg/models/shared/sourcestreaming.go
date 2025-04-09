// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/types"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
)

type SourceStreamingType string

const (
	SourceStreamingTypeStreaming SourceStreamingType = "STREAMING"
)

func (e SourceStreamingType) ToPointer() *SourceStreamingType {
	return &e
}

func (e *SourceStreamingType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "STREAMING":
		*e = SourceStreamingType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceStreamingType: %v", v)
	}
}

// SourceStreamingNewFileBehavior - Specifies whether new files update, add to or replace existing files. See <a target="_blank" href="https://docs.etleap.com/docs/documentation/ZG9jOjI0NTQwNzI2-create-a-file-based-pipeline#update-method">the documentation</a> for more details.
type SourceStreamingNewFileBehavior string

const (
	SourceStreamingNewFileBehaviorUpdate  SourceStreamingNewFileBehavior = "UPDATE"
	SourceStreamingNewFileBehaviorAppend  SourceStreamingNewFileBehavior = "APPEND"
	SourceStreamingNewFileBehaviorReplace SourceStreamingNewFileBehavior = "REPLACE"
)

func (e SourceStreamingNewFileBehavior) ToPointer() *SourceStreamingNewFileBehavior {
	return &e
}

func (e *SourceStreamingNewFileBehavior) UnmarshalJSON(data []byte) error {
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
		*e = SourceStreamingNewFileBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceStreamingNewFileBehavior: %v", v)
	}
}

type SourceStreaming struct {
	// The universally unique identifier for the source.
	ConnectionID string              `json:"connectionId"`
	Type         SourceStreamingType `json:"type"`
	// Notify if we can't extract for `x` hours. Setting it to `null` disables the notification. Defaults to `null`.
	LatencyThreshold *int64 `json:"latencyThreshold,omitempty"`
	// Specifies whether new files update, add to or replace existing files. See <a target="_blank" href="https://docs.etleap.com/docs/documentation/ZG9jOjI0NTQwNzI2-create-a-file-based-pipeline#update-method">the documentation</a> for more details.
	NewFileBehavior SourceStreamingNewFileBehavior `json:"newFileBehavior"`
	// Regular expression matching the names of the files to be processed by this pipeline. A single value for `paths` is required when `fileNameFilter` is specified.
	FileNameFilter *string `json:"fileNameFilter,omitempty"`
	// Timestamp of the earliest modified file that should be processed by the pipeline. Only the files modified after this timestamp will be processed. Format of the timestamp: 'yyyy-MM-dd'.
	LowWatermark *types.Date `json:"lowWatermark,omitempty"`
	// File or folder paths for the files to be extracted from the source. In the case when `fileNameFilter` is specified exactly one folder path must be given here.
	Paths []string `json:"paths"`
}

func (s SourceStreaming) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceStreaming) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceStreaming) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *SourceStreaming) GetType() SourceStreamingType {
	if o == nil {
		return SourceStreamingType("")
	}
	return o.Type
}

func (o *SourceStreaming) GetLatencyThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LatencyThreshold
}

func (o *SourceStreaming) GetNewFileBehavior() SourceStreamingNewFileBehavior {
	if o == nil {
		return SourceStreamingNewFileBehavior("")
	}
	return o.NewFileBehavior
}

func (o *SourceStreaming) GetFileNameFilter() *string {
	if o == nil {
		return nil
	}
	return o.FileNameFilter
}

func (o *SourceStreaming) GetLowWatermark() *types.Date {
	if o == nil {
		return nil
	}
	return o.LowWatermark
}

func (o *SourceStreaming) GetPaths() []string {
	if o == nil {
		return []string{}
	}
	return o.Paths
}
