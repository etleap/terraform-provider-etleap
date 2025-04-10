// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateDbtScheduleRequest struct {
	// The id of the dbt schedule
	ID                string                    `pathParam:"style=simple,explode=false,name=id"`
	DbtScheduleUpdate *shared.DbtScheduleUpdate `request:"mediaType=application/json"`
}

func (o *UpdateDbtScheduleRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateDbtScheduleRequest) GetDbtScheduleUpdate() *shared.DbtScheduleUpdate {
	if o == nil {
		return nil
	}
	return o.DbtScheduleUpdate
}

type UpdateDbtScheduleResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	DbtScheduleOutput *shared.DbtScheduleOutput
	// Not Found
	Errors *shared.Errors
}

func (o *UpdateDbtScheduleResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateDbtScheduleResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateDbtScheduleResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateDbtScheduleResponse) GetDbtScheduleOutput() *shared.DbtScheduleOutput {
	if o == nil {
		return nil
	}
	return o.DbtScheduleOutput
}

func (o *UpdateDbtScheduleResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
