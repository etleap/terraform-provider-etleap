resource "etleap_connection_postgres_sharded" "my_connectionpostgres_sharded" {
  auto_replicate              = "...my_auto_replicate..."
  cdc_enabled                 = false
  deletion_of_export_products = false
  fetch_lobs_for_updated_rows = true
  name                        = "Dr. Eloise Hudson DVM"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "370 Towne Extensions"
      database = "...my_database..."
      password = "...my_password..."
      port     = 9
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "32316 Orn Prairie"
        port     = 8
        username = "Rosamond.Ruecker"
      }
      username = "Antwan.Barton"
    },
  ]
  type = "POSTGRES_SHARDED"
}