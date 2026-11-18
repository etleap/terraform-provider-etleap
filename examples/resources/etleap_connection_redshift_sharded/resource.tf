resource "etleap_connection_redshift_sharded" "my_connectionredshift_sharded" {
  deletion_of_export_products   = false
  dynamic_varchar_width_enabled = false
  name                          = "Miss Shelia Bayer"
  query_tags_enabled            = true
  schema                        = "...my_schema..."
  shards = [
    {
      address  = "7688 Turner Meadow"
      database = "...my_database..."
      password = "...my_password..."
      port     = 10
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "85523 Quitzon Fork"
        port     = 10
        username = "Tierra.Schmidt"
      }
      username = "Ebba_Johnson81"
    },
  ]
  source_only = true
  type        = "REDSHIFT_SHARDED"
}