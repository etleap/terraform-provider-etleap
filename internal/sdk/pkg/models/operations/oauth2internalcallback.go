// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type Oauth2InternalCallbackRequest struct {
	// Encrypted relay state
	State string `queryParam:"style=form,explode=true,name=state"`
	// OAuth2 code - returned for every connection except JIRA
	Code *string `queryParam:"style=form,explode=true,name=code"`
	// OAuth2 authentication error
	Error *string `queryParam:"style=form,explode=true,name=error"`
	// OAuth2 authentication error description
	ErrorDescription *string `queryParam:"style=form,explode=true,name=error_description"`
}

func (o *Oauth2InternalCallbackRequest) GetState() string {
	if o == nil {
		return ""
	}
	return o.State
}

func (o *Oauth2InternalCallbackRequest) GetCode() *string {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *Oauth2InternalCallbackRequest) GetError() *string {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *Oauth2InternalCallbackRequest) GetErrorDescription() *string {
	if o == nil {
		return nil
	}
	return o.ErrorDescription
}

type Oauth2InternalCallbackResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Bad Request
	Errors  *shared.Errors
	Headers map[string][]string
}

func (o *Oauth2InternalCallbackResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *Oauth2InternalCallbackResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *Oauth2InternalCallbackResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *Oauth2InternalCallbackResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}

func (o *Oauth2InternalCallbackResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}
