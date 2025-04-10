resource "etleap_connection_elluminate" "my_connectionelluminate" {
  api_key                     = "...my_api_key..."
  api_secret                  = "...my_api_secret..."
  base_url                    = "...my_base_url..."
  deletion_of_export_products = false
  name                        = "Erma Lakin DVM"
  type                        = "ELLUMINATE"
}