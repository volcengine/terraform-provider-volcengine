resource "volcengine_cr_repository" "foo"{
     registry = "tf-2"
     namespace = "namespace-1"
     name = "repository-1"
     description = "A test repository created by terraform."
     access_level = "Public"
}

# resource "volcengine_cr_repository" "foo1"{
#      registry = "tf-1"
#      namespace = "namespace-2"
#      name = "repository"
#      description = "A test repositoryaaa."
#      access_level = "Private"
# }

# resource "volcengine_cr_repository" "foo2"{
#      registry = "tf-1"
#      namespace = "namespace-3"
#      name = "repository"
#      description = "A test repository."
#      access_level = "Private"
# }