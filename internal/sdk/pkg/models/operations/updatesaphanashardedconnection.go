// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateSAPHANASHARDEDConnectionRequest struct {
	ID                             string                                 `pathParam:"style=simple,explode=false,name=id"`
	ConnectionSapHanaShardedUpdate *shared.ConnectionSapHanaShardedUpdate `request:"mediaType=application/json"`
}

func (o *UpdateSAPHANASHARDEDConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateSAPHANASHARDEDConnectionRequest) GetConnectionSapHanaShardedUpdate() *shared.ConnectionSapHanaShardedUpdate {
	if o == nil {
		return nil
	}
	return o.ConnectionSapHanaShardedUpdate
}

type UpdateSAPHANASHARDEDConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionSapHanaSharded *shared.ConnectionSapHanaShardedOutput
	// Connection for this id was not found.
	Errors *shared.Errors
}

func (o *UpdateSAPHANASHARDEDConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateSAPHANASHARDEDConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateSAPHANASHARDEDConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateSAPHANASHARDEDConnectionResponse) GetConnectionSapHanaSharded() *shared.ConnectionSapHanaShardedOutput {
	if o == nil {
		return nil
	}
	return o.ConnectionSapHanaSharded
}

func (o *UpdateSAPHANASHARDEDConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
