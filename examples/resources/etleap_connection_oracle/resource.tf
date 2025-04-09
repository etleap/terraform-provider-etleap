resource "etleap_connection_oracle" "my_connectionoracle" {
  address                              = "1868 Arielle Extension"
  cdc_enabled                          = true
  certificate                          = "...my_certificate..."
  database                             = "...my_database..."
  deletion_of_export_products          = false
  name                                 = "Gregg Legros DVM"
  password                             = "...my_password..."
  port                                 = 8
  require_ssl_and_validate_certificate = true
  schema                               = "...my_schema..."
  type                                 = "ORACLE"
  username                             = "Domenic.Tremblay"
}