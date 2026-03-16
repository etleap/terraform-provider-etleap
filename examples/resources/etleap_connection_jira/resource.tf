resource "etleap_connection_jira" "my_connectionjira" {
  base_url                    = "...my_base_url..."
  code                        = "...my_code..."
  deletion_of_export_products = false
  name                        = "Marian Schmeler"
  type                        = "JIRA"
}