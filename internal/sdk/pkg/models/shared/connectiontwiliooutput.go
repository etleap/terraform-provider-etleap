// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
	"time"
)

// ConnectionTwilioStatus - The current status of the connection.
type ConnectionTwilioStatus string

const (
	ConnectionTwilioStatusUnknown     ConnectionTwilioStatus = "UNKNOWN"
	ConnectionTwilioStatusUp          ConnectionTwilioStatus = "UP"
	ConnectionTwilioStatusDown        ConnectionTwilioStatus = "DOWN"
	ConnectionTwilioStatusResize      ConnectionTwilioStatus = "RESIZE"
	ConnectionTwilioStatusMaintenance ConnectionTwilioStatus = "MAINTENANCE"
	ConnectionTwilioStatusQuota       ConnectionTwilioStatus = "QUOTA"
	ConnectionTwilioStatusCreating    ConnectionTwilioStatus = "CREATING"
)

func (e ConnectionTwilioStatus) ToPointer() *ConnectionTwilioStatus {
	return &e
}

func (e *ConnectionTwilioStatus) UnmarshalJSON(data []byte) error {
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
		*e = ConnectionTwilioStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionTwilioStatus: %v", v)
	}
}

type ConnectionTwilioDefaultUpdateSchedule struct {
	// The pipeline mode refers to how the pipeline fetches data changes from the source and how those changes are applied to the destination table. See <a target="_blank" href="https://docs.etleap.com/docs/documentation/ZG9jOjIyMjE3ODA2-introduction">the documentation</a> for more details.
	PipelineMode *PipelineUpdateModes `json:"pipelineMode,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
}

func (o *ConnectionTwilioDefaultUpdateSchedule) GetPipelineMode() *PipelineUpdateModes {
	if o == nil {
		return nil
	}
	return o.PipelineMode
}

func (o *ConnectionTwilioDefaultUpdateSchedule) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionTwilioDefaultUpdateSchedule) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionTwilioDefaultUpdateSchedule) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionTwilioDefaultUpdateSchedule) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionTwilioDefaultUpdateSchedule) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionTwilioDefaultUpdateSchedule) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

type ConnectionTwilioType string

const (
	ConnectionTwilioTypeTwilio ConnectionTwilioType = "TWILIO"
)

func (e ConnectionTwilioType) ToPointer() *ConnectionTwilioType {
	return &e
}

func (e *ConnectionTwilioType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TWILIO":
		*e = ConnectionTwilioType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionTwilioType: %v", v)
	}
}

type ConnectionTwilioOutput struct {
	// The current status of the connection.
	Status ConnectionTwilioStatus `json:"status"`
	// The unique name of this connection.
	Name string `json:"name"`
	// The date and time when then the connection was created.
	CreateDate time.Time `json:"createDate"`
	// When an update schedule is not defined for a connection, the default schedule is used. The default defined individually per `pipelineMode` and may be subject to change.
	DefaultUpdateSchedule []ConnectionTwilioDefaultUpdateSchedule `json:"defaultUpdateSchedule"`
	// Whether this connection has been marked as active.
	Active bool                 `json:"active"`
	Type   ConnectionTwilioType `json:"type"`
	// The unique identifier of the connection.
	ID string `json:"id"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// You can find your Account SID in your <a target="_blank" href="https://twilio.com/console">Account Dashboard</a>.
	AccountSid string `json:"accountSid"`
	// You can find your API keys, or generate a new one, in the <a target="_blank" href="https://www.twilio.com/console/runtime/api-keys">Twilio Console</a>.
	APIKeySid string `json:"apiKeySid"`
}

func (c ConnectionTwilioOutput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionTwilioOutput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionTwilioOutput) GetStatus() ConnectionTwilioStatus {
	if o == nil {
		return ConnectionTwilioStatus("")
	}
	return o.Status
}

func (o *ConnectionTwilioOutput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionTwilioOutput) GetCreateDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreateDate
}

func (o *ConnectionTwilioOutput) GetDefaultUpdateSchedule() []ConnectionTwilioDefaultUpdateSchedule {
	if o == nil {
		return []ConnectionTwilioDefaultUpdateSchedule{}
	}
	return o.DefaultUpdateSchedule
}

func (o *ConnectionTwilioOutput) GetActive() bool {
	if o == nil {
		return false
	}
	return o.Active
}

func (o *ConnectionTwilioOutput) GetType() ConnectionTwilioType {
	if o == nil {
		return ConnectionTwilioType("")
	}
	return o.Type
}

func (o *ConnectionTwilioOutput) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ConnectionTwilioOutput) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionTwilioOutput) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionTwilioOutput) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionTwilioOutput) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionTwilioOutput) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionTwilioOutput) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionTwilioOutput) GetAccountSid() string {
	if o == nil {
		return ""
	}
	return o.AccountSid
}

func (o *ConnectionTwilioOutput) GetAPIKeySid() string {
	if o == nil {
		return ""
	}
	return o.APIKeySid
}

type ConnectionTwilio struct {
	// The unique name of this connection.
	Name string               `json:"name"`
	Type ConnectionTwilioType `json:"type"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// You can find your Account SID in your <a target="_blank" href="https://twilio.com/console">Account Dashboard</a>.
	AccountSid string `json:"accountSid"`
	// You can find your API keys, or generate a new one, in the <a target="_blank" href="https://www.twilio.com/console/runtime/api-keys">Twilio Console</a>.
	APIKeySid string `json:"apiKeySid"`
	// You can find your API keys, or generate a new one, in the <a target="_blank" href="https://www.twilio.com/console/runtime/api-keys">Twilio Console</a>.
	APIKeySecret string `json:"apiKeySecret"`
}

func (o *ConnectionTwilio) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionTwilio) GetType() ConnectionTwilioType {
	if o == nil {
		return ConnectionTwilioType("")
	}
	return o.Type
}

func (o *ConnectionTwilio) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionTwilio) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionTwilio) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionTwilio) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionTwilio) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionTwilio) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionTwilio) GetAccountSid() string {
	if o == nil {
		return ""
	}
	return o.AccountSid
}

func (o *ConnectionTwilio) GetAPIKeySid() string {
	if o == nil {
		return ""
	}
	return o.APIKeySid
}

func (o *ConnectionTwilio) GetAPIKeySecret() string {
	if o == nil {
		return ""
	}
	return o.APIKeySecret
}
