resource "etleap_connection_mongodb" "my_connectionmongodb" {
  auth_database_name          = "...my_auth_database_name..."
  database_name               = "...my_database_name..."
  deletion_of_export_products = true
  name                        = "Ms. Melba Hamill"
  password                    = "...my_password..."
  replica_set = [
    {
      address = "8379 Karianne Motorway"
      port    = 4
    },
  ]
  type     = "MONGODB"
  username = "Kyler.Paucek17"
  use_ssl  = true
}