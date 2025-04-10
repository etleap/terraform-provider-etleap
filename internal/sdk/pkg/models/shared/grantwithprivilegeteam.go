// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
	"time"
)

type GrantWithPrivilegeTeamType string

const (
	GrantWithPrivilegeTeamTypeTeam GrantWithPrivilegeTeamType = "TEAM"
)

func (e GrantWithPrivilegeTeamType) ToPointer() *GrantWithPrivilegeTeamType {
	return &e
}

func (e *GrantWithPrivilegeTeamType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TEAM":
		*e = GrantWithPrivilegeTeamType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GrantWithPrivilegeTeamType: %v", v)
	}
}

type GrantWithPrivilegeTeamPrivilege string

const (
	GrantWithPrivilegeTeamPrivilegeView GrantWithPrivilegeTeamPrivilege = "VIEW"
	GrantWithPrivilegeTeamPrivilegeEdit GrantWithPrivilegeTeamPrivilege = "EDIT"
)

func (e GrantWithPrivilegeTeamPrivilege) ToPointer() *GrantWithPrivilegeTeamPrivilege {
	return &e
}

func (e *GrantWithPrivilegeTeamPrivilege) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "VIEW":
		fallthrough
	case "EDIT":
		*e = GrantWithPrivilegeTeamPrivilege(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GrantWithPrivilegeTeamPrivilege: %v", v)
	}
}

// GrantWithPrivilegeTeam - A grant with privilege providing access to an object in Etleap for a team.
type GrantWithPrivilegeTeam struct {
	// The date and time when then the grant was created.
	CreateDate time.Time                       `json:"createDate"`
	Type       GrantWithPrivilegeTeamType      `json:"type"`
	ID         string                          `json:"id"`
	TeamID     string                          `json:"teamId"`
	Name       string                          `json:"name"`
	Privilege  GrantWithPrivilegeTeamPrivilege `json:"privilege"`
}

func (g GrantWithPrivilegeTeam) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GrantWithPrivilegeTeam) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *GrantWithPrivilegeTeam) GetCreateDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreateDate
}

func (o *GrantWithPrivilegeTeam) GetType() GrantWithPrivilegeTeamType {
	if o == nil {
		return GrantWithPrivilegeTeamType("")
	}
	return o.Type
}

func (o *GrantWithPrivilegeTeam) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GrantWithPrivilegeTeam) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *GrantWithPrivilegeTeam) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *GrantWithPrivilegeTeam) GetPrivilege() GrantWithPrivilegeTeamPrivilege {
	if o == nil {
		return GrantWithPrivilegeTeamPrivilege("")
	}
	return o.Privilege
}

// GrantWithPrivilegeTeamInput - A grant with privilege providing access to an object in Etleap for a team.
type GrantWithPrivilegeTeamInput struct {
	Type      GrantWithPrivilegeTeamType      `json:"type"`
	TeamID    string                          `json:"teamId"`
	Privilege GrantWithPrivilegeTeamPrivilege `json:"privilege"`
}

func (o *GrantWithPrivilegeTeamInput) GetType() GrantWithPrivilegeTeamType {
	if o == nil {
		return GrantWithPrivilegeTeamType("")
	}
	return o.Type
}

func (o *GrantWithPrivilegeTeamInput) GetTeamID() string {
	if o == nil {
		return ""
	}
	return o.TeamID
}

func (o *GrantWithPrivilegeTeamInput) GetPrivilege() GrantWithPrivilegeTeamPrivilege {
	if o == nil {
		return GrantWithPrivilegeTeamPrivilege("")
	}
	return o.Privilege
}
