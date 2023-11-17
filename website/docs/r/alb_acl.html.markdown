---
subcategory: "ALB"
layout: "volcengine"
page_title: "Volcengine: volcengine_alb_acl"
sidebar_current: "docs-volcengine-resource-alb_acl"
description: |-
  Provides a resource to manage alb acl
---
# volcengine_alb_acl
Provides a resource to manage alb acl
## Example Usage
```hcl
resource "volcengine_alb_acl" "foo" {
  acl_name    = "tf-test-1"
  description = "tftest"
  acl_entries {
    entry       = "172.20.1.0/24"
    description = "e1"
  }

  acl_entries {
    entry       = "172.20.3.0/24"
    description = "e2"
  }
}
```
## Argument Reference
The following arguments are supported:
* `acl_entries` - (Optional) The acl entry set of the Acl.
* `acl_name` - (Optional) The name of Acl.
* `description` - (Optional) The description of the Acl.
* `project_name` - (Optional, ForceNew) The project name of the Acl.

The `acl_entries` object supports the following:

* `entry` - (Required) The content of the AclEntry.
* `description` - (Optional) The description of the AclEntry.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `create_time` - Create time of Acl.


## Import
Acl can be imported using the id, e.g.
```
$ terraform import volcengine_alb_acl.default acl-mizl7m1kqccg5smt1bdpijuj
```

