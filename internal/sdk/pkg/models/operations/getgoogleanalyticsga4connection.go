// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetGOOGLEANALYTICSGa4ConnectionRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetGOOGLEANALYTICSGa4ConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetGOOGLEANALYTICSGa4ConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionGoogleAnalyticsGa4 *shared.ConnectionGoogleAnalyticsGa4
	// Forbidden. You don't have access to view this connection.
	Errors *shared.Errors
}

func (o *GetGOOGLEANALYTICSGa4ConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetGOOGLEANALYTICSGa4ConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetGOOGLEANALYTICSGa4ConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetGOOGLEANALYTICSGa4ConnectionResponse) GetConnectionGoogleAnalyticsGa4() *shared.ConnectionGoogleAnalyticsGa4 {
	if o == nil {
		return nil
	}
	return o.ConnectionGoogleAnalyticsGa4
}

func (o *GetGOOGLEANALYTICSGa4ConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
