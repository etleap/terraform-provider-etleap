// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
	"time"
)

type TeamOutput struct {
	// The unique identifier of the team.
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Members     []User `json:"members"`
	// The date and time when then the team was created.
	CreateDate time.Time `json:"createDate"`
}

func (t TeamOutput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TeamOutput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TeamOutput) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *TeamOutput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *TeamOutput) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *TeamOutput) GetMembers() []User {
	if o == nil {
		return []User{}
	}
	return o.Members
}

func (o *TeamOutput) GetCreateDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreateDate
}
