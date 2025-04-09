resource "etleap_connection_db2_sharded" "my_connectiondb2_sharded" {
  certificate                 = "...my_certificate..."
  deletion_of_export_products = true
  name                        = "Cary Monahan"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "480 Isaiah Park"
      database = "...my_database..."
      password = "...my_password..."
      port     = 4
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "05761 Schaefer Place"
        username = "Junius.Treutel71"
      }
      username = "Boris_Stoltenberg"
    },
  ]
  type = "DB2_SHARDED"
}