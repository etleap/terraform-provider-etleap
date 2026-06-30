resource "etleap_connection_oracle_sharded" "my_connectionoracle_sharded" {
  cdc_address                          = "...my_cdc_address..."
  cdc_enabled                          = true
  cdc_port                             = 3
  certificate                          = "...my_certificate..."
  deletion_of_export_products          = true
  name                                 = "Dr. Jimmy Rice"
  require_ssl_and_validate_certificate = true
  schema                               = "...my_schema..."
  shards = [
    {
      address  = "941 Schaefer Underpass"
      database = "...my_database..."
      password = "...my_password..."
      port     = 6
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "128 Vincenzo Divide"
        port     = 10
        username = "Nella_Pfannerstill22"
      }
      username = "Danny.Bosco24"
    },
  ]
  type = "ORACLE_SHARDED"
}