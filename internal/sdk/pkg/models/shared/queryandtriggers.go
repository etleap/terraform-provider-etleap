// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type QueryAndTriggers struct {
	// The SQL query used to build this model. To specify dependencies on pipelines or other models, replace the schema and table name of the dependency with the id of the dependency enclosed in `{{` and `}}`. The dependency must load data into the same Etleap connection as the one given in `warehouse.connectionId` for this model.
	//
	// **For Example**
	// Say there is a pipeline with the id `abcd1234` which loads data to the table "schema"."my_table". To create a model in Etleap that has a dependency on this pipeline, the following query:
	//
	// ```sql
	// SELECT col1, col2 FROM "schema"."my_table";
	// ```
	//
	// becomes:
	// ```sql
	// SELECT col1, col2 FROM {{abcd1234}};
	// ```
	//
	// [See the Model documentation](https://docs.etleap.com/docs/documentation/ZG9jOjI0MzU2NDY3-introduction-to-models#model-dependencies) for more information on Model dependencies.
	Query string `json:"query"`
	// A list of model dependency ids. An update will be automatically triggered in this model if any of the dependencies listed here get new data. Any ids given here must be present as dependencies in the `query`.
	Triggers []string `json:"triggers"`
}

func (o *QueryAndTriggers) GetQuery() string {
	if o == nil {
		return ""
	}
	return o.Query
}

func (o *QueryAndTriggers) GetTriggers() []string {
	if o == nil {
		return []string{}
	}
	return o.Triggers
}
