resource "etleap_connection_oracle_sharded" "my_connectionoracle_sharded" {
  cdc_enabled                 = true
  deletion_of_export_products = false
  name                        = "Mr. Brooke Schamberger"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "730 Rahul Crescent"
      database = "...my_database..."
      password = "...my_password..."
      port     = 6
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "57523 Daisha Parks"
        username = "Julie.Goodwin"
      }
      username = "Margarita_Emard49"
    },
  ]
  type = "ORACLE_SHARDED"
}