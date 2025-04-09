// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateSQLSERVERSHARDEDConnectionRequest struct {
	ID                               string                                   `pathParam:"style=simple,explode=false,name=id"`
	ConnectionSQLServerShardedUpdate *shared.ConnectionSQLServerShardedUpdate `request:"mediaType=application/json"`
}

func (o *UpdateSQLSERVERSHARDEDConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateSQLSERVERSHARDEDConnectionRequest) GetConnectionSQLServerShardedUpdate() *shared.ConnectionSQLServerShardedUpdate {
	if o == nil {
		return nil
	}
	return o.ConnectionSQLServerShardedUpdate
}

type UpdateSQLSERVERSHARDEDConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionSQLServerSharded *shared.ConnectionSQLServerShardedOutput
	// Connection for this id was not found.
	Errors *shared.Errors
}

func (o *UpdateSQLSERVERSHARDEDConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateSQLSERVERSHARDEDConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateSQLSERVERSHARDEDConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateSQLSERVERSHARDEDConnectionResponse) GetConnectionSQLServerSharded() *shared.ConnectionSQLServerShardedOutput {
	if o == nil {
		return nil
	}
	return o.ConnectionSQLServerSharded
}

func (o *UpdateSQLSERVERSHARDEDConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
