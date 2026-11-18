resource "etleap_connection_sap_hana_sharded" "my_connectionsap_hana_sharded" {
  cdc_enabled                 = false
  deletion_of_export_products = false
  name                        = "Peter Mraz"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "126 Renner Motorway"
      database = "...my_database..."
      password = "...my_password..."
      port     = 1
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "1539 Beer Shoals"
        port     = 2
        username = "Josiane.Kessler74"
      }
      username = "Corbin_Bergnaum59"
    },
  ]
  type = "SAP_HANA_SHARDED"
}