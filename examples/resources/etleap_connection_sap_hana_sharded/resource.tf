resource "etleap_connection_sap_hana_sharded" "my_connectionsap_hana_sharded" {
  cdc_enabled                 = true
  deletion_of_export_products = true
  name                        = "Johnnie Zboncak"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "932 Myriam Key"
      database = "...my_database..."
      password = "...my_password..."
      port     = 7
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "1497 Josh Ford"
        username = "Brennon_Botsford"
      }
      username = "Camren.Williamson65"
    },
  ]
  type = "SAP_HANA_SHARDED"
}