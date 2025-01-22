resource "etleap_connection_mongodb" "my_connectionmongodb" {
  auth_database_name          = "...my_auth_database_name..."
  database_name               = "...my_database_name..."
  deletion_of_export_products = false
  name                        = "Juan Grimes"
  password                    = "...my_password..."
  replica_set = [
    {
      address = "037 Adams Inlet"
      port    = 8
    },
  ]
  type     = "MONGODB"
  username = "Asha_Hoppe43"
  use_ssl  = true
}