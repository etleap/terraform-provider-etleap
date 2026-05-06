resource "etleap_connection_redshift" "my_connectionredshift" {
  address                       = "24185 Lebsack Ramp"
  database                      = "...my_database..."
  deletion_of_export_products   = true
  dynamic_varchar_width_enabled = true
  name                          = "Kristin Grady"
  password                      = "...my_password..."
  port                          = 7
  query_tags_enabled            = true
  schema                        = "...my_schema..."
  source_only                   = false
  type                          = "REDSHIFT"
  username                      = "Loy_Veum"
}