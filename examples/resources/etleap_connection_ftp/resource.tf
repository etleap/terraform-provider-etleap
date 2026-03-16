resource "etleap_connection_ftp" "my_connectionftp" {
  deletion_of_export_products = true
  hostname                    = "weepy-moron.biz"
  name                        = "Ronald Heller"
  passive_mode                = false
  password                    = "...my_password..."
  port                        = 9
  type                        = "FTP"
  username                    = "Hilbert.Veum"
}