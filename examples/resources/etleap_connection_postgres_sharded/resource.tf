resource "etleap_connection_postgres_sharded" "my_connectionpostgres_sharded" {
  auto_replicate              = "...my_auto_replicate..."
  cdc_enabled                 = false
  deletion_of_export_products = false
  name                        = "Pete Homenick"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "189 Dean Loaf"
      database = "...my_database..."
      password = "...my_password..."
      port     = 9
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "058 Weber Landing"
        username = "Felipe_Muller5"
      }
      username = "Tess86"
    },
  ]
  type = "POSTGRES_SHARDED"
}