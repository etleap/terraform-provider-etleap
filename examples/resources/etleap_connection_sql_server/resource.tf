resource "etleap_connection_sql_server" "my_connectionsql_server" {
  address                     = "102 Louisa Branch"
  cdc_enabled                 = false
  database                    = "...my_database..."
  deletion_of_export_products = true
  name                        = "Martha Nicolas"
  password                    = "...my_password..."
  port                        = 6
  schema                      = "...my_schema..."
  type                        = "SQL_SERVER"
  username                    = "Kari.Labadie88"
}