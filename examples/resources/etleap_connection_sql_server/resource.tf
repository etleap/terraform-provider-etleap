resource "etleap_connection_sql_server" "my_connectionsql_server" {
  address                     = "53980 Durgan Manors"
  cdc_enabled                 = false
  database                    = "...my_database..."
  deletion_of_export_products = false
  name                        = "Dr. Shelia Durgan"
  password                    = "...my_password..."
  port                        = 6
  schema                      = "...my_schema..."
  type                        = "SQL_SERVER"
  username                    = "Chaya_Collins10"
}