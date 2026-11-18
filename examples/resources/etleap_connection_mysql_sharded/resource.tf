resource "etleap_connection_mysql_sharded" "my_connectionmysql_sharded" {
  auto_replicate                       = "...my_auto_replicate..."
  cdc_address                          = "...my_cdc_address..."
  cdc_enabled                          = false
  cdc_port                             = 5
  certificate                          = "...my_certificate..."
  database                             = "...my_database..."
  deletion_of_export_products          = false
  name                                 = "Mrs. Dave Labadie"
  require_ssl_and_validate_certificate = false
  shards = [
    {
      address  = "6424 Gottlieb Well"
      database = "...my_database..."
      password = "...my_password..."
      port     = 6
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "0971 Krystel Knolls"
        port     = 8
        username = "Earline1"
      }
      username = "Thora.Renner"
    },
  ]
  tiny_int1_is_boolean = true
  type                 = "MYSQL_SHARDED"
}