resource "etleap_connection_sftp" "my_connectionsftp" {
  base_directory              = "...my_base_directory..."
  deletion_of_export_products = false
  hostname                    = "honorable-saving.biz"
  name                        = "Mack Willms"
  password                    = "...my_password..."
  pgp_secret_key              = "...my_pgp_secret_key..."
  port                        = 2
  type                        = "SFTP"
  username                    = "Stanton39"
}