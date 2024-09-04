resource "etleap_connection_sap_hana_sharded" "my_connectionsap_hana_sharded" {
  cdc_enabled                 = false
  deletion_of_export_products = true
  name                        = "Grace Kshlerin"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "8561 Block Key"
      database = "...my_database..."
      password = "...my_password..."
      port     = 3
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "974 Nienow Unions"
        username = "Keeley18"
      }
      username = "Brigitte_Nader"
    },
  ]
  type = "SAP_HANA_SHARDED"
}