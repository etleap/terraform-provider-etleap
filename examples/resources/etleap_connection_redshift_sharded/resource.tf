resource "etleap_connection_redshift_sharded" "my_connectionredshift_sharded" {
  deletion_of_export_products   = true
  dynamic_varchar_width_enabled = false
  name                          = "Earl Bayer"
  query_tags_enabled            = false
  schema                        = "...my_schema..."
  shards = [
    {
      address  = "08034 Crona Ville"
      database = "...my_database..."
      password = "...my_password..."
      port     = 8
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "59572 Mayert Wells"
        username = "Kelley.Farrell91"
      }
      username = "Jacynthe_Kshlerin"
    },
  ]
  source_only = true
  type        = "REDSHIFT_SHARDED"
}