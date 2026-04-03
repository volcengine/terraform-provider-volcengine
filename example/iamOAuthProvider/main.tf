resource "volcengine_iam_oauth_provider" "foo" {
  oauth_provider_name = "acc-test-oauth"
  sso_type            = 2
  status              = 2
  description         = "acc-test-modify"
  client_id           = "test_client_id_1"
  client_secret       = ""
  user_info_url       = "https://example.com/user_info_modify"
  token_url           = "https://example.com/access_token_modify"
  authorize_url       = "https://example.com/authorize_modify"
  authorize_template  = "$${authEndpoint}?client_id=$${clientId}&scope=$${scope}&response_type=code&state=123456"
  scope               = "openid"
  identity_map_type   = 2
  idp_identity_key    = "username-1"
}
