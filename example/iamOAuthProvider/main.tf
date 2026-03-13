resource "volcengine_iam_oauth_provider" "foo" {
  oauth_provider_name = "acc-test-oauth"
  sso_type            = 2
  status              = 1
  description         = "acc-test"
  client_id           = "test_client_id"
  client_secret       = ""
  user_info_url       = "https://example.com/user_info"
  token_url           = "https://example.com/access_token"
  authorize_url       = "https://example.com/authorize"
  authorize_template  = "$${authEndpoint}?client_id=$${clientId}&scope=$${scope}&response_type=code&state=12345"
  scope               = "openid"
  identity_map_type   = 1
  idp_identity_key    = "username"
}
