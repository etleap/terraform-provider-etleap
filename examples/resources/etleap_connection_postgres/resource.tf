resource "etleap_connection_postgres" "my_connectionpostgres" {
  address                     = "484 Sterling Curve"
  auto_replicate              = "...my_auto_replicate..."
  cdc_enabled                 = true
  database                    = "...my_database..."
  deletion_of_export_products = false
  fetch_lobs_for_updated_rows = false
  name                        = "Walter Walker"
  password                    = "...my_password..."
  port                        = 8
  schema                      = "...my_schema..."
  type                        = "POSTGRES"
  username                    = "Adell.West89"
}