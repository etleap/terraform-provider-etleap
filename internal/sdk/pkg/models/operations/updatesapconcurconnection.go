// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateSAPCONCURConnectionRequest struct {
	ID                        string                            `pathParam:"style=simple,explode=false,name=id"`
	ConnectionSapConcurUpdate *shared.ConnectionSapConcurUpdate `request:"mediaType=application/json"`
}

func (o *UpdateSAPCONCURConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateSAPCONCURConnectionRequest) GetConnectionSapConcurUpdate() *shared.ConnectionSapConcurUpdate {
	if o == nil {
		return nil
	}
	return o.ConnectionSapConcurUpdate
}

type UpdateSAPCONCURConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionSapConcur *shared.ConnectionSapConcur
	// Bad Request
	Errors *shared.Errors
}

func (o *UpdateSAPCONCURConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateSAPCONCURConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateSAPCONCURConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateSAPCONCURConnectionResponse) GetConnectionSapConcur() *shared.ConnectionSapConcur {
	if o == nil {
		return nil
	}
	return o.ConnectionSapConcur
}

func (o *UpdateSAPCONCURConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
