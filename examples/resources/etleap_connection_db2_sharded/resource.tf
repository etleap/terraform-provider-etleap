resource "etleap_connection_db2_sharded" "my_connectiondb2_sharded" {
  deletion_of_export_products = true
  name                        = "Marilyn Abbott"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "08361 Buckridge Forks"
      database = "...my_database..."
      password = "...my_password..."
      port     = 2
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "774 Maybell Hills"
        username = "Lindsay80"
      }
      username = "Marina65"
    },
  ]
  type = "DB2_SHARDED"
}