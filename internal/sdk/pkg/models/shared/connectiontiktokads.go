// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
	"time"
)

type ConnectionTikTokAdsType string

const (
	ConnectionTikTokAdsTypeTikTokAds ConnectionTikTokAdsType = "TIK_TOK_ADS"
)

func (e ConnectionTikTokAdsType) ToPointer() *ConnectionTikTokAdsType {
	return &e
}

func (e *ConnectionTikTokAdsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TIK_TOK_ADS":
		*e = ConnectionTikTokAdsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionTikTokAdsType: %v", v)
	}
}

// ConnectionTikTokAdsStatus - The current status of the connection.
type ConnectionTikTokAdsStatus string

const (
	ConnectionTikTokAdsStatusUnknown     ConnectionTikTokAdsStatus = "UNKNOWN"
	ConnectionTikTokAdsStatusUp          ConnectionTikTokAdsStatus = "UP"
	ConnectionTikTokAdsStatusDown        ConnectionTikTokAdsStatus = "DOWN"
	ConnectionTikTokAdsStatusResize      ConnectionTikTokAdsStatus = "RESIZE"
	ConnectionTikTokAdsStatusMaintenance ConnectionTikTokAdsStatus = "MAINTENANCE"
	ConnectionTikTokAdsStatusQuota       ConnectionTikTokAdsStatus = "QUOTA"
	ConnectionTikTokAdsStatusCreating    ConnectionTikTokAdsStatus = "CREATING"
)

func (e ConnectionTikTokAdsStatus) ToPointer() *ConnectionTikTokAdsStatus {
	return &e
}

func (e *ConnectionTikTokAdsStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "UNKNOWN":
		fallthrough
	case "UP":
		fallthrough
	case "DOWN":
		fallthrough
	case "RESIZE":
		fallthrough
	case "MAINTENANCE":
		fallthrough
	case "QUOTA":
		fallthrough
	case "CREATING":
		*e = ConnectionTikTokAdsStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionTikTokAdsStatus: %v", v)
	}
}

type ConnectionTikTokAdsDefaultUpdateSchedule struct {
	// The pipeline mode refers to how the pipeline fetches data changes from the source and how those changes are applied to the destination table. See <a target="_blank" href="https://docs.etleap.com/docs/documentation/ZG9jOjIyMjE3ODA2-introduction">the documentation</a> for more details.
	PipelineMode *PipelineUpdateModes `json:"pipelineMode,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
}

func (o *ConnectionTikTokAdsDefaultUpdateSchedule) GetPipelineMode() *PipelineUpdateModes {
	if o == nil {
		return nil
	}
	return o.PipelineMode
}

func (o *ConnectionTikTokAdsDefaultUpdateSchedule) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionTikTokAdsDefaultUpdateSchedule) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionTikTokAdsDefaultUpdateSchedule) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionTikTokAdsDefaultUpdateSchedule) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionTikTokAdsDefaultUpdateSchedule) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionTikTokAdsDefaultUpdateSchedule) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

type ConnectionTikTokAds struct {
	// The unique identifier of the connection.
	ID string `json:"id"`
	// The unique name of this connection.
	Name string                  `json:"name"`
	Type ConnectionTikTokAdsType `json:"type"`
	// Whether this connection has been marked as active.
	Active bool `json:"active"`
	// The current status of the connection.
	Status ConnectionTikTokAdsStatus `json:"status"`
	// The date and time when then the connection was created.
	CreateDate time.Time `json:"createDate"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// When an update schedule is not defined for a connection, the default schedule is used. The default defined individually per `pipelineMode` and may be subject to change.
	DefaultUpdateSchedule []ConnectionTikTokAdsDefaultUpdateSchedule `json:"defaultUpdateSchedule"`
}

func (c ConnectionTikTokAds) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionTikTokAds) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionTikTokAds) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ConnectionTikTokAds) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionTikTokAds) GetType() ConnectionTikTokAdsType {
	if o == nil {
		return ConnectionTikTokAdsType("")
	}
	return o.Type
}

func (o *ConnectionTikTokAds) GetActive() bool {
	if o == nil {
		return false
	}
	return o.Active
}

func (o *ConnectionTikTokAds) GetStatus() ConnectionTikTokAdsStatus {
	if o == nil {
		return ConnectionTikTokAdsStatus("")
	}
	return o.Status
}

func (o *ConnectionTikTokAds) GetCreateDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreateDate
}

func (o *ConnectionTikTokAds) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionTikTokAds) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionTikTokAds) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionTikTokAds) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionTikTokAds) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionTikTokAds) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionTikTokAds) GetDefaultUpdateSchedule() []ConnectionTikTokAdsDefaultUpdateSchedule {
	if o == nil {
		return []ConnectionTikTokAdsDefaultUpdateSchedule{}
	}
	return o.DefaultUpdateSchedule
}

type ConnectionTikTokAdsInput struct {
	// The unique name of this connection.
	Name string                  `json:"name"`
	Type ConnectionTikTokAdsType `json:"type"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// Code retrieved from `/connections/oauth2-initiation`. **Note:** it is short-lived, therefore the connection creation should be done as soon as code is returned.
	Code string `json:"code"`
}

func (o *ConnectionTikTokAdsInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionTikTokAdsInput) GetType() ConnectionTikTokAdsType {
	if o == nil {
		return ConnectionTikTokAdsType("")
	}
	return o.Type
}

func (o *ConnectionTikTokAdsInput) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionTikTokAdsInput) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionTikTokAdsInput) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionTikTokAdsInput) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionTikTokAdsInput) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionTikTokAdsInput) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionTikTokAdsInput) GetCode() string {
	if o == nil {
		return ""
	}
	return o.Code
}
