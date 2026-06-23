resource "etleap_connection_oracle" "my_connectionoracle" {
  address                              = "24880 Strosin Ridges"
  cdc_address                          = "...my_cdc_address..."
  cdc_enabled                          = true
  cdc_port                             = 5
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