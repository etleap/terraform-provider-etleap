// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/types"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
)

type RetentionDayRowCount struct {
	// Format of the timestamp: 'yyyy-MM-dd'
	Date     *types.Date `json:"date,omitempty"`
	RowCount int64       `json:"rowCount"`
}

func (r RetentionDayRowCount) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RetentionDayRowCount) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RetentionDayRowCount) GetDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.Date
}

func (o *RetentionDayRowCount) GetRowCount() int64 {
	if o == nil {
		return 0
	}
	return o.RowCount
}
