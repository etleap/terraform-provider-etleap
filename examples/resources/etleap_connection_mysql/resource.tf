resource "etleap_connection_mysql" "my_connectionmysql" {
  address                              = "7681 Nienow Vista"
  auto_replicate                       = "...my_auto_replicate..."
  cdc_enabled                          = false
  certificate                          = "...my_certificate..."
  database                             = "...my_database..."
  deletion_of_export_products          = true
  name                                 = "Crystal Watsica"
  password                             = "...my_password..."
  port                                 = 3
  require_ssl_and_validate_certificate = true
  tiny_int1_is_boolean                 = true
  type                                 = "MYSQL"
  username                             = "Jonas42"
}