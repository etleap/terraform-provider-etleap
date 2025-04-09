// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateSQUAREConnectionRequest struct {
	ID                     string                         `pathParam:"style=simple,explode=false,name=id"`
	ConnectionSquareUpdate *shared.ConnectionSquareUpdate `request:"mediaType=application/json"`
}

func (o *UpdateSQUAREConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateSQUAREConnectionRequest) GetConnectionSquareUpdate() *shared.ConnectionSquareUpdate {
	if o == nil {
		return nil
	}
	return o.ConnectionSquareUpdate
}

type UpdateSQUAREConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionSquare *shared.ConnectionSquareOutput
	// Connection for this id was not found.
	Errors *shared.Errors
}

func (o *UpdateSQUAREConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateSQUAREConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateSQUAREConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateSQUAREConnectionResponse) GetConnectionSquare() *shared.ConnectionSquareOutput {
	if o == nil {
		return nil
	}
	return o.ConnectionSquare
}

func (o *UpdateSQUAREConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
