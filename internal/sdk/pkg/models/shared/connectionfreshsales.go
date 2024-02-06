// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
	"time"
)

type ConnectionFreshsalesType string

const (
	ConnectionFreshsalesTypeFreshsales ConnectionFreshsalesType = "FRESHSALES"
)

func (e ConnectionFreshsalesType) ToPointer() *ConnectionFreshsalesType {
	return &e
}

func (e *ConnectionFreshsalesType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "FRESHSALES":
		*e = ConnectionFreshsalesType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionFreshsalesType: %v", v)
	}
}

// ConnectionFreshsalesStatus - The current status of the connection.
type ConnectionFreshsalesStatus string

const (
	ConnectionFreshsalesStatusUnknown     ConnectionFreshsalesStatus = "UNKNOWN"
	ConnectionFreshsalesStatusUp          ConnectionFreshsalesStatus = "UP"
	ConnectionFreshsalesStatusDown        ConnectionFreshsalesStatus = "DOWN"
	ConnectionFreshsalesStatusResize      ConnectionFreshsalesStatus = "RESIZE"
	ConnectionFreshsalesStatusMaintenance ConnectionFreshsalesStatus = "MAINTENANCE"
	ConnectionFreshsalesStatusQuota       ConnectionFreshsalesStatus = "QUOTA"
	ConnectionFreshsalesStatusCreating    ConnectionFreshsalesStatus = "CREATING"
)

func (e ConnectionFreshsalesStatus) ToPointer() *ConnectionFreshsalesStatus {
	return &e
}

func (e *ConnectionFreshsalesStatus) UnmarshalJSON(data []byte) error {
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
		*e = ConnectionFreshsalesStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionFreshsalesStatus: %v", v)
	}
}

type ConnectionFreshsalesDefaultUpdateSchedule struct {
	// The pipeline mode refers to how the pipeline fetches data changes from the source and how those changes are applied to the destination table. See <a target="_blank" href="https://docs.etleap.com/docs/documentation/ZG9jOjIyMjE3ODA2-introduction">the documentation</a> for more details.
	PipelineMode *PipelineUpdateModes `json:"pipelineMode,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
}

func (o *ConnectionFreshsalesDefaultUpdateSchedule) GetPipelineMode() *PipelineUpdateModes {
	if o == nil {
		return nil
	}
	return o.PipelineMode
}

func (o *ConnectionFreshsalesDefaultUpdateSchedule) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionFreshsalesDefaultUpdateSchedule) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionFreshsalesDefaultUpdateSchedule) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionFreshsalesDefaultUpdateSchedule) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionFreshsalesDefaultUpdateSchedule) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionFreshsalesDefaultUpdateSchedule) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

type ConnectionFreshsales struct {
	// The unique identifier of the connection.
	ID string `json:"id"`
	// The unique name of this connection.
	Name string                   `json:"name"`
	Type ConnectionFreshsalesType `json:"type"`
	// Whether this connection has been marked as active.
	Active bool `json:"active"`
	// The current status of the connection.
	Status ConnectionFreshsalesStatus `json:"status"`
	// The date and time when then the connection was created.
	CreateDate time.Time `json:"createDate"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// When an update schedule is not defined for a connection, the default schedule is used. The default defined individually per `pipelineMode` and may be subject to change.
	DefaultUpdateSchedule []ConnectionFreshsalesDefaultUpdateSchedule `json:"defaultUpdateSchedule"`
	// Your Freshsales domain. Can be found under Admin Settings -> API Settings under the "bundle alias" label.
	Domain string `json:"domain"`
}

func (c ConnectionFreshsales) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionFreshsales) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionFreshsales) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ConnectionFreshsales) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionFreshsales) GetType() ConnectionFreshsalesType {
	if o == nil {
		return ConnectionFreshsalesType("")
	}
	return o.Type
}

func (o *ConnectionFreshsales) GetActive() bool {
	if o == nil {
		return false
	}
	return o.Active
}

func (o *ConnectionFreshsales) GetStatus() ConnectionFreshsalesStatus {
	if o == nil {
		return ConnectionFreshsalesStatus("")
	}
	return o.Status
}

func (o *ConnectionFreshsales) GetCreateDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreateDate
}

func (o *ConnectionFreshsales) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionFreshsales) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionFreshsales) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionFreshsales) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionFreshsales) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionFreshsales) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionFreshsales) GetDefaultUpdateSchedule() []ConnectionFreshsalesDefaultUpdateSchedule {
	if o == nil {
		return []ConnectionFreshsalesDefaultUpdateSchedule{}
	}
	return o.DefaultUpdateSchedule
}

func (o *ConnectionFreshsales) GetDomain() string {
	if o == nil {
		return ""
	}
	return o.Domain
}

type ConnectionFreshsalesInput struct {
	// The unique name of this connection.
	Name string                   `json:"name"`
	Type ConnectionFreshsalesType `json:"type"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// Your Freshsales domain. Can be found under Admin Settings -> API Settings under the "bundle alias" label.
	Domain string `json:"domain"`
	// Your Freshsales API Key. Can be found under Admin Settings -> API Settings under the "API Key" label.
	APIKey string `json:"apiKey"`
}

func (o *ConnectionFreshsalesInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionFreshsalesInput) GetType() ConnectionFreshsalesType {
	if o == nil {
		return ConnectionFreshsalesType("")
	}
	return o.Type
}

func (o *ConnectionFreshsalesInput) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionFreshsalesInput) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionFreshsalesInput) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionFreshsalesInput) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionFreshsalesInput) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionFreshsalesInput) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionFreshsalesInput) GetDomain() string {
	if o == nil {
		return ""
	}
	return o.Domain
}

func (o *ConnectionFreshsalesInput) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}
