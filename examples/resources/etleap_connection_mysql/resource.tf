resource "etleap_connection_mysql" "my_connectionmysql" {
  address                              = "533 Toy Spur"
  auto_replicate                       = "...my_auto_replicate..."
  cdc_enabled                          = false
  certificate                          = "...my_certificate..."
  database                             = "...my_database..."
  deletion_of_export_products          = true
  name                                 = "Aaron Langworth III"
  password                             = "...my_password..."
  port                                 = 9
  require_ssl_and_validate_certificate = true
  tiny_int1_is_boolean                 = true
  type                                 = "MYSQL"
  username                             = "Marquis.Sipes"
}