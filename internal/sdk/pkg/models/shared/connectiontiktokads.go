// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
	"time"
)

type ConnectionTikTokAdsType string

const (
	ConnectionTikTokAdsTypeTikTokAds ConnectionTikTokAdsType = "TIK_TOK_ADS"
)

func (e ConnectionTikTokAdsType) ToPointer() *ConnectionTikTokAdsType {
	return &e
}

func (e *ConnectionTikTokAdsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TIK_TOK_ADS":
		*e = ConnectionTikTokAdsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionTikTokAdsType: %v", v)
	}
}

// ConnectionTikTokAdsStatus - The current status of the connection.
type ConnectionTikTokAdsStatus string

const (
	ConnectionTikTokAdsStatusUnknown     ConnectionTikTokAdsStatus = "UNKNOWN"
	ConnectionTikTokAdsStatusUp          ConnectionTikTokAdsStatus = "UP"
	ConnectionTikTokAdsStatusDown        ConnectionTikTokAdsStatus = "DOWN"
	ConnectionTikTokAdsStatusResize      ConnectionTikTokAdsStatus = "RESIZE"
	ConnectionTikTokAdsStatusMaintenance ConnectionTikTokAdsStatus = "MAINTENANCE"
	ConnectionTikTokAdsStatusQuota       ConnectionTikTokAdsStatus = "QUOTA"
	ConnectionTikTokAdsStatusCreating    ConnectionTikTokAdsStatus = "CREATING"
)

func (e ConnectionTikTokAdsStatus) ToPointer() *ConnectionTikTokAdsStatus {
	return &e
}

func (e *ConnectionTikTokAdsStatus) UnmarshalJSON(data []byte) error {
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
		*e = ConnectionTikTokAdsStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionTikTokAdsStatus: %v", v)
	}
}

type ConnectionTikTokAds struct {
	// The unique identifier of the connection.
	ID string `json:"id"`
	// The unique name of this connection.
	Name string                  `json:"name"`
	Type ConnectionTikTokAdsType `json:"type"`
	// Whether this connection has been marked as active.
	Active bool `json:"active"`
	// The current status of the connection.
	Status ConnectionTikTokAdsStatus `json:"status"`
	// The date and time when then the connection was created.
	CreateDate time.Time `json:"createDate"`
}

func (c ConnectionTikTokAds) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionTikTokAds) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionTikTokAds) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ConnectionTikTokAds) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionTikTokAds) GetType() ConnectionTikTokAdsType {
	if o == nil {
		return ConnectionTikTokAdsType("")
	}
	return o.Type
}

func (o *ConnectionTikTokAds) GetActive() bool {
	if o == nil {
		return false
	}
	return o.Active
}

func (o *ConnectionTikTokAds) GetStatus() ConnectionTikTokAdsStatus {
	if o == nil {
		return ConnectionTikTokAdsStatus("")
	}
	return o.Status
}

func (o *ConnectionTikTokAds) GetCreateDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreateDate
}

type ConnectionTikTokAdsInput struct {
	// The unique name of this connection.
	Name string                  `json:"name"`
	Type ConnectionTikTokAdsType `json:"type"`
	// Code retrieved from `/connections/oauth2-initiation`. **Note:** it is short-lived, therefore the connection creation should be done as soon as code is returned.
	Code string `json:"code"`
}

func (o *ConnectionTikTokAdsInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionTikTokAdsInput) GetType() ConnectionTikTokAdsType {
	if o == nil {
		return ConnectionTikTokAdsType("")
	}
	return o.Type
}

func (o *ConnectionTikTokAdsInput) GetCode() string {
	if o == nil {
		return ""
	}
	return o.Code
}
