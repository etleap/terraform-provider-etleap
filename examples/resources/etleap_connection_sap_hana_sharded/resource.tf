resource "etleap_connection_sap_hana_sharded" "my_connectionsap_hana_sharded" {
  cdc_enabled                 = true
  deletion_of_export_products = false
  name                        = "Sadie King"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "46503 Ivah Crossing"
      database = "...my_database..."
      password = "...my_password..."
      port     = 6
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "823 Koss Road"
        port     = 2
        username = "Lukas_Skiles"
      }
      username = "Arvel69"
    },
  ]
  type = "SAP_HANA_SHARDED"
}