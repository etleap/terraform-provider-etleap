resource "etleap_connection_postgres_sharded" "my_connectionpostgres_sharded" {
  auto_replicate              = "...my_auto_replicate..."
  cdc_enabled                 = true
  deletion_of_export_products = false
  name                        = "Kent McLaughlin"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "091 Lavina Ranch"
      database = "...my_database..."
      password = "...my_password..."
      port     = 2
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "097 Pattie Well"
        username = "Javier_Dare"
      }
      username = "Stephania.Medhurst20"
    },
  ]
  type = "POSTGRES_SHARDED"
}