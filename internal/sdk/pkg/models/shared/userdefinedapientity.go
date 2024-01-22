// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/etleap/terraform-provider-etleap/internal/sdk/pkg/utils"
)

// RestMethod - The HTTP method used to call the apiUrl.
type RestMethod string

const (
	RestMethodGet  RestMethod = "GET"
	RestMethodPost RestMethod = "POST"
)

func (e RestMethod) ToPointer() *RestMethod {
	return &e
}

func (e *RestMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "GET":
		fallthrough
	case "POST":
		*e = RestMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RestMethod: %v", v)
	}
}

// UserDefinedAPIEntity - Extends [Entity](entity-schema.v1.json) with properties required by the connector for extracting all its data, such as a paging strategy.
type UserDefinedAPIEntity struct {
	// The unique identifier of the entity.
	ID string `json:"id"`
	// The columns of the entity.
	Columns []Column `json:"columns"`
	// The name of the entity.
	DisplayName string `json:"displayName"`
	// The base URL to fetch data. Note that query parameters for things like pagination and sorting will be appended.
	APIURL string `json:"apiUrl"`
	// The [JMESPath](https://jmespath.org/) expression that converts the API response into an array containing one JSON object per record.
	PathToResults string `json:"pathToResults"`
	// The pipeline mode.
	PipelineMode UserDefinedAPIPipelineMode `json:"pipelineMode"`
	// The paging strategy.
	PagingStrategy *PagingStrategy `json:"pagingStrategy,omitempty"`
	// A list of query parameters to be passed with all the requests.
	QueryParameters []QueryParameters `json:"queryParameters,omitempty"`
	// A list of headers to be passed with all the requests.
	HeaderParameters []HeaderParameters `json:"headerParameters,omitempty"`
	// The HTTP method used to call the apiUrl.
	RestMethod *RestMethod `default:"GET" json:"restMethod"`
}

func (u UserDefinedAPIEntity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UserDefinedAPIEntity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UserDefinedAPIEntity) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UserDefinedAPIEntity) GetColumns() []Column {
	if o == nil {
		return []Column{}
	}
	return o.Columns
}

func (o *UserDefinedAPIEntity) GetDisplayName() string {
	if o == nil {
		return ""
	}
	return o.DisplayName
}

func (o *UserDefinedAPIEntity) GetAPIURL() string {
	if o == nil {
		return ""
	}
	return o.APIURL
}

func (o *UserDefinedAPIEntity) GetPathToResults() string {
	if o == nil {
		return ""
	}
	return o.PathToResults
}

func (o *UserDefinedAPIEntity) GetPipelineMode() UserDefinedAPIPipelineMode {
	if o == nil {
		return UserDefinedAPIPipelineMode{}
	}
	return o.PipelineMode
}

func (o *UserDefinedAPIEntity) GetPagingStrategy() *PagingStrategy {
	if o == nil {
		return nil
	}
	return o.PagingStrategy
}

func (o *UserDefinedAPIEntity) GetPagingStrategyCursorURI() *CursorURIPagingStrategy {
	if v := o.GetPagingStrategy(); v != nil {
		return v.CursorURIPagingStrategy
	}
	return nil
}

func (o *UserDefinedAPIEntity) GetPagingStrategyOffset() *OffsetPagingStrategy {
	if v := o.GetPagingStrategy(); v != nil {
		return v.OffsetPagingStrategy
	}
	return nil
}

func (o *UserDefinedAPIEntity) GetQueryParameters() []QueryParameters {
	if o == nil {
		return nil
	}
	return o.QueryParameters
}

func (o *UserDefinedAPIEntity) GetHeaderParameters() []HeaderParameters {
	if o == nil {
		return nil
	}
	return o.HeaderParameters
}

func (o *UserDefinedAPIEntity) GetRestMethod() *RestMethod {
	if o == nil {
		return nil
	}
	return o.RestMethod
}
