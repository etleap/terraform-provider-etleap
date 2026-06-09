resource "etleap_connection_gitlab" "my_connectiongitlab" {
  access_token                = "...my_access_token..."
  api_url                     = "...my_api_url..."
  deletion_of_export_products = true
  name                        = "Miss Lamar Hickle Jr."
  type                        = "GITLAB"
}