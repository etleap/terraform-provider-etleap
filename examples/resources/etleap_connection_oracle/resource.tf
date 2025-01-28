resource "etleap_connection_oracle" "my_connectionoracle" {
  address                     = "1575 Geoffrey Corners"
  cdc_enabled                 = false
  database                    = "...my_database..."
  deletion_of_export_products = true
  name                        = "Ricky Herman"
  password                    = "...my_password..."
  port                        = 2
  schema                      = "...my_schema..."
  type                        = "ORACLE"
  username                    = "Talon.Yost"
}