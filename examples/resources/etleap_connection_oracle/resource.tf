resource "etleap_connection_oracle" "my_connectionoracle" {
  address                              = "8805 Reynolds View"
  cdc_enabled                          = true
  certificate                          = "...my_certificate..."
  database                             = "...my_database..."
  deletion_of_export_products          = true
  name                                 = "Dr. Emmett Powlowski"
  password                             = "...my_password..."
  port                                 = 0
  require_ssl_and_validate_certificate = true
  schema                               = "...my_schema..."
  type                                 = "ORACLE"
  username                             = "Ludie_Schuppe"
}