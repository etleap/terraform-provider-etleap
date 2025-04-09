// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateELASTICSEARCHConnectionRequest struct {
	ID                            string                                `pathParam:"style=simple,explode=false,name=id"`
	ConnectionElasticSearchUpdate *shared.ConnectionElasticSearchUpdate `request:"mediaType=application/json"`
}

func (o *UpdateELASTICSEARCHConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateELASTICSEARCHConnectionRequest) GetConnectionElasticSearchUpdate() *shared.ConnectionElasticSearchUpdate {
	if o == nil {
		return nil
	}
	return o.ConnectionElasticSearchUpdate
}

type UpdateELASTICSEARCHConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionElasticSearch *shared.ConnectionElasticSearch
	// Connection for this id was not found.
	Errors *shared.Errors
}

func (o *UpdateELASTICSEARCHConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateELASTICSEARCHConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateELASTICSEARCHConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateELASTICSEARCHConnectionResponse) GetConnectionElasticSearch() *shared.ConnectionElasticSearch {
	if o == nil {
		return nil
	}
	return o.ConnectionElasticSearch
}

func (o *UpdateELASTICSEARCHConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
