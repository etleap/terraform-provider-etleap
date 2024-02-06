resource "etleap_connection_redshift" "my_connectionredshift" {
  address                       = "1579 Trantow Dale"
  database                      = "...my_database..."
  deletion_of_export_products   = false
  dynamic_varchar_width_enabled = true
  name                          = "Lucia Dickens"
  password                      = "...my_password..."
  port                          = 7
  query_tags_enabled            = true
  schema                        = "...my_schema..."
  source_only                   = false
  type                          = "REDSHIFT"
  username                      = "Dallas77"
}