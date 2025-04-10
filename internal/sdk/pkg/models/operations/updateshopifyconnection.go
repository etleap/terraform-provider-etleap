// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateSHOPIFYConnectionRequest struct {
	ID                      string                          `pathParam:"style=simple,explode=false,name=id"`
	ConnectionShopifyUpdate *shared.ConnectionShopifyUpdate `request:"mediaType=application/json"`
}

func (o *UpdateSHOPIFYConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateSHOPIFYConnectionRequest) GetConnectionShopifyUpdate() *shared.ConnectionShopifyUpdate {
	if o == nil {
		return nil
	}
	return o.ConnectionShopifyUpdate
}

type UpdateSHOPIFYConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionShopify *shared.ConnectionShopifyOutput
	// Connection for this id was not found.
	Errors *shared.Errors
}

func (o *UpdateSHOPIFYConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateSHOPIFYConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateSHOPIFYConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateSHOPIFYConnectionResponse) GetConnectionShopify() *shared.ConnectionShopifyOutput {
	if o == nil {
		return nil
	}
	return o.ConnectionShopify
}

func (o *UpdateSHOPIFYConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
