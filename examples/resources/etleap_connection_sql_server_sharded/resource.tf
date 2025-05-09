resource "etleap_connection_sql_server_sharded" "my_connectionsql_server_sharded" {
  cdc_enabled                 = false
  deletion_of_export_products = false
  name                        = "Lewis Batz"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "7178 Schaden Ridge"
      database = "...my_database..."
      password = "...my_password..."
      port     = 2
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "433 Lonzo Valleys"
        username = "Lelia_Gorczany51"
      }
      username = "Lauryn_Will6"
    },
  ]
  type = "SQL_SERVER_SHARDED"
}