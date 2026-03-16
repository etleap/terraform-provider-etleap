resource "etleap_connection_snowflake" "my_connectionsnowflake" {
  address                     = "178 Schaden Ridge"
  database                    = "...my_database..."
  deletion_of_export_products = false
  name                        = "Karl Haag"
  password                    = "...my_password..."
  role                        = "...my_role..."
  schema                      = "...my_schema..."
  source_only                 = true
  storage_integration         = "...my_storage_integration..."
  type                        = "SNOWFLAKE"
  username                    = "Thea69"
  warehouse                   = "...my_warehouse..."
}