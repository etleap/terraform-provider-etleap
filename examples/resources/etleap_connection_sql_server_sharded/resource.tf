resource "etleap_connection_sql_server_sharded" "my_connectionsql_server_sharded" {
  cdc_enabled                 = false
  deletion_of_export_products = false
  name                        = "Dr. Shelia Durgan"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "1178 Ephraim Wall"
      database = "...my_database..."
      password = "...my_password..."
      port     = 6
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "0543 Daugherty River"
        username = "Ricardo_Willms19"
      }
      username = "Stanton39"
    },
  ]
  type = "SQL_SERVER_SHARDED"
}