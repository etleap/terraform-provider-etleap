// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type DeletePipelineRequest struct {
	ID             string                 `pathParam:"style=simple,explode=false,name=id"`
	PipelineDelete *shared.PipelineDelete `request:"mediaType=application/json"`
}

func (o *DeletePipelineRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DeletePipelineRequest) GetPipelineDelete() *shared.PipelineDelete {
	if o == nil {
		return nil
	}
	return o.PipelineDelete
}

type DeletePipelineResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Bad Request. You must specify the `deletionOfExportProducts` property.
	Errors *shared.Errors
}

func (o *DeletePipelineResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeletePipelineResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeletePipelineResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeletePipelineResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
