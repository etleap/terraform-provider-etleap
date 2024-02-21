// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionCriteoUpdateType string

const (
	ConnectionCriteoUpdateTypeCriteo ConnectionCriteoUpdateType = "CRITEO"
)

func (e ConnectionCriteoUpdateType) ToPointer() *ConnectionCriteoUpdateType {
	return &e
}

func (e *ConnectionCriteoUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CRITEO":
		*e = ConnectionCriteoUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionCriteoUpdateType: %v", v)
	}
}

type ConnectionCriteoUpdate struct {
	// The unique name of this connection.
	Name *string                     `json:"name,omitempty"`
	Type *ConnectionCriteoUpdateType `json:"type"`
	// Whether this connection should be marked as active.
	Active *bool `json:"active,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// Your Client Id can be found under 'Developer Dashboard' > 'My apps'
	ClientID *string `json:"clientId,omitempty"`
	// Your Client Secret can be found under 'Developer Dashboard' > 'My apps'
	ClientSecret *string `json:"clientSecret,omitempty"`
}

func (o *ConnectionCriteoUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionCriteoUpdate) GetType() *ConnectionCriteoUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ConnectionCriteoUpdate) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ConnectionCriteoUpdate) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionCriteoUpdate) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionCriteoUpdate) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionCriteoUpdate) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionCriteoUpdate) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionCriteoUpdate) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionCriteoUpdate) GetClientID() *string {
	if o == nil {
		return nil
	}
	return o.ClientID
}

func (o *ConnectionCriteoUpdate) GetClientSecret() *string {
	if o == nil {
		return nil
	}
	return o.ClientSecret
}