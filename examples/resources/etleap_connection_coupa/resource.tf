resource "etleap_connection_coupa" "my_connectioncoupa" {
  client_id                   = "...my_client_id..."
  client_secret               = "...my_client_secret..."
  deletion_of_export_products = true
  name                        = "Courtney Block"
  subdomain                   = "...my_subdomain..."
  type                        = "COUPA"
}