resource "etleap_connection_oracle" "my_connectionoracle" {
  address                              = "943 Zboncak Radial"
  cdc_address                          = "...my_cdc_address..."
  cdc_enabled                          = false
  cdc_port                             = 10
  certificate                          = "...my_certificate..."
  database                             = "...my_database..."
  deletion_of_export_products          = true
  name                                 = "Lynette Nitzsche"
  password                             = "...my_password..."
  port                                 = 6
  require_ssl_and_validate_certificate = true
  schema                               = "...my_schema..."
  type                                 = "ORACLE"
  username                             = "Fabian97"
}