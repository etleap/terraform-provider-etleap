// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetS3DATALAKEConnectionRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetS3DATALAKEConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetS3DATALAKEConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionS3DataLake *shared.ConnectionS3DataLake
	// Forbidden. You don't have access to view this connection.
	Errors *shared.Errors
}

func (o *GetS3DATALAKEConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetS3DATALAKEConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetS3DATALAKEConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetS3DATALAKEConnectionResponse) GetConnectionS3DataLake() *shared.ConnectionS3DataLake {
	if o == nil {
		return nil
	}
	return o.ConnectionS3DataLake
}

func (o *GetS3DATALAKEConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}