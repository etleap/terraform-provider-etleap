// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateWORKDAYREPORTConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionWorkdayReport *shared.ConnectionWorkdayReport
	// Bad Request
	Errors *shared.Errors
}

func (o *CreateWORKDAYREPORTConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateWORKDAYREPORTConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateWORKDAYREPORTConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateWORKDAYREPORTConnectionResponse) GetConnectionWorkdayReport() *shared.ConnectionWorkdayReport {
	if o == nil {
		return nil
	}
	return o.ConnectionWorkdayReport
}

func (o *CreateWORKDAYREPORTConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
