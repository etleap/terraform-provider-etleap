resource "etleap_connection_postgres" "my_connectionpostgres" {
  address                     = "91830 Bode Via"
  auto_replicate              = "...my_auto_replicate..."
  cdc_address                 = "...my_cdc_address..."
  cdc_enabled                 = true
  cdc_port                    = 3
  database                    = "...my_database..."
  deletion_of_export_products = true
  fetch_lobs_for_updated_rows = false
  name                        = "Stewart Ernser"
  password                    = "...my_password..."
  port                        = 3
  schema                      = "...my_schema..."
  type                        = "POSTGRES"
  username                    = "Eliane66"
}