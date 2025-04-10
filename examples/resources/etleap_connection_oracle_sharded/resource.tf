resource "etleap_connection_oracle_sharded" "my_connectionoracle_sharded" {
  cdc_enabled                          = true
  certificate                          = "...my_certificate..."
  deletion_of_export_products          = true
  name                                 = "Ms. Leona Cummerata"
  require_ssl_and_validate_certificate = false
  schema                               = "...my_schema..."
  shards = [
    {
      address  = "489 Ernser Square"
      database = "...my_database..."
      password = "...my_password..."
      port     = 4
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "805 Albertha Locks"
        username = "Brennan_Greenfelder86"
      }
      username = "Bernita31"
    },
  ]
  type = "ORACLE_SHARDED"
}