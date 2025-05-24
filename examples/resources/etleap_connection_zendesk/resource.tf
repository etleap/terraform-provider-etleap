resource "etleap_connection_zendesk" "my_connectionzendesk" {
  code                        = "...my_code..."
  deletion_of_export_products = false
  name                        = "Molly Hammes"
  subdomain                   = "...my_subdomain..."
  type                        = "ZENDESK"
}