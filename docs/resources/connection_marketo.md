---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "etleap_connection_marketo Resource - terraform-provider-etleap"
subcategory: ""
description: |-
  ConnectionMARKETO Resource
---

# etleap_connection_marketo (Resource)

ConnectionMARKETO Resource

## Example Usage

```terraform
resource "etleap_connection_marketo" "my_connectionmarketo" {
  deletion_of_export_products = false
  name                        = "Brittany Greenfelder"
  quota_limit                 = 9
  rest_client_id              = "...my_rest_client_id..."
  rest_client_secret          = "...my_rest_client_secret..."
  rest_endpoint               = "...my_rest_endpoint..."
  soap_encryption_key         = "...my_soap_encryption_key..."
  soap_endpoint               = "...my_soap_endpoint..."
  soap_user_id                = "...my_soap_user_id..."
  type                        = "MARKETO"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The unique name of this connection.
- `quota_limit` (Number) The maximum number of requests Etleap will use per day. Your Marketo account's daily quota is the 'Daily Request Limit' number under Admin -> Integration -> Web Services. We recommend setting this to 75% of your Marketo limit.
- `rest_client_id` (String) Under Admin -> Integration -> LaunchPoint, you can find this value by clicking 'View Details'.
- `rest_client_secret` (String) Under Admin -> Integration -> LaunchPoint, you can find this value by clicking 'View Details'.
- `rest_endpoint` (String) E.g. 'https://259-ZDK-675.mktorest.com/rest'. In the Marketo UI this is the 'Endpoint' value in the 'REST API' section.
- `soap_encryption_key` (String) In the Marketo UI this is the 'Encryption Key' value in the 'SOAP API' section.
- `soap_endpoint` (String) E.g. 'https://259-ZDK-675.mktoapi.com/soap/mktows/2_9'. In the Marketo UI this is the 'Endpoint' value in the 'SOAP API' section.
- `soap_user_id` (String) E.g. 'MKTOWS_259-ZDK-675_1'. In the Marketo UI this is the 'User ID' value in the 'SOAP API' section.
- `type` (String) must be one of ["MARKETO"]

### Optional

- `deletion_of_export_products` (Boolean) Applicable for REDSHIFT and SNOWFLAKE connections only in the case when there are pipelines that use this connection as a destination, and these pipelines have been migrated to use a different destination. Specifies whether any tables created by these pipelines in this destination should be deleted. Defaults to `false`. Default: false
- `update_schedule` (Attributes) The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. When undefined, the pipeline will default to the schedule set on the source connection. (see [below for nested schema](#nestedatt--update_schedule))

### Read-Only

- `active` (Boolean) Whether this connection has been marked as active.
- `create_date` (String) The date and time when then the connection was created.
- `default_update_schedule` (Attributes List) When an update schedule is not defined for a connection, the default schedule is used. The default defined individually per `pipelineMode` and may be subject to change. (see [below for nested schema](#nestedatt--default_update_schedule))
- `id` (String) The unique identifier of the connection.
- `status` (String) The current status of the connection. must be one of ["UNKNOWN", "UP", "DOWN", "RESIZE", "MAINTENANCE", "QUOTA", "CREATING"]

<a id="nestedatt--update_schedule"></a>
### Nested Schema for `update_schedule`

Optional:

- `daily` (Attributes) The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. (see [below for nested schema](#nestedatt--update_schedule--daily))
- `hourly` (Attributes) The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. (see [below for nested schema](#nestedatt--update_schedule--hourly))
- `interval` (Attributes) The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. (see [below for nested schema](#nestedatt--update_schedule--interval))
- `monthly` (Attributes) The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. (see [below for nested schema](#nestedatt--update_schedule--monthly))
- `weekly` (Attributes) The update schedule defines when Etleap should automatically check the source for new data. See <a href= "https://support.etleap.com/hc/en-us/articles/360019768853-What-is-the-difference-between-a-Refresh-and-an-Update-" target="_blank" rel="noopener">Updates &amp; Refreshes</a> for more information. (see [below for nested schema](#nestedatt--update_schedule--weekly))

<a id="nestedatt--update_schedule--daily"></a>
### Nested Schema for `update_schedule.daily`

Optional:

- `hour_of_day` (Number) Hour of day the  pipeline update should be started at (in UTC). Not Null
- `mode` (String) Not Null; must be one of ["DAILY"]


<a id="nestedatt--update_schedule--hourly"></a>
### Nested Schema for `update_schedule.hourly`

Optional:

- `mode` (String) Not Null; must be one of ["HOURLY"]


<a id="nestedatt--update_schedule--interval"></a>
### Nested Schema for `update_schedule.interval`

Optional:

- `interval_minutes` (Number) Time to wait before new data is pulled (in minutes). Not Null
- `mode` (String) Not Null; must be one of ["INTERVAL"]


<a id="nestedatt--update_schedule--monthly"></a>
### Nested Schema for `update_schedule.monthly`

Optional:

- `day_of_month` (Number) Not Null
- `hour_of_day` (Number) Hour of day the  pipeline update should be started at (in UTC). Not Null
- `mode` (String) Not Null; must be one of ["MONTHLY"]


<a id="nestedatt--update_schedule--weekly"></a>
### Nested Schema for `update_schedule.weekly`

Optional:

- `day_of_week` (Number) Not Null
- `hour_of_day` (Number) Hour of day the  pipeline update should be started at (in UTC). Not Null
- `mode` (String) Not Null; must be one of ["WEEKLY"]



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

