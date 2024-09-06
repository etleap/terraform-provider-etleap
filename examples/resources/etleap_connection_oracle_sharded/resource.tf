resource "etleap_connection_oracle_sharded" "my_connectionoracle_sharded" {
  cdc_enabled                 = true
  deletion_of_export_products = true
  name                        = "Ella Schowalter"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "57523 Daisha Parks"
      database = "...my_database..."
      password = "...my_password..."
      port     = 5
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "33629 Kutch Freeway"
        username = "Marisol_Cormier11"
      }
      username = "Ralph80"
    },
  ]
  type = "ORACLE_SHARDED"
}