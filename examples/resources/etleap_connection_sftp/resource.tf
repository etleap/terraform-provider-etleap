resource "etleap_connection_sftp" "my_connectionsftp" {
  base_directory              = "...my_base_directory..."
  deletion_of_export_products = false
  hostname                    = "overdue-doggie.net"
  name                        = "Jimmie Metz"
  password                    = "...my_password..."
  pgp_secret_key              = "...my_pgp_secret_key..."
  port                        = 9
  type                        = "SFTP"
  username                    = "Angelica_Dietrich15"
}