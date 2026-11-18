resource "etleap_connection_oracle_sharded" "my_connectionoracle_sharded" {
  cdc_address                          = "...my_cdc_address..."
  cdc_enabled                          = true
  cdc_port                             = 0
  certificate                          = "...my_certificate..."
  deletion_of_export_products          = false
  name                                 = "Ted Dare"
  require_ssl_and_validate_certificate = true
  schema                               = "...my_schema..."
  shards = [
    {
      address  = "8128 Vincenzo Divide"
      database = "...my_database..."
      password = "...my_password..."
      port     = 10
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "689 Danny Cliff"
        port     = 8
        username = "Dejah51"
      }
      username = "Jeanie64"
    },
  ]
  type = "ORACLE_SHARDED"
}