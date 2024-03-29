// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
	"time"
)

type ConnectionActiveCampaignType string

const (
	ConnectionActiveCampaignTypeActiveCampaign ConnectionActiveCampaignType = "ACTIVE_CAMPAIGN"
)

func (e ConnectionActiveCampaignType) ToPointer() *ConnectionActiveCampaignType {
	return &e
}

func (e *ConnectionActiveCampaignType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ACTIVE_CAMPAIGN":
		*e = ConnectionActiveCampaignType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionActiveCampaignType: %v", v)
	}
}

// ConnectionActiveCampaignStatus - The current status of the connection.
type ConnectionActiveCampaignStatus string

const (
	ConnectionActiveCampaignStatusUnknown     ConnectionActiveCampaignStatus = "UNKNOWN"
	ConnectionActiveCampaignStatusUp          ConnectionActiveCampaignStatus = "UP"
	ConnectionActiveCampaignStatusDown        ConnectionActiveCampaignStatus = "DOWN"
	ConnectionActiveCampaignStatusResize      ConnectionActiveCampaignStatus = "RESIZE"
	ConnectionActiveCampaignStatusMaintenance ConnectionActiveCampaignStatus = "MAINTENANCE"
	ConnectionActiveCampaignStatusQuota       ConnectionActiveCampaignStatus = "QUOTA"
	ConnectionActiveCampaignStatusCreating    ConnectionActiveCampaignStatus = "CREATING"
)

func (e ConnectionActiveCampaignStatus) ToPointer() *ConnectionActiveCampaignStatus {
	return &e
}

func (e *ConnectionActiveCampaignStatus) UnmarshalJSON(data []byte) error {
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
		*e = ConnectionActiveCampaignStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionActiveCampaignStatus: %v", v)
	}
}

type DefaultUpdateSchedule struct {
	// The pipeline mode refers to how the pipeline fetches data changes from the source and how those changes are applied to the destination table. See <a target="_blank" href="https://docs.etleap.com/docs/documentation/ZG9jOjIyMjE3ODA2-introduction">the documentation</a> for more details.
	PipelineMode *PipelineUpdateModes `json:"pipelineMode,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
}

func (o *DefaultUpdateSchedule) GetPipelineMode() *PipelineUpdateModes {
	if o == nil {
		return nil
	}
	return o.PipelineMode
}

func (o *DefaultUpdateSchedule) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *DefaultUpdateSchedule) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *DefaultUpdateSchedule) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *DefaultUpdateSchedule) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *DefaultUpdateSchedule) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *DefaultUpdateSchedule) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

type ConnectionActiveCampaign struct {
	// The unique identifier of the connection.
	ID string `json:"id"`
	// The unique name of this connection.
	Name string                       `json:"name"`
	Type ConnectionActiveCampaignType `json:"type"`
	// Whether this connection has been marked as active.
	Active bool `json:"active"`
	// The current status of the connection.
	Status ConnectionActiveCampaignStatus `json:"status"`
	// The date and time when then the connection was created.
	CreateDate time.Time `json:"createDate"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// When an update schedule is not defined for a connection, the default schedule is used. The default defined individually per `pipelineMode` and may be subject to change.
	DefaultUpdateSchedule []DefaultUpdateSchedule `json:"defaultUpdateSchedule"`
	// The base URL is specific to your account. Your API URL can be found in your account on the My Settings page under the "Developer" tab.
	BaseURL string `json:"baseUrl"`
}

func (c ConnectionActiveCampaign) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionActiveCampaign) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionActiveCampaign) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ConnectionActiveCampaign) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionActiveCampaign) GetType() ConnectionActiveCampaignType {
	if o == nil {
		return ConnectionActiveCampaignType("")
	}
	return o.Type
}

func (o *ConnectionActiveCampaign) GetActive() bool {
	if o == nil {
		return false
	}
	return o.Active
}

func (o *ConnectionActiveCampaign) GetStatus() ConnectionActiveCampaignStatus {
	if o == nil {
		return ConnectionActiveCampaignStatus("")
	}
	return o.Status
}

func (o *ConnectionActiveCampaign) GetCreateDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreateDate
}

func (o *ConnectionActiveCampaign) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionActiveCampaign) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionActiveCampaign) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionActiveCampaign) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionActiveCampaign) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionActiveCampaign) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionActiveCampaign) GetDefaultUpdateSchedule() []DefaultUpdateSchedule {
	if o == nil {
		return []DefaultUpdateSchedule{}
	}
	return o.DefaultUpdateSchedule
}

func (o *ConnectionActiveCampaign) GetBaseURL() string {
	if o == nil {
		return ""
	}
	return o.BaseURL
}

type ConnectionActiveCampaignInput struct {
	// The unique name of this connection.
	Name string                       `json:"name"`
	Type ConnectionActiveCampaignType `json:"type"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// The base URL is specific to your account. Your API URL can be found in your account on the My Settings page under the "Developer" tab.
	BaseURL string `json:"baseUrl"`
	// Your API key can be found in your account on the Settings page under the "Developer" tab. Each user in your ActiveCampaign account has their own unique API key.
	APIKey string `json:"apiKey"`
}

func (o *ConnectionActiveCampaignInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionActiveCampaignInput) GetType() ConnectionActiveCampaignType {
	if o == nil {
		return ConnectionActiveCampaignType("")
	}
	return o.Type
}

func (o *ConnectionActiveCampaignInput) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionActiveCampaignInput) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionActiveCampaignInput) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionActiveCampaignInput) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionActiveCampaignInput) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionActiveCampaignInput) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionActiveCampaignInput) GetBaseURL() string {
	if o == nil {
		return ""
	}
	return o.BaseURL
}

func (o *ConnectionActiveCampaignInput) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}
