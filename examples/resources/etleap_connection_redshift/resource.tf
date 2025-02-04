resource "etleap_connection_redshift" "my_connectionredshift" {
  address                       = "848 Emanuel Underpass"
  database                      = "...my_database..."
  deletion_of_export_products   = true
  dynamic_varchar_width_enabled = true
  name                          = "Dr. Joyce Stehr"
  password                      = "...my_password..."
  port                          = 3
  query_tags_enabled            = true
  schema                        = "...my_schema..."
  source_only                   = false
  type                          = "REDSHIFT"
  username                      = "Tyrique89"
}