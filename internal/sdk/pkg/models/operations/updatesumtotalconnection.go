// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateSUMTOTALConnectionRequest struct {
	ID                       string                           `pathParam:"style=simple,explode=false,name=id"`
	ConnectionSumTotalUpdate *shared.ConnectionSumTotalUpdate `request:"mediaType=application/json"`
}

func (o *UpdateSUMTOTALConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateSUMTOTALConnectionRequest) GetConnectionSumTotalUpdate() *shared.ConnectionSumTotalUpdate {
	if o == nil {
		return nil
	}
	return o.ConnectionSumTotalUpdate
}

type UpdateSUMTOTALConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionSumTotal *shared.ConnectionSumTotal
	// Bad Request
	Errors *shared.Errors
}

func (o *UpdateSUMTOTALConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateSUMTOTALConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateSUMTOTALConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateSUMTOTALConnectionResponse) GetConnectionSumTotal() *shared.ConnectionSumTotal {
	if o == nil {
		return nil
	}
	return o.ConnectionSumTotal
}

func (o *UpdateSUMTOTALConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}