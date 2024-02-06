// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetELOQUAConnectionRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetELOQUAConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetELOQUAConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionEloqua *shared.ConnectionEloqua
	// Forbidden. You don't have access to view this connection.
	Errors *shared.Errors
}

func (o *GetELOQUAConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetELOQUAConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetELOQUAConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetELOQUAConnectionResponse) GetConnectionEloqua() *shared.ConnectionEloqua {
	if o == nil {
		return nil
	}
	return o.ConnectionEloqua
}

func (o *GetELOQUAConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
