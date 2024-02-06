resource "etleap_connection_impact_radius" "my_connectionimpact_radius" {
  account_sid                 = "...my_account_sid..."
  auth_token                  = "...my_auth_token..."
  deletion_of_export_products = true
  name                        = "Sergio Daugherty II"
  type                        = "IMPACT_RADIUS"
}