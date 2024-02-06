// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
	"time"
)

type ConnectionRaveMedidataType string

const (
	ConnectionRaveMedidataTypeRaveMedidata ConnectionRaveMedidataType = "RAVE_MEDIDATA"
)

func (e ConnectionRaveMedidataType) ToPointer() *ConnectionRaveMedidataType {
	return &e
}

func (e *ConnectionRaveMedidataType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "RAVE_MEDIDATA":
		*e = ConnectionRaveMedidataType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionRaveMedidataType: %v", v)
	}
}

// ConnectionRaveMedidataStatus - The current status of the connection.
type ConnectionRaveMedidataStatus string

const (
	ConnectionRaveMedidataStatusUnknown     ConnectionRaveMedidataStatus = "UNKNOWN"
	ConnectionRaveMedidataStatusUp          ConnectionRaveMedidataStatus = "UP"
	ConnectionRaveMedidataStatusDown        ConnectionRaveMedidataStatus = "DOWN"
	ConnectionRaveMedidataStatusResize      ConnectionRaveMedidataStatus = "RESIZE"
	ConnectionRaveMedidataStatusMaintenance ConnectionRaveMedidataStatus = "MAINTENANCE"
	ConnectionRaveMedidataStatusQuota       ConnectionRaveMedidataStatus = "QUOTA"
	ConnectionRaveMedidataStatusCreating    ConnectionRaveMedidataStatus = "CREATING"
)

func (e ConnectionRaveMedidataStatus) ToPointer() *ConnectionRaveMedidataStatus {
	return &e
}

func (e *ConnectionRaveMedidataStatus) UnmarshalJSON(data []byte) error {
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
		*e = ConnectionRaveMedidataStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionRaveMedidataStatus: %v", v)
	}
}

type ConnectionRaveMedidataDefaultUpdateSchedule struct {
	// The pipeline mode refers to how the pipeline fetches data changes from the source and how those changes are applied to the destination table. See <a target="_blank" href="https://docs.etleap.com/docs/documentation/ZG9jOjIyMjE3ODA2-introduction">the documentation</a> for more details.
	PipelineMode *PipelineUpdateModes `json:"pipelineMode,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
}

func (o *ConnectionRaveMedidataDefaultUpdateSchedule) GetPipelineMode() *PipelineUpdateModes {
	if o == nil {
		return nil
	}
	return o.PipelineMode
}

func (o *ConnectionRaveMedidataDefaultUpdateSchedule) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionRaveMedidataDefaultUpdateSchedule) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionRaveMedidataDefaultUpdateSchedule) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionRaveMedidataDefaultUpdateSchedule) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionRaveMedidataDefaultUpdateSchedule) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionRaveMedidataDefaultUpdateSchedule) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

type ConnectionRaveMedidata struct {
	// The unique identifier of the connection.
	ID string `json:"id"`
	// The unique name of this connection.
	Name string                     `json:"name"`
	Type ConnectionRaveMedidataType `json:"type"`
	// Whether this connection has been marked as active.
	Active bool `json:"active"`
	// The current status of the connection.
	Status ConnectionRaveMedidataStatus `json:"status"`
	// The date and time when then the connection was created.
	CreateDate time.Time `json:"createDate"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// When an update schedule is not defined for a connection, the default schedule is used. The default defined individually per `pipelineMode` and may be subject to change.
	DefaultUpdateSchedule []ConnectionRaveMedidataDefaultUpdateSchedule `json:"defaultUpdateSchedule"`
}

func (c ConnectionRaveMedidata) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionRaveMedidata) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionRaveMedidata) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ConnectionRaveMedidata) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionRaveMedidata) GetType() ConnectionRaveMedidataType {
	if o == nil {
		return ConnectionRaveMedidataType("")
	}
	return o.Type
}

func (o *ConnectionRaveMedidata) GetActive() bool {
	if o == nil {
		return false
	}
	return o.Active
}

func (o *ConnectionRaveMedidata) GetStatus() ConnectionRaveMedidataStatus {
	if o == nil {
		return ConnectionRaveMedidataStatus("")
	}
	return o.Status
}

func (o *ConnectionRaveMedidata) GetCreateDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreateDate
}

func (o *ConnectionRaveMedidata) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionRaveMedidata) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionRaveMedidata) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionRaveMedidata) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionRaveMedidata) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionRaveMedidata) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionRaveMedidata) GetDefaultUpdateSchedule() []ConnectionRaveMedidataDefaultUpdateSchedule {
	if o == nil {
		return []ConnectionRaveMedidataDefaultUpdateSchedule{}
	}
	return o.DefaultUpdateSchedule
}

type ConnectionRaveMedidataInput struct {
	// The unique name of this connection.
	Name string                     `json:"name"`
	Type ConnectionRaveMedidataType `json:"type"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// E.g., my-org.mdsol.com
	Hostname string `json:"hostname"`
	// Rave Account Login
	Username string `json:"username"`
	// Rave Account Password
	Password string `json:"password"`
}

func (o *ConnectionRaveMedidataInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionRaveMedidataInput) GetType() ConnectionRaveMedidataType {
	if o == nil {
		return ConnectionRaveMedidataType("")
	}
	return o.Type
}

func (o *ConnectionRaveMedidataInput) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionRaveMedidataInput) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionRaveMedidataInput) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionRaveMedidataInput) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionRaveMedidataInput) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionRaveMedidataInput) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionRaveMedidataInput) GetHostname() string {
	if o == nil {
		return ""
	}
	return o.Hostname
}

func (o *ConnectionRaveMedidataInput) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *ConnectionRaveMedidataInput) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}
