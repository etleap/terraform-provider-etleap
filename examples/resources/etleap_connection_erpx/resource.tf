resource "etleap_connection_erpx" "my_connectionerpx" {
  api_url                     = "...my_api_url..."
  client_id                   = "...my_client_id..."
  client_secret               = "...my_client_secret..."
  deletion_of_export_products = false
  name                        = "Guy Schroeder"
  token_url                   = "...my_token_url..."
  type                        = "ERPX"
}