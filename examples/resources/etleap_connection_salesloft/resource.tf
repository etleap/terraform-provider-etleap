resource "etleap_connection_salesloft" "my_connectionsalesloft" {
  api_key                     = "...my_api_key..."
  deletion_of_export_products = false
  name                        = "Austin Witting"
  type                        = "SALESLOFT"
}