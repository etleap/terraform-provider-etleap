resource "etleap_connection_mongodb" "my_connectionmongodb" {
  auth_database_name          = "...my_auth_database_name..."
  database_name               = "...my_database_name..."
  deletion_of_export_products = true
  name                        = "Winston Deckow"
  password                    = "...my_password..."
  replica_set = [
    {
      address = "24861 Parisian Divide"
      port    = 2
    },
  ]
  type     = "MONGODB"
  username = "Genesis18"
  use_ssl  = true
}