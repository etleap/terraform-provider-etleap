// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
	"time"
)

// ConnectionLdapVirtualListViewStatus - The current status of the connection.
type ConnectionLdapVirtualListViewStatus string

const (
	ConnectionLdapVirtualListViewStatusUnknown     ConnectionLdapVirtualListViewStatus = "UNKNOWN"
	ConnectionLdapVirtualListViewStatusUp          ConnectionLdapVirtualListViewStatus = "UP"
	ConnectionLdapVirtualListViewStatusDown        ConnectionLdapVirtualListViewStatus = "DOWN"
	ConnectionLdapVirtualListViewStatusResize      ConnectionLdapVirtualListViewStatus = "RESIZE"
	ConnectionLdapVirtualListViewStatusMaintenance ConnectionLdapVirtualListViewStatus = "MAINTENANCE"
	ConnectionLdapVirtualListViewStatusQuota       ConnectionLdapVirtualListViewStatus = "QUOTA"
	ConnectionLdapVirtualListViewStatusCreating    ConnectionLdapVirtualListViewStatus = "CREATING"
)

func (e ConnectionLdapVirtualListViewStatus) ToPointer() *ConnectionLdapVirtualListViewStatus {
	return &e
}

func (e *ConnectionLdapVirtualListViewStatus) UnmarshalJSON(data []byte) error {
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
		*e = ConnectionLdapVirtualListViewStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionLdapVirtualListViewStatus: %v", v)
	}
}

type ConnectionLdapVirtualListViewDefaultUpdateSchedule struct {
	// The pipeline mode refers to how the pipeline fetches data changes from the source and how those changes are applied to the destination table. See <a target="_blank" href="https://docs.etleap.com/docs/documentation/ZG9jOjIyMjE3ODA2-introduction">the documentation</a> for more details.
	PipelineMode *PipelineUpdateModes `json:"pipelineMode,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
}

func (o *ConnectionLdapVirtualListViewDefaultUpdateSchedule) GetPipelineMode() *PipelineUpdateModes {
	if o == nil {
		return nil
	}
	return o.PipelineMode
}

func (o *ConnectionLdapVirtualListViewDefaultUpdateSchedule) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionLdapVirtualListViewDefaultUpdateSchedule) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionLdapVirtualListViewDefaultUpdateSchedule) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionLdapVirtualListViewDefaultUpdateSchedule) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionLdapVirtualListViewDefaultUpdateSchedule) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionLdapVirtualListViewDefaultUpdateSchedule) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

type ConnectionLdapVirtualListViewType string

const (
	ConnectionLdapVirtualListViewTypeLdapVirtualListView ConnectionLdapVirtualListViewType = "LDAP_VIRTUAL_LIST_VIEW"
)

func (e ConnectionLdapVirtualListViewType) ToPointer() *ConnectionLdapVirtualListViewType {
	return &e
}

func (e *ConnectionLdapVirtualListViewType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "LDAP_VIRTUAL_LIST_VIEW":
		*e = ConnectionLdapVirtualListViewType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionLdapVirtualListViewType: %v", v)
	}
}

// Scope - Indicates the set of entries that may be considered potential matches for the search.
type Scope string

const (
	ScopeBase        Scope = "Base"
	ScopeSingleLevel Scope = "Single-level"
	ScopeSubtree     Scope = "Subtree"
)

func (e Scope) ToPointer() *Scope {
	return &e
}

func (e *Scope) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Base":
		fallthrough
	case "Single-level":
		fallthrough
	case "Subtree":
		*e = Scope(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Scope: %v", v)
	}
}

type ConnectionLdapVirtualListView struct {
	// The current status of the connection.
	Status ConnectionLdapVirtualListViewStatus `json:"status"`
	// The unique name of this connection.
	Name string `json:"name"`
	// The date and time when then the connection was created.
	CreateDate time.Time `json:"createDate"`
	// When an update schedule is not defined for a connection, the default schedule is used. The default defined individually per `pipelineMode` and may be subject to change.
	DefaultUpdateSchedule []ConnectionLdapVirtualListViewDefaultUpdateSchedule `json:"defaultUpdateSchedule"`
	// Whether this connection has been marked as active.
	Active bool                              `json:"active"`
	Type   ConnectionLdapVirtualListViewType `json:"type"`
	// The unique identifier of the connection.
	ID string `json:"id"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// The path of your DIT, typically a DC, OU, or O entry.
	BaseDn string `json:"baseDn"`
	// Enable this if you are using a secure port to connect to LDAP. Usually, this should be enabled if you are connecting via port 636.
	UseSsl bool `json:"useSsl"`
	// LDAP server name or ip address.
	Hostname string `json:"hostname"`
	// The filter to be used in the search, e.g. (uid=*)
	Filter string `json:"filter"`
	// The login string to use, similar to 'uid=admin,ou=system'.
	User string `json:"user"`
	// Indicates the set of entries that may be considered potential matches for the search.
	Scope Scope `json:"scope"`
	Port  int64 `json:"port"`
	// The fields (comma-delimited) to be used for sorting the result. For descending order, you should append "-" to the field. E.g.: "uid,-cn".
	SortOrder string `json:"sortOrder"`
}

