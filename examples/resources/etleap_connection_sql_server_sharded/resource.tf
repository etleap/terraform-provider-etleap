resource "etleap_connection_sql_server_sharded" "my_connectionsql_server_sharded" {
  cdc_enabled                 = true
  deletion_of_export_products = true
  name                        = "Dr. Raul Flatley"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "0366 Lea Wall"
      database = "...my_database..."
      password = "...my_password..."
      port     = 8
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "9081 Braun Light"
        port     = 3
        username = "Nettie61"
      }
      username = "Granville24"
    },
  ]
  type = "SQL_SERVER_SHARDED"
}