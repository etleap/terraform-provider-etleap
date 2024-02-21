// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionUserVoiceUpdateType string

const (
	ConnectionUserVoiceUpdateTypeUservoice ConnectionUserVoiceUpdateType = "USERVOICE"
)

func (e ConnectionUserVoiceUpdateType) ToPointer() *ConnectionUserVoiceUpdateType {
	return &e
}

func (e *ConnectionUserVoiceUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "USERVOICE":
		*e = ConnectionUserVoiceUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionUserVoiceUpdateType: %v", v)
	}
}

type ConnectionUserVoiceUpdate struct {
	// The unique name of this connection.
	Name *string                        `json:"name,omitempty"`
	Type *ConnectionUserVoiceUpdateType `json:"type"`
	// Whether this connection should be marked as active.
	Active *bool `json:"active,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// Your current UserVoice subdomain (i.e. SUBDOMAIN.uservoice.com)
	Subdomain *string `json:"subdomain,omitempty"`
	// A client access token to connect to UserVoice API. It can be obtained from your UserVoice API key. If you need help to create a client access token, <a target="_blank" href="https://developer.uservoice.com/docs/api/api-key/">follow this article</a>.
	AccessToken *string `json:"accessToken,omitempty"`
}

func (o *ConnectionUserVoiceUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionUserVoiceUpdate) GetType() *ConnectionUserVoiceUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ConnectionUserVoiceUpdate) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ConnectionUserVoiceUpdate) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionUserVoiceUpdate) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionUserVoiceUpdate) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionUserVoiceUpdate) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionUserVoiceUpdate) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionUserVoiceUpdate) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionUserVoiceUpdate) GetSubdomain() *string {
	if o == nil {
		return nil
	}
	return o.Subdomain
}

func (o *ConnectionUserVoiceUpdate) GetAccessToken() *string {
	if o == nil {
		return nil
	}
	return o.AccessToken
}