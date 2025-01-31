resource "etleap_connection_sftp" "my_connectionsftp" {
  base_directory              = "...my_base_directory..."
  deletion_of_export_products = false
  hostname                    = "starry-metallurgist.com"
  name                        = "Tonya Schaden"
  password                    = "...my_password..."
  port                        = 0
  type                        = "SFTP"
  username                    = "Helena71"
}