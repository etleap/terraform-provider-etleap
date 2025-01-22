resource "etleap_connection_redshift_sharded" "my_connectionredshift_sharded" {
  deletion_of_export_products   = true
  dynamic_varchar_width_enabled = true
  name                          = "Brittany Dickens"
  query_tags_enabled            = false
  schema                        = "...my_schema..."
  shards = [
    {
      address  = "787 Dominique Centers"
      database = "...my_database..."
      password = "...my_password..."
      port     = 0
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "1095 Abernathy Alley"
        username = "Rosalia.Beahan"
      }
      username = "Mikel_Crona9"
    },
  ]
  source_only = false
  type        = "REDSHIFT_SHARDED"
}