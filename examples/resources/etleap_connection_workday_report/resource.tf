resource "etleap_connection_workday_report" "my_connectionworkday_report" {
  deletion_of_export_products = false
  name                        = "Maggie Wuckert"
  password                    = "...my_password..."
  report_url                  = "...my_report_url..."
  type                        = "WORKDAY_REPORT"
  username                    = "Jaqueline.Morissette24"
}