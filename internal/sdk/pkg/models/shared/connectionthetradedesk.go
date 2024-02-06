// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
	"time"
)

type ConnectionTheTradeDeskType string

const (
	ConnectionTheTradeDeskTypeTheTradeDesk ConnectionTheTradeDeskType = "THE_TRADE_DESK"
)

func (e ConnectionTheTradeDeskType) ToPointer() *ConnectionTheTradeDeskType {
	return &e
}

func (e *ConnectionTheTradeDeskType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "THE_TRADE_DESK":
		*e = ConnectionTheTradeDeskType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionTheTradeDeskType: %v", v)
	}
}

// ConnectionTheTradeDeskStatus - The current status of the connection.
type ConnectionTheTradeDeskStatus string

const (
	ConnectionTheTradeDeskStatusUnknown     ConnectionTheTradeDeskStatus = "UNKNOWN"
	ConnectionTheTradeDeskStatusUp          ConnectionTheTradeDeskStatus = "UP"
	ConnectionTheTradeDeskStatusDown        ConnectionTheTradeDeskStatus = "DOWN"
	ConnectionTheTradeDeskStatusResize      ConnectionTheTradeDeskStatus = "RESIZE"
	ConnectionTheTradeDeskStatusMaintenance ConnectionTheTradeDeskStatus = "MAINTENANCE"
	ConnectionTheTradeDeskStatusQuota       ConnectionTheTradeDeskStatus = "QUOTA"
	ConnectionTheTradeDeskStatusCreating    ConnectionTheTradeDeskStatus = "CREATING"
)

func (e ConnectionTheTradeDeskStatus) ToPointer() *ConnectionTheTradeDeskStatus {
	return &e
}

func (e *ConnectionTheTradeDeskStatus) UnmarshalJSON(data []byte) error {
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
		*e = ConnectionTheTradeDeskStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionTheTradeDeskStatus: %v", v)
	}
}

type ConnectionTheTradeDeskDefaultUpdateSchedule struct {
	// The pipeline mode refers to how the pipeline fetches data changes from the source and how those changes are applied to the destination table. See <a target="_blank" href="https://docs.etleap.com/docs/documentation/ZG9jOjIyMjE3ODA2-introduction">the documentation</a> for more details.
	PipelineMode *PipelineUpdateModes `json:"pipelineMode,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
}

func (o *ConnectionTheTradeDeskDefaultUpdateSchedule) GetPipelineMode() *PipelineUpdateModes {
	if o == nil {
		return nil
	}
	return o.PipelineMode
}

func (o *ConnectionTheTradeDeskDefaultUpdateSchedule) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionTheTradeDeskDefaultUpdateSchedule) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionTheTradeDeskDefaultUpdateSchedule) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionTheTradeDeskDefaultUpdateSchedule) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionTheTradeDeskDefaultUpdateSchedule) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionTheTradeDeskDefaultUpdateSchedule) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

type ConnectionTheTradeDesk struct {
	// The unique identifier of the connection.
	ID string `json:"id"`
	// The unique name of this connection.
	Name string                     `json:"name"`
	Type ConnectionTheTradeDeskType `json:"type"`
	// Whether this connection has been marked as active.
	Active bool `json:"active"`
	// The current status of the connection.
	Status ConnectionTheTradeDeskStatus `json:"status"`
	// The date and time when then the connection was created.
	CreateDate time.Time `json:"createDate"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// When an update schedule is not defined for a connection, the default schedule is used. The default defined individually per `pipelineMode` and may be subject to change.
	DefaultUpdateSchedule []ConnectionTheTradeDeskDefaultUpdateSchedule `json:"defaultUpdateSchedule"`
	// Typically an email address.
	Username string `json:"username"`
	// Your Partner ID at The Trade Desk.
	PartnerID string `json:"partnerId"`
	// Whether this is a sandbox account.
	Sandbox bool `json:"sandbox"`
}

func (c ConnectionTheTradeDesk) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionTheTradeDesk) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionTheTradeDesk) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ConnectionTheTradeDesk) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionTheTradeDesk) GetType() ConnectionTheTradeDeskType {
	if o == nil {
		return ConnectionTheTradeDeskType("")
	}
	return o.Type
}

func (o *ConnectionTheTradeDesk) GetActive() bool {
	if o == nil {
		return false
	}
	return o.Active
}

func (o *ConnectionTheTradeDesk) GetStatus() ConnectionTheTradeDeskStatus {
	if o == nil {
		return ConnectionTheTradeDeskStatus("")
	}
	return o.Status
}

func (o *ConnectionTheTradeDesk) GetCreateDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreateDate
}

func (o *ConnectionTheTradeDesk) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionTheTradeDesk) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionTheTradeDesk) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionTheTradeDesk) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionTheTradeDesk) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionTheTradeDesk) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionTheTradeDesk) GetDefaultUpdateSchedule() []ConnectionTheTradeDeskDefaultUpdateSchedule {
	if o == nil {
		return []ConnectionTheTradeDeskDefaultUpdateSchedule{}
	}
	return o.DefaultUpdateSchedule
}

func (o *ConnectionTheTradeDesk) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *ConnectionTheTradeDesk) GetPartnerID() string {
	if o == nil {
		return ""
	}
	return o.PartnerID
}

func (o *ConnectionTheTradeDesk) GetSandbox() bool {
	if o == nil {
		return false
	}
	return o.Sandbox
}

type ConnectionTheTradeDeskInput struct {
	// The unique name of this connection.
	Name string                     `json:"name"`
	Type ConnectionTheTradeDeskType `json:"type"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// Typically an email address.
	Username string `json:"username"`
	Password string `json:"password"`
	// Your Partner ID at The Trade Desk.
	PartnerID string `json:"partnerId"`
	// Whether this is a sandbox account.
	Sandbox bool `json:"sandbox"`
}

func (o *ConnectionTheTradeDeskInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionTheTradeDeskInput) GetType() ConnectionTheTradeDeskType {
	if o == nil {
		return ConnectionTheTradeDeskType("")
	}
	return o.Type
}

func (o *ConnectionTheTradeDeskInput) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionTheTradeDeskInput) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionTheTradeDeskInput) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionTheTradeDeskInput) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionTheTradeDeskInput) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionTheTradeDeskInput) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionTheTradeDeskInput) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *ConnectionTheTradeDeskInput) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *ConnectionTheTradeDeskInput) GetPartnerID() string {
	if o == nil {
		return ""
	}
	return o.PartnerID
}

func (o *ConnectionTheTradeDeskInput) GetSandbox() bool {
	if o == nil {
		return false
	}
	return o.Sandbox
}
