resource "etleap_connection_oracle" "my_connectionoracle" {
  address                     = "915 Jerde Oval"
  cdc_enabled                 = false
  database                    = "...my_database..."
  deletion_of_export_products = false
  name                        = "Wanda Kertzmann DDS"
  password                    = "...my_password..."
  port                        = 8
  schema                      = "...my_schema..."
  type                        = "ORACLE"
  username                    = "Amalia_Cartwright"
}