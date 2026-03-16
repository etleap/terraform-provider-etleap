resource "etleap_connection_mysql" "my_connectionmysql" {
  address                              = "077 Bonnie Manors"
  auto_replicate                       = "...my_auto_replicate..."
  cdc_enabled                          = false
  certificate                          = "...my_certificate..."
  database                             = "...my_database..."
  deletion_of_export_products          = true
  name                                 = "Betty Johnson"
  password                             = "...my_password..."
  port                                 = 6
  require_ssl_and_validate_certificate = true
  tiny_int1_is_boolean                 = false
  type                                 = "MYSQL"
  username                             = "Meredith38"
}