resource "etleap_connection_oracle" "my_connectionoracle" {
  address                              = "099 Sarah Pines"
  cdc_address                          = "...my_cdc_address..."
  cdc_enabled                          = false
  cdc_port                             = 6
  certificate                          = "...my_certificate..."
  database                             = "...my_database..."
  deletion_of_export_products          = false
  name                                 = "Dr. Yvonne Schultz"
  password                             = "...my_password..."
  port                                 = 3
  require_ssl_and_validate_certificate = true
  schema                               = "...my_schema..."
  type                                 = "ORACLE"
  username                             = "Meaghan.Donnelly90"
}