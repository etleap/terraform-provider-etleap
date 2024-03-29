// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// GrantsWithPrivilegeListInput - A list of grants with privilege on an object
type GrantsWithPrivilegeListInput struct {
	Grants []GrantWithPrivilegeInput `json:"grants"`
}

func (o *GrantsWithPrivilegeListInput) GetGrants() []GrantWithPrivilegeInput {
	if o == nil {
		return []GrantWithPrivilegeInput{}
	}
	return o.Grants
}
