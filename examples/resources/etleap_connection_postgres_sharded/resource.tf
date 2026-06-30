resource "etleap_connection_postgres_sharded" "my_connectionpostgres_sharded" {
  auto_replicate              = "...my_auto_replicate..."
  cdc_address                 = "...my_cdc_address..."
  cdc_enabled                 = false
  cdc_port                    = 8
  deletion_of_export_products = false
  fetch_lobs_for_updated_rows = true
  name                        = "Ms. Anita Bins"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "95700 Beahan Haven"
      database = "...my_database..."
      password = "...my_password..."
      port     = 4
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "198 Kelsi Unions"
        port     = 5
        username = "Patricia_Franey99"
      }
      username = "Kelley.Farrell91"
    },
  ]
  type = "POSTGRES_SHARDED"
}