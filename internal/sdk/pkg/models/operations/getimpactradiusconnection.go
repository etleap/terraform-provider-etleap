// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetIMPACTRADIUSConnectionRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetIMPACTRADIUSConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetIMPACTRADIUSConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionImpactRadius *shared.ConnectionImpactRadius
	// Forbidden. You don't have access to view this connection.
	Errors *shared.Errors
}

func (o *GetIMPACTRADIUSConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetIMPACTRADIUSConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetIMPACTRADIUSConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetIMPACTRADIUSConnectionResponse) GetConnectionImpactRadius() *shared.ConnectionImpactRadius {
	if o == nil {
		return nil
	}
	return o.ConnectionImpactRadius
}

func (o *GetIMPACTRADIUSConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}