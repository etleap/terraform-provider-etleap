resource "etleap_connection_kinesis" "my_connectionkinesis" {
  deletion_of_export_products = false
  iam_role_arn                = "...my_iam_role_arn..."
  name                        = "Katie Robel"
  region                      = "...my_region..."
  type                        = "KINESIS"
}