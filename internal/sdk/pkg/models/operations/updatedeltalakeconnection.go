// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateDELTALAKEConnectionRequest struct {
	ID                        string                            `pathParam:"style=simple,explode=false,name=id"`
	ConnectionDeltaLakeUpdate *shared.ConnectionDeltaLakeUpdate `request:"mediaType=application/json"`
}

func (o *UpdateDELTALAKEConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateDELTALAKEConnectionRequest) GetConnectionDeltaLakeUpdate() *shared.ConnectionDeltaLakeUpdate {
	if o == nil {
		return nil
	}
	return o.ConnectionDeltaLakeUpdate
}

type UpdateDELTALAKEConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionDeltaLake *shared.ConnectionDeltaLakeOutput
	// Connection for this id was not found.
	Errors *shared.Errors
}

func (o *UpdateDELTALAKEConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateDELTALAKEConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateDELTALAKEConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateDELTALAKEConnectionResponse) GetConnectionDeltaLake() *shared.ConnectionDeltaLakeOutput {
	if o == nil {
		return nil
	}
	return o.ConnectionDeltaLake
}

func (o *UpdateDELTALAKEConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
