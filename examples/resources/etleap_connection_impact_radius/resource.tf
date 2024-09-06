resource "etleap_connection_impact_radius" "my_connectionimpact_radius" {
  account_sid                 = "...my_account_sid..."
  auth_token                  = "...my_auth_token..."
  deletion_of_export_products = false
  name                        = "Arnold Turcotte"
  type                        = "IMPACT_RADIUS"
}