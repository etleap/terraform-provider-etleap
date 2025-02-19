resource "etleap_connection_postgres_sharded" "my_connectionpostgres_sharded" {
  auto_replicate              = "...my_auto_replicate..."
  cdc_enabled                 = false
  deletion_of_export_products = false
  name                        = "Johnny Gusikowski"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "70970 Zemlak Lights"
      database = "...my_database..."
      password = "...my_password..."
      port     = 1
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "7958 Constantin Trace"
        username = "Jalyn93"
      }
      username = "Nella_Pfannerstill22"
    },
  ]
  type = "POSTGRES_SHARDED"
}