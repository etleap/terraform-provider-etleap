// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
	"time"
)

// LegacyScript - To be used only for copying a script exactly.
type LegacyScript struct {
	// The serialization of a script. Not meant to be human-readable.
	LegacyScript string `json:"legacyScript"`
	// The date and time when then the script was created.
	CreateDate time.Time `json:"createDate"`
}

func (l LegacyScript) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LegacyScript) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *LegacyScript) GetLegacyScript() string {
	if o == nil {
		return ""
	}
	return o.LegacyScript
}

func (o *LegacyScript) GetCreateDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreateDate
}
