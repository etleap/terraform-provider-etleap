// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateConnectionGrantPrivilegeRequest struct {
	ID             string                `pathParam:"style=simple,explode=false,name=id"`
	GrantID        string                `pathParam:"style=simple,explode=false,name=grantId"`
	GrantPrivilege shared.GrantPrivilege `request:"mediaType=application/json"`
}

func (o *UpdateConnectionGrantPrivilegeRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateConnectionGrantPrivilegeRequest) GetGrantID() string {
	if o == nil {
		return ""
	}
	return o.GrantID
}

func (o *UpdateConnectionGrantPrivilegeRequest) GetGrantPrivilege() shared.GrantPrivilege {
	if o == nil {
		return shared.GrantPrivilege{}
	}
	return o.GrantPrivilege
}

type UpdateConnectionGrantPrivilegeResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	GrantWithPrivilege *shared.GrantWithPrivilege
	// A connection or grant with the given id was not found.
	Errors *shared.Errors
}

func (o *UpdateConnectionGrantPrivilegeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateConnectionGrantPrivilegeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateConnectionGrantPrivilegeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateConnectionGrantPrivilegeResponse) GetGrantWithPrivilege() *shared.GrantWithPrivilege {
	if o == nil {
		return nil
	}
	return o.GrantWithPrivilege
}

func (o *UpdateConnectionGrantPrivilegeResponse) GetGrantWithPrivilegeUser() *shared.GrantWithPrivilegeUser {
	if v := o.GetGrantWithPrivilege(); v != nil {
		return v.GrantWithPrivilegeUser
	}
	return nil
}

func (o *UpdateConnectionGrantPrivilegeResponse) GetGrantWithPrivilegeTeam() *shared.GrantWithPrivilegeTeam {
	if v := o.GetGrantWithPrivilege(); v != nil {
		return v.GrantWithPrivilegeTeam
	}
	return nil
}

func (o *UpdateConnectionGrantPrivilegeResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
