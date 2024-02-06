// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateDb2SHARDEDConnectionRequest struct {
	ID                         string                             `pathParam:"style=simple,explode=false,name=id"`
	ConnectionDb2ShardedUpdate *shared.ConnectionDb2ShardedUpdate `request:"mediaType=application/json"`
}

func (o *UpdateDb2SHARDEDConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateDb2SHARDEDConnectionRequest) GetConnectionDb2ShardedUpdate() *shared.ConnectionDb2ShardedUpdate {
	if o == nil {
		return nil
	}
	return o.ConnectionDb2ShardedUpdate
}

type UpdateDb2SHARDEDConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionDb2Sharded *shared.ConnectionDb2Sharded
	// Bad Request
	Errors *shared.Errors
}

func (o *UpdateDb2SHARDEDConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateDb2SHARDEDConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateDb2SHARDEDConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateDb2SHARDEDConnectionResponse) GetConnectionDb2Sharded() *shared.ConnectionDb2Sharded {
	if o == nil {
		return nil
	}
	return o.ConnectionDb2Sharded
}

func (o *UpdateDb2SHARDEDConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
