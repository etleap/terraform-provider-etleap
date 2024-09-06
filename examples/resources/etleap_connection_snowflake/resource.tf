resource "etleap_connection_snowflake" "my_connectionsnowflake" {
  address                     = "095 Brakus Streets"
  database                    = "...my_database..."
  deletion_of_export_products = true
  name                        = "Nathan Hansen"
  password                    = "...my_password..."
  role                        = "...my_role..."
  schema                      = "...my_schema..."
  source_only                 = false
  type                        = "SNOWFLAKE"
  username                    = "Trycia_Walter2"
  warehouse                   = "...my_warehouse..."
}