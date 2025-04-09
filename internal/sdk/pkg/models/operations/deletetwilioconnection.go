// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type DeleteTWILIOConnectionRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Body required only for REDSHIFT and SNOWFLAKE destinations.
	ConnectionDelete *shared.ConnectionDelete `request:"mediaType=application/json"`
}

func (o *DeleteTWILIOConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DeleteTWILIOConnectionRequest) GetConnectionDelete() *shared.ConnectionDelete {
	if o == nil {
		return nil
	}
	return o.ConnectionDelete
}

type DeleteTWILIOConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Connection for this id was not found.
	Errors *shared.Errors
}

func (o *DeleteTWILIOConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteTWILIOConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteTWILIOConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteTWILIOConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
