// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
	"time"
)

type ConnectionMysqlType string

const (
	ConnectionMysqlTypeMysql ConnectionMysqlType = "MYSQL"
)

func (e ConnectionMysqlType) ToPointer() *ConnectionMysqlType {
	return &e
}

func (e *ConnectionMysqlType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MYSQL":
		*e = ConnectionMysqlType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionMysqlType: %v", v)
	}
}

// ConnectionMysqlStatus - The current status of the connection.
type ConnectionMysqlStatus string

const (
	ConnectionMysqlStatusUnknown     ConnectionMysqlStatus = "UNKNOWN"
	ConnectionMysqlStatusUp          ConnectionMysqlStatus = "UP"
	ConnectionMysqlStatusDown        ConnectionMysqlStatus = "DOWN"
	ConnectionMysqlStatusResize      ConnectionMysqlStatus = "RESIZE"
	ConnectionMysqlStatusMaintenance ConnectionMysqlStatus = "MAINTENANCE"
	ConnectionMysqlStatusQuota       ConnectionMysqlStatus = "QUOTA"
	ConnectionMysqlStatusCreating    ConnectionMysqlStatus = "CREATING"
)

func (e ConnectionMysqlStatus) ToPointer() *ConnectionMysqlStatus {
	return &e
}

func (e *ConnectionMysqlStatus) UnmarshalJSON(data []byte) error {
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
		*e = ConnectionMysqlStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionMysqlStatus: %v", v)
	}
}

// ConnectionMysql - Specifies the location of a database.
type ConnectionMysql struct {
	// The unique identifier of the connection.
	ID string `json:"id"`
	// The unique name of this connection.
	Name string              `json:"name"`
	Type ConnectionMysqlType `json:"type"`
	// Whether this connection has been marked as active.
	Active bool `json:"active"`
	// The current status of the connection.
	Status ConnectionMysqlStatus `json:"status"`
	// The date and time when then the connection was created.
	CreateDate      time.Time `json:"createDate"`
	ValidateSslCert *bool     `default:"false" json:"validateSslCert"`
	// Should Etleap use MySQL binlogs to capture changes from this database? This setting cannot be changed later.
	CdcEnabled *bool `default:"false" json:"cdcEnabled"`
	// If you want Etleap to create pipelines for each source table automatically, specify the id of an Etleap destination connection here. If you want to create pipelines manually, omit this property. Note that only the connection owner can change this setting.
	AutoReplicate *string `json:"autoReplicate,omitempty"`
	// Should Etleap interpret columns with type Tinyint(1) as Boolean (i.e. true/false)?
	TinyInt1IsBoolean *bool      `default:"false" json:"tinyInt1IsBoolean"`
	Address           string     `json:"address"`
	Port              int64      `json:"port"`
	Database          string     `json:"database"`
	Username          string     `json:"username"`
	SSHConfig         *SSHConfig `json:"sshConfig,omitempty"`
}

func (c ConnectionMysql) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionMysql) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionMysql) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ConnectionMysql) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionMysql) GetType() ConnectionMysqlType {
	if o == nil {
		return ConnectionMysqlType("")
	}
	return o.Type
}

func (o *ConnectionMysql) GetActive() bool {
	if o == nil {
		return false
	}
	return o.Active
}

func (o *ConnectionMysql) GetStatus() ConnectionMysqlStatus {
	if o == nil {
		return ConnectionMysqlStatus("")
	}
	return o.Status
}

func (o *ConnectionMysql) GetCreateDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreateDate
}

func (o *ConnectionMysql) GetValidateSslCert() *bool {
	if o == nil {
		return nil
	}
	return o.ValidateSslCert
}

func (o *ConnectionMysql) GetCdcEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.CdcEnabled
}

func (o *ConnectionMysql) GetAutoReplicate() *string {
	if o == nil {
		return nil
	}
	return o.AutoReplicate
}

func (o *ConnectionMysql) GetTinyInt1IsBoolean() *bool {
	if o == nil {
		return nil
	}
	return o.TinyInt1IsBoolean
}

func (o *ConnectionMysql) GetAddress() string {
	if o == nil {
		return ""
	}
	return o.Address
}

func (o *ConnectionMysql) GetPort() int64 {
	if o == nil {
		return 0
	}
	return o.Port
}

func (o *ConnectionMysql) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *ConnectionMysql) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *ConnectionMysql) GetSSHConfig() *SSHConfig {
	if o == nil {
		return nil
	}
	return o.SSHConfig
}

// ConnectionMysqlInput - Specifies the location of a database.
type ConnectionMysqlInput struct {
	// The unique name of this connection.
	Name            string              `json:"name"`
	Type            ConnectionMysqlType `json:"type"`
	ValidateSslCert *bool               `default:"false" json:"validateSslCert"`
	// Should Etleap use MySQL binlogs to capture changes from this database? This setting cannot be changed later.
	CdcEnabled *bool `default:"false" json:"cdcEnabled"`
	// If you want Etleap to create pipelines for each source table automatically, specify the id of an Etleap destination connection here. If you want to create pipelines manually, omit this property. Note that only the connection owner can change this setting.
	AutoReplicate *string `json:"autoReplicate,omitempty"`
	// Should Etleap interpret columns with type Tinyint(1) as Boolean (i.e. true/false)?
	TinyInt1IsBoolean *bool      `default:"false" json:"tinyInt1IsBoolean"`
	Address           string     `json:"address"`
	Port              int64      `json:"port"`
	Database          string     `json:"database"`
	Username          string     `json:"username"`
	Password          string     `json:"password"`
	SSHConfig         *SSHConfig `json:"sshConfig,omitempty"`
}

func (c ConnectionMysqlInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionMysqlInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionMysqlInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConnectionMysqlInput) GetType() ConnectionMysqlType {
	if o == nil {
		return ConnectionMysqlType("")
	}
	return o.Type
}

func (o *ConnectionMysqlInput) GetValidateSslCert() *bool {
	if o == nil {
		return nil
	}
	return o.ValidateSslCert
}

func (o *ConnectionMysqlInput) GetCdcEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.CdcEnabled
}

func (o *ConnectionMysqlInput) GetAutoReplicate() *string {
	if o == nil {
		return nil
	}
	return o.AutoReplicate
}

func (o *ConnectionMysqlInput) GetTinyInt1IsBoolean() *bool {
	if o == nil {
		return nil
	}
	return o.TinyInt1IsBoolean
}

func (o *ConnectionMysqlInput) GetAddress() string {
	if o == nil {
		return ""
	}
	return o.Address
}

func (o *ConnectionMysqlInput) GetPort() int64 {
	if o == nil {
		return 0
	}
	return o.Port
}

func (o *ConnectionMysqlInput) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *ConnectionMysqlInput) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *ConnectionMysqlInput) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *ConnectionMysqlInput) GetSSHConfig() *SSHConfig {
	if o == nil {
		return nil
	}
	return o.SSHConfig
}
