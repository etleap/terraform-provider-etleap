resource "etleap_connection_sql_server_sharded" "my_connectionsql_server_sharded" {
  cdc_enabled                 = false
  deletion_of_export_products = false
  name                        = "Wilson Pollich"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "9376 Torp Key"
      database = "...my_database..."
      password = "...my_password..."
      port     = 3
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "410 Joey Falls"
        username = "Cassandre.Kautzer64"
      }
      username = "Sam.Marvin"
    },
  ]
  type = "SQL_SERVER_SHARDED"
}