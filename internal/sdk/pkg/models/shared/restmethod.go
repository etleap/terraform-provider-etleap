// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
)

type RestMethod1 string

const (
	RestMethod1Get RestMethod1 = "GET"
)

func (e RestMethod1) ToPointer() *RestMethod1 {
	return &e
}

func (e *RestMethod1) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "GET":
		*e = RestMethod1(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RestMethod1: %v", v)
	}
}

type RestMethodType string

const (
	RestMethodTypeRestMethod1 RestMethodType = "rest_method_1"
	RestMethodTypePostMethod  RestMethodType = "post_method"
)

// RestMethod - Can be either the string `GET` or `POST`, which includes body parameters.
type RestMethod struct {
	RestMethod1 *RestMethod1
	PostMethod  *PostMethod

	Type RestMethodType
}

func CreateRestMethodRestMethod1(restMethod1 RestMethod1) RestMethod {
	typ := RestMethodTypeRestMethod1

	return RestMethod{
		RestMethod1: &restMethod1,
		Type:        typ,
	}
}

func CreateRestMethodPostMethod(postMethod PostMethod) RestMethod {
	typ := RestMethodTypePostMethod

	return RestMethod{
		PostMethod: &postMethod,
		Type:       typ,
	}
}

func (u *RestMethod) UnmarshalJSON(data []byte) error {

	postMethod := new(PostMethod)
	if err := utils.UnmarshalJSON(data, &postMethod, "", true, true); err == nil {
		u.PostMethod = postMethod
		u.Type = RestMethodTypePostMethod
		return nil
	}

	restMethod1 := new(RestMethod1)
	if err := utils.UnmarshalJSON(data, &restMethod1, "", true, true); err == nil {
		u.RestMethod1 = restMethod1
		u.Type = RestMethodTypeRestMethod1
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u RestMethod) MarshalJSON() ([]byte, error) {
	if u.RestMethod1 != nil {
		return utils.MarshalJSON(u.RestMethod1, "", true)
	}

	if u.PostMethod != nil {
		return utils.MarshalJSON(u.PostMethod, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}
