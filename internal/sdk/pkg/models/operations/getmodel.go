// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetModelRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetModelRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetModelResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	ModelOutput *shared.ModelOutput
	// Bad Request
	Errors *shared.Errors
}

func (o *GetModelResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetModelResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetModelResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetModelResponse) GetModelOutput() *shared.ModelOutput {
	if o == nil {
		return nil
	}
	return o.ModelOutput
}

func (o *GetModelResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
