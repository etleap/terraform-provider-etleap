// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionMarketoUpdateType string

const (
	ConnectionMarketoUpdateTypeMarketo ConnectionMarketoUpdateType = "MARKETO"
)

func (e ConnectionMarketoUpdateType) ToPointer() *ConnectionMarketoUpdateType {
	return &e
}

func (e *ConnectionMarketoUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MARKETO":
		*e = ConnectionMarketoUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionMarketoUpdateType: %v", v)
	}
}

type ConnectionMarketoUpdate struct {
	// Whether this connection should be marked as active.
	Active *bool                        `json:"active,omitempty"`
	Type   *ConnectionMarketoUpdateType `json:"type"`
	// The unique name of this connection.
	Name *string `json:"name,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// The maximum number of requests Etleap will use per day. Your Marketo account's daily quota is the 'Daily Request Limit' number under Admin -> Integration -> Web Services. We recommend setting this to 75% of your Marketo limit.
	QuotaLimit *int64 `json:"quotaLimit,omitempty"`
	// Under Admin -> Integration -> LaunchPoint, you can find this value by clicking 'View Details'.
	RestClientID *string `json:"restClientId,omitempty"`
	// E.g. 'https://259-ZDK-675.mktoapi.com/soap/mktows/2_9'. In the Marketo UI this is the 'Endpoint' value in the 'SOAP API' section.
	SoapEndpoint *string `json:"soapEndpoint,omitempty"`
	// E.g. 'https://259-ZDK-675.mktorest.com/rest'. In the Marketo UI this is the 'Endpoint' value in the 'REST API' section.
	RestEndpoint *string `json:"restEndpoint,omitempty"`
	// Under Admin -> Integration -> LaunchPoint, you can find this value by clicking 'View Details'.
	RestClientSecret *string `json:"restClientSecret,omitempty"`
	// E.g. 'MKTOWS_259-ZDK-675_1'. In the Marketo UI this is the 'User ID' value in the 'SOAP API' section.
	SoapUserID *string `json:"soapUserId,omitempty"`
	// In the Marketo UI this is the 'Encryption Key' value in the 'SOAP API' section.
	SoapEncryptionKey *string `json:"soapEncryptionKey,omitempty"`
}

func (o *ConnectionMarketoUpdate) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ConnectionMarketoUpdate) GetType() *ConnectionMarketoUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ConnectionMarketoUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionMarketoUpdate) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionMarketoUpdate) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionMarketoUpdate) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionMarketoUpdate) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionMarketoUpdate) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionMarketoUpdate) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionMarketoUpdate) GetQuotaLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.QuotaLimit
}

func (o *ConnectionMarketoUpdate) GetRestClientID() *string {
	if o == nil {
		return nil
	}
	return o.RestClientID
}

func (o *ConnectionMarketoUpdate) GetSoapEndpoint() *string {
	if o == nil {
		return nil
	}
	return o.SoapEndpoint
}

func (o *ConnectionMarketoUpdate) GetRestEndpoint() *string {
	if o == nil {
		return nil
	}
	return o.RestEndpoint
}

func (o *ConnectionMarketoUpdate) GetRestClientSecret() *string {
	if o == nil {
		return nil
	}
	return o.RestClientSecret
}

func (o *ConnectionMarketoUpdate) GetSoapUserID() *string {
	if o == nil {
		return nil
	}
	return o.SoapUserID
}

func (o *ConnectionMarketoUpdate) GetSoapEncryptionKey() *string {
	if o == nil {
		return nil
	}
	return o.SoapEncryptionKey
}
