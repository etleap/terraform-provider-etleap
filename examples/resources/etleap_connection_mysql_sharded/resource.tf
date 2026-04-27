resource "etleap_connection_mysql_sharded" "my_connectionmysql_sharded" {
  auto_replicate                       = "...my_auto_replicate..."
  cdc_enabled                          = true
  certificate                          = "...my_certificate..."
  database                             = "...my_database..."
  deletion_of_export_products          = false
  name                                 = "Wallace Goodwin"
  require_ssl_and_validate_certificate = false
  shards = [
    {
      address  = "99436 Tamia Mountain"
      database = "...my_database..."
      password = "...my_password..."
      port     = 1
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "608 Jolie Prairie"
        port     = 4
        username = "Earnest_Terry"
      }
      username = "Elinor_Kuhn"
    },
  ]
  tiny_int1_is_boolean = true
  type                 = "MYSQL_SHARDED"
}