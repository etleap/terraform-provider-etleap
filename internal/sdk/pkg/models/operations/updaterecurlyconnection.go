// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateRECURLYConnectionRequest struct {
	ID                      string                          `pathParam:"style=simple,explode=false,name=id"`
	ConnectionRecurlyUpdate *shared.ConnectionRecurlyUpdate `request:"mediaType=application/json"`
}

func (o *UpdateRECURLYConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateRECURLYConnectionRequest) GetConnectionRecurlyUpdate() *shared.ConnectionRecurlyUpdate {
	if o == nil {
		return nil
	}
	return o.ConnectionRecurlyUpdate
}

type UpdateRECURLYConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionRecurly *shared.ConnectionRecurlyOutput
	// Connection for this id was not found.
	Errors *shared.Errors
}

func (o *UpdateRECURLYConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateRECURLYConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateRECURLYConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateRECURLYConnectionResponse) GetConnectionRecurly() *shared.ConnectionRecurlyOutput {
	if o == nil {
		return nil
	}
	return o.ConnectionRecurly
}

func (o *UpdateRECURLYConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
