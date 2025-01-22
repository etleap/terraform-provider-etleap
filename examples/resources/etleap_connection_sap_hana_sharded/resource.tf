resource "etleap_connection_sap_hana_sharded" "my_connectionsap_hana_sharded" {
  cdc_enabled                 = false
  deletion_of_export_products = true
  name                        = "Lora Maggio"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "2063 Alba Orchard"
      database = "...my_database..."
      password = "...my_password..."
      port     = 0
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "6880 Dameon Groves"
        username = "Gertrude_Terry"
      }
      username = "Jay_McKenzie"
    },
  ]
  type = "SAP_HANA_SHARDED"
}