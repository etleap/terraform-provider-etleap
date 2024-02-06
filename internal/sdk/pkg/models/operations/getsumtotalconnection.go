// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetSUMTOTALConnectionRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetSUMTOTALConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetSUMTOTALConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionSumTotal *shared.ConnectionSumTotal
	// Forbidden. You don't have access to view this connection.
	Errors *shared.Errors
}

func (o *GetSUMTOTALConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSUMTOTALConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSUMTOTALConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSUMTOTALConnectionResponse) GetConnectionSumTotal() *shared.ConnectionSumTotal {
	if o == nil {
		return nil
	}
	return o.ConnectionSumTotal
}

func (o *GetSUMTOTALConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