func (c ConnectionLdapVirtualListView) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionLdapVirtualListView) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionLdapVirtualListView) GetStatus() ConnectionLdapVirtualListViewStatus {
	if o == nil {
		return ConnectionLdapVirtualListViewStatus("")
	}
	return o.Status
}

func (o *ConnectionLdapVirtualListView) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionLdapVirtualListView) GetCreateDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreateDate
}

func (o *ConnectionLdapVirtualListView) GetDefaultUpdateSchedule() []ConnectionLdapVirtualListViewDefaultUpdateSchedule {
	if o == nil {
		return []ConnectionLdapVirtualListViewDefaultUpdateSchedule{}
	}
	return o.DefaultUpdateSchedule
}

func (o *ConnectionLdapVirtualListView) GetActive() bool {
	if o == nil {
		return false
	}
	return o.Active
}

func (o *ConnectionLdapVirtualListView) GetType() ConnectionLdapVirtualListViewType {
	if o == nil {
		return ConnectionLdapVirtualListViewType("")
	}
	return o.Type
}

func (o *ConnectionLdapVirtualListView) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ConnectionLdapVirtualListView) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionLdapVirtualListView) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionLdapVirtualListView) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionLdapVirtualListView) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionLdapVirtualListView) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionLdapVirtualListView) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionLdapVirtualListView) GetBaseDn() string {
	if o == nil {
		return ""
	}
	return o.BaseDn
}

func (o *ConnectionLdapVirtualListView) GetUseSsl() bool {
	if o == nil {
		return false
	}
	return o.UseSsl
}

func (o *ConnectionLdapVirtualListView) GetHostname() string {
	if o == nil {
		return ""
	}
	return o.Hostname
}

func (o *ConnectionLdapVirtualListView) GetFilter() string {
	if o == nil {
		return ""
	}
	return o.Filter
}

func (o *ConnectionLdapVirtualListView) GetUser() string {
	if o == nil {
		return ""
	}
	return o.User
}

func (o *ConnectionLdapVirtualListView) GetScope() Scope {
	if o == nil {
		return Scope("")
	}
	return o.Scope
}

func (o *ConnectionLdapVirtualListView) GetPort() int64 {
	if o == nil {
		return 0
	}
	return o.Port
}

func (o *ConnectionLdapVirtualListView) GetSortOrder() string {
	if o == nil {
		return ""
	}
	return o.SortOrder
}

type ConnectionLdapVirtualListViewInput struct {
	// The unique name of this connection.
	Name string                            `json:"name"`
	Type ConnectionLdapVirtualListViewType `json:"type"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// The path of your DIT, typically a DC, OU, or O entry.
	BaseDn string `json:"baseDn"`
	// Enable this if you are using a secure port to connect to LDAP. Usually, this should be enabled if you are connecting via port 636.
	UseSsl bool `json:"useSsl"`
	// LDAP server name or ip address.
	Hostname string `json:"hostname"`
	// The filter to be used in the search, e.g. (uid=*)
	Filter string `json:"filter"`
	// The login string to use, similar to 'uid=admin,ou=system'.
	User string `json:"user"`
	// Indicates the set of entries that may be considered potential matches for the search.
	Scope    Scope  `json:"scope"`
	Password string `json:"password"`
	Port     int64  `json:"port"`
	// The fields (comma-delimited) to be used for sorting the result. For descending order, you should append "-" to the field. E.g.: "uid,-cn".
	SortOrder string `json:"sortOrder"`
}

func (o *ConnectionLdapVirtualListViewInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionLdapVirtualListViewInput) GetType() ConnectionLdapVirtualListViewType {
	if o == nil {
		return ConnectionLdapVirtualListViewType("")
	}
	return o.Type
}

func (o *ConnectionLdapVirtualListViewInput) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionLdapVirtualListViewInput) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionLdapVirtualListViewInput) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionLdapVirtualListViewInput) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionLdapVirtualListViewInput) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionLdapVirtualListViewInput) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionLdapVirtualListViewInput) GetBaseDn() string {
	if o == nil {
		return ""
	}
	return o.BaseDn
}

func (o *ConnectionLdapVirtualListViewInput) GetUseSsl() bool {
	if o == nil {
		return false
	}
	return o.UseSsl
}

func (o *ConnectionLdapVirtualListViewInput) GetHostname() string {
	if o == nil {
		return ""
	}
	return o.Hostname
}

func (o *ConnectionLdapVirtualListViewInput) GetFilter() string {
	if o == nil {
		return ""
	}
	return o.Filter
}

func (o *ConnectionLdapVirtualListViewInput) GetUser() string {
	if o == nil {
		return ""
	}
	return o.User
}

func (o *ConnectionLdapVirtualListViewInput) GetScope() Scope {
	if o == nil {
		return Scope("")
	}
	return o.Scope
}

func (o *ConnectionLdapVirtualListViewInput) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *ConnectionLdapVirtualListViewInput) GetPort() int64 {
	if o == nil {
		return 0
	}
	return o.Port
}

func (o *ConnectionLdapVirtualListViewInput) GetSortOrder() string {
	if o == nil {
		return ""
	}
	return o.SortOrder
}
