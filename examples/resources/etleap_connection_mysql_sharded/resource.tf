resource "etleap_connection_mysql_sharded" "my_connectionmysql_sharded" {
  auto_replicate                       = "...my_auto_replicate..."
  cdc_enabled                          = true
  certificate                          = "...my_certificate..."
  database                             = "...my_database..."
  deletion_of_export_products          = true
  name                                 = "Ms. Melba Hamill"
  require_ssl_and_validate_certificate = false
  shards = [
    {
      address  = "379 Karianne Motorway"
      database = "...my_database..."
      password = "...my_password..."
      port     = 4
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "6841 Osvaldo Brooks"
        username = "Peggie12"
      }
      username = "Joyce.Veum2"
    },
  ]
  tiny_int1_is_boolean = false
  type                 = "MYSQL_SHARDED"
}