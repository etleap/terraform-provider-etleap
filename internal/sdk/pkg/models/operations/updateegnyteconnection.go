// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateEGNYTEConnectionRequest struct {
	ID                     string                         `pathParam:"style=simple,explode=false,name=id"`
	ConnectionEgnyteUpdate *shared.ConnectionEgnyteUpdate `request:"mediaType=application/json"`
}

func (o *UpdateEGNYTEConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateEGNYTEConnectionRequest) GetConnectionEgnyteUpdate() *shared.ConnectionEgnyteUpdate {
	if o == nil {
		return nil
	}
	return o.ConnectionEgnyteUpdate
}

type UpdateEGNYTEConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionEgnyte *shared.ConnectionEgnyte
	// Connection for this id was not found.
	Errors *shared.Errors
}

func (o *UpdateEGNYTEConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateEGNYTEConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateEGNYTEConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateEGNYTEConnectionResponse) GetConnectionEgnyte() *shared.ConnectionEgnyte {
	if o == nil {
		return nil
	}
	return o.ConnectionEgnyte
}

func (o *UpdateEGNYTEConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
