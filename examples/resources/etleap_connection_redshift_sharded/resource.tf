resource "etleap_connection_redshift_sharded" "my_connectionredshift_sharded" {
  deletion_of_export_products   = true
  dynamic_varchar_width_enabled = false
  name                          = "Yvette Konopelski"
  query_tags_enabled            = false
  schema                        = "...my_schema..."
  shards = [
    {
      address  = "910 Schuster Key"
      database = "...my_database..."
      password = "...my_password..."
      port     = 3
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "812 Broderick Points"
        username = "Amaya.Schneider76"
      }
      username = "Sabryna86"
    },
  ]
  source_only = false
  type        = "REDSHIFT_SHARDED"
}