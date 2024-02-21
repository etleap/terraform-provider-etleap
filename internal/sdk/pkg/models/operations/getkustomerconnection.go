// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetKUSTOMERConnectionRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetKUSTOMERConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetKUSTOMERConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionKustomer *shared.ConnectionKustomer
	// Forbidden. You don't have access to view this connection.
	Errors *shared.Errors
}

func (o *GetKUSTOMERConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetKUSTOMERConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetKUSTOMERConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetKUSTOMERConnectionResponse) GetConnectionKustomer() *shared.ConnectionKustomer {
	if o == nil {
		return nil
	}
	return o.ConnectionKustomer
}

func (o *GetKUSTOMERConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}