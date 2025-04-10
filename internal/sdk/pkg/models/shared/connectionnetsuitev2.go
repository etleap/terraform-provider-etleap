// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
	"time"
)

// ConnectionNetsuiteV2Status - The current status of the connection.
type ConnectionNetsuiteV2Status string

const (
	ConnectionNetsuiteV2StatusUnknown     ConnectionNetsuiteV2Status = "UNKNOWN"
	ConnectionNetsuiteV2StatusUp          ConnectionNetsuiteV2Status = "UP"
	ConnectionNetsuiteV2StatusDown        ConnectionNetsuiteV2Status = "DOWN"
	ConnectionNetsuiteV2StatusResize      ConnectionNetsuiteV2Status = "RESIZE"
	ConnectionNetsuiteV2StatusMaintenance ConnectionNetsuiteV2Status = "MAINTENANCE"
	ConnectionNetsuiteV2StatusQuota       ConnectionNetsuiteV2Status = "QUOTA"
	ConnectionNetsuiteV2StatusCreating    ConnectionNetsuiteV2Status = "CREATING"
)

func (e ConnectionNetsuiteV2Status) ToPointer() *ConnectionNetsuiteV2Status {
	return &e
}

func (e *ConnectionNetsuiteV2Status) UnmarshalJSON(data []byte) error {
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
		*e = ConnectionNetsuiteV2Status(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionNetsuiteV2Status: %v", v)
	}
}

type ConnectionNetsuiteV2DefaultUpdateSchedule struct {
	// The pipeline mode refers to how the pipeline fetches data changes from the source and how those changes are applied to the destination table. See <a target="_blank" href="https://docs.etleap.com/docs/documentation/ZG9jOjIyMjE3ODA2-introduction">the documentation</a> for more details.
	PipelineMode *PipelineUpdateModes `json:"pipelineMode,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
}

func (o *ConnectionNetsuiteV2DefaultUpdateSchedule) GetPipelineMode() *PipelineUpdateModes {
	if o == nil {
		return nil
	}
	return o.PipelineMode
}

func (o *ConnectionNetsuiteV2DefaultUpdateSchedule) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionNetsuiteV2DefaultUpdateSchedule) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionNetsuiteV2DefaultUpdateSchedule) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionNetsuiteV2DefaultUpdateSchedule) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionNetsuiteV2DefaultUpdateSchedule) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionNetsuiteV2DefaultUpdateSchedule) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

type ConnectionNetsuiteV2Type string

const (
	ConnectionNetsuiteV2TypeNetsuiteV2 ConnectionNetsuiteV2Type = "NETSUITE_V2"
)

func (e ConnectionNetsuiteV2Type) ToPointer() *ConnectionNetsuiteV2Type {
	return &e
}

func (e *ConnectionNetsuiteV2Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NETSUITE_V2":
		*e = ConnectionNetsuiteV2Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionNetsuiteV2Type: %v", v)
	}
}

type ConnectionNetsuiteV2 struct {
	// The current status of the connection.
	Status ConnectionNetsuiteV2Status `json:"status"`
	// The unique name of this connection.
	Name string `json:"name"`
	// The date and time when then the connection was created.
	CreateDate time.Time `json:"createDate"`
	// When an update schedule is not defined for a connection, the default schedule is used. The default defined individually per `pipelineMode` and may be subject to change.
	DefaultUpdateSchedule []ConnectionNetsuiteV2DefaultUpdateSchedule `json:"defaultUpdateSchedule"`
	// Whether this connection has been marked as active.
	Active bool                     `json:"active"`
	Type   ConnectionNetsuiteV2Type `json:"type"`
	// The unique identifier of the connection.
	ID string `json:"id"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// Token Id value displayed after creating a new access token.
	TokenID string `json:"tokenId"`
	// Consumer Key value displayed after creating a new integration.
	ConsumerKey string `json:"consumerKey"`
	// Under Setup > Company > Company Information
	AccountID string `json:"accountId"`
}

func (c ConnectionNetsuiteV2) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionNetsuiteV2) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionNetsuiteV2) GetStatus() ConnectionNetsuiteV2Status {
	if o == nil {
		return ConnectionNetsuiteV2Status("")
	}
	return o.Status
}

func (o *ConnectionNetsuiteV2) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionNetsuiteV2) GetCreateDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreateDate
}

func (o *ConnectionNetsuiteV2) GetDefaultUpdateSchedule() []ConnectionNetsuiteV2DefaultUpdateSchedule {
	if o == nil {
		return []ConnectionNetsuiteV2DefaultUpdateSchedule{}
	}
	return o.DefaultUpdateSchedule
}

func (o *ConnectionNetsuiteV2) GetActive() bool {
	if o == nil {
		return false
	}
	return o.Active
}

func (o *ConnectionNetsuiteV2) GetType() ConnectionNetsuiteV2Type {
	if o == nil {
		return ConnectionNetsuiteV2Type("")
	}
	return o.Type
}

func (o *ConnectionNetsuiteV2) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ConnectionNetsuiteV2) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionNetsuiteV2) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionNetsuiteV2) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionNetsuiteV2) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionNetsuiteV2) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionNetsuiteV2) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionNetsuiteV2) GetTokenID() string {
	if o == nil {
		return ""
	}
	return o.TokenID
}

func (o *ConnectionNetsuiteV2) GetConsumerKey() string {
	if o == nil {
		return ""
	}
	return o.ConsumerKey
}

func (o *ConnectionNetsuiteV2) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

type ConnectionNetsuiteV2Input struct {
	// The unique name of this connection.
	Name string                   `json:"name"`
	Type ConnectionNetsuiteV2Type `json:"type"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// Token Secret value displayed after creating a new access token.
	TokenSecret string `json:"tokenSecret"`
	// Token Id value displayed after creating a new access token.
	TokenID string `json:"tokenId"`
	// Consumer Key value displayed after creating a new integration.
	ConsumerKey string `json:"consumerKey"`
	// Consumer Secret value displayed after creating a new integration.
	ConsumerSecret string `json:"consumerSecret"`
	// Under Setup > Company > Company Information
	AccountID string `json:"accountId"`
}

func (o *ConnectionNetsuiteV2Input) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionNetsuiteV2Input) GetType() ConnectionNetsuiteV2Type {
	if o == nil {
		return ConnectionNetsuiteV2Type("")
	}
	return o.Type
}

func (o *ConnectionNetsuiteV2Input) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionNetsuiteV2Input) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionNetsuiteV2Input) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionNetsuiteV2Input) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionNetsuiteV2Input) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionNetsuiteV2Input) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionNetsuiteV2Input) GetTokenSecret() string {
	if o == nil {
		return ""
	}
	return o.TokenSecret
}

func (o *ConnectionNetsuiteV2Input) GetTokenID() string {
	if o == nil {
		return ""
	}
	return o.TokenID
}

func (o *ConnectionNetsuiteV2Input) GetConsumerKey() string {
	if o == nil {
		return ""
	}
	return o.ConsumerKey
}

func (o *ConnectionNetsuiteV2Input) GetConsumerSecret() string {
	if o == nil {
		return ""
	}
	return o.ConsumerSecret
}

func (o *ConnectionNetsuiteV2Input) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}
