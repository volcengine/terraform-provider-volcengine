data "volcengine_cr_repositories" "foo"{
     registry = "tf-1"
     # access_levels = ["Private"]
     # namespaces = ["namespace*"]
     names = ["repo*"]
}