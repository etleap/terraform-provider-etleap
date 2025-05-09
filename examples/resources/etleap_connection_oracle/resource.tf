resource "etleap_connection_oracle" "my_connectionoracle" {
  address                              = "943 Cormier Unions"
  cdc_enabled                          = false
  certificate                          = "...my_certificate..."
  database                             = "...my_database..."
  deletion_of_export_products          = true
  name                                 = "Luke Bartoletti"
  password                             = "...my_password..."
  port                                 = 5
  require_ssl_and_validate_certificate = false
  schema                               = "...my_schema..."
  type                                 = "ORACLE"
  username                             = "Guadalupe_Grady93"
}