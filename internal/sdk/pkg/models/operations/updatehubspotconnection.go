// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateHUBSPOTConnectionRequest struct {
	ID                      string                          `pathParam:"style=simple,explode=false,name=id"`
	ConnectionHubspotUpdate *shared.ConnectionHubspotUpdate `request:"mediaType=application/json"`
}

func (o *UpdateHUBSPOTConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateHUBSPOTConnectionRequest) GetConnectionHubspotUpdate() *shared.ConnectionHubspotUpdate {
	if o == nil {
		return nil
	}
	return o.ConnectionHubspotUpdate
}

type UpdateHUBSPOTConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionHubspot *shared.ConnectionHubspotOutput
	// Connection for this id was not found.
	Errors *shared.Errors
}

func (o *UpdateHUBSPOTConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateHUBSPOTConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateHUBSPOTConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateHUBSPOTConnectionResponse) GetConnectionHubspot() *shared.ConnectionHubspotOutput {
	if o == nil {
		return nil
	}
	return o.ConnectionHubspot
}

func (o *UpdateHUBSPOTConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
