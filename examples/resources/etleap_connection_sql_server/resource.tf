resource "etleap_connection_sql_server" "my_connectionsql_server" {
  address                     = "1691 Margie Wells"
  cdc_enabled                 = false
  database                    = "...my_database..."
  deletion_of_export_products = false
  name                        = "Alberta Deckow"
  password                    = "...my_password..."
  port                        = 10
  schema                      = "...my_schema..."
  type                        = "SQL_SERVER"
  username                    = "Sally.Beer51"
}