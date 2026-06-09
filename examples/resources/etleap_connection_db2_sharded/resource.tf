resource "etleap_connection_db2_sharded" "my_connectiondb2_sharded" {
  certificate                 = "...my_certificate..."
  deletion_of_export_products = true
  name                        = "Henry Beatty"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "2625 Lenore Meadow"
      database = "...my_database..."
      password = "...my_password..."
      port     = 8
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "31764 Baylee Pass"
        port     = 7
        username = "Cleveland90"
      }
      username = "Sandrine84"
    },
  ]
  type = "DB2_SHARDED"
}