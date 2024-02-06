// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateSNOWFLAKEConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionSnowflake *shared.ConnectionSnowflake
	// Bad Request
	Errors *shared.Errors
}

func (o *CreateSNOWFLAKEConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateSNOWFLAKEConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateSNOWFLAKEConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateSNOWFLAKEConnectionResponse) GetConnectionSnowflake() *shared.ConnectionSnowflake {
	if o == nil {
		return nil
	}
	return o.ConnectionSnowflake
}

func (o *CreateSNOWFLAKEConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
