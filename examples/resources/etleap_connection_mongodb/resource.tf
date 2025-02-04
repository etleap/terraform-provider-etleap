resource "etleap_connection_mongodb" "my_connectionmongodb" {
  auth_database_name          = "...my_auth_database_name..."
  database_name               = "...my_database_name..."
  deletion_of_export_products = true
  name                        = "Michael Jacobi III"
  password                    = "...my_password..."
  replica_set = [
    {
      address = "544 Emerson Greens"
      port    = 2
    },
  ]
  type     = "MONGODB"
  username = "Ernestine11"
  use_ssl  = false
}