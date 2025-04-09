resource "etleap_connection_postgres" "my_connectionpostgres" {
  address                     = "0998 Tillman Ranch"
  auto_replicate              = "...my_auto_replicate..."
  cdc_enabled                 = false
  database                    = "...my_database..."
  deletion_of_export_products = true
  name                        = "Darryl Jacobi"
  password                    = "...my_password..."
  port                        = 7
  schema                      = "...my_schema..."
  type                        = "POSTGRES"
  username                    = "Heber_Hodkiewicz"
}