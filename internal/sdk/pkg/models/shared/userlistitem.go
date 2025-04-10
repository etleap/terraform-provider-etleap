// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
	"time"
)

// UserListItem - This model is returned in a list by the /users endpoint
type UserListItem struct {
	ID           string `json:"id"`
	LastName     string `json:"lastName"`
	EmailAddress string `json:"emailAddress"`
	FirstName    string `json:"firstName"`
	// The date and time when then the user was created.
	CreateDate time.Time `json:"createDate"`
}

func (u UserListItem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UserListItem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UserListItem) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UserListItem) GetLastName() string {
	if o == nil {
		return ""
	}
	return o.LastName
}

func (o *UserListItem) GetEmailAddress() string {
	if o == nil {
		return ""
	}
	return o.EmailAddress
}

func (o *UserListItem) GetFirstName() string {
	if o == nil {
		return ""
	}
	return o.FirstName
}

func (o *UserListItem) GetCreateDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreateDate
}
