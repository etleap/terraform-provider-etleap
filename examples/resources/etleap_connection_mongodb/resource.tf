resource "etleap_connection_mongodb" "my_connectionmongodb" {
  auth_database_name          = "...my_auth_database_name..."
  database_name               = "...my_database_name..."
  deletion_of_export_products = true
  name                        = "Ms. Devin Beahan"
  password                    = "...my_password..."
  replica_set = [
    {
      address = "304 Corwin Path"
      port    = 2
    },
  ]
  type     = "MONGODB"
  username = "Kitty.Rath"
  use_ssl  = false
}