resource "volcengine_iam_oidc_provider_client" "foo" {
  oidc_provider_name = "oidc_provider"
  client_id          = "test_client_id_2"
}
