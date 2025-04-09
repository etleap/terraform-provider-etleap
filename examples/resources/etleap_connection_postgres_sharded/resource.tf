resource "etleap_connection_postgres_sharded" "my_connectionpostgres_sharded" {
  auto_replicate              = "...my_auto_replicate..."
  cdc_enabled                 = false
  deletion_of_export_products = false
  name                        = "Jenna Daniel V"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "068 Ledner Harbors"
      database = "...my_database..."
      password = "...my_password..."
      port     = 8
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "437 Marvin Run"
        username = "Dino_Haley"
      }
      username = "Kara_Ruecker20"
    },
  ]
  type = "POSTGRES_SHARDED"
}