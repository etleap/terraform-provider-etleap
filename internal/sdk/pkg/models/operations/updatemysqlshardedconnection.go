// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateMYSQLSHARDEDConnectionRequest struct {
	ID                           string                               `pathParam:"style=simple,explode=false,name=id"`
	ConnectionMysqlShardedUpdate *shared.ConnectionMysqlShardedUpdate `request:"mediaType=application/json"`
}

func (o *UpdateMYSQLSHARDEDConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateMYSQLSHARDEDConnectionRequest) GetConnectionMysqlShardedUpdate() *shared.ConnectionMysqlShardedUpdate {
	if o == nil {
		return nil
	}
	return o.ConnectionMysqlShardedUpdate
}

type UpdateMYSQLSHARDEDConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionMysqlSharded *shared.ConnectionMysqlShardedOutput
	// Connection for this id was not found.
	Errors *shared.Errors
}

func (o *UpdateMYSQLSHARDEDConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateMYSQLSHARDEDConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateMYSQLSHARDEDConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateMYSQLSHARDEDConnectionResponse) GetConnectionMysqlSharded() *shared.ConnectionMysqlShardedOutput {
	if o == nil {
		return nil
	}
	return o.ConnectionMysqlSharded
}

func (o *UpdateMYSQLSHARDEDConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
