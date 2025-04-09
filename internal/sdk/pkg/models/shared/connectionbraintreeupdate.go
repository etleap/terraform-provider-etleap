// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionBraintreeUpdateType string

const (
	ConnectionBraintreeUpdateTypeBraintree ConnectionBraintreeUpdateType = "BRAINTREE"
)

func (e ConnectionBraintreeUpdateType) ToPointer() *ConnectionBraintreeUpdateType {
	return &e
}

func (e *ConnectionBraintreeUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "BRAINTREE":
		*e = ConnectionBraintreeUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionBraintreeUpdateType: %v", v)
	}
}

type ConnectionBraintreeUpdate struct {
	// Whether this connection should be marked as active.
	Active *bool                          `json:"active,omitempty"`
	Type   *ConnectionBraintreeUpdateType `json:"type"`
	// The unique name of this connection.
	Name *string `json:"name,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	// The private API key.
	PrivateKey *string `json:"privateKey,omitempty"`
	// The public API key.
	PublicKey *string `json:"publicKey,omitempty"`
	// Check this box if this is a sandbox account.
	Sandbox *bool `json:"sandbox,omitempty"`
	// Your Braintree tenant Merchant ID.
	MerchantID *string `json:"merchantId,omitempty"`
}

func (o *ConnectionBraintreeUpdate) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ConnectionBraintreeUpdate) GetType() *ConnectionBraintreeUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ConnectionBraintreeUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionBraintreeUpdate) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionBraintreeUpdate) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionBraintreeUpdate) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionBraintreeUpdate) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionBraintreeUpdate) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionBraintreeUpdate) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionBraintreeUpdate) GetPrivateKey() *string {
	if o == nil {
		return nil
	}
	return o.PrivateKey
}

func (o *ConnectionBraintreeUpdate) GetPublicKey() *string {
	if o == nil {
		return nil
	}
	return o.PublicKey
}

func (o *ConnectionBraintreeUpdate) GetSandbox() *bool {
	if o == nil {
		return nil
	}
	return o.Sandbox
}

func (o *ConnectionBraintreeUpdate) GetMerchantID() *string {
	if o == nil {
		return nil
	}
	return o.MerchantID
}
