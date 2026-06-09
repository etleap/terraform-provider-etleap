resource "etleap_connection_sql_server" "my_connectionsql_server" {
  address                     = "62315 Ferry Drives"
  cdc_enabled                 = false
  database                    = "...my_database..."
  deletion_of_export_products = true
  name                        = "Sidney Bailey"
  password                    = "...my_password..."
  port                        = 10
  schema                      = "...my_schema..."
  type                        = "SQL_SERVER"
  username                    = "Dexter.Pacocha89"
}