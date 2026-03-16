resource "etleap_connection_sql_server" "my_connectionsql_server" {
  address                     = "861 Dulce Corner"
  cdc_enabled                 = true
  database                    = "...my_database..."
  deletion_of_export_products = false
  name                        = "Miss Kelly Howell"
  password                    = "...my_password..."
  port                        = 2
  schema                      = "...my_schema..."
  type                        = "SQL_SERVER"
  username                    = "Kole_Littel"
}