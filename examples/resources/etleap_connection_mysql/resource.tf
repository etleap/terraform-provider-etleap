resource "etleap_connection_mysql" "my_connectionmysql" {
  address                     = "619 Dibbert Ford"
  auto_replicate              = "...my_auto_replicate..."
  cdc_enabled                 = true
  database                    = "...my_database..."
  deletion_of_export_products = false
  name                        = "Mr. Norma Sawayn"
  password                    = "...my_password..."
  port                        = 10
  tiny_int1_is_boolean        = false
  type                        = "MYSQL"
  username                    = "Eleazar_Hoeger2"
  validate_ssl_cert           = false
}