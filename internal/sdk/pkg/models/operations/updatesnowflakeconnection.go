// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateSNOWFLAKEConnectionRequest struct {
	ID                        string                            `pathParam:"style=simple,explode=false,name=id"`
	ConnectionSnowflakeUpdate *shared.ConnectionSnowflakeUpdate `request:"mediaType=application/json"`
}

func (o *UpdateSNOWFLAKEConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateSNOWFLAKEConnectionRequest) GetConnectionSnowflakeUpdate() *shared.ConnectionSnowflakeUpdate {
	if o == nil {
		return nil
	}
	return o.ConnectionSnowflakeUpdate
}

type UpdateSNOWFLAKEConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionSnowflake *shared.ConnectionSnowflake
	// Connection for this id was not found.
	Errors *shared.Errors
}

func (o *UpdateSNOWFLAKEConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateSNOWFLAKEConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateSNOWFLAKEConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateSNOWFLAKEConnectionResponse) GetConnectionSnowflake() *shared.ConnectionSnowflake {
	if o == nil {
		return nil
	}
	return o.ConnectionSnowflake
}

func (o *UpdateSNOWFLAKEConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
