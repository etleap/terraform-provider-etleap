resource "etleap_connection_freshcaller" "my_connectionfreshcaller" {
  api_key                     = "...my_api_key..."
  deletion_of_export_products = true
  domain                      = "...my_domain..."
  name                        = "Bonnie Lockman IV"
  type                        = "FRESHCALLER"
}