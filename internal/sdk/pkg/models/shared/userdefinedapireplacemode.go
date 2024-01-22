// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"errors"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
)

type Two struct {
}

type UserDefinedAPIReplaceMode1Type string

const (
	UserDefinedAPIReplaceMode1TypeStr UserDefinedAPIReplaceMode1Type = "str"
	UserDefinedAPIReplaceMode1TypeTwo UserDefinedAPIReplaceMode1Type = "2"
)

type UserDefinedAPIReplaceMode1 struct {
	Str *string
	Two *Two

	Type UserDefinedAPIReplaceMode1Type
}

func CreateUserDefinedAPIReplaceMode1Str(str string) UserDefinedAPIReplaceMode1 {
	typ := UserDefinedAPIReplaceMode1TypeStr

	return UserDefinedAPIReplaceMode1{
		Str:  &str,
		Type: typ,
	}
}

func CreateUserDefinedAPIReplaceMode1Two(two Two) UserDefinedAPIReplaceMode1 {
	typ := UserDefinedAPIReplaceMode1TypeTwo

	return UserDefinedAPIReplaceMode1{
		Two:  &two,
		Type: typ,
	}
}

func (u *UserDefinedAPIReplaceMode1) UnmarshalJSON(data []byte) error {

	two := new(Two)
	if err := utils.UnmarshalJSON(data, &two, "", true, true); err == nil {
		u.Two = two
		u.Type = UserDefinedAPIReplaceMode1TypeTwo
		return nil
	}

	str := new(string)
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = str
		u.Type = UserDefinedAPIReplaceMode1TypeStr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u UserDefinedAPIReplaceMode1) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Two != nil {
		return utils.MarshalJSON(u.Two, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type UserDefinedAPIReplaceModeType string

const (
	UserDefinedAPIReplaceModeTypeUserDefinedAPIReplaceMode1 UserDefinedAPIReplaceModeType = "user_defined_api_replace_mode_1"
	UserDefinedAPIReplaceModeTypeExtendedReplaceMode        UserDefinedAPIReplaceModeType = "ExtendedReplaceMode"
)

type UserDefinedAPIReplaceMode struct {
	UserDefinedAPIReplaceMode1 *UserDefinedAPIReplaceMode1
	ExtendedReplaceMode        *ExtendedReplaceMode

	Type UserDefinedAPIReplaceModeType
}

func CreateUserDefinedAPIReplaceModeUserDefinedAPIReplaceMode1(userDefinedAPIReplaceMode1 UserDefinedAPIReplaceMode1) UserDefinedAPIReplaceMode {
	typ := UserDefinedAPIReplaceModeTypeUserDefinedAPIReplaceMode1

	return UserDefinedAPIReplaceMode{
		UserDefinedAPIReplaceMode1: &userDefinedAPIReplaceMode1,
		Type:                       typ,
	}
}

func CreateUserDefinedAPIReplaceModeExtendedReplaceMode(extendedReplaceMode ExtendedReplaceMode) UserDefinedAPIReplaceMode {
	typ := UserDefinedAPIReplaceModeTypeExtendedReplaceMode

	return UserDefinedAPIReplaceMode{
		ExtendedReplaceMode: &extendedReplaceMode,
		Type:                typ,
	}
}

func (u *UserDefinedAPIReplaceMode) UnmarshalJSON(data []byte) error {

	extendedReplaceMode := new(ExtendedReplaceMode)
	if err := utils.UnmarshalJSON(data, &extendedReplaceMode, "", true, true); err == nil {
		u.ExtendedReplaceMode = extendedReplaceMode
		u.Type = UserDefinedAPIReplaceModeTypeExtendedReplaceMode
		return nil
	}

	userDefinedAPIReplaceMode1 := new(UserDefinedAPIReplaceMode1)
	if err := utils.UnmarshalJSON(data, &userDefinedAPIReplaceMode1, "", true, true); err == nil {
		u.UserDefinedAPIReplaceMode1 = userDefinedAPIReplaceMode1
		u.Type = UserDefinedAPIReplaceModeTypeUserDefinedAPIReplaceMode1
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u UserDefinedAPIReplaceMode) MarshalJSON() ([]byte, error) {
	if u.UserDefinedAPIReplaceMode1 != nil {
		return utils.MarshalJSON(u.UserDefinedAPIReplaceMode1, "", true)
	}

	if u.ExtendedReplaceMode != nil {
		return utils.MarshalJSON(u.ExtendedReplaceMode, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}
