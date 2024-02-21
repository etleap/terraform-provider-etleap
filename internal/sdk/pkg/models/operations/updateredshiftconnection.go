// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateREDSHIFTConnectionRequest struct {
	ID                       string                           `pathParam:"style=simple,explode=false,name=id"`
	ConnectionRedshiftUpdate *shared.ConnectionRedshiftUpdate `request:"mediaType=application/json"`
}

func (o *UpdateREDSHIFTConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateREDSHIFTConnectionRequest) GetConnectionRedshiftUpdate() *shared.ConnectionRedshiftUpdate {
	if o == nil {
		return nil
	}
	return o.ConnectionRedshiftUpdate
}

type UpdateREDSHIFTConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionRedshift *shared.ConnectionRedshift
	// Bad Request
	Errors *shared.Errors
}

func (o *UpdateREDSHIFTConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateREDSHIFTConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateREDSHIFTConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateREDSHIFTConnectionResponse) GetConnectionRedshift() *shared.ConnectionRedshift {
	if o == nil {
		return nil
	}
	return o.ConnectionRedshift
}

func (o *UpdateREDSHIFTConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}