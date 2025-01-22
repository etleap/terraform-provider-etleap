resource "etleap_connection_stripe" "my_connectionstripe" {
  api_secret_key              = "...my_api_secret_key..."
  deletion_of_export_products = true
  name                        = "Elaine Nader"
  type                        = "STRIPE"
}