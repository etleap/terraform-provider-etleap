resource "etleap_connection_egnyte" "my_connectionegnyte" {
  base_directory              = "...my_base_directory..."
  code                        = "...my_code..."
  deletion_of_export_products = true
  domain_name                 = "...my_domain_name..."
  name                        = "Lorenzo Jacobson"
  type                        = "EGNYTE"
}