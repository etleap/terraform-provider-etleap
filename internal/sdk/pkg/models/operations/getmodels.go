// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetModelsRequest struct {
	// The Base64URL encoded pageToken for the page to fetch. If not present, then get the first page. Pagination is not currently supported, but will be added in a future release.
	PageToken *string `queryParam:"style=form,explode=true,name=pageToken"`
	// The size of the pages returned. If the number of elements to return is greater than the page size, then a nextPageToken will be returned with a non-null value. A pageSize of 0 indicates the server default, which is currently infinite, but may change in the future. 0 is currently the only allowed value.
	PageSize int64 `queryParam:"style=form,explode=true,name=pageSize"`
}

func (o *GetModelsRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

func (o *GetModelsRequest) GetPageSize() int64 {
	if o == nil {
		return 0
	}
	return o.PageSize
}

type GetModelsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	ModelsOutput *shared.ModelsOutput
	// Bad Request
	Errors *shared.Errors
}

func (o *GetModelsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetModelsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetModelsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetModelsResponse) GetModelsOutput() *shared.ModelsOutput {
	if o == nil {
		return nil
	}
	return o.ModelsOutput
}

func (o *GetModelsResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
