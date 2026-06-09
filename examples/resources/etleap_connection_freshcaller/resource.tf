resource "etleap_connection_freshcaller" "my_connectionfreshcaller" {
  api_key                     = "...my_api_key..."
  deletion_of_export_products = false
  domain                      = "...my_domain..."
  name                        = "Dr. Willie Schaefer"
  type                        = "FRESHCALLER"
}