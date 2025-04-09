resource "etleap_connection_sql_server" "my_connectionsql_server" {
  address                     = "4993 Zemlak Rapids"
  cdc_enabled                 = false
  database                    = "...my_database..."
  deletion_of_export_products = false
  name                        = "Mrs. Evelyn Bogan"
  password                    = "...my_password..."
  port                        = 0
  schema                      = "...my_schema..."
  type                        = "SQL_SERVER"
  username                    = "Marlee24"
}