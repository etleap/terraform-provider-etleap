// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateLDAPConnectionRequest struct {
	ID                   string                       `pathParam:"style=simple,explode=false,name=id"`
	ConnectionLdapUpdate *shared.ConnectionLdapUpdate `request:"mediaType=application/json"`
}

func (o *UpdateLDAPConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateLDAPConnectionRequest) GetConnectionLdapUpdate() *shared.ConnectionLdapUpdate {
	if o == nil {
		return nil
	}
	return o.ConnectionLdapUpdate
}

type UpdateLDAPConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionLdap *shared.ConnectionLdap
	// Bad Request
	Errors *shared.Errors
}

func (o *UpdateLDAPConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateLDAPConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateLDAPConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateLDAPConnectionResponse) GetConnectionLdap() *shared.ConnectionLdap {
	if o == nil {
		return nil
	}
	return o.ConnectionLdap
}

func (o *UpdateLDAPConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}