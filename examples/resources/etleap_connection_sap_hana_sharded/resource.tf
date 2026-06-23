resource "etleap_connection_sap_hana_sharded" "my_connectionsap_hana_sharded" {
  cdc_enabled                 = false
  deletion_of_export_products = false
  name                        = "Lee Nolan"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "0616 Bruen Fields"
      database = "...my_database..."
      password = "...my_password..."
      port     = 7
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "751 Deckow Mews"
        port     = 3
        username = "Zachary77"
      }
      username = "Cordelia_Leannon"
    },
  ]
  type = "SAP_HANA_SHARDED"
}