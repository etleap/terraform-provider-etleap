resource "etleap_connection_redshift_sharded" "my_connectionredshift_sharded" {
  deletion_of_export_products   = false
  dynamic_varchar_width_enabled = false
  name                          = "Silvia Mertz Jr."
  query_tags_enabled            = true
  schema                        = "...my_schema..."
  shards = [
    {
      address  = "34719 Boyle Mount"
      database = "...my_database..."
      password = "...my_password..."
      port     = 9
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "7275 MacGyver Fall"
        username = "Zander.Gerlach48"
      }
      username = "Johnny82"
    },
  ]
  source_only = false
  type        = "REDSHIFT_SHARDED"
}