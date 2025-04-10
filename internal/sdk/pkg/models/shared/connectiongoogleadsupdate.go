// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionGoogleAdsUpdateType string

const (
	ConnectionGoogleAdsUpdateTypeGoogleAds ConnectionGoogleAdsUpdateType = "GOOGLE_ADS"
)

func (e ConnectionGoogleAdsUpdateType) ToPointer() *ConnectionGoogleAdsUpdateType {
	return &e
}

func (e *ConnectionGoogleAdsUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "GOOGLE_ADS":
		*e = ConnectionGoogleAdsUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionGoogleAdsUpdateType: %v", v)
	}
}

type ConnectionGoogleAdsUpdate struct {
	// Whether this connection should be marked as active.
	Active *bool                         `json:"active,omitempty"`
	Type   ConnectionGoogleAdsUpdateType `json:"type"`
	// The unique name of this connection.
	Name *string `json:"name,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// Code retrieved from `/connections/oauth2-initiation`. **Note:** it is short-lived, therefore the connection creation should be done as soon as code is returned.
	Code *string `json:"code,omitempty"`
	// The ID of the customer you want to log in as. Etleap will extract the data of all customers that this login customer has access to. 10 digits without hyphens.
	CustomerID *string `json:"customerId,omitempty"`
}

func (o *ConnectionGoogleAdsUpdate) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ConnectionGoogleAdsUpdate) GetType() ConnectionGoogleAdsUpdateType {
	if o == nil {
		return ConnectionGoogleAdsUpdateType("")
	}
	return o.Type
}

func (o *ConnectionGoogleAdsUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionGoogleAdsUpdate) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionGoogleAdsUpdate) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionGoogleAdsUpdate) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionGoogleAdsUpdate) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionGoogleAdsUpdate) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionGoogleAdsUpdate) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionGoogleAdsUpdate) GetCode() *string {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *ConnectionGoogleAdsUpdate) GetCustomerID() *string {
	if o == nil {
		return nil
	}
	return o.CustomerID
}
