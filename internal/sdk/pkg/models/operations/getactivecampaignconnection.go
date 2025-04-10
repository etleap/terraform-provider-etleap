// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetACTIVECAMPAIGNConnectionRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetACTIVECAMPAIGNConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetACTIVECAMPAIGNConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionActiveCampaign *shared.ConnectionActiveCampaignOutput
	// Not Found.
	Errors *shared.Errors
}

func (o *GetACTIVECAMPAIGNConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetACTIVECAMPAIGNConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetACTIVECAMPAIGNConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetACTIVECAMPAIGNConnectionResponse) GetConnectionActiveCampaign() *shared.ConnectionActiveCampaignOutput {
	if o == nil {
		return nil
	}
	return o.ConnectionActiveCampaign
}

func (o *GetACTIVECAMPAIGNConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
