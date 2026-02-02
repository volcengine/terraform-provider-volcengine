resource "volcengine_tls_project" "foo" {
  project_name     = "tf-project-m"
  description      = "tf-desc-modify"
  region           = "cn-guilin-boe"
  iam_project_name = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
  tags {
      key   = "k2"
      value = "v3"
    }

}
