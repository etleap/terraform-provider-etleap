// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetZENDESKConnectionRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetZENDESKConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetZENDESKConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionZendesk *shared.ConnectionZendesk
	// Forbidden. You don't have access to view this connection.
	Errors *shared.Errors
}

func (o *GetZENDESKConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetZENDESKConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetZENDESKConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetZENDESKConnectionResponse) GetConnectionZendesk() *shared.ConnectionZendesk {
	if o == nil {
		return nil
	}
	return o.ConnectionZendesk
}

func (o *GetZENDESKConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
