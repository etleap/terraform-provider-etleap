resource "etleap_connection_mysql_sharded" "my_connectionmysql_sharded" {
  auto_replicate                       = "...my_auto_replicate..."
  cdc_enabled                          = false
  certificate                          = "...my_certificate..."
  database                             = "...my_database..."
  deletion_of_export_products          = false
  name                                 = "Jasmine King"
  require_ssl_and_validate_certificate = true
  shards = [
    {
      address  = "85924 Schulist Mill"
      database = "...my_database..."
      password = "...my_password..."
      port     = 5
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "25546 Eleanora Lane"
        port     = 8
        username = "Doyle_Zemlak"
      }
      username = "Austin.Weissnat22"
    },
  ]
  tiny_int1_is_boolean = true
  type                 = "MYSQL_SHARDED"
}