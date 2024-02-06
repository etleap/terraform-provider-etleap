resource "etleap_connection_snowflake" "my_connectionsnowflake" {
  address                     = "688 Turner Meadow"
  database                    = "...my_database..."
  deletion_of_export_products = false
  name                        = "Antonia Lesch"
  password                    = "...my_password..."
  role                        = "...my_role..."
  schema                      = "...my_schema..."
  source_only                 = false
  type                        = "SNOWFLAKE"
  username                    = "Martine94"
  warehouse                   = "...my_warehouse..."
}