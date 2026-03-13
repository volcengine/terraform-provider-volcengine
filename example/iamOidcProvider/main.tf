resource "volcengine_iam_oidc_provider" "foo" {
  oidc_provider_name  = "oidc_provider"
  issuer_url          = "https://security-api.snssdk.com/qa/sso/oidc/6c505fb67d32417c8de287ee1fa89fc1"
  description         = "acc-test-oidc"
  issuance_limit_time = 6
  client_ids          = ["6c505fb67d32417c8de287ee1fa89fc1"]
  thumbprints         = ["9b1afaa2dfca349fe38c5ef3e72ee03cb0696d65ea2e11f597ea9aa55fcff33a"]
}
