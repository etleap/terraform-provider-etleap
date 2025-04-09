// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetOUTLOOKConnectionRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetOUTLOOKConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetOUTLOOKConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionOutlook *shared.ConnectionOutlookOutput
	// Not Found.
	Errors *shared.Errors
}

func (o *GetOUTLOOKConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetOUTLOOKConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetOUTLOOKConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetOUTLOOKConnectionResponse) GetConnectionOutlook() *shared.ConnectionOutlookOutput {
	if o == nil {
		return nil
	}
	return o.ConnectionOutlook
}

func (o *GetOUTLOOKConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
