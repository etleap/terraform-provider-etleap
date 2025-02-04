resource "etleap_connection_oracle_sharded" "my_connectionoracle_sharded" {
  cdc_enabled                 = false
  deletion_of_export_products = false
  name                        = "Jasmine King"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "2859 Jennie Manor"
      database = "...my_database..."
      password = "...my_password..."
      port     = 8
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "5025 Kirlin Plain"
        username = "Hoyt_Grimes"
      }
      username = "Doyle_Zemlak"
    },
  ]
  type = "ORACLE_SHARDED"
}