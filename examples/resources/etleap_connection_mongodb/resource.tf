resource "etleap_connection_mongodb" "my_connectionmongodb" {
  auth_database_name          = "...my_auth_database_name..."
  database_name               = "...my_database_name..."
  deletion_of_export_products = true
  name                        = "Mrs. Julia Heaney"
  password                    = "...my_password..."
  replica_set = [
    {
      address = "20066 Forrest Port"
      port    = 0
    },
  ]
  type     = "MONGODB"
  username = "Hyman_Purdy"
  use_ssl  = true
}