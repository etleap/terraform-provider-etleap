// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionLdapVirtualListViewUpdateType string

const (
	ConnectionLdapVirtualListViewUpdateTypeLdapVirtualListView ConnectionLdapVirtualListViewUpdateType = "LDAP_VIRTUAL_LIST_VIEW"
)

func (e ConnectionLdapVirtualListViewUpdateType) ToPointer() *ConnectionLdapVirtualListViewUpdateType {
	return &e
}

func (e *ConnectionLdapVirtualListViewUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "LDAP_VIRTUAL_LIST_VIEW":
		*e = ConnectionLdapVirtualListViewUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionLdapVirtualListViewUpdateType: %v", v)
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

type ConnectionLdapVirtualListViewUpdate struct {
	// The unique name of this connection.
	Name *string                                  `json:"name,omitempty"`
	Type *ConnectionLdapVirtualListViewUpdateType `json:"type"`
	// Whether this connection should be marked as active.
	Active *bool `json:"active,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// LDAP server name or ip address.
	Hostname *string `json:"hostname,omitempty"`
	Port     *int64  `json:"port,omitempty"`
	// Enable this if you are using a secure port to connect to LDAP. Usually, this should be enabled if you are connecting via port 636.
	UseSsl *bool `json:"useSsl,omitempty"`
	// The login string to use, similar to 'uid=admin,ou=system'.
	User     *string `json:"user,omitempty"`
	Password *string `json:"password,omitempty"`
	// The path of your DIT, typically a DC, OU, or O entry.
	BaseDn *string `json:"baseDn,omitempty"`
	// Indicates the set of entries that may be considered potential matches for the search.
	Scope *Scope `json:"scope,omitempty"`
	// The filter to be used in the search, e.g. (uid=*)
	Filter *string `json:"filter,omitempty"`
	// The fields (comma-delimited) to be used for sorting the result. For descending order, you should append "-" to the field. E.g.: "uid,-cn".
	SortOrder *string `json:"sortOrder,omitempty"`
}

func (o *ConnectionLdapVirtualListViewUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionLdapVirtualListViewUpdate) GetType() *ConnectionLdapVirtualListViewUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ConnectionLdapVirtualListViewUpdate) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ConnectionLdapVirtualListViewUpdate) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionLdapVirtualListViewUpdate) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionLdapVirtualListViewUpdate) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionLdapVirtualListViewUpdate) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionLdapVirtualListViewUpdate) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionLdapVirtualListViewUpdate) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionLdapVirtualListViewUpdate) GetHostname() *string {
	if o == nil {
		return nil
	}
	return o.Hostname
}

func (o *ConnectionLdapVirtualListViewUpdate) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *ConnectionLdapVirtualListViewUpdate) GetUseSsl() *bool {
	if o == nil {
		return nil
	}
	return o.UseSsl
}

func (o *ConnectionLdapVirtualListViewUpdate) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}

func (o *ConnectionLdapVirtualListViewUpdate) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *ConnectionLdapVirtualListViewUpdate) GetBaseDn() *string {
	if o == nil {
		return nil
	}
	return o.BaseDn
}

func (o *ConnectionLdapVirtualListViewUpdate) GetScope() *Scope {
	if o == nil {
		return nil
	}
	return o.Scope
}

func (o *ConnectionLdapVirtualListViewUpdate) GetFilter() *string {
	if o == nil {
		return nil
	}
	return o.Filter
}

func (o *ConnectionLdapVirtualListViewUpdate) GetSortOrder() *string {
	if o == nil {
		return nil
	}
	return o.SortOrder
}