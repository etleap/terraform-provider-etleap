resource "etleap_connection_sql_server" "my_connectionsql_server" {
  address                     = "432 Little Flats"
  cdc_enabled                 = false
  database                    = "...my_database..."
  deletion_of_export_products = false
  name                        = "Karla West"
  password                    = "...my_password..."
  port                        = 1
  schema                      = "...my_schema..."
  type                        = "SQL_SERVER"
  username                    = "Kaitlyn62"
}