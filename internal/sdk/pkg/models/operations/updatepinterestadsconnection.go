// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdatePINTERESTADSConnectionRequest struct {
	ID                           string                               `pathParam:"style=simple,explode=false,name=id"`
	ConnectionPinterestAdsUpdate *shared.ConnectionPinterestAdsUpdate `request:"mediaType=application/json"`
}

func (o *UpdatePINTERESTADSConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdatePINTERESTADSConnectionRequest) GetConnectionPinterestAdsUpdate() *shared.ConnectionPinterestAdsUpdate {
	if o == nil {
		return nil
	}
	return o.ConnectionPinterestAdsUpdate
}

type UpdatePINTERESTADSConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionPinterestAds *shared.ConnectionPinterestAds
	// Connection for this id was not found.
	Errors *shared.Errors
}

func (o *UpdatePINTERESTADSConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdatePINTERESTADSConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdatePINTERESTADSConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdatePINTERESTADSConnectionResponse) GetConnectionPinterestAds() *shared.ConnectionPinterestAds {
	if o == nil {
		return nil
	}
	return o.ConnectionPinterestAds
}

func (o *UpdatePINTERESTADSConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
