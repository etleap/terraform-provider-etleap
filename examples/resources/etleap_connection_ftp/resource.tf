resource "etleap_connection_ftp" "my_connectionftp" {
  deletion_of_export_products = true
  hostname                    = "flimsy-tank-top.org"
  name                        = "Leroy Paucek"
  passive_mode                = true
  password                    = "...my_password..."
  port                        = 10
  type                        = "FTP"
  username                    = "Deshawn_McLaughlin78"
}