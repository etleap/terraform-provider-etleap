// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateFACEBOOKADSConnectionRequest struct {
	ID                 string                     `pathParam:"style=simple,explode=false,name=id"`
	ConnectionFbUpdate *shared.ConnectionFbUpdate `request:"mediaType=application/json"`
}

func (o *UpdateFACEBOOKADSConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateFACEBOOKADSConnectionRequest) GetConnectionFbUpdate() *shared.ConnectionFbUpdate {
	if o == nil {
		return nil
	}
	return o.ConnectionFbUpdate
}

type UpdateFACEBOOKADSConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionFb *shared.ConnectionFb
	// Bad Request
	Errors *shared.Errors
}

func (o *UpdateFACEBOOKADSConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateFACEBOOKADSConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateFACEBOOKADSConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateFACEBOOKADSConnectionResponse) GetConnectionFb() *shared.ConnectionFb {
	if o == nil {
		return nil
	}
	return o.ConnectionFb
}

func (o *UpdateFACEBOOKADSConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}