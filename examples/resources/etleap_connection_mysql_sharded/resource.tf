resource "etleap_connection_mysql_sharded" "my_connectionmysql_sharded" {
  auto_replicate                       = "...my_auto_replicate..."
  cdc_enabled                          = true
  database                             = "...my_database..."
  deletion_of_export_products          = true
  name                                 = "Ms. Marsha Krajcik"
  require_ssl_and_validate_certificate = true
  shards = [
    {
      address  = "68333 Simonis Park"
      database = "...my_database..."
      password = "...my_password..."
      port     = 2
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "50930 Feeney Land"
        username = "Ryley_Harber11"
      }
      username = "Karianne_Luettgen"
    },
  ]
  tiny_int1_is_boolean = false
  type                 = "MYSQL_SHARDED"
}