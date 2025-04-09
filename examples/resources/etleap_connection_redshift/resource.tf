resource "etleap_connection_redshift" "my_connectionredshift" {
  address                       = "1087 Mia Ranch"
  database                      = "...my_database..."
  deletion_of_export_products   = false
  dynamic_varchar_width_enabled = false
  name                          = "Kristopher Hamill"
  password                      = "...my_password..."
  port                          = 7
  query_tags_enabled            = false
  schema                        = "...my_schema..."
  source_only                   = false
  type                          = "REDSHIFT"
  username                      = "Frederick_Cummings"
}