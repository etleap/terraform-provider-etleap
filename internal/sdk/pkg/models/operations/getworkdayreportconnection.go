// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetWORKDAYREPORTConnectionRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetWORKDAYREPORTConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetWORKDAYREPORTConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionWorkdayReport *shared.ConnectionWorkdayReport
	// Forbidden. You don't have access to view this connection.
	Errors *shared.Errors
}

func (o *GetWORKDAYREPORTConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetWORKDAYREPORTConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetWORKDAYREPORTConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetWORKDAYREPORTConnectionResponse) GetConnectionWorkdayReport() *shared.ConnectionWorkdayReport {
	if o == nil {
		return nil
	}
	return o.ConnectionWorkdayReport
}

func (o *GetWORKDAYREPORTConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
