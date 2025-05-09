resource "etleap_connection_stripe" "my_connectionstripe" {
  api_secret_key              = "...my_api_secret_key..."
  deletion_of_export_products = false
  name                        = "Mr. Elvira Douglas"
  type                        = "STRIPE"
}