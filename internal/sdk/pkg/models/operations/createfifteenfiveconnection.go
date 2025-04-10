// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateFIFTEENFIVEConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionFifteenFive *shared.ConnectionFifteenFiveOutput
	// Bad Request
	Errors *shared.Errors
}

func (o *CreateFIFTEENFIVEConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateFIFTEENFIVEConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateFIFTEENFIVEConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateFIFTEENFIVEConnectionResponse) GetConnectionFifteenFive() *shared.ConnectionFifteenFiveOutput {
	if o == nil {
		return nil
	}
	return o.ConnectionFifteenFive
}

func (o *CreateFIFTEENFIVEConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
