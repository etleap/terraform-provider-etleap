// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetTWITTERADSConnectionRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetTWITTERADSConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetTWITTERADSConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionTwitter *shared.ConnectionTwitter
	// Not Found.
	Errors *shared.Errors
}

func (o *GetTWITTERADSConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTWITTERADSConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTWITTERADSConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTWITTERADSConnectionResponse) GetConnectionTwitter() *shared.ConnectionTwitter {
	if o == nil {
		return nil
	}
	return o.ConnectionTwitter
}

func (o *GetTWITTERADSConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
