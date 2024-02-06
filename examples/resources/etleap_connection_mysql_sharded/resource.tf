resource "etleap_connection_mysql_sharded" "my_connectionmysql_sharded" {
  auto_replicate              = "...my_auto_replicate..."
  cdc_enabled                 = true
  deletion_of_export_products = true
  name                        = "Mrs. Jeanne Rutherford"
  shards = [
    {
      address  = "4423 Dereck Glens"
      database = "...my_database..."
      password = "...my_password..."
      port     = 3
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "10576 Dickinson Trail"
        username = "Lourdes.Williamson"
      }
      username = "Emanuel_Doyle13"
    },
  ]
  tiny_int1_is_boolean = true
  type                 = "MYSQL_SHARDED"
  validate_ssl_cert    = true
}