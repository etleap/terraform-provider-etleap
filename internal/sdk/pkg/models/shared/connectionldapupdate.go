// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionLdapUpdateType string

const (
	ConnectionLdapUpdateTypeLdap ConnectionLdapUpdateType = "LDAP"
)

func (e ConnectionLdapUpdateType) ToPointer() *ConnectionLdapUpdateType {
	return &e
}

func (e *ConnectionLdapUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "LDAP":
		*e = ConnectionLdapUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionLdapUpdateType: %v", v)
	}
}

type ConnectionLdapUpdate struct {
	// The unique name of this connection.
	Name *string                   `json:"name,omitempty"`
	Type *ConnectionLdapUpdateType `json:"type"`
	// Whether this connection should be marked as active.
	Active *bool `json:"active,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// LDAP server name or ip address
	Hostname *string `json:"hostname,omitempty"`
	Port     *int64  `json:"port,omitempty"`
	// Enable this if you are using a secure port to connect to LDAP. Usually, this should be enabled if you are connecting via port 636.
	UseSsl *bool `json:"useSsl,omitempty"`
	// The login string to use, similar to 'uid=admin,ou=system'
	User     *string `json:"user,omitempty"`
	Password *string `json:"password,omitempty"`
	// The IANA-assigned number for your custom schema entities
	Pen *int64 `json:"pen,omitempty"`
	// The path of your DIT, typically a DC, OU, or O entry
	BaseDn *string `json:"baseDn,omitempty"`
}

func (o *ConnectionLdapUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionLdapUpdate) GetType() *ConnectionLdapUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ConnectionLdapUpdate) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ConnectionLdapUpdate) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionLdapUpdate) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionLdapUpdate) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionLdapUpdate) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionLdapUpdate) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionLdapUpdate) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionLdapUpdate) GetHostname() *string {
	if o == nil {
		return nil
	}
	return o.Hostname
}

func (o *ConnectionLdapUpdate) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *ConnectionLdapUpdate) GetUseSsl() *bool {
	if o == nil {
		return nil
	}
	return o.UseSsl
}

func (o *ConnectionLdapUpdate) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}

func (o *ConnectionLdapUpdate) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *ConnectionLdapUpdate) GetPen() *int64 {
	if o == nil {
		return nil
	}
	return o.Pen
}

func (o *ConnectionLdapUpdate) GetBaseDn() *string {
	if o == nil {
		return nil
	}
	return o.BaseDn
}