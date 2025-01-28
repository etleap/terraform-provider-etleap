resource "etleap_connection_postgres_sharded" "my_connectionpostgres_sharded" {
  auto_replicate              = "...my_auto_replicate..."
  cdc_enabled                 = false
  deletion_of_export_products = true
  name                        = "Lynette Nitzsche"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "3380 Dickinson Fort"
      database = "...my_database..."
      password = "...my_password..."
      port     = 6
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "170 Schmidt Alley"
        username = "Pattie.Zemlak"
      }
      username = "Jon_Schaefer89"
    },
  ]
  type = "POSTGRES_SHARDED"
}