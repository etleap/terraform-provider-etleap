resource "etleap_connection_oracle" "my_connectionoracle" {
  address                     = "50930 Feeney Land"
  cdc_enabled                 = false
  database                    = "...my_database..."
  deletion_of_export_products = true
  name                        = "Santos Carroll"
  password                    = "...my_password..."
  port                        = 4
  schema                      = "...my_schema..."
  type                        = "ORACLE"
  username                    = "Kyler.Paucek17"
}