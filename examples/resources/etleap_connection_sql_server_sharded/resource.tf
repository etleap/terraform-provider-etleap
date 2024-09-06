resource "etleap_connection_sql_server_sharded" "my_connectionsql_server_sharded" {
  cdc_enabled                 = true
  deletion_of_export_products = true
  name                        = "Raquel Sauer"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "503 Ivah Crossing"
      database = "...my_database..."
      password = "...my_password..."
      port     = 6
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "823 Koss Road"
        username = "Delaney_Nolan65"
      }
      username = "Arvel69"
    },
  ]
  type = "SQL_SERVER_SHARDED"
}