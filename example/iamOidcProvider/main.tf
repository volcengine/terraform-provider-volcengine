resource "volcengine_iam_oidc_provider" "foo" {
  oidc_provider_name  = "oidc_provider"
  issuer_url          = "test-issuer-url"
  description         = "acc-test-oidc"
  issuance_limit_time = 6
  client_ids          = ["test-client-id-1", "test-client-id-2"]
  thumbprints         = ["test-thumbprint-1", "test-thumbprint-2"]
}
