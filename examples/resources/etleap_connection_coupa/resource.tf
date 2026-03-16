resource "etleap_connection_coupa" "my_connectioncoupa" {
  client_id                   = "...my_client_id..."
  client_secret               = "...my_client_secret..."
  deletion_of_export_products = false
  name                        = "Margie Rippin"
  subdomain                   = "...my_subdomain..."
  type                        = "COUPA"
}