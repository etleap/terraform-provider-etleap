resource "etleap_connection_mysql_sharded" "my_connectionmysql_sharded" {
  auto_replicate                       = "...my_auto_replicate..."
  cdc_enabled                          = false
  certificate                          = "...my_certificate..."
  database                             = "...my_database..."
  deletion_of_export_products          = true
  name                                 = "Micheal Rogahn"
  require_ssl_and_validate_certificate = true
  shards = [
    {
      address  = "15751 Hahn Roads"
      database = "...my_database..."
      password = "...my_password..."
      port     = 10
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "487 Heathcote Ranch"
        username = "Allene99"
      }
      username = "Esteban57"
    },
  ]
  tiny_int1_is_boolean = true
  type                 = "MYSQL_SHARDED"
}