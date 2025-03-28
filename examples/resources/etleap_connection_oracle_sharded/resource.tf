resource "etleap_connection_oracle_sharded" "my_connectionoracle_sharded" {
  cdc_enabled                          = false
  certificate                          = "...my_certificate..."
  deletion_of_export_products          = true
  name                                 = "Leigh Metz"
  require_ssl_and_validate_certificate = true
  schema                               = "...my_schema..."
  shards = [
    {
      address  = "8550 Kathryn Overpass"
      database = "...my_database..."
      password = "...my_password..."
      port     = 4
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "4247 Yvonne Neck"
        username = "Levi_Boyer14"
      }
      username = "Damaris_McDermott"
    },
  ]
  type = "ORACLE_SHARDED"
}