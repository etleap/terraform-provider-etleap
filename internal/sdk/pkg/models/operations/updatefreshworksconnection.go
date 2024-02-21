// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateFRESHWORKSConnectionRequest struct {
	ID                         string                             `pathParam:"style=simple,explode=false,name=id"`
	ConnectionFreshworksUpdate *shared.ConnectionFreshworksUpdate `request:"mediaType=application/json"`
}

func (o *UpdateFRESHWORKSConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateFRESHWORKSConnectionRequest) GetConnectionFreshworksUpdate() *shared.ConnectionFreshworksUpdate {
	if o == nil {
		return nil
	}
	return o.ConnectionFreshworksUpdate
}

type UpdateFRESHWORKSConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionFreshworks *shared.ConnectionFreshworks
	// Bad Request
	Errors *shared.Errors
}

func (o *UpdateFRESHWORKSConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateFRESHWORKSConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateFRESHWORKSConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateFRESHWORKSConnectionResponse) GetConnectionFreshworks() *shared.ConnectionFreshworks {
	if o == nil {
		return nil
	}
	return o.ConnectionFreshworks
}

func (o *UpdateFRESHWORKSConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}