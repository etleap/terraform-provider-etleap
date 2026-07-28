resource "etleap_connection_redshift_sharded" "my_connectionredshift_sharded" {
  deletion_of_export_products   = true
  dynamic_varchar_width_enabled = true
  name                          = "Faith Kris"
  query_tags_enabled            = false
  schema                        = "...my_schema..."
  shards = [
    {
      address  = "47007 Whitney Motorway"
      database = "...my_database..."
      password = "...my_password..."
      port     = 7
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "8809 West Common"
        port     = 9
        username = "Kadin32"
      }
      username = "Martine94"
    },
  ]
  source_only = false
  type        = "REDSHIFT_SHARDED"
}