resource "etleap_connection_redshift" "my_connectionredshift" {
  address                       = "87214 Danika Lights"
  database                      = "...my_database..."
  deletion_of_export_products   = true
  dynamic_varchar_width_enabled = true
  name                          = "Lionel Lebsack"
  password                      = "...my_password..."
  port                          = 9
  query_tags_enabled            = true
  schema                        = "...my_schema..."
  source_only                   = true
  type                          = "REDSHIFT"
  username                      = "Earline_Wilkinson42"
}