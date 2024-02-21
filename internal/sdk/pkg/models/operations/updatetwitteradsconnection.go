// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateTWITTERADSConnectionRequest struct {
	ID                      string                          `pathParam:"style=simple,explode=false,name=id"`
	ConnectionTwitterUpdate *shared.ConnectionTwitterUpdate `request:"mediaType=application/json"`
}

func (o *UpdateTWITTERADSConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateTWITTERADSConnectionRequest) GetConnectionTwitterUpdate() *shared.ConnectionTwitterUpdate {
	if o == nil {
		return nil
	}
	return o.ConnectionTwitterUpdate
}

type UpdateTWITTERADSConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionTwitter *shared.ConnectionTwitter
	// Bad Request
	Errors *shared.Errors
}

func (o *UpdateTWITTERADSConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateTWITTERADSConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateTWITTERADSConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateTWITTERADSConnectionResponse) GetConnectionTwitter() *shared.ConnectionTwitter {
	if o == nil {
		return nil
	}
	return o.ConnectionTwitter
}

func (o *UpdateTWITTERADSConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}