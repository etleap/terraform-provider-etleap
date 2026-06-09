resource "etleap_connection_redshift_sharded" "my_connectionredshift_sharded" {
  deletion_of_export_products   = false
  dynamic_varchar_width_enabled = true
  name                          = "Mae Bartell"
  query_tags_enabled            = true
  schema                        = "...my_schema..."
  shards = [
    {
      address  = "022 Gertrude Track"
      database = "...my_database..."
      password = "...my_password..."
      port     = 4
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "455 Bins Junctions"
        port     = 8
        username = "Alyce.Anderson97"
      }
      username = "Keaton_Prosacco"
    },
  ]
  source_only = true
  type        = "REDSHIFT_SHARDED"
}