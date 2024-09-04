resource "etleap_connection_mysql_sharded" "my_connectionmysql_sharded" {
  auto_replicate              = "...my_auto_replicate..."
  cdc_enabled                 = true
  database                    = "...my_database..."
  deletion_of_export_products = false
  name                        = "Lorraine Herzog Jr."
  shards = [
    {
      address  = "7681 Nienow Vista"
      database = "...my_database..."
      password = "...my_password..."
      port     = 5
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "3298 Favian Turnpike"
        username = "Camryn.Langworth"
      }
      username = "Hilma3"
    },
  ]
  tiny_int1_is_boolean = true
  type                 = "MYSQL_SHARDED"
  validate_ssl_cert    = true
}