// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type RetractEventRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// The pipeline version, either the currentVersion or refreshVersion.
	Version             int64                      `pathParam:"style=simple,explode=false,name=version"`
	RetractEventRequest shared.RetractEventRequest `request:"mediaType=application/json"`
}

func (o *RetractEventRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RetractEventRequest) GetVersion() int64 {
	if o == nil {
		return 0
	}
	return o.Version
}

func (o *RetractEventRequest) GetRetractEventRequest() shared.RetractEventRequest {
	if o == nil {
		return shared.RetractEventRequest{}
	}
	return o.RetractEventRequest
}

type RetractEventResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The specified externalBatchId cannot be retracted.
	Errors *shared.Errors
}

func (o *RetractEventResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RetractEventResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RetractEventResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *RetractEventResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
