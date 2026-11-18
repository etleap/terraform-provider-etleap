resource "etleap_connection_redshift" "my_connectionredshift" {
  address                       = "06310 Bartell Key"
  database                      = "...my_database..."
  deletion_of_export_products   = false
  dynamic_varchar_width_enabled = true
  name                          = "Joseph Fisher"
  password                      = "...my_password..."
  port                          = 4
  query_tags_enabled            = true
  schema                        = "...my_schema..."
  source_only                   = true
  type                          = "REDSHIFT"
  username                      = "Virginia_Kris"
}