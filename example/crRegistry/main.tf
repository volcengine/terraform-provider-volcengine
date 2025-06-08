# create cr registry
resource "volcengine_cr_registry" "foo" {
  name               = "acc-test-cr"
  delete_immediately = false
  password           = "1qaz!QAZ"
  project            = "default"
}

# create cr namespace
resource "volcengine_cr_namespace" "foo" {
  registry = volcengine_cr_registry.foo.id
  name     = "acc-test-namespace"
  project  = "default"
}

# create cr repository
resource "volcengine_cr_repository" "foo" {
  registry     = volcengine_cr_registry.foo.id
  namespace    = volcengine_cr_namespace.foo.name
  name         = "acc-test-repository"
  description  = "A test repository created by terraform."
  access_level = "Public"
}
