resource "etleap_connection_db2_sharded" "my_connectiondb2_sharded" {
  certificate                 = "...my_certificate..."
  deletion_of_export_products = true
  name                        = "Gerald Greenholt"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "744 Howe Passage"
      database = "...my_database..."
      password = "...my_password..."
      port     = 1
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "86526 Khalid Bypass"
        port     = 6
        username = "Judson39"
      }
      username = "Ceasar.Reilly"
    },
  ]
  type = "DB2_SHARDED"
}