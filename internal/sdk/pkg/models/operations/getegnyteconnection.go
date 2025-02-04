// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetEGNYTEConnectionRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetEGNYTEConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetEGNYTEConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionEgnyte *shared.ConnectionEgnyte
	// Forbidden. You don't have access to view this connection.
	Errors *shared.Errors
}

func (o *GetEGNYTEConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetEGNYTEConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetEGNYTEConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetEGNYTEConnectionResponse) GetConnectionEgnyte() *shared.ConnectionEgnyte {
	if o == nil {
		return nil
	}
	return o.ConnectionEgnyte
}

func (o *GetEGNYTEConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
