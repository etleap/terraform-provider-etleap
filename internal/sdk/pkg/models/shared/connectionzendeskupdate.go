// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionZendeskUpdateType string

const (
	ConnectionZendeskUpdateTypeZendesk ConnectionZendeskUpdateType = "ZENDESK"
)

func (e ConnectionZendeskUpdateType) ToPointer() *ConnectionZendeskUpdateType {
	return &e
}

func (e *ConnectionZendeskUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ZENDESK":
		*e = ConnectionZendeskUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionZendeskUpdateType: %v", v)
	}
}

type ConnectionZendeskUpdate struct {
	// Whether this connection should be marked as active.
	Active *bool                       `json:"active,omitempty"`
	Type   ConnectionZendeskUpdateType `json:"type"`
	// The unique name of this connection.
	Name *string `json:"name,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// Code retrieved from `/connections/oauth2-initiation`. **Note:** it is short-lived, therefore the connection creation should be done as soon as code is returned.
	Code *string `json:"code,omitempty"`
	// The name of the subdomain. This is in the URL when you use zendesk: [subdomain].zendesk.com
	Subdomain *string `json:"subdomain,omitempty"`
}

func (o *ConnectionZendeskUpdate) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ConnectionZendeskUpdate) GetType() ConnectionZendeskUpdateType {
	if o == nil {
		return ConnectionZendeskUpdateType("")
	}
	return o.Type
}

func (o *ConnectionZendeskUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionZendeskUpdate) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionZendeskUpdate) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionZendeskUpdate) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionZendeskUpdate) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionZendeskUpdate) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionZendeskUpdate) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionZendeskUpdate) GetCode() *string {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *ConnectionZendeskUpdate) GetSubdomain() *string {
	if o == nil {
		return nil
	}
	return o.Subdomain
}
