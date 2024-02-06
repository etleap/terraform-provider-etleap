// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionSalesforceMarketingCloudUpdateType string

const (
	ConnectionSalesforceMarketingCloudUpdateTypeSalesforceMarketingCloud ConnectionSalesforceMarketingCloudUpdateType = "SALESFORCE_MARKETING_CLOUD"
)

func (e ConnectionSalesforceMarketingCloudUpdateType) ToPointer() *ConnectionSalesforceMarketingCloudUpdateType {
	return &e
}

func (e *ConnectionSalesforceMarketingCloudUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SALESFORCE_MARKETING_CLOUD":
		*e = ConnectionSalesforceMarketingCloudUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionSalesforceMarketingCloudUpdateType: %v", v)
	}
}

type ConnectionSalesforceMarketingCloudUpdate struct {
	// The unique name of this connection.
	Name *string                                       `json:"name,omitempty"`
	Type *ConnectionSalesforceMarketingCloudUpdateType `json:"type"`
	// Whether this connection should be marked as active.
	Active *bool `json:"active,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// You can find your Account Subdomain with the following <a target="_blank" href="https://developer.salesforce.com/docs/marketing/marketing-cloud/guide/your-subdomain-tenant-specific-endpoints.html#locate-your-subdomain-and-endpoints">guide</a>.
	Subdomain *string `json:"subdomain,omitempty"`
	// You can retrieve your credentials, by creating an <a target="_blank" href="https://developer.salesforce.com/docs/marketing/marketing-cloud/guide/create-integration-enhanced.html#related-items">integration</a>.
	ClientID *string `json:"clientId,omitempty"`
	// You can retrieve your credentials, by creating an <a target="_blank" href="https://developer.salesforce.com/docs/marketing/marketing-cloud/guide/create-integration-enhanced.html#related-items">integration</a>.
	ClientSecret *string `json:"clientSecret,omitempty"`
}

func (o *ConnectionSalesforceMarketingCloudUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionSalesforceMarketingCloudUpdate) GetType() *ConnectionSalesforceMarketingCloudUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ConnectionSalesforceMarketingCloudUpdate) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ConnectionSalesforceMarketingCloudUpdate) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionSalesforceMarketingCloudUpdate) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionSalesforceMarketingCloudUpdate) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionSalesforceMarketingCloudUpdate) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionSalesforceMarketingCloudUpdate) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionSalesforceMarketingCloudUpdate) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionSalesforceMarketingCloudUpdate) GetSubdomain() *string {
	if o == nil {
		return nil
	}
	return o.Subdomain
}

func (o *ConnectionSalesforceMarketingCloudUpdate) GetClientID() *string {
	if o == nil {
		return nil
	}
	return o.ClientID
}

func (o *ConnectionSalesforceMarketingCloudUpdate) GetClientSecret() *string {
	if o == nil {
		return nil
	}
	return o.ClientSecret
}
