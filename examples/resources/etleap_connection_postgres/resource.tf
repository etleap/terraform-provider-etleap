resource "etleap_connection_postgres" "my_connectionpostgres" {
  address                     = "224 Sporer Bridge"
  auto_replicate              = "...my_auto_replicate..."
  cdc_enabled                 = true
  database                    = "...my_database..."
  deletion_of_export_products = false
  name                        = "Merle Kohler"
  password                    = "...my_password..."
  port                        = 10
  schema                      = "...my_schema..."
  type                        = "POSTGRES"
  username                    = "Maureen_Becker6"
}