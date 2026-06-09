resource "etleap_connection_postgres_sharded" "my_connectionpostgres_sharded" {
  auto_replicate              = "...my_auto_replicate..."
  cdc_enabled                 = true
  deletion_of_export_products = true
  fetch_lobs_for_updated_rows = true
  name                        = "Paula Ortiz"
  schema                      = "...my_schema..."
  shards = [
    {
      address  = "787 Dominique Centers"
      database = "...my_database..."
      password = "...my_password..."
      port     = 0
      shard_id = "...my_shard_id..."
      ssh_config = {
        address  = "1095 Abernathy Alley"
        port     = 9
        username = "Alycia_Hilll"
      }
      username = "Carlotta.Wilkinson56"
    },
  ]
  type = "POSTGRES_SHARDED"
}