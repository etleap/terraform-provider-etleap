// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetFRESHSALESConnectionRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetFRESHSALESConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetFRESHSALESConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionFreshsales *shared.ConnectionFreshsalesOutput
	// Not Found.
	Errors *shared.Errors
}

func (o *GetFRESHSALESConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetFRESHSALESConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetFRESHSALESConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetFRESHSALESConnectionResponse) GetConnectionFreshsales() *shared.ConnectionFreshsalesOutput {
	if o == nil {
		return nil
	}
	return o.ConnectionFreshsales
}

func (o *GetFRESHSALESConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
