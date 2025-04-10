// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetREDSHIFTConnectionRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetREDSHIFTConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetREDSHIFTConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionRedshift *shared.ConnectionRedshift
	// Not Found.
	Errors *shared.Errors
}

func (o *GetREDSHIFTConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetREDSHIFTConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetREDSHIFTConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetREDSHIFTConnectionResponse) GetConnectionRedshift() *shared.ConnectionRedshift {
	if o == nil {
		return nil
	}
	return o.ConnectionRedshift
}

func (o *GetREDSHIFTConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
