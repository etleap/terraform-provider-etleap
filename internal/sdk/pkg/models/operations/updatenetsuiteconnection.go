// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateNETSUITEConnectionRequest struct {
	ID                       string                           `pathParam:"style=simple,explode=false,name=id"`
	ConnectionNetsuiteUpdate *shared.ConnectionNetsuiteUpdate `request:"mediaType=application/json"`
}

func (o *UpdateNETSUITEConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateNETSUITEConnectionRequest) GetConnectionNetsuiteUpdate() *shared.ConnectionNetsuiteUpdate {
	if o == nil {
		return nil
	}
	return o.ConnectionNetsuiteUpdate
}

type UpdateNETSUITEConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionNetsuite *shared.ConnectionNetsuite
	// Connection for this id was not found.
	Errors *shared.Errors
}

func (o *UpdateNETSUITEConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateNETSUITEConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateNETSUITEConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateNETSUITEConnectionResponse) GetConnectionNetsuite() *shared.ConnectionNetsuite {
	if o == nil {
		return nil
	}
	return o.ConnectionNetsuite
}

func (o *UpdateNETSUITEConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
