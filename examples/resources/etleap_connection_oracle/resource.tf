resource "etleap_connection_oracle" "my_connectionoracle" {
  address                              = "5233 Emard Unions"
  cdc_enabled                          = true
  certificate                          = "...my_certificate..."
  database                             = "...my_database..."
  deletion_of_export_products          = true
  name                                 = "Kristine Cormier"
  password                             = "...my_password..."
  port                                 = 1
  require_ssl_and_validate_certificate = false
  schema                               = "...my_schema..."
  type                                 = "ORACLE"
  username                             = "Loy.Bartoletti51"
}