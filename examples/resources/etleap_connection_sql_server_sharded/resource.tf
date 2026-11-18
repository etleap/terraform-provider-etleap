resource "etleap_connection_sql_server_sharded" "my_connectionsql_server_sharded" {
  cdc_address                 = "...my_cdc_address..."
  cdc_enabled                 = false
  cdc_port                    = 2
  deletion_of_export_products = true
  name                        = "Dixie Bashirian"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "6017 Bertha Haven"
      database = "...my_database..."
      password = "...my_password..."
      port     = 3
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "95892 Witting Gateway"
        port     = 4
        username = "Monty_Wolff"
      }
      username = "Ocie.Bailey"
    },
  ]
  type = "SQL_SERVER_SHARDED"
}