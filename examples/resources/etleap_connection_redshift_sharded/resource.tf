resource "etleap_connection_redshift_sharded" "my_connectionredshift_sharded" {
  deletion_of_export_products   = false
  dynamic_varchar_width_enabled = false
  name                          = "Willard Lowe"
  query_tags_enabled            = false
  schema                        = "...my_schema..."
  shards = [
    {
      address  = "4857 Heller Extensions"
      database = "...my_database..."
      password = "...my_password..."
      port     = 0
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "3105 Hiram Ramp"
        port     = 9
        username = "Richmond22"
      }
      username = "Fern.Howe96"
    },
  ]
  source_only = true
  type        = "REDSHIFT_SHARDED"
}