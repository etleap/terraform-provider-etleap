resource "etleap_connection_oracle_sharded" "my_connectionoracle_sharded" {
  cdc_enabled                          = true
  certificate                          = "...my_certificate..."
  deletion_of_export_products          = false
  name                                 = "Brian Wolff"
  require_ssl_and_validate_certificate = true
  schema                               = "...my_schema..."
  shards = [
    {
      address  = "170 Schmidt Alley"
      database = "...my_database..."
      password = "...my_password..."
      port     = 8
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "415 Tromp Overpass"
        port     = 9
        username = "Charlene.DuBuque96"
      }
      username = "Clara_Watsica83"
    },
  ]
  type = "ORACLE_SHARDED"
}