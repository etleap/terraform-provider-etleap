resource "etleap_connection_mysql" "my_connectionmysql" {
  address                              = "813 Corwin Mall"
  auto_replicate                       = "...my_auto_replicate..."
  cdc_enabled                          = true
  certificate                          = "...my_certificate..."
  database                             = "...my_database..."
  deletion_of_export_products          = true
  name                                 = "Miss Jenna Hansen"
  password                             = "...my_password..."
  port                                 = 4
  require_ssl_and_validate_certificate = false
  tiny_int1_is_boolean                 = false
  type                                 = "MYSQL"
  username                             = "Leone.Simonis"
}