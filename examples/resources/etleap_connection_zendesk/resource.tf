resource "etleap_connection_zendesk" "my_connectionzendesk" {
  code                        = "...my_code..."
  deletion_of_export_products = true
  name                        = "Ms. Kurt Schaefer"
  subdomain                   = "...my_subdomain..."
  type                        = "ZENDESK"
}