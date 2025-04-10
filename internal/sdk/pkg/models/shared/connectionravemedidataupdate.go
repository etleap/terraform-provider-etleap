// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionRaveMedidataUpdateType string

const (
	ConnectionRaveMedidataUpdateTypeRaveMedidata ConnectionRaveMedidataUpdateType = "RAVE_MEDIDATA"
)

func (e ConnectionRaveMedidataUpdateType) ToPointer() *ConnectionRaveMedidataUpdateType {
	return &e
}

func (e *ConnectionRaveMedidataUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "RAVE_MEDIDATA":
		*e = ConnectionRaveMedidataUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionRaveMedidataUpdateType: %v", v)
	}
}

type ConnectionRaveMedidataUpdate struct {
	// Whether this connection should be marked as active.
	Active *bool                             `json:"active,omitempty"`
	Type   *ConnectionRaveMedidataUpdateType `json:"type"`
	// The unique name of this connection.
	Name *string `json:"name,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// Rave Account Login
	Username *string `json:"username,omitempty"`
	// Rave Account Password
	Password *string `json:"password,omitempty"`
	// E.g., my-org.mdsol.com
	Hostname *string `json:"hostname,omitempty"`
}

func (o *ConnectionRaveMedidataUpdate) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ConnectionRaveMedidataUpdate) GetType() *ConnectionRaveMedidataUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ConnectionRaveMedidataUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionRaveMedidataUpdate) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionRaveMedidataUpdate) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionRaveMedidataUpdate) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionRaveMedidataUpdate) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionRaveMedidataUpdate) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionRaveMedidataUpdate) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionRaveMedidataUpdate) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *ConnectionRaveMedidataUpdate) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *ConnectionRaveMedidataUpdate) GetHostname() *string {
	if o == nil {
		return nil
	}
	return o.Hostname
}
