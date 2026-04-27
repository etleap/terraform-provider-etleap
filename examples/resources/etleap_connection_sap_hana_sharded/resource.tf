resource "etleap_connection_sap_hana_sharded" "my_connectionsap_hana_sharded" {
  cdc_enabled                 = false
  deletion_of_export_products = false
  name                        = "Karla West"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "57169 Terry Run"
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