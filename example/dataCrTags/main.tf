data "volcengine_cr_tags" "foo"{
     registry = "enterprise-1"
     namespace = "test"
     repository = "repo"
     types = ["Image"]
}