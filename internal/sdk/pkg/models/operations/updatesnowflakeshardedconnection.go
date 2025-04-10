// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateSNOWFLAKESHARDEDConnectionRequest struct {
	ID                               string                                   `pathParam:"style=simple,explode=false,name=id"`
	ConnectionSnowflakeShardedUpdate *shared.ConnectionSnowflakeShardedUpdate `request:"mediaType=application/json"`
}

func (o *UpdateSNOWFLAKESHARDEDConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateSNOWFLAKESHARDEDConnectionRequest) GetConnectionSnowflakeShardedUpdate() *shared.ConnectionSnowflakeShardedUpdate {
	if o == nil {
		return nil
	}
	return o.ConnectionSnowflakeShardedUpdate
}

type UpdateSNOWFLAKESHARDEDConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionSnowflakeSharded *shared.ConnectionSnowflakeSharded
	// Connection for this id was not found.
	Errors *shared.Errors
}

func (o *UpdateSNOWFLAKESHARDEDConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateSNOWFLAKESHARDEDConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateSNOWFLAKESHARDEDConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateSNOWFLAKESHARDEDConnectionResponse) GetConnectionSnowflakeSharded() *shared.ConnectionSnowflakeSharded {
	if o == nil {
		return nil
	}
	return o.ConnectionSnowflakeSharded
}

func (o *UpdateSNOWFLAKESHARDEDConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
