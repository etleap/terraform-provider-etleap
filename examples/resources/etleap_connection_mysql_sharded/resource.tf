resource "etleap_connection_mysql_sharded" "my_connectionmysql_sharded" {
  auto_replicate                       = "...my_auto_replicate..."
  cdc_enabled                          = false
  certificate                          = "...my_certificate..."
  database                             = "...my_database..."
  deletion_of_export_products          = true
  name                                 = "Enrique Gislason"
  require_ssl_and_validate_certificate = true
  shards = [
    {
      address  = "305 Jailyn Terrace"
      database = "...my_database..."
      password = "...my_password..."
      port     = 3
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "915 Jerde Oval"
        username = "Marlee.Towne"
      }
      username = "Hulda79"
    },
  ]
  tiny_int1_is_boolean = false
  type                 = "MYSQL_SHARDED"
}