resource "etleap_connection_snowflake_sharded" "my_connectionsnowflake_sharded" {
  deletion_of_export_products = true
  name                        = "Maurice Vandervort"
  schema                      = "...my_schema..."
  shards = [
    {
      address = "284 Tromp Junctions"
      authentication = {
        key_pair = {
          private_key            = "...my_private_key..."
          private_key_passphrase = "...my_private_key_passphrase..."
          public_key             = "...my_public_key..."
          type                   = "PASSWORD"
        }
      }
      database  = "...my_database..."
      password  = "...my_password..."
      role      = "...my_role..."
      shard_id  = "...my_shard_id..."
      username  = "Brando.Volkman0"
      warehouse = "...my_warehouse..."
    },
  ]
  source_only         = false
  storage_integration = "...my_storage_integration..."
  type                = "SNOWFLAKE_SHARDED"
}