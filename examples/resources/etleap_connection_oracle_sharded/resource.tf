resource "etleap_connection_oracle_sharded" "my_connectionoracle_sharded" {
  cdc_enabled                 = false
  deletion_of_export_products = true
  name                        = "Virgil Brown"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "84564 Serenity Park"
      database = "...my_database..."
      password = "...my_password..."
      port     = 10
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "45855 Cruz Mission"
        username = "Lauriane.Kirlin"
      }
      username = "Eleanora.Koss99"
    },
  ]
  type = "ORACLE_SHARDED"
}