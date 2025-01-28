resource "etleap_connection_sql_server_sharded" "my_connectionsql_server_sharded" {
  cdc_enabled                 = true
  deletion_of_export_products = true
  name                        = "Tim Zemlak DDS"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "54917 Aniyah Springs"
      database = "...my_database..."
      password = "...my_password..."
      port     = 9
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "1178 Ephraim Wall"
        username = "Kevon57"
      }
      username = "Guido.Hudson73"
    },
  ]
  type = "SQL_SERVER_SHARDED"
}