resource "etleap_connection_sql_server_sharded" "my_connectionsql_server_sharded" {
  cdc_enabled                 = false
  deletion_of_export_products = true
  name                        = "Terry Moen"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "58219 Schiller Mills"
      database = "...my_database..."
      password = "...my_password..."
      port     = 0
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "3926 Franey Tunnel"
        username = "Jaqueline.Morissette24"
      }
      username = "Brando.Volkman0"
    },
  ]
  type = "SQL_SERVER_SHARDED"
}