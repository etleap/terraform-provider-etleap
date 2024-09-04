// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SnowflakeAuthenticationKeyPairType string

const (
	SnowflakeAuthenticationKeyPairTypePassword SnowflakeAuthenticationKeyPairType = "PASSWORD"
	SnowflakeAuthenticationKeyPairTypeKeyPair  SnowflakeAuthenticationKeyPairType = "KEY_PAIR"
)

func (e SnowflakeAuthenticationKeyPairType) ToPointer() *SnowflakeAuthenticationKeyPairType {
	return &e
}

func (e *SnowflakeAuthenticationKeyPairType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "PASSWORD":
		fallthrough
	case "KEY_PAIR":
		*e = SnowflakeAuthenticationKeyPairType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SnowflakeAuthenticationKeyPairType: %v", v)
	}
}

// SnowflakeAuthenticationKeyPairOutput - Snowflake Key Pair Authentication
type SnowflakeAuthenticationKeyPairOutput struct {
	Type      SnowflakeAuthenticationKeyPairType `json:"type"`
	PublicKey string                             `json:"publicKey"`
}

func (o *SnowflakeAuthenticationKeyPairOutput) GetType() SnowflakeAuthenticationKeyPairType {
	if o == nil {
		return SnowflakeAuthenticationKeyPairType("")
	}
	return o.Type
}

func (o *SnowflakeAuthenticationKeyPairOutput) GetPublicKey() string {
	if o == nil {
		return ""
	}
	return o.PublicKey
}

// SnowflakeAuthenticationKeyPair - Snowflake Key Pair Authentication
type SnowflakeAuthenticationKeyPair struct {
	Type       SnowflakeAuthenticationKeyPairType `json:"type"`
	PrivateKey string                             `json:"privateKey"`
}

func (o *SnowflakeAuthenticationKeyPair) GetType() SnowflakeAuthenticationKeyPairType {
	if o == nil {
		return SnowflakeAuthenticationKeyPairType("")
	}
	return o.Type
}

func (o *SnowflakeAuthenticationKeyPair) GetPrivateKey() string {
	if o == nil {
		return ""
	}
	return o.PrivateKey
}
