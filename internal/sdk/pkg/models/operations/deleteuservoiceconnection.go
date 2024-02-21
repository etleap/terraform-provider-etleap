// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type DeleteUSERVOICEConnectionRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Body required only for REDSHIFT and SNOWFLAKE destinations.
	ConnectionDelete *shared.ConnectionDelete `request:"mediaType=application/json"`
}

func (o *DeleteUSERVOICEConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DeleteUSERVOICEConnectionRequest) GetConnectionDelete() *shared.ConnectionDelete {
	if o == nil {
		return nil
	}
	return o.ConnectionDelete
}

type DeleteUSERVOICEConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Bad Request.
	Errors *shared.Errors
}

func (o *DeleteUSERVOICEConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteUSERVOICEConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteUSERVOICEConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteUSERVOICEConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}