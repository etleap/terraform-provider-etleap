resource "etleap_connection_dayforce" "my_connectiondayforce" {
  company_id                  = "...my_company_id..."
  deletion_of_export_products = false
  domain                      = "...my_domain..."
  is_test_environment         = "...my_is_test_environment..."
  name                        = "Ms. Jasmine Windler"
  password                    = "...my_password..."
  type                        = "DAYFORCE"
  username                    = "Herbert14"
}