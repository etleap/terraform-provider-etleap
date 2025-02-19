resource "etleap_connection_oracle_sharded" "my_connectionoracle_sharded" {
  cdc_enabled                 = true
  deletion_of_export_products = false
  name                        = "Reginald Grady"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "245 Lindgren Loop"
      database = "...my_database..."
      password = "...my_password..."
      port     = 0
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "55464 Jammie River"
        username = "Doyle_Zemlak"
      }
      username = "Austin.Weissnat22"
    },
  ]
  type = "ORACLE_SHARDED"
}