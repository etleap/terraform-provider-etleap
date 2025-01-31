resource "etleap_connection_sap_hana_sharded" "my_connectionsap_hana_sharded" {
  cdc_enabled                 = true
  deletion_of_export_products = false
  name                        = "Ms. Shane Collier III"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "880 Dameon Groves"
      database = "...my_database..."
      password = "...my_password..."
      port     = 4
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "494 Marks Viaduct"
        username = "Antonietta.Kassulke1"
      }
      username = "Mellie_Hoeger70"
    },
  ]
  type = "SAP_HANA_SHARDED"
}