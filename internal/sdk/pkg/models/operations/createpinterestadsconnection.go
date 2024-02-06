// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreatePINTERESTADSConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionPinterestAds *shared.ConnectionPinterestAds
	// Bad Request
	Errors *shared.Errors
}

func (o *CreatePINTERESTADSConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreatePINTERESTADSConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreatePINTERESTADSConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreatePINTERESTADSConnectionResponse) GetConnectionPinterestAds() *shared.ConnectionPinterestAds {
	if o == nil {
		return nil
	}
	return o.ConnectionPinterestAds
}

func (o *CreatePINTERESTADSConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
