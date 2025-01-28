resource "etleap_connection_ftp" "my_connectionftp" {
  deletion_of_export_products = true
  hostname                    = "awkward-coalition.name"
  name                        = "Ronald Kreiger"
  passive_mode                = true
  password                    = "...my_password..."
  port                        = 6
  type                        = "FTP"
  username                    = "Loyce88"
}