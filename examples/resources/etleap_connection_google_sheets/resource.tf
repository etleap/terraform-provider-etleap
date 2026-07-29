resource "etleap_connection_google_sheets" "my_connectiongoogle_sheets" {
  authentication_method       = "OAUTH"
  code                        = "...my_code..."
  deletion_of_export_products = false
  name                        = "Victoria Nienow Jr."
  type                        = "GOOGLE_SHEETS"
}