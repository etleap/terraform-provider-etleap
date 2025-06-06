// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateSQSConnectionRequest struct {
	ID                  string                      `pathParam:"style=simple,explode=false,name=id"`
	ConnectionSqsUpdate *shared.ConnectionSqsUpdate `request:"mediaType=application/json"`
}

func (o *UpdateSQSConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateSQSConnectionRequest) GetConnectionSqsUpdate() *shared.ConnectionSqsUpdate {
	if o == nil {
		return nil
	}
	return o.ConnectionSqsUpdate
}

type UpdateSQSConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionSqs *shared.ConnectionSqs
	// Bad Request
	Errors *shared.Errors
}

func (o *UpdateSQSConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateSQSConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateSQSConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateSQSConnectionResponse) GetConnectionSqs() *shared.ConnectionSqs {
	if o == nil {
		return nil
	}
	return o.ConnectionSqs
}

func (o *UpdateSQSConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
