resource "etleap_connection_mysql" "my_connectionmysql" {
  address                     = "12912 Selina Road"
  auto_replicate              = "...my_auto_replicate..."
  cdc_enabled                 = true
  database                    = "...my_database..."
  deletion_of_export_products = false
  name                        = "Rufus Adams"
  password                    = "...my_password..."
  port                        = 0
  tiny_int1_is_boolean        = false
  type                        = "MYSQL"
  username                    = "Nadia.Leannon"
  validate_ssl_cert           = false
}