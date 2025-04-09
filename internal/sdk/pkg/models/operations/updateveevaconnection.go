// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateVEEVAConnectionRequest struct {
	ID                    string                        `pathParam:"style=simple,explode=false,name=id"`
	ConnectionVeevaUpdate *shared.ConnectionVeevaUpdate `request:"mediaType=application/json"`
}

func (o *UpdateVEEVAConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateVEEVAConnectionRequest) GetConnectionVeevaUpdate() *shared.ConnectionVeevaUpdate {
	if o == nil {
		return nil
	}
	return o.ConnectionVeevaUpdate
}

type UpdateVEEVAConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionVeeva *shared.ConnectionVeeva
	// Bad Request
	Errors *shared.Errors
}

func (o *UpdateVEEVAConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateVEEVAConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateVEEVAConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateVEEVAConnectionResponse) GetConnectionVeeva() *shared.ConnectionVeeva {
	if o == nil {
		return nil
	}
	return o.ConnectionVeeva
}

func (o *UpdateVEEVAConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
