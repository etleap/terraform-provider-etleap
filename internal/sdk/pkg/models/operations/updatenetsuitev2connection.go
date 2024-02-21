// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateNETSUITEV2ConnectionRequest struct {
	ID                         string                             `pathParam:"style=simple,explode=false,name=id"`
	ConnectionNetsuiteV2Update *shared.ConnectionNetsuiteV2Update `request:"mediaType=application/json"`
}

func (o *UpdateNETSUITEV2ConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateNETSUITEV2ConnectionRequest) GetConnectionNetsuiteV2Update() *shared.ConnectionNetsuiteV2Update {
	if o == nil {
		return nil
	}
	return o.ConnectionNetsuiteV2Update
}

type UpdateNETSUITEV2ConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionNetsuiteV2 *shared.ConnectionNetsuiteV2
	// Bad Request
	Errors *shared.Errors
}

func (o *UpdateNETSUITEV2ConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateNETSUITEV2ConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateNETSUITEV2ConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateNETSUITEV2ConnectionResponse) GetConnectionNetsuiteV2() *shared.ConnectionNetsuiteV2 {
	if o == nil {
		return nil
	}
	return o.ConnectionNetsuiteV2
}

func (o *UpdateNETSUITEV2ConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}