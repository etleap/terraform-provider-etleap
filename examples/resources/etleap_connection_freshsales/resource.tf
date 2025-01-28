resource "etleap_connection_freshsales" "my_connectionfreshsales" {
  api_key                     = "...my_api_key..."
  deletion_of_export_products = true
  domain                      = "...my_domain..."
  name                        = "Genevieve Koch DVM"
  quota_limit                 = 40.9
  type                        = "FRESHSALES"
}