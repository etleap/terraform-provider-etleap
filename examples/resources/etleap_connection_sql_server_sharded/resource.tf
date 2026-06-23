resource "etleap_connection_sql_server_sharded" "my_connectionsql_server_sharded" {
  cdc_address                 = "...my_cdc_address..."
  cdc_enabled                 = false
  cdc_port                    = 0
  deletion_of_export_products = false
  name                        = "Madeline Braun"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "706 Monserrate Crest"
      database = "...my_database..."
      password = "...my_password..."
      port     = 2
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "390 Goodwin Grove"
        port     = 6
        username = "Annetta.Crist11"
      }
      username = "Frederick91"
    },
  ]
  type = "SQL_SERVER_SHARDED"
}