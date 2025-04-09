// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetSHOPIFYConnectionRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetSHOPIFYConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetSHOPIFYConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionShopify *shared.ConnectionShopify
	// Forbidden. You don't have access to view this connection.
	Errors *shared.Errors
}

func (o *GetSHOPIFYConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSHOPIFYConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSHOPIFYConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSHOPIFYConnectionResponse) GetConnectionShopify() *shared.ConnectionShopify {
	if o == nil {
		return nil
	}
	return o.ConnectionShopify
}

func (o *GetSHOPIFYConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
