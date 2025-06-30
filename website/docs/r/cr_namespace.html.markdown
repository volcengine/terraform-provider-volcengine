---
subcategory: "CR"
layout: "volcengine"
page_title: "Volcengine: volcengine_cr_namespace"
sidebar_current: "docs-volcengine-resource-cr_namespace"
description: |-
  Provides a resource to manage cr namespace
---
# volcengine_cr_namespace
Provides a resource to manage cr namespace
## Example Usage
```hcl
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
```
## Argument Reference
The following arguments are supported:
* `name` - (Required, ForceNew) The name of CrNamespace.
* `registry` - (Required, ForceNew) The registry name.
* `project` - (Optional) The ProjectName of the CrNamespace.
* `repository_default_access_level` - (Optional, ForceNew) The default access level of repository. Valid values: `Private`, `Public`. Default is `Private`.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `create_time` - The time when namespace created.


## Import
CR namespace can be imported using the registry:name, e.g.
```
$ terraform import volcengine_cr_namespace.default cr-basic:namespace-1
```

