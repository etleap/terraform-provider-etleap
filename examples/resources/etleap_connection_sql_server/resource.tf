resource "etleap_connection_sql_server" "my_connectionsql_server" {
  address                     = "69284 Tromp Junctions"
  cdc_address                 = "...my_cdc_address..."
  cdc_enabled                 = false
  cdc_port                    = 1
  database                    = "...my_database..."
  deletion_of_export_products = true
  name                        = "Lester Abbott"
  password                    = "...my_password..."
  port                        = 6
  schema                      = "...my_schema..."
  type                        = "SQL_SERVER"
  username                    = "Adrianna_Mills54"
}