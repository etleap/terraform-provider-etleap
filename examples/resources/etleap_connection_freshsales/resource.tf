resource "etleap_connection_freshsales" "my_connectionfreshsales" {
  api_key                     = "...my_api_key..."
  deletion_of_export_products = false
  domain                      = "...my_domain..."
  name                        = "Arnold Turcotte"
  quota_limit                 = 58.58
  type                        = "FRESHSALES"
}