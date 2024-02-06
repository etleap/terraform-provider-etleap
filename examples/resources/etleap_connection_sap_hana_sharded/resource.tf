resource "etleap_connection_sap_hana_sharded" "my_connectionsap_hana_sharded" {
  cdc_enabled                 = true
  deletion_of_export_products = true
  name                        = "Mario Daugherty"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "41858 Price Crescent"
      database = "...my_database..."
      password = "...my_password..."
      port     = 9
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "42997 Johan Pine"
        username = "Tanner.Macejkovic"
      }
      username = "Chesley.Collins"
    },
  ]
  type = "SAP_HANA_SHARDED"
}