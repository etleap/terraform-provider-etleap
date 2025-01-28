resource "etleap_connection_kustomer" "my_connectionkustomer" {
  api_key                     = "...my_api_key..."
  deletion_of_export_products = false
  name                        = "Marty Dooley"
  type                        = "KUSTOMER"
}