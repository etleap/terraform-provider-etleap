resource "etleap_connection_mysql" "my_connectionmysql" {
  address                              = "7523 Daisha Parks"
  auto_replicate                       = "...my_auto_replicate..."
  cdc_address                          = "...my_cdc_address..."
  cdc_enabled                          = false
  cdc_port                             = 3
  certificate                          = "...my_certificate..."
  database                             = "...my_database..."
  deletion_of_export_products          = true
  name                                 = "Maggie Emard"
  password                             = "...my_password..."
  port                                 = 5
  require_ssl_and_validate_certificate = false
  tiny_int1_is_boolean                 = false
  type                                 = "MYSQL"
  username                             = "Calista.Vandervort"
}