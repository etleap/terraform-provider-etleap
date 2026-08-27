resource "etleap_connection_snowflake" "my_connectionsnowflake" {
  address                     = "84306 Fahey Neck"
  database                    = "...my_database..."
  deletion_of_export_products = false
  name                        = "Cesar Brakus"
  password                    = "...my_password..."
  role                        = "...my_role..."
  schema                      = "...my_schema..."
  source_only                 = false
  storage_integration         = "...my_storage_integration..."
  type                        = "SNOWFLAKE"
  username                    = "Lenna.Grady"
  warehouse                   = "...my_warehouse..."
}