// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
	"time"
)

// ConnectionRecurlyStatus - The current status of the connection.
type ConnectionRecurlyStatus string

const (
	ConnectionRecurlyStatusUnknown     ConnectionRecurlyStatus = "UNKNOWN"
	ConnectionRecurlyStatusUp          ConnectionRecurlyStatus = "UP"
	ConnectionRecurlyStatusDown        ConnectionRecurlyStatus = "DOWN"
	ConnectionRecurlyStatusResize      ConnectionRecurlyStatus = "RESIZE"
	ConnectionRecurlyStatusMaintenance ConnectionRecurlyStatus = "MAINTENANCE"
	ConnectionRecurlyStatusQuota       ConnectionRecurlyStatus = "QUOTA"
	ConnectionRecurlyStatusCreating    ConnectionRecurlyStatus = "CREATING"
)

func (e ConnectionRecurlyStatus) ToPointer() *ConnectionRecurlyStatus {
	return &e
}

func (e *ConnectionRecurlyStatus) UnmarshalJSON(data []byte) error {
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
		*e = ConnectionRecurlyStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionRecurlyStatus: %v", v)
	}
}

type ConnectionRecurlyDefaultUpdateSchedule struct {
	// The pipeline mode refers to how the pipeline fetches data changes from the source and how those changes are applied to the destination table. See <a target="_blank" href="https://docs.etleap.com/docs/documentation/ZG9jOjIyMjE3ODA2-introduction">the documentation</a> for more details.
	PipelineMode *PipelineUpdateModes `json:"pipelineMode,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
}

func (o *ConnectionRecurlyDefaultUpdateSchedule) GetPipelineMode() *PipelineUpdateModes {
	if o == nil {
		return nil
	}
	return o.PipelineMode
}

func (o *ConnectionRecurlyDefaultUpdateSchedule) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionRecurlyDefaultUpdateSchedule) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionRecurlyDefaultUpdateSchedule) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionRecurlyDefaultUpdateSchedule) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionRecurlyDefaultUpdateSchedule) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionRecurlyDefaultUpdateSchedule) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

type ConnectionRecurlyType string

const (
	ConnectionRecurlyTypeRecurly ConnectionRecurlyType = "RECURLY"
)

func (e ConnectionRecurlyType) ToPointer() *ConnectionRecurlyType {
	return &e
}

func (e *ConnectionRecurlyType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "RECURLY":
		*e = ConnectionRecurlyType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionRecurlyType: %v", v)
	}
}

type ConnectionRecurlyOutput struct {
	// The current status of the connection.
	Status ConnectionRecurlyStatus `json:"status"`
	// The unique name of this connection.
	Name string `json:"name"`
	// The date and time when then the connection was created.
	CreateDate time.Time `json:"createDate"`
	// When an update schedule is not defined for a connection, the default schedule is used. The default defined individually per `pipelineMode` and may be subject to change.
	DefaultUpdateSchedule []ConnectionRecurlyDefaultUpdateSchedule `json:"defaultUpdateSchedule"`
	// Whether this connection has been marked as active.
	Active bool                  `json:"active"`
	Type   ConnectionRecurlyType `json:"type"`
	// The unique identifier of the connection.
	ID string `json:"id"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// Your site subdomain at Recurly.
	Subdomain string `json:"subdomain"`
}

func (c ConnectionRecurlyOutput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionRecurlyOutput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionRecurlyOutput) GetStatus() ConnectionRecurlyStatus {
	if o == nil {
		return ConnectionRecurlyStatus("")
	}
	return o.Status
}

func (o *ConnectionRecurlyOutput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionRecurlyOutput) GetCreateDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreateDate
}

func (o *ConnectionRecurlyOutput) GetDefaultUpdateSchedule() []ConnectionRecurlyDefaultUpdateSchedule {
	if o == nil {
		return []ConnectionRecurlyDefaultUpdateSchedule{}
	}
	return o.DefaultUpdateSchedule
}

func (o *ConnectionRecurlyOutput) GetActive() bool {
	if o == nil {
		return false
	}
	return o.Active
}

func (o *ConnectionRecurlyOutput) GetType() ConnectionRecurlyType {
	if o == nil {
		return ConnectionRecurlyType("")
	}
	return o.Type
}

func (o *ConnectionRecurlyOutput) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ConnectionRecurlyOutput) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionRecurlyOutput) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionRecurlyOutput) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionRecurlyOutput) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionRecurlyOutput) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionRecurlyOutput) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionRecurlyOutput) GetSubdomain() string {
	if o == nil {
		return ""
	}
	return o.Subdomain
}

type ConnectionRecurly struct {
	// The unique name of this connection.
	Name string                `json:"name"`
	Type ConnectionRecurlyType `json:"type"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// Your site subdomain at Recurly.
	Subdomain string `json:"subdomain"`
	APIKey    string `json:"apiKey"`
}

func (o *ConnectionRecurly) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionRecurly) GetType() ConnectionRecurlyType {
	if o == nil {
		return ConnectionRecurlyType("")
	}
	return o.Type
}

func (o *ConnectionRecurly) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionRecurly) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionRecurly) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionRecurly) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionRecurly) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionRecurly) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionRecurly) GetSubdomain() string {
	if o == nil {
		return ""
	}
	return o.Subdomain
}

func (o *ConnectionRecurly) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}
