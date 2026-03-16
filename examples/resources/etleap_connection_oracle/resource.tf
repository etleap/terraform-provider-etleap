resource "etleap_connection_oracle" "my_connectionoracle" {
  address                              = "25482 Eladio Branch"
  cdc_enabled                          = false
  certificate                          = "...my_certificate..."
  database                             = "...my_database..."
  deletion_of_export_products          = false
  name                                 = "Thelma Hamill PhD"
  password                             = "...my_password..."
  port                                 = 10
  require_ssl_and_validate_certificate = false
  schema                               = "...my_schema..."
  type                                 = "ORACLE"
  username                             = "Dean_Kub1"
}