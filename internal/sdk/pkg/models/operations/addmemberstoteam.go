// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/models/shared"
	"net/http"
)

type AddMembersToTeamAddMembersToTeamRequest struct {
	Members []shared.UserInput `json:"members"`
}

func (o *AddMembersToTeamAddMembersToTeamRequest) GetMembers() []shared.UserInput {
	if o == nil {
		return []shared.UserInput{}
	}
	return o.Members
}

type AddMembersToTeamRequest struct {
	ID          string                                   `pathParam:"style=simple,explode=false,name=id"`
	RequestBody *AddMembersToTeamAddMembersToTeamRequest `request:"mediaType=application/json"`
}

func (o *AddMembersToTeamRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AddMembersToTeamRequest) GetRequestBody() *AddMembersToTeamAddMembersToTeamRequest {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type AddMembersToTeamResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	TeamOutput *shared.TeamOutput
	// Team for this id was not found.
	Errors *shared.Errors
}

func (o *AddMembersToTeamResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AddMembersToTeamResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AddMembersToTeamResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *AddMembersToTeamResponse) GetTeamOutput() *shared.TeamOutput {
	if o == nil {
		return nil
	}
	return o.TeamOutput
}

func (o *AddMembersToTeamResponse) GetErrors() *shared.Errors {
	if o == nil {
		return nil
	}
	return o.Errors
}
