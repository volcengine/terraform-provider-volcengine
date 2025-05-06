resource "volcengine_cr_namespace" "foo" {
  registry = "tf-test-cr"
  name     = "test-namespace-1"
  project  = "default"
}

resource "volcengine_cr_namespace" "foo1" {
  registry = "tf-test-cr"
  name     = "test-namespace-2"
  project  = "default"
}
