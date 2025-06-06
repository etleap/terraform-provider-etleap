---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "etleap_connection_redshift_sharded Data Source - terraform-provider-etleap"
subcategory: ""
description: |-
  ConnectionREDSHIFTSHARDED DataSource
---

# etleap_connection_redshift_sharded (Data Source)

ConnectionREDSHIFTSHARDED DataSource

## Example Usage

```terraform
data "etleap_connection_redshift_sharded" "my_connectionredshift_sharded" {
  id = "18c08e44-8ddb-4223-b4a7-2481b6269371"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `active` (Boolean) Whether this connection has been marked as active.
- `create_date` (String) The date and time when then the connection was created.
- `data_sharing_destinations` (List of String) The id of another Etleap Redshift connection. If specified, Etleap will make the data loaded available to the other cluster via Redshift Data Sharing.
- `default_update_schedule` (Attributes List) When an update schedule is not defined for a connection, the default schedule is used. The default defined individually per `pipelineMode` and may be subject to change. (see [below for nested schema](#nestedatt--default_update_schedule))
- `dynamic_varchar_width_enabled` (Boolean) Etleap will create VARCHAR columns with the minimal required width based on the data it's loading, and expand the column width as required. This can improve performance but there are <a target="_blank" href="https://docs.etleap.com/docs/documentation/ba7744fcf6114-redshift-optional-connection-settings#enable-dynamic-varchar-widths">some limitations</a>. Note: if set to `true`, it can't later be updated to `false`.
- `id` (String) The ID of this resource.
- `name` (String) The unique name of this connection.
- `query_tags_enabled` (Boolean) Should Etleap prefix each load query with metadata? More info can be found <a href="https://docs.etleap.com/docs/documentation/ba7744fcf6114-redshift-optional-connection-settings#include-query-tags">here</a>.
- `schema` (String) If not specified, the default schema will be used.
- `shards` (Attributes List) (see [below for nested schema](#nestedatt--shards))
- `source_only` (Boolean) Are you going to use this connection only as a source for pipelines? When `true`, this connection will only be available as an ETL source only, and Etleap will skip the creation of an audit table in the database.
- `status` (String) The current status of the connection. must be one of ["UNKNOWN", "UP", "DOWN", "RESIZE", "MAINTENANCE", "QUOTA", "CREATING"]
- `type` (String) must be one of ["REDSHIFT_SHARDED"]
- `update_schedule` (Attributes) The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection. (see [below for nested schema](#nestedatt--update_schedule))
- `user_groups` (List of String) When Etleap creates Redshift tables, SELECT privileges will be granted to user groups specified here.

<a id="nestedatt--default_update_schedule"></a>
### Nested Schema for `default_update_schedule`

Read-Only:

- `pipeline_mode` (String) The pipeline mode refers to how the pipeline fetches data changes from the source and how those changes are applied to the destination table. See <a target="_blank" href="https://docs.etleap.com/docs/documentation/ZG9jOjIyMjE3ODA2-introduction">the documentation</a> for more details. must be one of ["APPEND", "REPLACE", "UPDATE", "QUERY"]
- `update_schedule` (Attributes) The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection. (see [below for nested schema](#nestedatt--default_update_schedule--update_schedule))

<a id="nestedatt--default_update_schedule--update_schedule"></a>
### Nested Schema for `default_update_schedule.update_schedule`

Read-Only:

- `daily` (Attributes) The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. (see [below for nested schema](#nestedatt--default_update_schedule--update_schedule--daily))
- `hourly` (Attributes) The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. (see [below for nested schema](#nestedatt--default_update_schedule--update_schedule--hourly))
- `interval` (Attributes) The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. (see [below for nested schema](#nestedatt--default_update_schedule--update_schedule--interval))
- `monthly` (Attributes) The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. (see [below for nested schema](#nestedatt--default_update_schedule--update_schedule--monthly))
- `weekly` (Attributes) The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. (see [below for nested schema](#nestedatt--default_update_schedule--update_schedule--weekly))

<a id="nestedatt--default_update_schedule--update_schedule--daily"></a>
### Nested Schema for `default_update_schedule.update_schedule.daily`

Read-Only:

- `hour_of_day` (Number) Hour of day the  pipeline update should be started at (in UTC).
- `mode` (String) must be one of ["DAILY"]


<a id="nestedatt--default_update_schedule--update_schedule--hourly"></a>
### Nested Schema for `default_update_schedule.update_schedule.hourly`

Read-Only:

- `mode` (String) must be one of ["HOURLY"]


<a id="nestedatt--default_update_schedule--update_schedule--interval"></a>
### Nested Schema for `default_update_schedule.update_schedule.interval`

Read-Only:

- `interval_minutes` (Number) Time to wait before new data is pulled (in minutes).
- `mode` (String) must be one of ["INTERVAL"]


<a id="nestedatt--default_update_schedule--update_schedule--monthly"></a>
### Nested Schema for `default_update_schedule.update_schedule.monthly`

Read-Only:

- `day_of_month` (Number)
- `hour_of_day` (Number) Hour of day the  pipeline update should be started at (in UTC).
- `mode` (String) must be one of ["MONTHLY"]


<a id="nestedatt--default_update_schedule--update_schedule--weekly"></a>
### Nested Schema for `default_update_schedule.update_schedule.weekly`

Read-Only:

- `day_of_week` (Number)
- `hour_of_day` (Number) Hour of day the  pipeline update should be started at (in UTC).
- `mode` (String) must be one of ["WEEKLY"]




<a id="nestedatt--shards"></a>
### Nested Schema for `shards`

Read-Only:

- `address` (String)
- `database` (String)
- `port` (Number)
- `shard_id` (String)
- `ssh_config` (Attributes) (see [below for nested schema](#nestedatt--shards--ssh_config))
- `username` (String)

<a id="nestedatt--shards--ssh_config"></a>
### Nested Schema for `shards.ssh_config`

Read-Only:

- `address` (String) The server address for the SSH connection.
- `username` (String) The username for the SSH connection.



<a id="nestedatt--update_schedule"></a>
### Nested Schema for `update_schedule`

Read-Only:

- `daily` (Attributes) The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. (see [below for nested schema](#nestedatt--update_schedule--daily))
- `hourly` (Attributes) The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. (see [below for nested schema](#nestedatt--update_schedule--hourly))
- `interval` (Attributes) The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. (see [below for nested schema](#nestedatt--update_schedule--interval))
- `monthly` (Attributes) The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. (see [below for nested schema](#nestedatt--update_schedule--monthly))
- `weekly` (Attributes) The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. (see [below for nested schema](#nestedatt--update_schedule--weekly))

<a id="nestedatt--update_schedule--daily"></a>
### Nested Schema for `update_schedule.daily`

Read-Only:

- `hour_of_day` (Number) Hour of day the  pipeline update should be started at (in UTC).
- `mode` (String) must be one of ["DAILY"]


<a id="nestedatt--update_schedule--hourly"></a>
### Nested Schema for `update_schedule.hourly`

Read-Only:

- `mode` (String) must be one of ["HOURLY"]


<a id="nestedatt--update_schedule--interval"></a>
### Nested Schema for `update_schedule.interval`

Read-Only:

- `interval_minutes` (Number) Time to wait before new data is pulled (in minutes).
- `mode` (String) must be one of ["INTERVAL"]


<a id="nestedatt--update_schedule--monthly"></a>
### Nested Schema for `update_schedule.monthly`

Read-Only:

- `day_of_month` (Number)
- `hour_of_day` (Number) Hour of day the  pipeline update should be started at (in UTC).
- `mode` (String) must be one of ["MONTHLY"]


<a id="nestedatt--update_schedule--weekly"></a>
### Nested Schema for `update_schedule.weekly`

Read-Only:

- `day_of_week` (Number)
- `hour_of_day` (Number) Hour of day the  pipeline update should be started at (in UTC).
- `mode` (String) must be one of ["WEEKLY"]


