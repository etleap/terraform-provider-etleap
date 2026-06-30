resource "etleap_connection_postgres" "my_connectionpostgres" {
  address                     = "308 Theron Ways"
  auto_replicate              = "...my_auto_replicate..."
  cdc_address                 = "...my_cdc_address..."
  cdc_enabled                 = true
  cdc_port                    = 8
  database                    = "...my_database..."
  deletion_of_export_products = false
  fetch_lobs_for_updated_rows = false
  name                        = "Rodney Trantow"
  password                    = "...my_password..."
  port                        = 3
  schema                      = "...my_schema..."
  type                        = "POSTGRES"
  username                    = "Gaetano_Dickens66"
}