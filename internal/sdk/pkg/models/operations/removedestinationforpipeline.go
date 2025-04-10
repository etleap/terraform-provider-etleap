// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type RemoveDestinationForPipelineRequest struct {
	ID             string                 `pathParam:"style=simple,explode=false,name=id"`
	ConnectionID   string                 `pathParam:"style=simple,explode=false,name=connectionId"`
	PipelineDelete *shared.PipelineDelete `request:"mediaType=application/json"`
}

func (o *RemoveDestinationForPipelineRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RemoveDestinationForPipelineRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *RemoveDestinationForPipelineRequest) GetPipelineDelete() *shared.PipelineDelete {
	if o == nil {
		return nil
	}
	return o.PipelineDelete
}

type RemoveDestinationForPipelineResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The pipeline with this id was not found, or the connection with this id is not a destination for this pipeline.
	PipelineDelete *shared.PipelineDelete
	// You don't have the required permission to remove destinations from this pipeline. Please contact your organization admin or the pipeline's owner.
	Errors *shared.Errors
}

func (o *RemoveDestinationForPipelineResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RemoveDestinationForPipelineResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RemoveDestinationForPipelineResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *RemoveDestinationForPipelineResponse) GetPipelineDelete() *shared.PipelineDelete {
	if o == nil {
		return nil
	}
	return o.PipelineDelete
}

func (o *RemoveDestinationForPipelineResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
