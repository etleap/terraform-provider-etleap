resource "etleap_connection_postgres" "my_connectionpostgres" {
  address                     = "6424 Gottlieb Well"
  auto_replicate              = "...my_auto_replicate..."
  cdc_enabled                 = true
  database                    = "...my_database..."
  deletion_of_export_products = true
  name                        = "Darla Schimmel"
  password                    = "...my_password..."
  port                        = 6
  schema                      = "...my_schema..."
  type                        = "POSTGRES"
  username                    = "Ines2"
}