resource "etleap_connection_oracle" "my_connectionoracle" {
  address                     = "5233 Emard Unions"
  cdc_enabled                 = true
  database                    = "...my_database..."
  deletion_of_export_products = true
  name                        = "Kristine Cormier"
  password                    = "...my_password..."
  port                        = 1
  schema                      = "...my_schema..."
  type                        = "ORACLE"
  username                    = "Ralph80"
}