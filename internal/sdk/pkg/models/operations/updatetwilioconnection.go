// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateTWILIOConnectionRequest struct {
	ID                     string                         `pathParam:"style=simple,explode=false,name=id"`
	ConnectionTwilioUpdate *shared.ConnectionTwilioUpdate `request:"mediaType=application/json"`
}

func (o *UpdateTWILIOConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateTWILIOConnectionRequest) GetConnectionTwilioUpdate() *shared.ConnectionTwilioUpdate {
	if o == nil {
		return nil
	}
	return o.ConnectionTwilioUpdate
}

type UpdateTWILIOConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionTwilio *shared.ConnectionTwilio
	// Bad Request
	Errors *shared.Errors
}

func (o *UpdateTWILIOConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateTWILIOConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateTWILIOConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateTWILIOConnectionResponse) GetConnectionTwilio() *shared.ConnectionTwilio {
	if o == nil {
		return nil
	}
	return o.ConnectionTwilio
}

func (o *UpdateTWILIOConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}