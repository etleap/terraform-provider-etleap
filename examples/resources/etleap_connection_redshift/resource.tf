resource "etleap_connection_redshift" "my_connectionredshift" {
  address                       = "70982 Denesik Glen"
  database                      = "...my_database..."
  deletion_of_export_products   = true
  dynamic_varchar_width_enabled = true
  name                          = "Jodi Nitzsche"
  password                      = "...my_password..."
  port                          = 8
  query_tags_enabled            = false
  schema                        = "...my_schema..."
  source_only                   = true
  type                          = "REDSHIFT"
  username                      = "Jaycee3"
}