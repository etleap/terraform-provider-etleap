resource "etleap_connection_sql_server_sharded" "my_connectionsql_server_sharded" {
  cdc_enabled                 = true
  deletion_of_export_products = false
  name                        = "Bernard Herman"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "219 Schiller Mills"
      database = "...my_database..."
      password = "...my_password..."
      port     = 0
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "3926 Franey Tunnel"
        port     = 5
        username = "Lew.Tromp"
      }
      username = "Brando.Volkman0"
    },
  ]
  type = "SQL_SERVER_SHARDED"
}