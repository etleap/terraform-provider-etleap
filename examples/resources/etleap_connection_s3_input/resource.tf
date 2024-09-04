resource "etleap_connection_s3_input" "my_connections3_input" {
  base_directory              = "...my_base_directory..."
  deletion_of_export_products = false
  iam_role                    = "...my_iam_role..."
  input_bucket                = "...my_input_bucket..."
  name                        = "Pamela Barton"
  type                        = "S3_INPUT"
}