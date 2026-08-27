resource "etleap_connection_sql_server" "my_connectionsql_server" {
  address                     = "2198 Aaron Harbors"
  cdc_address                 = "...my_cdc_address..."
  cdc_enabled                 = false
  cdc_port                    = 6
  database                    = "...my_database..."
  deletion_of_export_products = false
  name                        = "Emilio Roob"
  password                    = "...my_password..."
  port                        = 0
  schema                      = "...my_schema..."
  type                        = "SQL_SERVER"
  username                    = "Roel.Bradtke"
}