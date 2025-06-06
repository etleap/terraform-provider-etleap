// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceMongodbUpdateType string

const (
	SourceMongodbUpdateTypeMongodb SourceMongodbUpdateType = "MONGODB"
)

func (e SourceMongodbUpdateType) ToPointer() *SourceMongodbUpdateType {
	return &e
}

func (e *SourceMongodbUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MONGODB":
		*e = SourceMongodbUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMongodbUpdateType: %v", v)
	}
}

type SourceMongodbUpdate struct {
	// Notify if we can't extract for `x` hours. Setting it to `null` disables the notification. Defaults to `null`.
	LatencyThreshold *int64                   `json:"latencyThreshold,omitempty"`
	Type             *SourceMongodbUpdateType `json:"type,omitempty"`
}

func (o *SourceMongodbUpdate) GetLatencyThreshold() *int64 {
	if o == nil {
		return nil
	}
	return o.LatencyThreshold
}

func (o *SourceMongodbUpdate) GetType() *SourceMongodbUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}
