resource "etleap_connection_redshift" "my_connectionredshift" {
  address                       = "8577 Dallin Center"
  database                      = "...my_database..."
  deletion_of_export_products   = true
  dynamic_varchar_width_enabled = false
  name                          = "Ms. Nancy McDermott"
  password                      = "...my_password..."
  port                          = 9
  query_tags_enabled            = false
  schema                        = "...my_schema..."
  source_only                   = false
  type                          = "REDSHIFT"
  username                      = "Dayana.Fadel"
}