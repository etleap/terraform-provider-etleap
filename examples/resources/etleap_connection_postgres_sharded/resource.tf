resource "etleap_connection_postgres_sharded" "my_connectionpostgres_sharded" {
  auto_replicate              = "...my_auto_replicate..."
  cdc_enabled                 = true
  deletion_of_export_products = false
  name                        = "Frank Vandervort DDS"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "415 Tromp Overpass"
      database = "...my_database..."
      password = "...my_password..."
      port     = 9
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "28491 Rohan Ports"
        username = "Ricky22"
      }
      username = "Audra24"
    },
  ]
  type = "POSTGRES_SHARDED"
}