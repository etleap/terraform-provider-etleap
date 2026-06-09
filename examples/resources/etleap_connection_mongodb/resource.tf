resource "etleap_connection_mongodb" "my_connectionmongodb" {
  auth_database_name          = "...my_auth_database_name..."
  database_name               = "...my_database_name..."
  deletion_of_export_products = false
  name                        = "Wanda Kertzmann DDS"
  password                    = "...my_password..."
  replica_set = [
    {
      address = "015 Rolfson Green"
      port    = 0
    },
  ]
  type     = "MONGODB"
  username = "Hazle62"
  use_ssl  = true
}