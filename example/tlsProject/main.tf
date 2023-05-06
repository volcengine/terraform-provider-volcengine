resource "volcengine_tls_project" "foo" {
  project_name = "tf-test"
  description = "tf-desc"
  iam_project_name = "default"
  tags {
    key = "k1"
    value = "v1"
  }
}