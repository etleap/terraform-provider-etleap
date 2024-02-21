// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateELOQUAConnectionRequest struct {
	ID                     string                         `pathParam:"style=simple,explode=false,name=id"`
	ConnectionEloquaUpdate *shared.ConnectionEloquaUpdate `request:"mediaType=application/json"`
}

func (o *UpdateELOQUAConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateELOQUAConnectionRequest) GetConnectionEloquaUpdate() *shared.ConnectionEloquaUpdate {
	if o == nil {
		return nil
	}
	return o.ConnectionEloquaUpdate
}

type UpdateELOQUAConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionEloqua *shared.ConnectionEloqua
	// Bad Request
	Errors *shared.Errors
}

func (o *UpdateELOQUAConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateELOQUAConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateELOQUAConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateELOQUAConnectionResponse) GetConnectionEloqua() *shared.ConnectionEloqua {
	if o == nil {
		return nil
	}
	return o.ConnectionEloqua
}

func (o *UpdateELOQUAConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}