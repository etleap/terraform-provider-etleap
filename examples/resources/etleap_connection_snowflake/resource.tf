resource "etleap_connection_snowflake" "my_connectionsnowflake" {
  address                     = "2748 Halvorson Knolls"
  database                    = "...my_database..."
  deletion_of_export_products = true
  name                        = "Ms. Gregory Howe"
  password                    = "...my_password..."
  role                        = "...my_role..."
  schema                      = "...my_schema..."
  source_only                 = true
  type                        = "SNOWFLAKE"
  username                    = "Neva39"
  warehouse                   = "...my_warehouse..."
}