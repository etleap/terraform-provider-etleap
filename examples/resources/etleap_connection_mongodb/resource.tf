resource "etleap_connection_mongodb" "my_connectionmongodb" {
  auth_database_name          = "...my_auth_database_name..."
  database_name               = "...my_database_name..."
  deletion_of_export_products = true
  name                        = "Betty Johnson"
  password                    = "...my_password..."
  replica_set = [
    {
      address = "1575 Geoffrey Corners"
      port    = 2
    },
  ]
  type     = "MONGODB"
  username = "Leann35"
  use_ssl  = false
}