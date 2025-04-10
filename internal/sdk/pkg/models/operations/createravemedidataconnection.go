// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateRAVEMEDIDATAConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionRaveMedidata *shared.ConnectionRaveMedidataOutput
	// Bad Request
	Errors *shared.Errors
}

func (o *CreateRAVEMEDIDATAConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateRAVEMEDIDATAConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateRAVEMEDIDATAConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateRAVEMEDIDATAConnectionResponse) GetConnectionRaveMedidata() *shared.ConnectionRaveMedidataOutput {
	if o == nil {
		return nil
	}
	return o.ConnectionRaveMedidata
}

func (o *CreateRAVEMEDIDATAConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
