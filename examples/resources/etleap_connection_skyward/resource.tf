resource "etleap_connection_skyward" "my_connectionskyward" {
  client_id                   = "...my_client_id..."
  client_secret               = "...my_client_secret..."
  deletion_of_export_products = true
  name                        = "Lola Fay"
  type                        = "SKYWARD"
}