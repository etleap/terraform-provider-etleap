resource "etleap_connection_jira_cloud" "my_connectionjira_cloud" {
  code                        = "...my_code..."
  deletion_of_export_products = true
  name                        = "Pat Bogan"
  type                        = "JIRA_CLOUD"
}