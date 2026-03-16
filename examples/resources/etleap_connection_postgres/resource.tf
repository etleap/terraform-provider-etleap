resource "etleap_connection_postgres" "my_connectionpostgres" {
  address                     = "220 Frami Fields"
  auto_replicate              = "...my_auto_replicate..."
  cdc_enabled                 = true
  database                    = "...my_database..."
  deletion_of_export_products = true
  fetch_lobs_for_updated_rows = true
  name                        = "Martin Nitzsche"
  password                    = "...my_password..."
  port                        = 8
  schema                      = "...my_schema..."
  type                        = "POSTGRES"
  username                    = "Floy.Weimann83"
}