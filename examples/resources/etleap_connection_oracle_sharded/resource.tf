resource "etleap_connection_oracle_sharded" "my_connectionoracle_sharded" {
  cdc_address                          = "...my_cdc_address..."
  cdc_enabled                          = true
  cdc_port                             = 3
  certificate                          = "...my_certificate..."
  deletion_of_export_products          = false
  name                                 = "Shawna Dickinson"
  require_ssl_and_validate_certificate = true
  schema                               = "...my_schema..."
  shards = [
    {
      address  = "70970 Zemlak Lights"
      database = "...my_database..."
      password = "...my_password..."
      port     = 1
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "7958 Constantin Trace"
        port     = 5
        username = "Vincenzo_Dickens68"
      }
      username = "Ricky22"
    },
  ]
  type = "ORACLE_SHARDED"
}