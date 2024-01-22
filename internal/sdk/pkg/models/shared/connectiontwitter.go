// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
	"time"
)

type ConnectionTwitterType string

const (
	ConnectionTwitterTypeTwitterAds ConnectionTwitterType = "TWITTER_ADS"
)

func (e ConnectionTwitterType) ToPointer() *ConnectionTwitterType {
	return &e
}

func (e *ConnectionTwitterType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TWITTER_ADS":
		*e = ConnectionTwitterType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionTwitterType: %v", v)
	}
}

// ConnectionTwitterStatus - The current status of the connection.
type ConnectionTwitterStatus string

const (
	ConnectionTwitterStatusUnknown     ConnectionTwitterStatus = "UNKNOWN"
	ConnectionTwitterStatusUp          ConnectionTwitterStatus = "UP"
	ConnectionTwitterStatusDown        ConnectionTwitterStatus = "DOWN"
	ConnectionTwitterStatusResize      ConnectionTwitterStatus = "RESIZE"
	ConnectionTwitterStatusMaintenance ConnectionTwitterStatus = "MAINTENANCE"
	ConnectionTwitterStatusQuota       ConnectionTwitterStatus = "QUOTA"
	ConnectionTwitterStatusCreating    ConnectionTwitterStatus = "CREATING"
)

func (e ConnectionTwitterStatus) ToPointer() *ConnectionTwitterStatus {
	return &e
}

func (e *ConnectionTwitterStatus) UnmarshalJSON(data []byte) error {
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
		*e = ConnectionTwitterStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionTwitterStatus: %v", v)
	}
}

type ConnectionTwitter struct {
	// The unique identifier of the connection.
	ID string `json:"id"`
	// The unique name of this connection.
	Name string                `json:"name"`
	Type ConnectionTwitterType `json:"type"`
	// Whether this connection has been marked as active.
	Active bool `json:"active"`
	// The current status of the connection.
	Status ConnectionTwitterStatus `json:"status"`
	// The date and time when then the connection was created.
	CreateDate time.Time `json:"createDate"`
	// This represents your Twitter developer app when making API requests. Generated under 'Consumer API keys'.
	AppKey string `json:"appKey"`
	// If you want to create pipelines from entities that uses Twitter API V2 endpoints you need to specify which Twitter accounts you want to retrieve data from. The usernames must be separated by comma and without the @
	TwitterUsernames *string `json:"twitterUsernames,omitempty"`
}

func (c ConnectionTwitter) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionTwitter) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionTwitter) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ConnectionTwitter) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionTwitter) GetType() ConnectionTwitterType {
	if o == nil {
		return ConnectionTwitterType("")
	}
	return o.Type
}

func (o *ConnectionTwitter) GetActive() bool {
	if o == nil {
		return false
	}
	return o.Active
}

func (o *ConnectionTwitter) GetStatus() ConnectionTwitterStatus {
	if o == nil {
		return ConnectionTwitterStatus("")
	}
	return o.Status
}

func (o *ConnectionTwitter) GetCreateDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreateDate
}

func (o *ConnectionTwitter) GetAppKey() string {
	if o == nil {
		return ""
	}
	return o.AppKey
}

func (o *ConnectionTwitter) GetTwitterUsernames() *string {
	if o == nil {
		return nil
	}
	return o.TwitterUsernames
}

type ConnectionTwitterInput struct {
	// The unique name of this connection.
	Name string                `json:"name"`
	Type ConnectionTwitterType `json:"type"`
	// This represents your Twitter developer app when making API requests. Generated under 'Consumer API keys'.
	AppKey string `json:"appKey"`
	// This authenticates your Twitter developer app when making API requests. Generated under 'Consumer API keys'.
	AppSecretKey string `json:"appSecretKey"`
	// This identifies your Twitter account when generating Ad reports. Generated under 'Access token & access token secret'.
	AccessToken string `json:"accessToken"`
	// This authenticates your Twitter account when generating Ad reports. Generated under 'Access token & access token secret'.
	AccessTokenSecret string `json:"accessTokenSecret"`
	// If you want to create pipelines from entities that uses Twitter API V2 endpoints you need to specify which Twitter accounts you want to retrieve data from. The usernames must be separated by comma and without the @
	TwitterUsernames *string `json:"twitterUsernames,omitempty"`
}

func (o *ConnectionTwitterInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionTwitterInput) GetType() ConnectionTwitterType {
	if o == nil {
		return ConnectionTwitterType("")
	}
	return o.Type
}

func (o *ConnectionTwitterInput) GetAppKey() string {
	if o == nil {
		return ""
	}
	return o.AppKey
}

func (o *ConnectionTwitterInput) GetAppSecretKey() string {
	if o == nil {
		return ""
	}
	return o.AppSecretKey
}

func (o *ConnectionTwitterInput) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}

func (o *ConnectionTwitterInput) GetAccessTokenSecret() string {
	if o == nil {
		return ""
	}
	return o.AccessTokenSecret
}

func (o *ConnectionTwitterInput) GetTwitterUsernames() *string {
	if o == nil {
		return nil
	}
	return o.TwitterUsernames
}
