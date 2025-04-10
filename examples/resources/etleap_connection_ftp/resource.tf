resource "etleap_connection_ftp" "my_connectionftp" {
  deletion_of_export_products = true
  hostname                    = "aching-base.name"
  name                        = "Sara Smith"
  passive_mode                = false
  password                    = "...my_password..."
  port                        = 1
  type                        = "FTP"
  username                    = "Mia.Harris"
}