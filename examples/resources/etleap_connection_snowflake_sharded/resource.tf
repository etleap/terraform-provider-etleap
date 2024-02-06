resource "etleap_connection_snowflake_sharded" "my_connectionsnowflake_sharded" {
  deletion_of_export_products = false
  name                        = "Clifton Barrows"
  schema                      = "...my_schema..."
  shards = [
    {
      address   = "880 Mia Cape"
      database  = "...my_database..."
      password  = "...my_password..."
      role      = "...my_role..."
      shard_id  = "...my_shard_id..."
      username  = "Michel70"
      warehouse = "...my_warehouse..."
    },
  ]
  source_only = true
  type        = "SNOWFLAKE_SHARDED"
}