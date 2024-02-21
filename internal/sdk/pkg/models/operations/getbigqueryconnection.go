// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetBIGQUERYConnectionRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetBIGQUERYConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetBIGQUERYConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionBigQuery *shared.ConnectionBigQuery
	// Forbidden. You don't have access to view this connection.
	Errors *shared.Errors
}

func (o *GetBIGQUERYConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetBIGQUERYConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetBIGQUERYConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetBIGQUERYConnectionResponse) GetConnectionBigQuery() *shared.ConnectionBigQuery {
	if o == nil {
		return nil
	}
	return o.ConnectionBigQuery
}

func (o *GetBIGQUERYConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}