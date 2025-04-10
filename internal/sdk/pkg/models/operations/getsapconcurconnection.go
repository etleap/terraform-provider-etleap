// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetSAPCONCURConnectionRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetSAPCONCURConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetSAPCONCURConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionSapConcur *shared.ConnectionSapConcurOutput
	// Not Found.
	Errors *shared.Errors
}

func (o *GetSAPCONCURConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSAPCONCURConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSAPCONCURConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSAPCONCURConnectionResponse) GetConnectionSapConcur() *shared.ConnectionSapConcurOutput {
	if o == nil {
		return nil
	}
	return o.ConnectionSapConcur
}

func (o *GetSAPCONCURConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
