resource "etleap_connection_redshift_sharded" "my_connectionredshift_sharded" {
  deletion_of_export_products   = false
  dynamic_varchar_width_enabled = false
  name                          = "Molly Koss"
  query_tags_enabled            = false
  schema                        = "...my_schema..."
  shards = [
    {
      address  = "320 Hodkiewicz Cove"
      database = "...my_database..."
      password = "...my_password..."
      port     = 0
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "0468 Auer Ferry"
        port     = 2
        username = "Fern.Howe96"
      }
      username = "Jay_McKenzie"
    },
  ]
  source_only = false
  type        = "REDSHIFT_SHARDED"
}