resource "etleap_connection_user_defined_api" "my_connectionuser_defined_api" {
  authentication = {
    basic = {
      password = "...my_password..."
      type     = "BASIC"
      username = "Jaqueline.Morissette24"
    }
  }
  deletion_of_export_products = false
  entities = [
    {
      api_url = "...my_api_url..."
      columns = [
        "{ \"see\": \"documentation\" }",
      ]
      display_name = "...my_display_name..."
      header_parameters = [
        {
          key   = "...my_key..."
          value = "...my_value..."
        },
      ]
      id = "ee605990-9fc8-4f1d-9a17-4cc096b23d6f"
      paging_strategy = {
        cursor_uri = {
          max_page_size        = 0
          page_size_field_name = "...my_page_size_field_name..."
          path_to_cursor       = "...my_path_to_cursor..."
          type                 = "CURSOR_URI"
          url_prefix           = "...my_url_prefix..."
        }
      }
      path_to_results = "...my_path_to_results..."
      pipeline_mode = {
        one = "REPLACE"
      }
      query_parameters = [
        {
          key   = "...my_key..."
          value = "...my_value..."
        },
      ]
      rest_method = {
        one = "GET"
      }
    },
  ]
  name = "Ricky Hegmann Sr."
  type = "USER_DEFINED_API"
}