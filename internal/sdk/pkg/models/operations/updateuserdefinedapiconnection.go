// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateUSERDEFINEDAPIConnectionRequest struct {
	ID                             string                                 `pathParam:"style=simple,explode=false,name=id"`
	ConnectionUserDefinedAPIUpdate *shared.ConnectionUserDefinedAPIUpdate `request:"mediaType=application/json"`
}

func (o *UpdateUSERDEFINEDAPIConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateUSERDEFINEDAPIConnectionRequest) GetConnectionUserDefinedAPIUpdate() *shared.ConnectionUserDefinedAPIUpdate {
	if o == nil {
		return nil
	}
	return o.ConnectionUserDefinedAPIUpdate
}

type UpdateUSERDEFINEDAPIConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionUserDefinedAPI *shared.ConnectionUserDefinedAPI
	// Connection for this id was not found.
	Errors *shared.Errors
}

func (o *UpdateUSERDEFINEDAPIConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateUSERDEFINEDAPIConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateUSERDEFINEDAPIConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateUSERDEFINEDAPIConnectionResponse) GetConnectionUserDefinedAPI() *shared.ConnectionUserDefinedAPI {
	if o == nil {
		return nil
	}
	return o.ConnectionUserDefinedAPI
}

func (o *UpdateUSERDEFINEDAPIConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
