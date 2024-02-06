// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateACTIVECAMPAIGNConnectionRequest struct {
	ID                             string                                 `pathParam:"style=simple,explode=false,name=id"`
	ConnectionActiveCampaignUpdate *shared.ConnectionActiveCampaignUpdate `request:"mediaType=application/json"`
}

func (o *UpdateACTIVECAMPAIGNConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateACTIVECAMPAIGNConnectionRequest) GetConnectionActiveCampaignUpdate() *shared.ConnectionActiveCampaignUpdate {
	if o == nil {
		return nil
	}
	return o.ConnectionActiveCampaignUpdate
}

type UpdateACTIVECAMPAIGNConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionActiveCampaign *shared.ConnectionActiveCampaign
	// Bad Request
	Errors *shared.Errors
}

func (o *UpdateACTIVECAMPAIGNConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateACTIVECAMPAIGNConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateACTIVECAMPAIGNConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateACTIVECAMPAIGNConnectionResponse) GetConnectionActiveCampaign() *shared.ConnectionActiveCampaign {
	if o == nil {
		return nil
	}
	return o.ConnectionActiveCampaign
}

func (o *UpdateACTIVECAMPAIGNConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
