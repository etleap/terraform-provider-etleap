resource "etleap_connection_oracle_sharded" "my_connectionoracle_sharded" {
  cdc_enabled                          = true
  certificate                          = "...my_certificate..."
  deletion_of_export_products          = false
  name                                 = "Merle Kohler"
  require_ssl_and_validate_certificate = true
  schema                               = "...my_schema..."
  shards = [
    {
      address  = "099 Sarah Pines"
      database = "...my_database..."
      password = "...my_password..."
      port     = 9
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "5338 Wiley Drive"
        port     = 3
        username = "Lavina73"
      }
      username = "Andy.Vandervort79"
    },
  ]
  type = "ORACLE_SHARDED"
}