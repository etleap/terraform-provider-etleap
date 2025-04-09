// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
	"time"
)

// ConnectionVeevaStatus - The current status of the connection.
type ConnectionVeevaStatus string

const (
	ConnectionVeevaStatusUnknown     ConnectionVeevaStatus = "UNKNOWN"
	ConnectionVeevaStatusUp          ConnectionVeevaStatus = "UP"
	ConnectionVeevaStatusDown        ConnectionVeevaStatus = "DOWN"
	ConnectionVeevaStatusResize      ConnectionVeevaStatus = "RESIZE"
	ConnectionVeevaStatusMaintenance ConnectionVeevaStatus = "MAINTENANCE"
	ConnectionVeevaStatusQuota       ConnectionVeevaStatus = "QUOTA"
	ConnectionVeevaStatusCreating    ConnectionVeevaStatus = "CREATING"
)

func (e ConnectionVeevaStatus) ToPointer() *ConnectionVeevaStatus {
	return &e
}

func (e *ConnectionVeevaStatus) UnmarshalJSON(data []byte) error {
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
		*e = ConnectionVeevaStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionVeevaStatus: %v", v)
	}
}

type ConnectionVeevaDefaultUpdateSchedule struct {
	// The pipeline mode refers to how the pipeline fetches data changes from the source and how those changes are applied to the destination table. See <a target="_blank" href="https://docs.etleap.com/docs/documentation/ZG9jOjIyMjE3ODA2-introduction">the documentation</a> for more details.
	PipelineMode *PipelineUpdateModes `json:"pipelineMode,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
}

func (o *ConnectionVeevaDefaultUpdateSchedule) GetPipelineMode() *PipelineUpdateModes {
	if o == nil {
		return nil
	}
	return o.PipelineMode
}

func (o *ConnectionVeevaDefaultUpdateSchedule) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionVeevaDefaultUpdateSchedule) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionVeevaDefaultUpdateSchedule) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionVeevaDefaultUpdateSchedule) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionVeevaDefaultUpdateSchedule) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionVeevaDefaultUpdateSchedule) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

type ConnectionVeevaType string

const (
	ConnectionVeevaTypeVeeva ConnectionVeevaType = "VEEVA"
)

func (e ConnectionVeevaType) ToPointer() *ConnectionVeevaType {
	return &e
}

func (e *ConnectionVeevaType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "VEEVA":
		*e = ConnectionVeevaType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionVeevaType: %v", v)
	}
}

// VaultType - Your Veeva Vault type. Currently supported types are: QUALITY, CTMS, and RIM. If a value is not provided it will fallback to the default QUALITY
type VaultType string

const (
	VaultTypeQuality   VaultType = "QUALITY"
	VaultTypeCtms      VaultType = "CTMS"
	VaultTypeRim       VaultType = "RIM"
	VaultTypePromomats VaultType = "PROMOMATS"
)

func (e VaultType) ToPointer() *VaultType {
	return &e
}

func (e *VaultType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "QUALITY":
		fallthrough
	case "CTMS":
		fallthrough
	case "RIM":
		fallthrough
	case "PROMOMATS":
		*e = VaultType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for VaultType: %v", v)
	}
}

type ConnectionVeeva struct {
	// The current status of the connection.
	Status ConnectionVeevaStatus `json:"status"`
	// The unique name of this connection.
	Name string `json:"name"`
	// The date and time when then the connection was created.
	CreateDate time.Time `json:"createDate"`
	// When an update schedule is not defined for a connection, the default schedule is used. The default defined individually per `pipelineMode` and may be subject to change.
	DefaultUpdateSchedule []ConnectionVeevaDefaultUpdateSchedule `json:"defaultUpdateSchedule"`
	// Whether this connection has been marked as active.
	Active bool                `json:"active"`
	Type   ConnectionVeevaType `json:"type"`
	// The unique identifier of the connection.
	ID string `json:"id"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	Username       string               `json:"username"`
	// Your Veeva Vault type. Currently supported types are: QUALITY, CTMS, and RIM. If a value is not provided it will fallback to the default QUALITY
	VaultType *VaultType `default:"QUALITY" json:"vaultType"`
	// The vault domain name is part of the URL that you use to access your Veeva Vault. You can follow the<a target="_blank" href="https://developer.veevavault.com/docs/#authentication">'Structuring the Endpoint'</a> instructions to find the URL.
	VaultDomainName string `json:"vaultDomainName"`
}

func (c ConnectionVeeva) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionVeeva) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionVeeva) GetStatus() ConnectionVeevaStatus {
	if o == nil {
		return ConnectionVeevaStatus("")
	}
	return o.Status
}

func (o *ConnectionVeeva) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionVeeva) GetCreateDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreateDate
}

func (o *ConnectionVeeva) GetDefaultUpdateSchedule() []ConnectionVeevaDefaultUpdateSchedule {
	if o == nil {
		return []ConnectionVeevaDefaultUpdateSchedule{}
	}
	return o.DefaultUpdateSchedule
}

func (o *ConnectionVeeva) GetActive() bool {
	if o == nil {
		return false
	}
	return o.Active
}

func (o *ConnectionVeeva) GetType() ConnectionVeevaType {
	if o == nil {
		return ConnectionVeevaType("")
	}
	return o.Type
}

func (o *ConnectionVeeva) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ConnectionVeeva) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionVeeva) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionVeeva) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionVeeva) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionVeeva) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionVeeva) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionVeeva) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *ConnectionVeeva) GetVaultType() *VaultType {
	if o == nil {
		return nil
	}
	return o.VaultType
}

func (o *ConnectionVeeva) GetVaultDomainName() string {
	if o == nil {
		return ""
	}
	return o.VaultDomainName
}

type ConnectionVeevaInput struct {
	// The unique name of this connection.
	Name string              `json:"name"`
	Type ConnectionVeevaType `json:"type"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	Username       string               `json:"username"`
	Password       string               `json:"password"`
	// Your Veeva Vault type. Currently supported types are: QUALITY, CTMS, and RIM. If a value is not provided it will fallback to the default QUALITY
	VaultType *VaultType `default:"QUALITY" json:"vaultType"`
	// The vault domain name is part of the URL that you use to access your Veeva Vault. You can follow the<a target="_blank" href="https://developer.veevavault.com/docs/#authentication">'Structuring the Endpoint'</a> instructions to find the URL.
	VaultDomainName string `json:"vaultDomainName"`
}

func (c ConnectionVeevaInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionVeevaInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionVeevaInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionVeevaInput) GetType() ConnectionVeevaType {
	if o == nil {
		return ConnectionVeevaType("")
	}
	return o.Type
}

func (o *ConnectionVeevaInput) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionVeevaInput) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionVeevaInput) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionVeevaInput) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionVeevaInput) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionVeevaInput) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionVeevaInput) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *ConnectionVeevaInput) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *ConnectionVeevaInput) GetVaultType() *VaultType {
	if o == nil {
		return nil
	}
	return o.VaultType
}

func (o *ConnectionVeevaInput) GetVaultDomainName() string {
	if o == nil {
		return ""
	}
	return o.VaultDomainName
}
