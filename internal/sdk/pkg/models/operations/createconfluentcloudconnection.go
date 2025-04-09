// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateCONFLUENTCLOUDConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionConfluentCloud *shared.ConnectionConfluentCloud
	// Bad Request
	Errors *shared.Errors
}

func (o *CreateCONFLUENTCLOUDConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateCONFLUENTCLOUDConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateCONFLUENTCLOUDConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateCONFLUENTCLOUDConnectionResponse) GetConnectionConfluentCloud() *shared.ConnectionConfluentCloud {
	if o == nil {
		return nil
	}
	return o.ConnectionConfluentCloud
}

func (o *CreateCONFLUENTCLOUDConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
