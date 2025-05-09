resource "etleap_connection_postgres" "my_connectionpostgres" {
  address                     = "609 Bernhard Trace"
  auto_replicate              = "...my_auto_replicate..."
  cdc_enabled                 = false
  database                    = "...my_database..."
  deletion_of_export_products = false
  name                        = "Eduardo Harber"
  password                    = "...my_password..."
  port                        = 0
  schema                      = "...my_schema..."
  type                        = "POSTGRES"
  username                    = "Wiley59"
}