resource "etleap_connection_mongodb" "my_connectionmongodb" {
  auth_database_name          = "...my_auth_database_name..."
  database_name               = "...my_database_name..."
  deletion_of_export_products = false
  name                        = "Karla Leannon"
  password                    = "...my_password..."
  replica_set = [
    {
      address = "33233 Betsy Burg"
      port    = 6
    },
  ]
  type     = "MONGODB"
  username = "Merlin.Murazik88"
  use_ssl  = false
}