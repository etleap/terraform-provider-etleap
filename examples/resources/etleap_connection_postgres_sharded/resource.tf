resource "etleap_connection_postgres_sharded" "my_connectionpostgres_sharded" {
  auto_replicate              = "...my_auto_replicate..."
  cdc_enabled                 = false
  deletion_of_export_products = false
  name                        = "Rosa McDermott"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "02097 Dedric Gateway"
      database = "...my_database..."
      password = "...my_password..."
      port     = 9
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "89224 Sporer Bridge"
        username = "Keith_Strosin46"
      }
      username = "Felipe_Muller5"
    },
  ]
  type = "POSTGRES_SHARDED"
}