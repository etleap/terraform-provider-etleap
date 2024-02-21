---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "etleap_connection_blackline Data Source - terraform-provider-etleap"
subcategory: ""
description: |-
  ConnectionBLACKLINE DataSource
---

# etleap_connection_blackline (Data Source)

ConnectionBLACKLINE DataSource

## Example Usage

```terraform
data "etleap_connection_blackline" "my_connectionblackline" {
  id = "e895185d-4b10-41a0-b642-ba38bc737d2b"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `active` (Boolean) Whether this connection has been marked as active.
- `api_key` (String) The Blackline API Key generated for your user
- `base_url` (String) Your Blackline instance base URL, i.e, https://<BASE_URL>.api.blackline.com
- `client_id` (String) Your Blackline instance Client ID
- `client_secret` (String) Your Blackline instance Client Secret
- `create_date` (String) The date and time when then the connection was created.
- `default_update_schedule` (Attributes List) When an update schedule is not defined for a connection, the default schedule is used. The default defined individually per `pipelineMode` and may be subject to change. (see [below for nested schema](#nestedatt--default_update_schedule))
- `id` (String) The ID of this resource.
- `instance_scope` (String) Your Blackline instance authorization scope.
- `name` (String) The unique name of this connection.
- `status` (String) The current status of the connection. must be one of ["UNKNOWN", "UP", "DOWN", "RESIZE", "MAINTENANCE", "QUOTA", "CREATING"]
- `type` (String) must be one of ["BLACKLINE"]
- `update_schedule` (Attributes) The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection. (see [below for nested schema](#nestedatt--update_schedule))
- `username` (String) Your Blackline username

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

