resource "volcengine_iam_saml_provider" "foo" {
    encoded_saml_metadata_document = "your document"
    saml_provider_name = "terraform"
    sso_type = 2
    status = 1
}