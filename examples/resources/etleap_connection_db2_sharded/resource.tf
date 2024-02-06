resource "etleap_connection_db2_sharded" "my_connectiondb2_sharded" {
  deletion_of_export_products = true
  name                        = "Daryl Rutherford"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "98390 Windler Brook"
      database = "...my_database..."
      password = "...my_password..."
      port     = 4
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "431 Carmine Alley"
        username = "Kathryn3"
      }
      username = "Sage.Hilll"
    },
  ]
  type = "DB2_SHARDED"
}