resource "etleap_connection_mysql" "my_connectionmysql" {
  address                              = "299 Elizabeth Ports"
  auto_replicate                       = "...my_auto_replicate..."
  cdc_address                          = "...my_cdc_address..."
  cdc_enabled                          = true
  cdc_port                             = 10
  certificate                          = "...my_certificate..."
  database                             = "...my_database..."
  deletion_of_export_products          = false
  name                                 = "Dr. Nichole Nienow"
  password                             = "...my_password..."
  port                                 = 5
  require_ssl_and_validate_certificate = false
  tiny_int1_is_boolean                 = true
  type                                 = "MYSQL"
  username                             = "Earnest_Terry"
}