---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "etleap_model Resource - terraform-provider-etleap"
subcategory: ""
description: |-
  Model Resource
---

# etleap_model (Resource)

Model Resource

## Example Usage

```terraform
resource "etleap_model" "my_model" {
  deletion_of_export_products = true
  name                        = "Aaron Fay"
  query_and_triggers = {
    query = "...my_query..."
    triggers = [
      "...",
    ]
  }
  update_schedule = {
    daily = {
      hour_of_day = 4
      mode        = "DAILY"
    }
  }
  warehouse = {
    redshift = {
      connection_id = "...my_connection_id..."
      distribution_style = {
        one = "AUTO"
      }
      materialized_view     = false
      pending_renamed_table = "...my_pending_renamed_table..."
      schema                = "...my_schema..."
      sort_columns = [
        "...",
      ]
      table                       = "...my_table..."
      type                        = "REDSHIFT"
      wait_for_update_preparation = false
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)
- `query_and_triggers` (Attributes) (see [below for nested schema](#nestedatt--query_and_triggers))
- `update_schedule` (Attributes) How often this model should update. Etleap will periodically update the model table in your warehouse according to this schedule. See [the Model Updates documentation](https://docs.etleap.com/docs/documentation/ZG9jOjI0MzU2NDY3-introduction-to-models#model-updates) for more information. (see [below for nested schema](#nestedatt--update_schedule))
- `warehouse` (Attributes) (see [below for nested schema](#nestedatt--warehouse))

### Optional

- `deletion_of_export_products` (Boolean) Specifies whether the table in the destination created by this model should be deleted. Defaults to `false`. Default: false
- `shares` (List of String) An array of users' emails to share the model with. Once shared, a model cannot be unshared, and future calls to `PATCH` can only add to this list.

### Read-Only

- `create_date` (String) The date and time when then the model was created.
- `dependencies` (Attributes List) (see [below for nested schema](#nestedatt--dependencies))
- `id` (String) The unique identifier of the model.
- `last_update_duration` (Number) How long the latest update took to complete, in milliseconds, or the duration of the current update if one is in progress.
- `last_update_time` (String) The date and time of the latest successful update for this model.
- `owner` (Attributes) (see [below for nested schema](#nestedatt--owner))
- `paused` (Boolean)

<a id="nestedatt--query_and_triggers"></a>
### Nested Schema for `query_and_triggers`

Required:

- `query` (String) The SQL query used to build this model. To specify dependencies on pipelines or other models, replace the schema and table name of the dependency with the id of the dependency enclosed in `{{` and `}}`. The dependency must load data into the same Etleap connection as the one given in `warehouse.connectionId` for this model.

**For Example**
Say there is a pipeline with the id `abcd1234` which loads data to the table "schema"."my_table". To create a model in Etleap that has a dependency on this pipeline, the following query:

```sql
SELECT col1, col2 FROM "schema"."my_table";
```

becomes:
```sql
SELECT col1, col2 FROM {{abcd1234}};
```

[See the Model documentation](https://docs.etleap.com/docs/documentation/ZG9jOjI0MzU2NDY3-introduction-to-models#model-dependencies) for more information on Model dependencies.
- `triggers` (List of String) A list of model dependency ids. An update will be automatically triggered in this model if any of the dependencies listed here get new data. Any ids given here must be present as dependencies in the `query`.


<a id="nestedatt--update_schedule"></a>
### Nested Schema for `update_schedule`

Optional:

- `daily` (Attributes) (see [below for nested schema](#nestedatt--update_schedule--daily))
- `hourly` (Attributes) (see [below for nested schema](#nestedatt--update_schedule--hourly))
- `monthly` (Attributes) (see [below for nested schema](#nestedatt--update_schedule--monthly))
- `never` (Attributes) (see [below for nested schema](#nestedatt--update_schedule--never))
- `weekly` (Attributes) (see [below for nested schema](#nestedatt--update_schedule--weekly))

<a id="nestedatt--update_schedule--daily"></a>
### Nested Schema for `update_schedule.daily`

Optional:

- `hour_of_day` (Number) Hour of day this schedule should trigger at (in UTC). Not Null
- `mode` (String) Not Null; must be one of ["DAILY"]


<a id="nestedatt--update_schedule--hourly"></a>
### Nested Schema for `update_schedule.hourly`

Optional:

- `mode` (String) Not Null; must be one of ["HOURLY"]


<a id="nestedatt--update_schedule--monthly"></a>
### Nested Schema for `update_schedule.monthly`

Optional:

- `day_of_month` (Number) Day of the month this schedule should trigger at (in UTC). Not Null
- `hour_of_day` (Number) Hour of day this schedule should trigger at (in UTC). Not Null
- `mode` (String) Not Null; must be one of ["MONTHLY"]


<a id="nestedatt--update_schedule--never"></a>
### Nested Schema for `update_schedule.never`

Optional:

- `mode` (String) Not Null; must be one of ["NEVER"]


<a id="nestedatt--update_schedule--weekly"></a>
### Nested Schema for `update_schedule.weekly`

Optional:

- `day_of_week` (Number) The day of the week this schedule should trigger at (in UTC). Not Null
- `hour_of_day` (Number) Hour of day this schedule should trigger at (in UTC). Not Null
- `mode` (String) Not Null; must be one of ["WEEKLY"]



<a id="nestedatt--warehouse"></a>
### Nested Schema for `warehouse`

Optional:

- `redshift` (Attributes) (see [below for nested schema](#nestedatt--warehouse--redshift))
- `snowflake` (Attributes) (see [below for nested schema](#nestedatt--warehouse--snowflake))

<a id="nestedatt--warehouse--redshift"></a>
### Nested Schema for `warehouse.redshift`

Optional:

- `connection_id` (String) Requires replacement if changed. ; Not Null
- `distribution_style` (Attributes) Can either be one the strings `ALL`, `AUTO` or `EVEN`, or an object for `KEY` distribution that specifies a column. Not Null (see [below for nested schema](#nestedatt--warehouse--redshift--distribution_style))
- `materialized_view` (Boolean) Requires replacement if changed. ; Not Null
- `schema` (String) Requires replacement if changed.
- `sort_columns` (List of String) The sort columns to use.
- `table` (String) Not Null
- `type` (String) Not Null; must be one of ["REDSHIFT"]
- `wait_for_update_preparation` (Boolean) Requires replacement if changed. ; Not Null

Read-Only:

- `pending_renamed_table` (String) Only set when a table rename was triggered but is not complete yet.

<a id="nestedatt--warehouse--redshift--distribution_style"></a>
### Nested Schema for `warehouse.redshift.distribution_style`

Optional:

- `distribution_style_key` (Attributes) (see [below for nested schema](#nestedatt--warehouse--redshift--distribution_style--distribution_style_key))
- `one` (String) must be one of ["ALL", "AUTO", "EVEN"]

<a id="nestedatt--warehouse--redshift--distribution_style--distribution_style_key"></a>
### Nested Schema for `warehouse.redshift.distribution_style.one`

Optional:

- `column` (String) Not Null
- `type` (String) Not Null; must be one of ["KEY"]




<a id="nestedatt--warehouse--snowflake"></a>
### Nested Schema for `warehouse.snowflake`

Optional:

- `connection_id` (String) Requires replacement if changed. ; Not Null
- `materialized_view` (Boolean) Requires replacement if changed. ; Not Null
- `schema` (String) Requires replacement if changed.
- `table` (String) Not Null
- `type` (String) Not Null; must be one of ["SNOWFLAKE"]
- `wait_for_update_preparation` (Boolean) Requires replacement if changed. ; Not Null

Read-Only:

- `pending_renamed_table` (String) Only set when a table rename was triggered but is not complete yet.



<a id="nestedatt--dependencies"></a>
### Nested Schema for `dependencies`

Read-Only:

- `id` (String) The unique identifier of the pipeline or model.
- `name` (String) The name of the pipeline or model.
- `type` (String) must be one of ["PIPELINE", "MODEL"]


<a id="nestedatt--owner"></a>
### Nested Schema for `owner`

Read-Only:

- `email_address` (String)
- `first_name` (String)
- `id` (String)
- `last_name` (String)


