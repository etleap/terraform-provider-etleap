resource "etleap_connection_snowflake" "my_connectionsnowflake" {
  address                     = "0702 Koelpin Park"
  database                    = "...my_database..."
  deletion_of_export_products = false
  name                        = "Lewis Hahn"
  password                    = "...my_password..."
  role                        = "...my_role..."
  schema                      = "...my_schema..."
  source_only                 = false
  type                        = "SNOWFLAKE"
  username                    = "Wilfred_Streich96"
  warehouse                   = "...my_warehouse..."
}