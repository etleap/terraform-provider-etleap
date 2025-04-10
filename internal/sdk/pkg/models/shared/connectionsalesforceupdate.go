// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionSalesforceUpdateType string

const (
	ConnectionSalesforceUpdateTypeSalesforce ConnectionSalesforceUpdateType = "SALESFORCE"
)

func (e ConnectionSalesforceUpdateType) ToPointer() *ConnectionSalesforceUpdateType {
	return &e
}

func (e *ConnectionSalesforceUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SALESFORCE":
		*e = ConnectionSalesforceUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionSalesforceUpdateType: %v", v)
	}
}

type ConnectionSalesforceUpdate struct {
	// Whether this connection should be marked as active.
	Active *bool                          `json:"active,omitempty"`
	Type   ConnectionSalesforceUpdateType `json:"type"`
	// The unique name of this connection.
	Name *string `json:"name,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// Code retrieved from `/connections/oauth2-initiation`. **Note:** it is short-lived, therefore the connection creation should be done as soon as code is returned.
	Code    *string `json:"code,omitempty"`
	Sandbox *bool   `json:"sandbox,omitempty"`
	// The maximum number of Bulk API batch requests Etleap will use per 24-hour window. Etleap will also stop if the remaining quota goes under 500 requests.
	QuotaLimit *int64 `json:"quotaLimit,omitempty"`
}

func (o *ConnectionSalesforceUpdate) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ConnectionSalesforceUpdate) GetType() ConnectionSalesforceUpdateType {
	if o == nil {
		return ConnectionSalesforceUpdateType("")
	}
	return o.Type
}

func (o *ConnectionSalesforceUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionSalesforceUpdate) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionSalesforceUpdate) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionSalesforceUpdate) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionSalesforceUpdate) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionSalesforceUpdate) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionSalesforceUpdate) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionSalesforceUpdate) GetCode() *string {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *ConnectionSalesforceUpdate) GetSandbox() *bool {
	if o == nil {
		return nil
	}
	return o.Sandbox
}

func (o *ConnectionSalesforceUpdate) GetQuotaLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.QuotaLimit
}
