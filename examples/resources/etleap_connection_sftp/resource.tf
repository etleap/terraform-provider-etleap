resource "etleap_connection_sftp" "my_connectionsftp" {
  base_directory              = "...my_base_directory..."
  deletion_of_export_products = true
  hostname                    = "third-major.name"
  name                        = "Joann Pfeffer"
  password                    = "...my_password..."
  port                        = 10
  type                        = "SFTP"
  username                    = "Tierra.Schmidt"
}