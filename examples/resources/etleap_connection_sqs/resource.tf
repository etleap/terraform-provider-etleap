resource "etleap_connection_sqs" "my_connectionsqs" {
  deletion_of_export_products = false
  iam_role_arn                = "...my_iam_role_arn..."
  name                        = "Irvin Friesen"
  region                      = "...my_region..."
  type                        = "SQS"
}