// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
	"time"
)

type ConnectionStripeType string

const (
	ConnectionStripeTypeStripe ConnectionStripeType = "STRIPE"
)

func (e ConnectionStripeType) ToPointer() *ConnectionStripeType {
	return &e
}

func (e *ConnectionStripeType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "STRIPE":
		*e = ConnectionStripeType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionStripeType: %v", v)
	}
}

// ConnectionStripeStatus - The current status of the connection.
type ConnectionStripeStatus string

const (
	ConnectionStripeStatusUnknown     ConnectionStripeStatus = "UNKNOWN"
	ConnectionStripeStatusUp          ConnectionStripeStatus = "UP"
	ConnectionStripeStatusDown        ConnectionStripeStatus = "DOWN"
	ConnectionStripeStatusResize      ConnectionStripeStatus = "RESIZE"
	ConnectionStripeStatusMaintenance ConnectionStripeStatus = "MAINTENANCE"
	ConnectionStripeStatusQuota       ConnectionStripeStatus = "QUOTA"
	ConnectionStripeStatusCreating    ConnectionStripeStatus = "CREATING"
)

func (e ConnectionStripeStatus) ToPointer() *ConnectionStripeStatus {
	return &e
}

func (e *ConnectionStripeStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "UNKNOWN":
		fallthrough
	case "UP":
		fallthrough
	case "DOWN":
		fallthrough
	case "RESIZE":
		fallthrough
	case "MAINTENANCE":
		fallthrough
	case "QUOTA":
		fallthrough
	case "CREATING":
		*e = ConnectionStripeStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionStripeStatus: %v", v)
	}
}

type ConnectionStripe struct {
	// The unique identifier of the connection.
	ID string `json:"id"`
	// The unique name of this connection.
	Name string               `json:"name"`
	Type ConnectionStripeType `json:"type"`
	// Whether this connection has been marked as active.
	Active bool `json:"active"`
	// The current status of the connection.
	Status ConnectionStripeStatus `json:"status"`
	// The date and time when then the connection was created.
	CreateDate time.Time `json:"createDate"`
}

func (c ConnectionStripe) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionStripe) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionStripe) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ConnectionStripe) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionStripe) GetType() ConnectionStripeType {
	if o == nil {
		return ConnectionStripeType("")
	}
	return o.Type
}

func (o *ConnectionStripe) GetActive() bool {
	if o == nil {
		return false
	}
	return o.Active
}

func (o *ConnectionStripe) GetStatus() ConnectionStripeStatus {
	if o == nil {
		return ConnectionStripeStatus("")
	}
	return o.Status
}

func (o *ConnectionStripe) GetCreateDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreateDate
}

type ConnectionStripeInput struct {
	// The unique name of this connection.
	Name string               `json:"name"`
	Type ConnectionStripeType `json:"type"`
	// Your API secret key can be found in your dashboard under the "Developer" tab.
	APISecretKey string `json:"apiSecretKey"`
}

func (o *ConnectionStripeInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionStripeInput) GetType() ConnectionStripeType {
	if o == nil {
		return ConnectionStripeType("")
	}
	return o.Type
}

func (o *ConnectionStripeInput) GetAPISecretKey() string {
	if o == nil {
		return ""
	}
	return o.APISecretKey
}
