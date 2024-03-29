// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// GrantsWithoutPrivilegeOutputPaginated - A list response for grants without privilege on an object
type GrantsWithoutPrivilegeOutputPaginated struct {
	// Contains a cursor to the next page of results.
	NextPageToken *string `json:"nextPageToken,omitempty"`
	// The total number of items in the entire collection (not just this page).
	TotalSize int64                   `json:"totalSize"`
	Grants    []GrantWithoutPrivilege `json:"grants"`
}

func (o *GrantsWithoutPrivilegeOutputPaginated) GetNextPageToken() *string {
	if o == nil {
		return nil
	}
	return o.NextPageToken
}

func (o *GrantsWithoutPrivilegeOutputPaginated) GetTotalSize() int64 {
	if o == nil {
		return 0
	}
	return o.TotalSize
}

func (o *GrantsWithoutPrivilegeOutputPaginated) GetGrants() []GrantWithoutPrivilege {
	if o == nil {
		return []GrantWithoutPrivilege{}
	}
	return o.Grants
}
