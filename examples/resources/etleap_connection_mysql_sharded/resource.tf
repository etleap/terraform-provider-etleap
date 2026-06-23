resource "etleap_connection_mysql_sharded" "my_connectionmysql_sharded" {
  auto_replicate                       = "...my_auto_replicate..."
  cdc_address                          = "...my_cdc_address..."
  cdc_enabled                          = false
  cdc_port                             = 7
  certificate                          = "...my_certificate..."
  database                             = "...my_database..."
  deletion_of_export_products          = false
  name                                 = "Harvey Langworth"
  require_ssl_and_validate_certificate = false
  shards = [
    {
      address  = "592 Josefina Spurs"
      database = "...my_database..."
      password = "...my_password..."
      port     = 5
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "0255 Mackenzie Keys"
        port     = 3
        username = "Jammie99"
      }
      username = "Kobe94"
    },
  ]
  tiny_int1_is_boolean = false
  type                 = "MYSQL_SHARDED"
}