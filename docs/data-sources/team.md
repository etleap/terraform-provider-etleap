---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "etleap_team Data Source - terraform-provider-etleap"
subcategory: ""
description: |-
  Team DataSource
---

# etleap_team (Data Source)

Team DataSource

## Example Usage

```terraform
data "etleap_team" "my_team" {
  id = "290e85d3-f67a-4bc4-841b-16fece898858"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `create_date` (String) The date and time when then the team was created.
- `description` (String)
- `id` (String) The ID of this resource.
- `members` (Attributes List) (see [below for nested schema](#nestedatt--members))
- `name` (String)

<a id="nestedatt--members"></a>
### Nested Schema for `members`

Read-Only:

- `email_address` (String)
- `first_name` (String)
- `id` (String)
- `last_name` (String)


