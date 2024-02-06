resource "etleap_connection_mongodb" "my_connectionmongodb" {
  auth_database_name          = "...my_auth_database_name..."
  database_name               = "...my_database_name..."
  deletion_of_export_products = true
  name                        = "Chelsea Casper"
  password                    = "...my_password..."
  replica_set = [
    {
      address = "148 Katrina Vista"
      port    = 2
    },
  ]
  type     = "MONGODB"
  username = "Francis29"
  use_ssl  = false
}