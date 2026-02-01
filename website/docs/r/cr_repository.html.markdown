---
subcategory: "CR"
layout: "volcengine"
page_title: "Volcengine: volcengine_cr_repository"
sidebar_current: "docs-volcengine-resource-cr_repository"
description: |-
  Provides a resource to manage cr repository
---
# volcengine_cr_repository
Provides a resource to manage cr repository
## Example Usage
```hcl
resource "volcengine_cr_repository" "foo" {
  registry     = "tf-2"
  namespace    = "namespace-1"
  name         = "repository-1"
  description  = "A test repository created by terraform."
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
```
## Argument Reference
The following arguments are supported:
* `name` - (Required, ForceNew) The name of CrRepository.
* `namespace` - (Required, ForceNew) The target namespace name.
* `registry` - (Required, ForceNew) The CrRegistry name.
* `access_level` - (Optional) The access level of CrRepository.
* `description` - (Optional) The description of CrRepository.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `create_time` - The creation time of repository.
* `update_time` - The last update time of repository.


## Import
CR Repository can be imported using the registry:namespace:name, e.g.
```
$ terraform import volcengine_cr_repository.default cr-basic:namespace-1:repo-1
```

