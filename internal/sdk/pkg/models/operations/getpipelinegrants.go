// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetPipelineGrantsRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// The Base64URL encoded pageToken for the page to fetch. If not present, then get the first page. Pagination is not currently supported, but will be added in a future release.
	PageToken *string `queryParam:"style=form,explode=true,name=pageToken"`
	// The size of the pages returned. If the number of elements to return is greater than the page size, then a nextPageToken will be returned with a non-null value. A pageSize of 0 indicates the server default, which is currently infinite, but may change in the future. 0 is currently the only allowed value.
	PageSize int64 `queryParam:"style=form,explode=true,name=pageSize"`
}

func (o *GetPipelineGrantsRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetPipelineGrantsRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

func (o *GetPipelineGrantsRequest) GetPageSize() int64 {
	if o == nil {
		return 0
	}
	return o.PageSize
}

type GetPipelineGrantsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	GrantsWithoutPrivilegeOutputPaginated *shared.GrantsWithoutPrivilegeOutputPaginated
	// A pipeline with the given id was not found.
	Errors *shared.Errors
}

func (o *GetPipelineGrantsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetPipelineGrantsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetPipelineGrantsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetPipelineGrantsResponse) GetGrantsWithoutPrivilegeOutputPaginated() *shared.GrantsWithoutPrivilegeOutputPaginated {
	if o == nil {
		return nil
	}
	return o.GrantsWithoutPrivilegeOutputPaginated
}

func (o *GetPipelineGrantsResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
