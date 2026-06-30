resource "etleap_connection_mysql_sharded" "my_connectionmysql_sharded" {
  auto_replicate                       = "...my_auto_replicate..."
  cdc_address                          = "...my_cdc_address..."
  cdc_enabled                          = false
  cdc_port                             = 4
  certificate                          = "...my_certificate..."
  database                             = "...my_database..."
  deletion_of_export_products          = false
  name                                 = "Freddie Weber"
  require_ssl_and_validate_certificate = false
  shards = [
    {
      address  = "550 Kathryn Overpass"
      database = "...my_database..."
      password = "...my_password..."
      port     = 4
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "4247 Yvonne Neck"
        port     = 6
        username = "Austin.Weissnat22"
      }
      username = "Krystel.Kihn2"
    },
  ]
  tiny_int1_is_boolean = false
  type                 = "MYSQL_SHARDED"
}