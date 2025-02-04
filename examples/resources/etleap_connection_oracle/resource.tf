resource "etleap_connection_oracle" "my_connectionoracle" {
  address                     = "2312 Legros Forest"
  cdc_enabled                 = true
  database                    = "...my_database..."
  deletion_of_export_products = false
  name                        = "Leonard Vandervort"
  password                    = "...my_password..."
  port                        = 3
  schema                      = "...my_schema..."
  type                        = "ORACLE"
  username                    = "Marisol_Cormier11"
}