resource "etleap_connection_postgres" "my_connectionpostgres" {
  address                     = "189 Dean Loaf"
  auto_replicate              = "...my_auto_replicate..."
  cdc_enabled                 = false
  database                    = "...my_database..."
  deletion_of_export_products = false
  name                        = "Leona Strosin"
  password                    = "...my_password..."
  port                        = 5
  schema                      = "...my_schema..."
  type                        = "POSTGRES"
  username                    = "Felipe_Muller5"
}