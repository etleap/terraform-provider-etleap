// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetRECURLYConnectionRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetRECURLYConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetRECURLYConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionRecurly *shared.ConnectionRecurlyOutput
	// Not Found.
	Errors *shared.Errors
}

func (o *GetRECURLYConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetRECURLYConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetRECURLYConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetRECURLYConnectionResponse) GetConnectionRecurly() *shared.ConnectionRecurlyOutput {
	if o == nil {
		return nil
	}
	return o.ConnectionRecurly
}

func (o *GetRECURLYConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
