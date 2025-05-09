resource "etleap_connection_mongodb" "my_connectionmongodb" {
  auth_database_name          = "...my_auth_database_name..."
  database_name               = "...my_database_name..."
  deletion_of_export_products = false
  name                        = "Dolores Funk"
  password                    = "...my_password..."
  replica_set = [
    {
      address = "10576 Dickinson Trail"
      port    = 7
    },
  ]
  type     = "MONGODB"
  username = "Vicente.Kutch"
  use_ssl  = true
}