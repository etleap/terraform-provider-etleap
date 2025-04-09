resource "etleap_connection_mysql" "my_connectionmysql" {
  address                              = "47522 Brown Burg"
  auto_replicate                       = "...my_auto_replicate..."
  cdc_enabled                          = false
  certificate                          = "...my_certificate..."
  database                             = "...my_database..."
  deletion_of_export_products          = false
  name                                 = "Jan Kuhlman"
  password                             = "...my_password..."
  port                                 = 4
  require_ssl_and_validate_certificate = false
  tiny_int1_is_boolean                 = false
  type                                 = "MYSQL"
  username                             = "Reyes.Pollich49"
}