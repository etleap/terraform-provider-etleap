resource "etleap_connection_db2" "my_connectiondb2" {
  address                     = "44538 Victor Mill"
  certificate                 = "...my_certificate..."
  database                    = "...my_database..."
  deletion_of_export_products = true
  name                        = "Katherine Altenwerth"
  password                    = "...my_password..."
  port                        = 7
  schema                      = "...my_schema..."
  type                        = "DB2"
  username                    = "Oceane_Spencer"
}