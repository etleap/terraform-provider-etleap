resource "etleap_connection_redshift" "my_connectionredshift" {
  address                       = "370 Towne Extensions"
  database                      = "...my_database..."
  deletion_of_export_products   = true
  dynamic_varchar_width_enabled = true
  name                          = "Danielle Hintz"
  password                      = "...my_password..."
  port                          = 7
  query_tags_enabled            = false
  schema                        = "...my_schema..."
  source_only                   = false
  type                          = "REDSHIFT"
  username                      = "Mozell.Sporer27"
}