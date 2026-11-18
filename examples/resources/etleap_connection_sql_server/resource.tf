resource "etleap_connection_sql_server" "my_connectionsql_server" {
  address                     = "6069 Littel Ville"
  cdc_address                 = "...my_cdc_address..."
  cdc_enabled                 = false
  cdc_port                    = 9
  database                    = "...my_database..."
  deletion_of_export_products = true
  name                        = "Jack Kozey"
  password                    = "...my_password..."
  port                        = 8
  schema                      = "...my_schema..."
  type                        = "SQL_SERVER"
  username                    = "Amely_Morar"
}