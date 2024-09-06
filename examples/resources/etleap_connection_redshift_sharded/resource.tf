resource "etleap_connection_redshift_sharded" "my_connectionredshift_sharded" {
  deletion_of_export_products   = false
  dynamic_varchar_width_enabled = true
  name                          = "Tom Ward"
  query_tags_enabled            = true
  schema                        = "...my_schema..."
  shards = [
    {
      address  = "484 Sterling Curve"
      database = "...my_database..."
      password = "...my_password..."
      port     = 9
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "0809 Harber Route"
        username = "Adell.West89"
      }
      username = "Christopher37"
    },
  ]
  source_only = false
  type        = "REDSHIFT_SHARDED"
}