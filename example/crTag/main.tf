# Tag cannot be created,please import by command `terraform import volcengine_cr_tag.default registry:namespace:repository:tag`
resource "volcengine_cr_tag" "default" {
  registry   = "enterprise-1"
  namespace  = "langyu"
  repository = "repo"
  name       = "v2"
}