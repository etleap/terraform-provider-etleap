// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetSKYWARDConnectionRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetSKYWARDConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetSKYWARDConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionSkyward *shared.ConnectionSkywardOutput
	// Not Found.
	Errors *shared.Errors
}

func (o *GetSKYWARDConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSKYWARDConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSKYWARDConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSKYWARDConnectionResponse) GetConnectionSkyward() *shared.ConnectionSkywardOutput {
	if o == nil {
		return nil
	}
	return o.ConnectionSkyward
}

func (o *GetSKYWARDConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
