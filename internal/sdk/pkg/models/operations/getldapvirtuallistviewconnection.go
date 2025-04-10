// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetLDAPVIRTUALLISTVIEWConnectionRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetLDAPVIRTUALLISTVIEWConnectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetLDAPVIRTUALLISTVIEWConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK.
	ConnectionLdapVirtualListView *shared.ConnectionLdapVirtualListView
	// Not Found.
	Errors *shared.Errors
}

func (o *GetLDAPVIRTUALLISTVIEWConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetLDAPVIRTUALLISTVIEWConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetLDAPVIRTUALLISTVIEWConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetLDAPVIRTUALLISTVIEWConnectionResponse) GetConnectionLdapVirtualListView() *shared.ConnectionLdapVirtualListView {
	if o == nil {
		return nil
	}
	return o.ConnectionLdapVirtualListView
}

func (o *GetLDAPVIRTUALLISTVIEWConnectionResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
