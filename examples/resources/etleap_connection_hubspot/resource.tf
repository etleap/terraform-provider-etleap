resource "etleap_connection_hubspot" "my_connectionhubspot" {
  code                        = "...my_code..."
  deletion_of_export_products = true
  name                        = "Johnny Ebert"
  type                        = "HUBSPOT"
}