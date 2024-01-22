// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
	"time"
)

type ConnectionZuoraType string

const (
	ConnectionZuoraTypeZuora ConnectionZuoraType = "ZUORA"
)

func (e ConnectionZuoraType) ToPointer() *ConnectionZuoraType {
	return &e
}

func (e *ConnectionZuoraType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ZUORA":
		*e = ConnectionZuoraType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionZuoraType: %v", v)
	}
}

// ConnectionZuoraStatus - The current status of the connection.
type ConnectionZuoraStatus string

const (
	ConnectionZuoraStatusUnknown     ConnectionZuoraStatus = "UNKNOWN"
	ConnectionZuoraStatusUp          ConnectionZuoraStatus = "UP"
	ConnectionZuoraStatusDown        ConnectionZuoraStatus = "DOWN"
	ConnectionZuoraStatusResize      ConnectionZuoraStatus = "RESIZE"
	ConnectionZuoraStatusMaintenance ConnectionZuoraStatus = "MAINTENANCE"
	ConnectionZuoraStatusQuota       ConnectionZuoraStatus = "QUOTA"
	ConnectionZuoraStatusCreating    ConnectionZuoraStatus = "CREATING"
)

func (e ConnectionZuoraStatus) ToPointer() *ConnectionZuoraStatus {
	return &e
}

func (e *ConnectionZuoraStatus) UnmarshalJSON(data []byte) error {
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
		*e = ConnectionZuoraStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionZuoraStatus: %v", v)
	}
}

type ConnectionZuora struct {
	// The unique identifier of the connection.
	ID string `json:"id"`
	// The unique name of this connection.
	Name string              `json:"name"`
	Type ConnectionZuoraType `json:"type"`
	// Whether this connection has been marked as active.
	Active bool `json:"active"`
	// The current status of the connection.
	Status ConnectionZuoraStatus `json:"status"`
	// The date and time when then the connection was created.
	CreateDate time.Time `json:"createDate"`
	// The Client ID displayed when you created the OAuth client. If you need help to create an OAuth Client, <a target="_blank" href="https://knowledgecenter.zuora.com/CF_Users_and_Administrators/A_Administrator_Settings/Manage_Users#Create_an_OAuth_Client_for_a_User">follow this article</a>.
	ClientID string `json:"clientId"`
	// The endpoint region.
	Endpoint string `json:"endpoint"`
	// Whether this is a sandbox account.
	Sandbox bool `json:"sandbox"`
	// Leave blank unless you have been assigned a specific endpoint, e.g. services42.zuora.com. Setting this overrides the 'Endpoint' and 'Sandbox' settings above.
	EndpointHostname *string `json:"endpointHostname,omitempty"`
}

func (c ConnectionZuora) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionZuora) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionZuora) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ConnectionZuora) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionZuora) GetType() ConnectionZuoraType {
	if o == nil {
		return ConnectionZuoraType("")
	}
	return o.Type
}

func (o *ConnectionZuora) GetActive() bool {
	if o == nil {
		return false
	}
	return o.Active
}

func (o *ConnectionZuora) GetStatus() ConnectionZuoraStatus {
	if o == nil {
		return ConnectionZuoraStatus("")
	}
	return o.Status
}

func (o *ConnectionZuora) GetCreateDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreateDate
}

func (o *ConnectionZuora) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *ConnectionZuora) GetEndpoint() string {
	if o == nil {
		return ""
	}
	return o.Endpoint
}

func (o *ConnectionZuora) GetSandbox() bool {
	if o == nil {
		return false
	}
	return o.Sandbox
}

func (o *ConnectionZuora) GetEndpointHostname() *string {
	if o == nil {
		return nil
	}
	return o.EndpointHostname
}

type ConnectionZuoraInput struct {
	// The unique name of this connection.
	Name string              `json:"name"`
	Type ConnectionZuoraType `json:"type"`
	// The Client ID displayed when you created the OAuth client. If you need help to create an OAuth Client, <a target="_blank" href="https://knowledgecenter.zuora.com/CF_Users_and_Administrators/A_Administrator_Settings/Manage_Users#Create_an_OAuth_Client_for_a_User">follow this article</a>.
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	// The endpoint region.
	Endpoint string `json:"endpoint"`
	// Whether this is a sandbox account.
	Sandbox bool `json:"sandbox"`
	// Leave blank unless you have been assigned a specific endpoint, e.g. services42.zuora.com. Setting this overrides the 'Endpoint' and 'Sandbox' settings above.
	EndpointHostname *string `json:"endpointHostname,omitempty"`
}

func (o *ConnectionZuoraInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionZuoraInput) GetType() ConnectionZuoraType {
	if o == nil {
		return ConnectionZuoraType("")
	}
	return o.Type
}

func (o *ConnectionZuoraInput) GetClientID() string {
	if o == nil {
		return ""
	}
	return o.ClientID
}

func (o *ConnectionZuoraInput) GetClientSecret() string {
	if o == nil {
		return ""
	}
	return o.ClientSecret
}

func (o *ConnectionZuoraInput) GetEndpoint() string {
	if o == nil {
		return ""
	}
	return o.Endpoint
}

func (o *ConnectionZuoraInput) GetSandbox() bool {
	if o == nil {
		return false
	}
	return o.Sandbox
}

func (o *ConnectionZuoraInput) GetEndpointHostname() *string {
	if o == nil {
		return nil
	}
	return o.EndpointHostname
}
