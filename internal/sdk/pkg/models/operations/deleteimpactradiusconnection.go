// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type DeleteIMPACTRADIUSConnectionRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Body required only for REDSHIFT and SNOWFLAKE destinations.
	ConnectionDelete *shared.ConnectionDelete `request:"mediaType=application/json"`
}

func (o *DeleteIMPACTRADIUSConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DeleteIMPACTRADIUSConnectionRequest) GetConnectionDelete() *shared.ConnectionDelete {
	if o == nil {
		return nil
	}
	return o.ConnectionDelete
}

type DeleteIMPACTRADIUSConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Bad Request.
	Errors *shared.Errors
}

func (o *DeleteIMPACTRADIUSConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteIMPACTRADIUSConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteIMPACTRADIUSConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteIMPACTRADIUSConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
