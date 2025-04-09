// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateLINKEDINADSConnectionRequest struct {
	ID                          string                              `pathParam:"style=simple,explode=false,name=id"`
	ConnectionLinkedInAdsUpdate *shared.ConnectionLinkedInAdsUpdate `request:"mediaType=application/json"`
}

func (o *UpdateLINKEDINADSConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateLINKEDINADSConnectionRequest) GetConnectionLinkedInAdsUpdate() *shared.ConnectionLinkedInAdsUpdate {
	if o == nil {
		return nil
	}
	return o.ConnectionLinkedInAdsUpdate
}

type UpdateLINKEDINADSConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionLinkedInAds *shared.ConnectionLinkedInAds
	// Bad Request
	Errors *shared.Errors
}

func (o *UpdateLINKEDINADSConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateLINKEDINADSConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateLINKEDINADSConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateLINKEDINADSConnectionResponse) GetConnectionLinkedInAds() *shared.ConnectionLinkedInAds {
	if o == nil {
		return nil
	}
	return o.ConnectionLinkedInAds
}

func (o *UpdateLINKEDINADSConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
