// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ConnectionVeevaUpdateType string

const (
	ConnectionVeevaUpdateTypeVeeva ConnectionVeevaUpdateType = "VEEVA"
)

func (e ConnectionVeevaUpdateType) ToPointer() *ConnectionVeevaUpdateType {
	return &e
}

func (e *ConnectionVeevaUpdateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "VEEVA":
		*e = ConnectionVeevaUpdateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionVeevaUpdateType: %v", v)
	}
}

// ConnectionVeevaUpdateVaultType - Your Veeva Vault type. Currently supported types are: QUALITY, CTMS, and RIM. If a value is not provided it will fallback to the default QUALITY
type ConnectionVeevaUpdateVaultType string

const (
	ConnectionVeevaUpdateVaultTypeQuality   ConnectionVeevaUpdateVaultType = "QUALITY"
	ConnectionVeevaUpdateVaultTypeCtms      ConnectionVeevaUpdateVaultType = "CTMS"
	ConnectionVeevaUpdateVaultTypeRim       ConnectionVeevaUpdateVaultType = "RIM"
	ConnectionVeevaUpdateVaultTypePromomats ConnectionVeevaUpdateVaultType = "PROMOMATS"
)

func (e ConnectionVeevaUpdateVaultType) ToPointer() *ConnectionVeevaUpdateVaultType {
	return &e
}

func (e *ConnectionVeevaUpdateVaultType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "QUALITY":
		fallthrough
	case "CTMS":
		fallthrough
	case "RIM":
		fallthrough
	case "PROMOMATS":
		*e = ConnectionVeevaUpdateVaultType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConnectionVeevaUpdateVaultType: %v", v)
	}
}

type ConnectionVeevaUpdate struct {
	// Whether this connection should be marked as active.
	Active *bool                      `json:"active,omitempty"`
	Type   *ConnectionVeevaUpdateType `json:"type"`
	// The unique name of this connection.
	Name *string `json:"name,omitempty"`
	// The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection.
	UpdateSchedule *UpdateScheduleTypes `json:"updateSchedule,omitempty"`
	Username       *string              `json:"username,omitempty"`
	Password       *string              `json:"password,omitempty"`
	// Your Veeva Vault type. Currently supported types are: QUALITY, CTMS, and RIM. If a value is not provided it will fallback to the default QUALITY
	VaultType *ConnectionVeevaUpdateVaultType `json:"vaultType,omitempty"`
	// The vault domain name is part of the URL that you use to access your Veeva Vault. You can follow the<a target="_blank" href="https://developer.veevavault.com/docs/#authentication">'Structuring the Endpoint'</a> instructions to find the URL.
	VaultDomainName *string `json:"vaultDomainName,omitempty"`
}

func (o *ConnectionVeevaUpdate) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *ConnectionVeevaUpdate) GetType() *ConnectionVeevaUpdateType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ConnectionVeevaUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionVeevaUpdate) GetUpdateSchedule() *UpdateScheduleTypes {
	if o == nil {
		return nil
	}
	return o.UpdateSchedule
}

func (o *ConnectionVeevaUpdate) GetUpdateScheduleMonthly() *UpdateScheduleModeMonthly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeMonthly
	}
	return nil
}

func (o *ConnectionVeevaUpdate) GetUpdateScheduleHourly() *UpdateScheduleModeHourly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeHourly
	}
	return nil
}

func (o *ConnectionVeevaUpdate) GetUpdateScheduleInterval() *UpdateScheduleModeInterval {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeInterval
	}
	return nil
}

func (o *ConnectionVeevaUpdate) GetUpdateScheduleDaily() *UpdateScheduleModeDaily {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeDaily
	}
	return nil
}

func (o *ConnectionVeevaUpdate) GetUpdateScheduleWeekly() *UpdateScheduleModeWeekly {
	if v := o.GetUpdateSchedule(); v != nil {
		return v.UpdateScheduleModeWeekly
	}
	return nil
}

func (o *ConnectionVeevaUpdate) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *ConnectionVeevaUpdate) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *ConnectionVeevaUpdate) GetVaultType() *ConnectionVeevaUpdateVaultType {
	if o == nil {
		return nil
	}
	return o.VaultType
}

func (o *ConnectionVeevaUpdate) GetVaultDomainName() *string {
	if o == nil {
		return nil
	}
	return o.VaultDomainName
}
