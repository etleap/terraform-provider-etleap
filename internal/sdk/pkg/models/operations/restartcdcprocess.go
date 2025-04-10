// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type RestartCdcProcessRequest struct {
	// The unique identifier of the connection you want to restart.
	ID          string              `pathParam:"style=simple,explode=false,name=id"`
	CdcRestarts *shared.CdcRestarts `request:"mediaType=application/json"`
}

func (o *RestartCdcProcessRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RestartCdcProcessRequest) GetCdcRestarts() *shared.CdcRestarts {
	if o == nil {
		return nil
	}
	return o.CdcRestarts
}

type RestartCdcProcessResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// User does not have permission to restart the CDC process for this connection.
	Errors *shared.Errors
}

func (o *RestartCdcProcessResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RestartCdcProcessResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RestartCdcProcessResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *RestartCdcProcessResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
