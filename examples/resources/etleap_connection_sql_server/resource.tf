resource "etleap_connection_sql_server" "my_connectionsql_server" {
  address                     = "750 Hilll Unions"
  cdc_enabled                 = false
  database                    = "...my_database..."
  deletion_of_export_products = false
  name                        = "Alfred Toy"
  password                    = "...my_password..."
  port                        = 9
  schema                      = "...my_schema..."
  type                        = "SQL_SERVER"
  username                    = "Helene91"
}