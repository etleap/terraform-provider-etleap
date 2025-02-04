resource "etleap_connection_ftp" "my_connectionftp" {
  deletion_of_export_products = false
  hostname                    = "political-narrative.name"
  name                        = "Pamela Torphy"
  passive_mode                = true
  password                    = "...my_password..."
  port                        = 1
  type                        = "FTP"
  username                    = "Pansy_Auer3"
}