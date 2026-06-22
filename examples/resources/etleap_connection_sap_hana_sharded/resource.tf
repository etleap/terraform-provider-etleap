resource "etleap_connection_sap_hana_sharded" "my_connectionsap_hana_sharded" {
  cdc_enabled                 = false
  deletion_of_export_products = true
  name                        = "Mrs. Maryann Ledner III"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "57823 Koss Road"
      database = "...my_database..."
      password = "...my_password..."
      port     = 2
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "8560 Cummerata Radial"
        port     = 9
        username = "Bessie_Franey72"
      }
      username = "Kayden.Daugherty"
    },
  ]
  type = "SAP_HANA_SHARDED"
}