// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/types"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
)

type SourceGoogleCloudStorageType string

const (
	SourceGoogleCloudStorageTypeGoogleCloudStorage SourceGoogleCloudStorageType = "GOOGLE_CLOUD_STORAGE"
)

func (e SourceGoogleCloudStorageType) ToPointer() *SourceGoogleCloudStorageType {
	return &e
}

func (e *SourceGoogleCloudStorageType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "GOOGLE_CLOUD_STORAGE":
		*e = SourceGoogleCloudStorageType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGoogleCloudStorageType: %v", v)
	}
}

// SourceGoogleCloudStorageNewFileBehavior - Specifies whether new files update, add to or replace existing files. See <a target="_blank" href="https://docs.etleap.com/docs/documentation/ZG9jOjI0NTQwNzI2-create-a-file-based-pipeline#update-method">the documentation</a> for more details.
type SourceGoogleCloudStorageNewFileBehavior string

const (
	SourceGoogleCloudStorageNewFileBehaviorUpdate  SourceGoogleCloudStorageNewFileBehavior = "UPDATE"
	SourceGoogleCloudStorageNewFileBehaviorAppend  SourceGoogleCloudStorageNewFileBehavior = "APPEND"
	SourceGoogleCloudStorageNewFileBehaviorReplace SourceGoogleCloudStorageNewFileBehavior = "REPLACE"
)

func (e SourceGoogleCloudStorageNewFileBehavior) ToPointer() *SourceGoogleCloudStorageNewFileBehavior {
	return &e
}

func (e *SourceGoogleCloudStorageNewFileBehavior) UnmarshalJSON(data []byte) error {
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
		*e = SourceGoogleCloudStorageNewFileBehavior(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGoogleCloudStorageNewFileBehavior: %v", v)
	}
}

type SourceGoogleCloudStorage struct {
	// The universally unique identifier for the source.
	ConnectionID string                       `json:"connectionId"`
	Type         SourceGoogleCloudStorageType `json:"type"`
	// Notify if we can't extract for `x` hours. Setting it to `null` disables the notification. Defaults to `null`.
	LatencyThreshold *int64 `json:"latencyThreshold,omitempty"`
	// Specifies whether new files update, add to or replace existing files. See <a target="_blank" href="https://docs.etleap.com/docs/documentation/ZG9jOjI0NTQwNzI2-create-a-file-based-pipeline#update-method">the documentation</a> for more details.
	NewFileBehavior SourceGoogleCloudStorageNewFileBehavior `json:"newFileBehavior"`
	// Regular expression matching the names of the files to be processed by this pipeline. A single value for `paths` is required when `fileNameFilter` is specified.
	FileNameFilter *string `json:"fileNameFilter,omitempty"`
	// Timestamp of the earliest modified file that should be processed by the pipeline. Only the files modified after this timestamp will be processed. Format of the timestamp: 'yyyy-MM-dd'.
	LowWatermark *types.Date `json:"lowWatermark,omitempty"`
	// File or folder paths for the files to be extracted from the source. In the case when `fileNameFilter` is specified exactly one folder path must be given here. `paths` can't be used when a `globPattern` is specified.
	Paths []string `json:"paths,omitempty"`
	// A glob pattern to be used as a path. Either `globPattern` or `paths` must be specified, but not both.
	GlobPattern *string `json:"globPattern,omitempty"`
}

func (s SourceGoogleCloudStorage) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceGoogleCloudStorage) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceGoogleCloudStorage) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *SourceGoogleCloudStorage) GetType() SourceGoogleCloudStorageType {
	if o == nil {
		return SourceGoogleCloudStorageType("")
	}
	return o.Type
}

func (o *SourceGoogleCloudStorage) GetLatencyThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LatencyThreshold
}

func (o *SourceGoogleCloudStorage) GetNewFileBehavior() SourceGoogleCloudStorageNewFileBehavior {
	if o == nil {
		return SourceGoogleCloudStorageNewFileBehavior("")
	}
	return o.NewFileBehavior
}

func (o *SourceGoogleCloudStorage) GetFileNameFilter() *string {
	if o == nil {
		return nil
	}
	return o.FileNameFilter
}

func (o *SourceGoogleCloudStorage) GetLowWatermark() *types.Date {
	if o == nil {
		return nil
	}
	return o.LowWatermark
}

func (o *SourceGoogleCloudStorage) GetPaths() []string {
	if o == nil {
		return nil
	}
	return o.Paths
}

func (o *SourceGoogleCloudStorage) GetGlobPattern() *string {
	if o == nil {
		return nil
	}
	return o.GlobPattern
}
