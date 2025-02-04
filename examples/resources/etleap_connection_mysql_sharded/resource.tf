resource "etleap_connection_mysql_sharded" "my_connectionmysql_sharded" {
  auto_replicate                       = "...my_auto_replicate..."
  cdc_enabled                          = false
  certificate                          = "...my_certificate..."
  database                             = "...my_database..."
  deletion_of_export_products          = true
  name                                 = "Fannie Sipes"
  require_ssl_and_validate_certificate = false
  shards = [
    {
      address  = "8523 Ankunding Ways"
      database = "...my_database..."
      password = "...my_password..."
      port     = 3
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "52483 Yost Corners"
        username = "Karianne_Luettgen"
      }
      username = "Marlee.Towne"
    },
  ]
  tiny_int1_is_boolean = true
  type                 = "MYSQL_SHARDED"
}