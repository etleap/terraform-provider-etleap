resource "etleap_connection_sql_server_sharded" "my_connectionsql_server_sharded" {
  cdc_enabled                 = false
  deletion_of_export_products = false
  name                        = "Stacy Monahan V"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "759 Roel Club"
      database = "...my_database..."
      password = "...my_password..."
      port     = 7
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "43770 Jacobi Ridges"
        port     = 1
        username = "Dejon_Steuber"
      }
      username = "Alfredo34"
    },
  ]
  type = "SQL_SERVER_SHARDED"
}