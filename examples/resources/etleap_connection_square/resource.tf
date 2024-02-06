resource "etleap_connection_square" "my_connectionsquare" {
  application_id              = "...my_application_id..."
  application_secret          = "...my_application_secret..."
  code                        = "...my_code..."
  deletion_of_export_products = false
  name                        = "Matt Bogan"
  sandbox_account             = false
  type                        = "SQUARE"
}