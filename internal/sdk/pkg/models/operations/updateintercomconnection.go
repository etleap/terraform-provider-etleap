// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateINTERCOMConnectionRequest struct {
	ID                       string                           `pathParam:"style=simple,explode=false,name=id"`
	ConnectionIntercomUpdate *shared.ConnectionIntercomUpdate `request:"mediaType=application/json"`
}

func (o *UpdateINTERCOMConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateINTERCOMConnectionRequest) GetConnectionIntercomUpdate() *shared.ConnectionIntercomUpdate {
	if o == nil {
		return nil
	}
	return o.ConnectionIntercomUpdate
}

type UpdateINTERCOMConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionIntercom *shared.ConnectionIntercom
	// Connection for this id was not found.
	Errors *shared.Errors
}

func (o *UpdateINTERCOMConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateINTERCOMConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateINTERCOMConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateINTERCOMConnectionResponse) GetConnectionIntercom() *shared.ConnectionIntercom {
	if o == nil {
		return nil
	}
	return o.ConnectionIntercom
}

func (o *UpdateINTERCOMConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
