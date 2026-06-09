resource "etleap_connection_snowflake_sharded" "my_connectionsnowflake_sharded" {
  deletion_of_export_products = true
  name                        = "Mrs. Neal Crist MD"
  schema                      = "...my_schema..."
  shards = [
    {
      address = "84306 Fahey Neck"
      authentication = {
        key_pair = {
          private_key            = "...my_private_key..."
          private_key_passphrase = "...my_private_key_passphrase..."
          public_key             = "...my_public_key..."
          type                   = "KEY_PAIR"
        }
      }
      database  = "...my_database..."
      password  = "...my_password..."
      role      = "...my_role..."
      shard_id  = "...my_shard_id..."
      username  = "Joanne88"
      warehouse = "...my_warehouse..."
    },
  ]
  source_only         = true
  storage_integration = "...my_storage_integration..."
  type                = "SNOWFLAKE_SHARDED"
}