resource "etleap_connection_redshift" "my_connectionredshift" {
  address                       = "745 Veum Motorway"
  database                      = "...my_database..."
  deletion_of_export_products   = false
  dynamic_varchar_width_enabled = false
  name                          = "Judy Nader"
  password                      = "...my_password..."
  port                          = 6
  query_tags_enabled            = false
  schema                        = "...my_schema..."
  source_only                   = false
  type                          = "REDSHIFT"
  username                      = "Fay64"
}