resource "volcengine_tls_project" "foo" {
  project_name     = "tf-project-m"
  description      = "tf-desc"
  region           = "cn-guilin-boe"
  iam_project_name = "default"
  tags {
    key   = "k2m"
    value = "v1"
  }
  tags {
      key   = "kt3"
      value = "v3"
    }

}
