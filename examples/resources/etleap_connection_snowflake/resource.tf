resource "etleap_connection_snowflake" "my_connectionsnowflake" {
  address                     = "2691 West Well"
  database                    = "...my_database..."
  deletion_of_export_products = true
  name                        = "Cassandra Quitzon"
  password                    = "...my_password..."
  role                        = "...my_role..."
  schema                      = "...my_schema..."
  source_only                 = true
  type                        = "SNOWFLAKE"
  username                    = "Treva_Terry98"
  warehouse                   = "...my_warehouse..."
}