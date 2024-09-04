resource "etleap_connection_redshift" "my_connectionredshift" {
  address                       = "8128 Vincenzo Divide"
  database                      = "...my_database..."
  deletion_of_export_products   = false
  dynamic_varchar_width_enabled = true
  name                          = "Levi Wolff"
  password                      = "...my_password..."
  port                          = 1
  query_tags_enabled            = false
  schema                        = "...my_schema..."
  source_only                   = false
  type                          = "REDSHIFT"
  username                      = "Delia.Becker"
}