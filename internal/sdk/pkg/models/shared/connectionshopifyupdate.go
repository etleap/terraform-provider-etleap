// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionShopifyUpdateType string

const (
	ConnectionShopifyUpdateTypeShopify ConnectionShopifyUpdateType = "SHOPIFY"
)

func (e ConnectionShopifyUpdateType) ToPointer() *ConnectionShopifyUpdateType {
	return &e
}

func (e *ConnectionShopifyUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SHOPIFY":
		*e = ConnectionShopifyUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionShopifyUpdateType: %v", v)
	}
}

type ConnectionShopifyUpdate struct {
	// Whether this connection should be marked as active.
	Active *bool                        `json:"active,omitempty"`
	Type   *ConnectionShopifyUpdateType `json:"type"`
	// The unique name of this connection.
	Name *string `json:"name,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// To find your API keys, or generate a new one, follow these instructions: <a target="_blank" href="https://shopify.dev/apps/auth/basic-http#2-generate-api-credentials">Generate Api Credentials</a>.
	Password *string `json:"password,omitempty"`
	// To find your API keys, or generate a new one, follow these instructions: <a target="_blank" href="https://shopify.dev/apps/auth/basic-http#2-generate-api-credentials">Generate Api Credentials</a>.
	APIKey *string `json:"apiKey,omitempty"`
	// The store name for your account is the name of your development store. You can find it in the url of your Shopify store: <b>storename</b>.myshopify.com/
	StoreName *string `json:"storeName,omitempty"`
}

func (o *ConnectionShopifyUpdate) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ConnectionShopifyUpdate) GetType() *ConnectionShopifyUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ConnectionShopifyUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionShopifyUpdate) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionShopifyUpdate) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionShopifyUpdate) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionShopifyUpdate) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionShopifyUpdate) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionShopifyUpdate) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionShopifyUpdate) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *ConnectionShopifyUpdate) GetAPIKey() *string {
	if o == nil {
		return nil
	}
	return o.APIKey
}

func (o *ConnectionShopifyUpdate) GetStoreName() *string {
	if o == nil {
		return nil
	}
	return o.StoreName
}
