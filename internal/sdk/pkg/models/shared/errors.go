// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Errors struct {
	// List of error messages.
	Errors []string `json:"errors"`
}

func (o *Errors) GetErrors() []string {
	if o == nil {
		return []string{}
	}
	return o.Errors
}
