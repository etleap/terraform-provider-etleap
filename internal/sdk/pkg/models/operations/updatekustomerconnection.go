// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateKUSTOMERConnectionRequest struct {
	ID                       string                           `pathParam:"style=simple,explode=false,name=id"`
	ConnectionKustomerUpdate *shared.ConnectionKustomerUpdate `request:"mediaType=application/json"`
}

func (o *UpdateKUSTOMERConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateKUSTOMERConnectionRequest) GetConnectionKustomerUpdate() *shared.ConnectionKustomerUpdate {
	if o == nil {
		return nil
	}
	return o.ConnectionKustomerUpdate
}

type UpdateKUSTOMERConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionKustomer *shared.ConnectionKustomer
	// Connection for this id was not found.
	Errors *shared.Errors
}

func (o *UpdateKUSTOMERConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateKUSTOMERConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateKUSTOMERConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateKUSTOMERConnectionResponse) GetConnectionKustomer() *shared.ConnectionKustomer {
	if o == nil {
		return nil
	}
	return o.ConnectionKustomer
}

func (o *UpdateKUSTOMERConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
