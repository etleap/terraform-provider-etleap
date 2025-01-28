resource "etleap_connection_snowflake_sharded" "my_connectionsnowflake_sharded" {
  deletion_of_export_products = true
  name                        = "Ricky Schmeler"
  schema                      = "...my_schema..."
  shards = [
    {
      address = "346 Bins Hills"
      authentication = {
        key_pair = {
          private_key = "...my_private_key..."
          public_key  = "...my_public_key..."
          type        = "PASSWORD"
        }
      }
      database  = "...my_database..."
      password  = "...my_password..."
      role      = "...my_role..."
      shard_id  = "...my_shard_id..."
      username  = "Ivah_Cruickshank"
      warehouse = "...my_warehouse..."
    },
  ]
  source_only = true
  type        = "SNOWFLAKE_SHARDED"
}