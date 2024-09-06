resource "etleap_connection_sql_server" "my_connectionsql_server" {
  address                     = "69179 Osinski Crescent"
  cdc_enabled                 = true
  database                    = "...my_database..."
  deletion_of_export_products = false
  name                        = "Marguerite Wuckert"
  password                    = "...my_password..."
  port                        = 8
  schema                      = "...my_schema..."
  type                        = "SQL_SERVER"
  username                    = "Maida_Wyman"
}