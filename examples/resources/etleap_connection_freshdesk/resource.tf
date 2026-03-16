resource "etleap_connection_freshdesk" "my_connectionfreshdesk" {
  api_key                     = "...my_api_key..."
  deletion_of_export_products = true
  domain                      = "...my_domain..."
  name                        = "Miss Sylvester Lockman DVM"
  type                        = "FRESHDESK"
}