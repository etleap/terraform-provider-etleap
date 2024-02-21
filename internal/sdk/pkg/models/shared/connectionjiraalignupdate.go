// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionJiraAlignUpdateType string

const (
	ConnectionJiraAlignUpdateTypeJiraAlign ConnectionJiraAlignUpdateType = "JIRA_ALIGN"
)

func (e ConnectionJiraAlignUpdateType) ToPointer() *ConnectionJiraAlignUpdateType {
	return &e
}

func (e *ConnectionJiraAlignUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "JIRA_ALIGN":
		*e = ConnectionJiraAlignUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionJiraAlignUpdateType: %v", v)
	}
}

type ConnectionJiraAlignUpdate struct {
	// The unique name of this connection.
	Name *string                        `json:"name,omitempty"`
	Type *ConnectionJiraAlignUpdateType `json:"type"`
	// Whether this connection should be marked as active.
	Active *bool `json:"active,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// Your JIRA Align subdomain (i.e. SUBDOMAIN.jiraalign.com)
	Subdomain *string `json:"subdomain,omitempty"`
	// The client API key required to authenticate with the JIRA Align API.
	APIKey *string `json:"apiKey,omitempty"`
}

func (o *ConnectionJiraAlignUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionJiraAlignUpdate) GetType() *ConnectionJiraAlignUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ConnectionJiraAlignUpdate) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ConnectionJiraAlignUpdate) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionJiraAlignUpdate) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionJiraAlignUpdate) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionJiraAlignUpdate) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionJiraAlignUpdate) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionJiraAlignUpdate) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionJiraAlignUpdate) GetSubdomain() *string {
	if o == nil {
		return nil
	}
	return o.Subdomain
}

func (o *ConnectionJiraAlignUpdate) GetAPIKey() *string {
	if o == nil {
		return nil
	}
	return o.APIKey
}