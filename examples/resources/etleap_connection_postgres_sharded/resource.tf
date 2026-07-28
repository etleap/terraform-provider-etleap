resource "etleap_connection_postgres_sharded" "my_connectionpostgres_sharded" {
  auto_replicate              = "...my_auto_replicate..."
  cdc_address                 = "...my_cdc_address..."
  cdc_enabled                 = false
  cdc_port                    = 7
  deletion_of_export_products = false
  fetch_lobs_for_updated_rows = false
  name                        = "Benny Kshlerin Jr."
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "1095 Abernathy Alley"
      database = "...my_database..."
      password = "...my_password..."
      port     = 9
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "34719 Boyle Mount"
        port     = 9
        username = "Joe77"
      }
      username = "Kobe.Zulauf"
    },
  ]
  type = "POSTGRES_SHARDED"
}