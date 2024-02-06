resource "etleap_connection_sftp" "my_connectionsftp" {
  base_directory              = "...my_base_directory..."
  deletion_of_export_products = false
  hostname                    = "potable-gather.com"
  name                        = "Mae Bartell"
  password                    = "...my_password..."
  port                        = 9
  type                        = "SFTP"
  username                    = "Richmond22"
}