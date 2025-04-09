// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateJIRACLOUDConnectionRequest struct {
	ID                        string                            `pathParam:"style=simple,explode=false,name=id"`
	ConnectionJiraCloudUpdate *shared.ConnectionJiraCloudUpdate `request:"mediaType=application/json"`
}

func (o *UpdateJIRACLOUDConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateJIRACLOUDConnectionRequest) GetConnectionJiraCloudUpdate() *shared.ConnectionJiraCloudUpdate {
	if o == nil {
		return nil
	}
	return o.ConnectionJiraCloudUpdate
}

type UpdateJIRACLOUDConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionJiraCloud *shared.ConnectionJiraCloud
	// Bad Request
	Errors *shared.Errors
}

func (o *UpdateJIRACLOUDConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateJIRACLOUDConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateJIRACLOUDConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateJIRACLOUDConnectionResponse) GetConnectionJiraCloud() *shared.ConnectionJiraCloud {
	if o == nil {
		return nil
	}
	return o.ConnectionJiraCloud
}

func (o *UpdateJIRACLOUDConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
