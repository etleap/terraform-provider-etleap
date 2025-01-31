resource "etleap_connection_workfront" "my_connectionworkfront" {
  api_key                     = "...my_api_key..."
  deletion_of_export_products = true
  name                        = "Jessie Renner"
  subdomain                   = "...my_subdomain..."
  type                        = "WORKFRONT"
}