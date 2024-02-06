resource "etleap_connection_redshift_sharded" "my_connectionredshift_sharded" {
  deletion_of_export_products   = false
  dynamic_varchar_width_enabled = false
  name                          = "Marsha Kuhic"
  query_tags_enabled            = true
  schema                        = "...my_schema..."
  shards = [
    {
      address  = "97398 River Locks"
      database = "...my_database..."
      password = "...my_password..."
      port     = 3
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "183 Ruby Circle"
        username = "Theron75"
      }
      username = "Adell.West89"
    },
  ]
  source_only = true
  type        = "REDSHIFT_SHARDED"
}