resource "etleap_connection_twitter_ads" "my_connectiontwitter_ads" {
  access_token                = "...my_access_token..."
  access_token_secret         = "...my_access_token_secret..."
  app_key                     = "...my_app_key..."
  app_secret_key              = "...my_app_secret_key..."
  deletion_of_export_products = false
  name                        = "Cesar Brakus"
  twitter_usernames           = "...my_twitter_usernames..."
  type                        = "TWITTER_ADS"
}