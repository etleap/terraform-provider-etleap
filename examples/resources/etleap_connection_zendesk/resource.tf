resource "etleap_connection_zendesk" "my_connectionzendesk" {
  client_id                   = "...my_client_id..."
  client_secret               = "...my_client_secret..."
  deletion_of_export_products = true
  name                        = "Rosalie Braun"
  subdomain                   = "...my_subdomain..."
  type                        = "ZENDESK"
}