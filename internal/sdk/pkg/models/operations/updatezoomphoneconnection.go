// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateZOOMPHONEConnectionRequest struct {
	ID                        string                            `pathParam:"style=simple,explode=false,name=id"`
	ConnectionZoomPhoneUpdate *shared.ConnectionZoomPhoneUpdate `request:"mediaType=application/json"`
}

func (o *UpdateZOOMPHONEConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateZOOMPHONEConnectionRequest) GetConnectionZoomPhoneUpdate() *shared.ConnectionZoomPhoneUpdate {
	if o == nil {
		return nil
	}
	return o.ConnectionZoomPhoneUpdate
}

type UpdateZOOMPHONEConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionZoomPhone *shared.ConnectionZoomPhone
	// Bad Request
	Errors *shared.Errors
}

func (o *UpdateZOOMPHONEConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateZOOMPHONEConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateZOOMPHONEConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateZOOMPHONEConnectionResponse) GetConnectionZoomPhone() *shared.ConnectionZoomPhone {
	if o == nil {
		return nil
	}
	return o.ConnectionZoomPhone
}

func (o *UpdateZOOMPHONEConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}