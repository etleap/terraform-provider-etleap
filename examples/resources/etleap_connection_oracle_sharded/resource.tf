resource "etleap_connection_oracle_sharded" "my_connectionoracle_sharded" {
  cdc_enabled                          = false
  certificate                          = "...my_certificate..."
  deletion_of_export_products          = true
  name                                 = "Irving Lindgren I"
  require_ssl_and_validate_certificate = false
  schema                               = "...my_schema..."
  shards = [
    {
      address  = "4642 Mustafa Forge"
      database = "...my_database..."
      password = "...my_password..."
      port     = 10
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "6097 Damaris Orchard"
        username = "Ines2"
      }
      username = "Eladio_Ankunding38"
    },
  ]
  type = "ORACLE_SHARDED"
}