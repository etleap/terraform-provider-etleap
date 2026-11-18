resource "etleap_connection_postgres" "my_connectionpostgres" {
  address                     = "70982 Denesik Glen"
  auto_replicate              = "...my_auto_replicate..."
  cdc_address                 = "...my_cdc_address..."
  cdc_enabled                 = true
  cdc_port                    = 4
  database                    = "...my_database..."
  deletion_of_export_products = false
  fetch_lobs_for_updated_rows = false
  name                        = "Carlton Pfeffer"
  password                    = "...my_password..."
  port                        = 8
  schema                      = "...my_schema..."
  type                        = "POSTGRES"
  username                    = "Jaycee3"
}