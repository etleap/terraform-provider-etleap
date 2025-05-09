resource "etleap_connection_sap_hana_sharded" "my_connectionsap_hana_sharded" {
  cdc_enabled                 = true
  deletion_of_export_products = true
  name                        = "Faith Kris"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "047 Adolph Rapid"
      database = "...my_database..."
      password = "...my_password..."
      port     = 4
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "576 Skiles Center"
        username = "Turner.Leffler85"
      }
      username = "Kadin32"
    },
  ]
  type = "SAP_HANA_SHARDED"
}