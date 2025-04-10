resource "etleap_connection_kustomer" "my_connectionkustomer" {
  api_key                     = "...my_api_key..."
  deletion_of_export_products = true
  name                        = "Ms. Mindy Price"
  type                        = "KUSTOMER"
}