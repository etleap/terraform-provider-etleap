// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetINTERCOMConnectionRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetINTERCOMConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetINTERCOMConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionIntercom *shared.ConnectionIntercom
	// Not Found.
	Errors *shared.Errors
}

func (o *GetINTERCOMConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetINTERCOMConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetINTERCOMConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetINTERCOMConnectionResponse) GetConnectionIntercom() *shared.ConnectionIntercom {
	if o == nil {
		return nil
	}
	return o.ConnectionIntercom
}

func (o *GetINTERCOMConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
