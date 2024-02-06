// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetUSERVOICEConnectionRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetUSERVOICEConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetUSERVOICEConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionUserVoice *shared.ConnectionUserVoice
	// Forbidden. You don't have access to view this connection.
	Errors *shared.Errors
}

func (o *GetUSERVOICEConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetUSERVOICEConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetUSERVOICEConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetUSERVOICEConnectionResponse) GetConnectionUserVoice() *shared.ConnectionUserVoice {
	if o == nil {
		return nil
	}
	return o.ConnectionUserVoice
}

func (o *GetUSERVOICEConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
