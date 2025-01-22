resource "etleap_connection_sql_server" "my_connectionsql_server" {
  address                     = "35472 Skiles Mount"
  cdc_enabled                 = false
  database                    = "...my_database..."
  deletion_of_export_products = false
  name                        = "Bobby Predovic I"
  password                    = "...my_password..."
  port                        = 7
  schema                      = "...my_schema..."
  type                        = "SQL_SERVER"
  username                    = "Zita.Renner"
}