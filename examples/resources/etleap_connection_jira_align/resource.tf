resource "etleap_connection_jira_align" "my_connectionjira_align" {
  api_key                     = "...my_api_key..."
  deletion_of_export_products = true
  name                        = "Miss Tricia Borer"
  subdomain                   = "...my_subdomain..."
  type                        = "JIRA_ALIGN"
}