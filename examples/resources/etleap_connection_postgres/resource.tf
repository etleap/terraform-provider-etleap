resource "etleap_connection_postgres" "my_connectionpostgres" {
  address                     = "794 Lindsey Well"
  auto_replicate              = "...my_auto_replicate..."
  cdc_enabled                 = true
  database                    = "...my_database..."
  deletion_of_export_products = false
  name                        = "Emmett Bernhard"
  password                    = "...my_password..."
  port                        = 9
  schema                      = "...my_schema..."
  type                        = "POSTGRES"
  username                    = "Kyle34"
}