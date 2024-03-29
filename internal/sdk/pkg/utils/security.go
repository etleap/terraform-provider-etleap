// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package utils

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

const (
	securityTagKey = "security"
)

type securityTag struct {
	Option  bool
	Scheme  bool
	Name    string
	Type    string
	SubType string
}

type securityConfig struct {
	headers     map[string]string
	queryParams map[string]string
}

type SecurityClient struct {
	HTTPClient
	security func(ctx context.Context) (interface{}, error)
}

func newSecurityClient(client HTTPClient, security func(ctx context.Context) (interface{}, error)) *SecurityClient {
	return &SecurityClient{
		HTTPClient: client,
		security:   security,
	}
}

func (c *SecurityClient) Do(req *http.Request) (*http.Response, error) {
	securityCtx, err := c.security(req.Context())
	if err != nil {
		return nil, err
	}

	ctx := securityConfig{
		headers:     make(map[string]string),
		queryParams: make(map[string]string),
	}
	parseSecurityStruct(&ctx, securityCtx)

	for k, v := range ctx.headers {
		req.Header.Set(k, v)
	}

	queryParams := req.URL.Query()

	for k, v := range ctx.queryParams {
		queryParams.Add(k, v)
	}

	req.URL.RawQuery = queryParams.Encode()

	return c.HTTPClient.Do(req)
}

func ConfigureSecurityClient(c HTTPClient, security func(ctx context.Context) (interface{}, error)) *SecurityClient {
	return newSecurityClient(c, security)
}

func trueReflectValue(val reflect.Value) reflect.Value {
	kind := val.Type().Kind()
	for kind == reflect.Interface || kind == reflect.Ptr {
		innerVal := val.Elem()
		if !innerVal.IsValid() {
			break
		}
		val = innerVal
		kind = val.Type().Kind()
	}
	return val
}

func parseSecurityStruct(c *securityConfig, security interface{}) {
	securityValType := trueReflectValue(reflect.ValueOf(security))
	securityStructType := securityValType.Type()

	if isNil(securityStructType, securityValType) {
		return
	}

	if securityStructType.Kind() == reflect.Ptr {
		securityStructType = securityStructType.Elem()
		securityValType = securityValType.Elem()
	}

	for i := 0; i < securityStructType.NumField(); i++ {
		fieldType := securityStructType.Field(i)
		valType := securityValType.Field(i)

		kind := valType.Kind()

		if isNil(fieldType.Type, valType) {
			continue
		}

		if fieldType.Type.Kind() == reflect.Pointer {
			kind = valType.Elem().Kind()
		}

		secTag := parseSecurityTag(fieldType)
		if secTag != nil {
			if secTag.Option {
				handleSecurityOption(c, valType.Interface())
			} else if secTag.Scheme {
				// Special case for basic auth which could be a flattened struct
				if secTag.SubType == "basic" && kind != reflect.Struct {
					parseSecurityScheme(c, secTag, security)
				} else {
					parseSecurityScheme(c, secTag, valType.Interface())
				}
			}
		}
	}
}

func handleSecurityOption(c *securityConfig, option interface{}) error {
	optionValType := trueReflectValue(reflect.ValueOf(option))
	optionStructType := optionValType.Type()

	if isNil(optionStructType, optionValType) {
		return nil
	}

	for i := 0; i < optionStructType.NumField(); i++ {
		fieldType := optionStructType.Field(i)
		valType := optionValType.Field(i)

		secTag := parseSecurityTag(fieldType)
		if secTag != nil && secTag.Scheme {
			parseSecurityScheme(c, secTag, valType.Interface())
		}
	}

	return nil
}

func parseSecurityScheme(client *securityConfig, schemeTag *securityTag, scheme interface{}) {
	schemeVal := trueReflectValue(reflect.ValueOf(scheme))
	schemeType := schemeVal.Type()

	if isNil(schemeType, schemeVal) {
		return
	}

	if schemeType.Kind() == reflect.Struct {
		if schemeTag.Type == "http" && schemeTag.SubType == "basic" {
			handleBasicAuthScheme(client, schemeVal.Interface())
			return
		}

		for i := 0; i < schemeType.NumField(); i++ {
			fieldType := schemeType.Field(i)
			valType := schemeVal.Field(i)

			if isNil(fieldType.Type, valType) {
				continue
			}

			if fieldType.Type.Kind() == reflect.Ptr {
				valType = valType.Elem()
			}

			secTag := parseSecurityTag(fieldType)
			if secTag == nil || secTag.Name == "" {
				return
			}

			parseSecuritySchemeValue(client, schemeTag, secTag, valType.Interface())
		}
	} else {
		parseSecuritySchemeValue(client, schemeTag, schemeTag, schemeVal.Interface())
	}
}

func parseSecuritySchemeValue(client *securityConfig, schemeTag *securityTag, secTag *securityTag, val interface{}) {
	switch schemeTag.Type {
	case "apiKey":
		switch schemeTag.SubType {
		case "header":
			client.headers[secTag.Name] = valToString(val)
		case "query":
			client.queryParams[secTag.Name] = valToString(val)
		case "cookie":
			client.headers["Cookie"] = fmt.Sprintf("%s=%s", secTag.Name, valToString(val))
		default:
			panic("not supported")
		}
	case "openIdConnect":
		client.headers[secTag.Name] = prefixBearer(valToString(val))
	case "oauth2":
		client.headers[secTag.Name] = prefixBearer(valToString(val))
	case "http":
		switch schemeTag.SubType {
		case "bearer":
			client.headers[secTag.Name] = prefixBearer(valToString(val))
		default:
			panic("not supported")
		}
	default:
		panic("not supported")
	}
}

func prefixBearer(authHeaderValue string) string {
	if strings.HasPrefix(strings.ToLower(authHeaderValue), "bearer ") {
		return authHeaderValue
	}

	return fmt.Sprintf("Bearer %s", authHeaderValue)
}

func handleBasicAuthScheme(client *securityConfig, scheme interface{}) {
	schemeStructType := reflect.TypeOf(scheme)
	schemeValType := reflect.ValueOf(scheme)

	var username, password string

	for i := 0; i < schemeStructType.NumField(); i++ {
		fieldType := schemeStructType.Field(i)
		valType := schemeValType.Field(i)

		secTag := parseSecurityTag(fieldType)
		if secTag == nil || secTag.Name == "" {
			continue
		}

		switch secTag.Name {
		case "username":
			username = valType.String()
		case "password":
			password = valType.String()
		}
	}

	client.headers["Authorization"] = fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password))))
}

func parseSecurityTag(field reflect.StructField) *securityTag {
	tag := field.Tag.Get(securityTagKey)
	if tag == "" {
		return nil
	}

	option := false
	scheme := false
	name := ""
	securityType := ""
	securitySubType := ""

	options := strings.Split(tag, ",")
	for _, optionConf := range options {
		parts := strings.Split(optionConf, "=")
		if len(parts) < 1 || len(parts) > 2 {
			continue
		}

		switch parts[0] {
		case "name":
			name = parts[1]
		case "type":
			securityType = parts[1]
		case "subtype":
			securitySubType = parts[1]
		case "option":
			option = true
		case "scheme":
			scheme = true
		}
	}

	// TODO: validate tag?

	return &securityTag{
		Option:  option,
		Scheme:  scheme,
		Name:    name,
		Type:    securityType,
		SubType: securitySubType,
	}
}
