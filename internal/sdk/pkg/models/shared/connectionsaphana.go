// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
	"time"
)

type ConnectionSapHanaType string

const (
	ConnectionSapHanaTypeSapHana ConnectionSapHanaType = "SAP_HANA"
)

func (e ConnectionSapHanaType) ToPointer() *ConnectionSapHanaType {
	return &e
}

func (e *ConnectionSapHanaType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SAP_HANA":
		*e = ConnectionSapHanaType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionSapHanaType: %v", v)
	}
}

// ConnectionSapHanaStatus - The current status of the connection.
type ConnectionSapHanaStatus string

const (
	ConnectionSapHanaStatusUnknown     ConnectionSapHanaStatus = "UNKNOWN"
	ConnectionSapHanaStatusUp          ConnectionSapHanaStatus = "UP"
	ConnectionSapHanaStatusDown        ConnectionSapHanaStatus = "DOWN"
	ConnectionSapHanaStatusResize      ConnectionSapHanaStatus = "RESIZE"
	ConnectionSapHanaStatusMaintenance ConnectionSapHanaStatus = "MAINTENANCE"
	ConnectionSapHanaStatusQuota       ConnectionSapHanaStatus = "QUOTA"
	ConnectionSapHanaStatusCreating    ConnectionSapHanaStatus = "CREATING"
)

func (e ConnectionSapHanaStatus) ToPointer() *ConnectionSapHanaStatus {
	return &e
}

func (e *ConnectionSapHanaStatus) UnmarshalJSON(data []byte) error {
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
		*e = ConnectionSapHanaStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionSapHanaStatus: %v", v)
	}
}

// ConnectionSapHana - Specifies the location of a database.
type ConnectionSapHana struct {
	// The unique identifier of the connection.
	ID string `json:"id"`
	// The unique name of this connection.
	Name string                `json:"name"`
	Type ConnectionSapHanaType `json:"type"`
	// Whether this connection has been marked as active.
	Active bool `json:"active"`
	// The current status of the connection.
	Status ConnectionSapHanaStatus `json:"status"`
	// The date and time when then the connection was created.
	CreateDate time.Time `json:"createDate"`
	// If not specified, the default schema will be used.
	Schema *string `json:"schema,omitempty"`
	// <i>Please note, this feature is currently in alpha and may be unstable.</i>
	//
	// Should Etleap use a change-tracking table and triggers defined on the source tables to capture changes from this database?
	//
	// For this setting to be enabled, the `ETLEAP_CTT` schema must be present in the source database, with the following privileges granted to the authenticating user: `SELECT`, `TRIGGER`, `CREATE ANY`, and `EXECUTE`.
	//
	// When enabled, Etleap will create a table and procedure in the `ETLEAP_CTT` schema and define triggers on any source table that has an Etleap pipeline. This setting cannot be changed later.
	CdcEnabled *bool      `default:"false" json:"cdcEnabled"`
	Address    string     `json:"address"`
	Port       int64      `json:"port"`
	Database   string     `json:"database"`
	Username   string     `json:"username"`
	SSHConfig  *SSHConfig `json:"sshConfig,omitempty"`
}

func (c ConnectionSapHana) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionSapHana) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionSapHana) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ConnectionSapHana) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionSapHana) GetType() ConnectionSapHanaType {
	if o == nil {
		return ConnectionSapHanaType("")
	}
	return o.Type
}

func (o *ConnectionSapHana) GetActive() bool {
	if o == nil {
		return false
	}
	return o.Active
}

func (o *ConnectionSapHana) GetStatus() ConnectionSapHanaStatus {
	if o == nil {
		return ConnectionSapHanaStatus("")
	}
	return o.Status
}

func (o *ConnectionSapHana) GetCreateDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreateDate
}

func (o *ConnectionSapHana) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *ConnectionSapHana) GetCdcEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.CdcEnabled
}

func (o *ConnectionSapHana) GetAddress() string {
	if o == nil {
		return ""
	}
	return o.Address
}

func (o *ConnectionSapHana) GetPort() int64 {
	if o == nil {
		return 0
	}
	return o.Port
}

func (o *ConnectionSapHana) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *ConnectionSapHana) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *ConnectionSapHana) GetSSHConfig() *SSHConfig {
	if o == nil {
		return nil
	}
	return o.SSHConfig
}

// ConnectionSapHanaInput - Specifies the location of a database.
type ConnectionSapHanaInput struct {
	// The unique name of this connection.
	Name string                `json:"name"`
	Type ConnectionSapHanaType `json:"type"`
	// If not specified, the default schema will be used.
	Schema *string `json:"schema,omitempty"`
	// <i>Please note, this feature is currently in alpha and may be unstable.</i>
	//
	// Should Etleap use a change-tracking table and triggers defined on the source tables to capture changes from this database?
	//
	// For this setting to be enabled, the `ETLEAP_CTT` schema must be present in the source database, with the following privileges granted to the authenticating user: `SELECT`, `TRIGGER`, `CREATE ANY`, and `EXECUTE`.
	//
	// When enabled, Etleap will create a table and procedure in the `ETLEAP_CTT` schema and define triggers on any source table that has an Etleap pipeline. This setting cannot be changed later.
	CdcEnabled *bool      `default:"false" json:"cdcEnabled"`
	Address    string     `json:"address"`
	Port       int64      `json:"port"`
	Database   string     `json:"database"`
	Username   string     `json:"username"`
	Password   string     `json:"password"`
	SSHConfig  *SSHConfig `json:"sshConfig,omitempty"`
}

func (c ConnectionSapHanaInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionSapHanaInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionSapHanaInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionSapHanaInput) GetType() ConnectionSapHanaType {
	if o == nil {
		return ConnectionSapHanaType("")
	}
	return o.Type
}

func (o *ConnectionSapHanaInput) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *ConnectionSapHanaInput) GetCdcEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.CdcEnabled
}

func (o *ConnectionSapHanaInput) GetAddress() string {
	if o == nil {
		return ""
	}
	return o.Address
}

func (o *ConnectionSapHanaInput) GetPort() int64 {
	if o == nil {
		return 0
	}
	return o.Port
}

func (o *ConnectionSapHanaInput) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *ConnectionSapHanaInput) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *ConnectionSapHanaInput) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *ConnectionSapHanaInput) GetSSHConfig() *SSHConfig {
	if o == nil {
		return nil
	}
	return o.SSHConfig
}
