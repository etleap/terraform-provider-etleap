// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
)

// CdcProcessType - Which CDC process to restart. Pipelines that load to Iceberg destinations ingest from the `STREAMING` CDC process, all other pipelines ingest from the `BATCH` CDC process.
type CdcProcessType string

const (
	CdcProcessTypeBatch     CdcProcessType = "BATCH"
	CdcProcessTypeStreaming CdcProcessType = "STREAMING"
)

func (e CdcProcessType) ToPointer() *CdcProcessType {
	return &e
}

func (e *CdcProcessType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "BATCH":
		fallthrough
	case "STREAMING":
		*e = CdcProcessType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CdcProcessType: %v", v)
	}
}

type CdcRestarts struct {
	// Which CDC process to restart. Pipelines that load to Iceberg destinations ingest from the `STREAMING` CDC process, all other pipelines ingest from the `BATCH` CDC process.
	CdcProcessType *CdcProcessType `default:"BATCH" json:"cdcProcessType"`
	// Enable this if you want to request the restart of a healthy CDC process.
	RestartHealthyCdcProcess *bool `default:"false" json:"restartHealthyCdcProcess"`
}

func (c CdcRestarts) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CdcRestarts) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CdcRestarts) GetCdcProcessType() *CdcProcessType {
	if o == nil {
		return nil
	}
	return o.CdcProcessType
}

func (o *CdcRestarts) GetRestartHealthyCdcProcess() *bool {
	if o == nil {
		return nil
	}
	return o.RestartHealthyCdcProcess
}
