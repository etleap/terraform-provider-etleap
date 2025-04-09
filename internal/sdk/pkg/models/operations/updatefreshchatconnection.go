// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateFRESHCHATConnectionRequest struct {
	ID                        string                            `pathParam:"style=simple,explode=false,name=id"`
	ConnectionFreshchatUpdate *shared.ConnectionFreshchatUpdate `request:"mediaType=application/json"`
}

func (o *UpdateFRESHCHATConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateFRESHCHATConnectionRequest) GetConnectionFreshchatUpdate() *shared.ConnectionFreshchatUpdate {
	if o == nil {
		return nil
	}
	return o.ConnectionFreshchatUpdate
}

type UpdateFRESHCHATConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionFreshchat *shared.ConnectionFreshchat
	// Bad Request
	Errors *shared.Errors
}

func (o *UpdateFRESHCHATConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateFRESHCHATConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateFRESHCHATConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateFRESHCHATConnectionResponse) GetConnectionFreshchat() *shared.ConnectionFreshchat {
	if o == nil {
		return nil
	}
	return o.ConnectionFreshchat
}

func (o *UpdateFRESHCHATConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
