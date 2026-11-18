resource "etleap_connection_postgres_sharded" "my_connectionpostgres_sharded" {
  auto_replicate              = "...my_auto_replicate..."
  cdc_address                 = "...my_cdc_address..."
  cdc_enabled                 = false
  cdc_port                    = 2
  deletion_of_export_products = true
  fetch_lobs_for_updated_rows = false
  name                        = "Dr. Loren Abernathy III"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "7198 Kelsi Unions"
      database = "...my_database..."
      password = "...my_password..."
      port     = 5
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "275 MacGyver Fall"
        port     = 10
        username = "Destiney.Volkman"
      }
      username = "Johnny82"
    },
  ]
  type = "POSTGRES_SHARDED"
}