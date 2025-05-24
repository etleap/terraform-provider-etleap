resource "etleap_connection_sqs" "my_connectionsqs" {
  deletion_of_export_products = true
  iam_role_arn                = "...my_iam_role_arn..."
  name                        = "Enrique Crist"
  region                      = "...my_region..."
  type                        = "SQS"
}