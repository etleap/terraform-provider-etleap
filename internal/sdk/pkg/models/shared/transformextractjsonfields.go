// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type TransformExtractJSONFieldsType string

const (
	TransformExtractJSONFieldsTypeFlattenJSONObject TransformExtractJSONFieldsType = "FLATTEN_JSON_OBJECT"
)

func (e TransformExtractJSONFieldsType) ToPointer() *TransformExtractJSONFieldsType {
	return &e
}

func (e *TransformExtractJSONFieldsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "FLATTEN_JSON_OBJECT":
		*e = TransformExtractJSONFieldsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TransformExtractJSONFieldsType: %v", v)
	}
}

type Keys struct {
	Type Type   `json:"type"`
	Name string `json:"name"`
}

func (o *Keys) GetType() Type {
	if o == nil {
		return Type{}
	}
	return o.Type
}

func (o *Keys) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

// TransformExtractJSONFields - Flattens a JSON object into columns.
type TransformExtractJSONFields struct {
	Type TransformExtractJSONFieldsType `json:"type"`
	// The input column containing the JSON to flatten.
	Column string `json:"column"`
	// Maps keys to extract from the JSON object to their types.
	Keys []Keys `json:"keys"`
	// If enabled, Etleap will discover new JSON keys at runtime and add these to the script. See the documentation for more details: https://support.etleap.com/hc/en-us/articles/360006296714-Auto-Discovery-of-JSON-Keys
	DiscoverNewKeys bool `json:"discoverNewKeys"`
	// If specified, Etleap will prepend this value to column names produced by this transform. This can be useful to add context for column names and avoid collisions between column names of keys extracted from different objects.
	Prefix *string `json:"prefix,omitempty"`
}

func (o *TransformExtractJSONFields) GetType() TransformExtractJSONFieldsType {
	if o == nil {
		return TransformExtractJSONFieldsType("")
	}
	return o.Type
}

func (o *TransformExtractJSONFields) GetColumn() string {
	if o == nil {
		return ""
	}
	return o.Column
}

func (o *TransformExtractJSONFields) GetKeys() []Keys {
	if o == nil {
		return []Keys{}
	}
	return o.Keys
}

func (o *TransformExtractJSONFields) GetDiscoverNewKeys() bool {
	if o == nil {
		return false
	}
	return o.DiscoverNewKeys
}

func (o *TransformExtractJSONFields) GetPrefix() *string {
	if o == nil {
		return nil
	}
	return o.Prefix
}
