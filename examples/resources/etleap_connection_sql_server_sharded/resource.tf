resource "etleap_connection_sql_server_sharded" "my_connectionsql_server_sharded" {
  cdc_enabled                 = true
  deletion_of_export_products = false
  name                        = "Spencer Olson"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "2748 Halvorson Knolls"
      database = "...my_database..."
      password = "...my_password..."
      port     = 7
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "0304 Kaylin Route"
        username = "Richard_Douglas"
      }
      username = "Jana64"
    },
  ]
  type = "SQL_SERVER_SHARDED"
}