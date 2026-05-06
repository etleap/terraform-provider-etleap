resource "etleap_connection_sap_hana_sharded" "my_connectionsap_hana_sharded" {
  cdc_enabled                 = true
  deletion_of_export_products = true
  name                        = "Emmett Osinski"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "169 Terry Run"
      database = "...my_database..."
      password = "...my_password..."
      port     = 7
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "562 King Tunnel"
        port     = 8
        username = "Emery_Kihn7"
      }
      username = "Gerry.Boyle"
    },
  ]
  type = "SAP_HANA_SHARDED"
}