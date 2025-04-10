resource "etleap_connection_snowflake_sharded" "my_connectionsnowflake_sharded" {
  deletion_of_export_products = false
  name                        = "Rodolfo Conroy"
  schema                      = "...my_schema..."
  shards = [
    {
      address = "3962 Alfonzo Dam"
      authentication = {
        key_pair = {
          private_key = "...my_private_key..."
          public_key  = "...my_public_key..."
          type        = "KEY_PAIR"
        }
      }
      database  = "...my_database..."
      password  = "...my_password..."
      role      = "...my_role..."
      shard_id  = "...my_shard_id..."
      username  = "Letitia.Turner"
      warehouse = "...my_warehouse..."
    },
  ]
  source_only = true
  type        = "SNOWFLAKE_SHARDED"
}