resource "etleap_connection_workfront" "my_connectionworkfront" {
  api_key                     = "...my_api_key..."
  deletion_of_export_products = false
  name                        = "Anne Champlin"
  subdomain                   = "...my_subdomain..."
  type                        = "WORKFRONT"
}