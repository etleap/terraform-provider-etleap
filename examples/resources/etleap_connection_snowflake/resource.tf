resource "etleap_connection_snowflake" "my_connectionsnowflake" {
  address                     = "19192 Amina Glens"
  database                    = "...my_database..."
  deletion_of_export_products = true
  name                        = "Mrs. Neal Crist MD"
  password                    = "...my_password..."
  role                        = "...my_role..."
  schema                      = "...my_schema..."
  source_only                 = false
  storage_integration         = "...my_storage_integration..."
  type                        = "SNOWFLAKE"
  username                    = "Raymond.Kunze"
  warehouse                   = "...my_warehouse..."
}