resource "etleap_connection_sql_server_sharded" "my_connectionsql_server_sharded" {
  cdc_address                 = "...my_cdc_address..."
  cdc_enabled                 = false
  cdc_port                    = 3
  deletion_of_export_products = true
  name                        = "Gary Morar"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "28390 Goodwin Grove"
      database = "...my_database..."
      password = "...my_password..."
      port     = 6
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "17113 Concepcion Valleys"
        port     = 5
        username = "Stanford75"
      }
      username = "Walter_Hammes"
    },
  ]
  type = "SQL_SERVER_SHARDED"
}