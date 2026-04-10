---
subcategory: "ECS"
layout: "volcengine"
page_title: "Volcengine: volcengine_ecs_hpc_cluster"
sidebar_current: "docs-volcengine-resource-ecs_hpc_cluster"
description: |-
  Provides a resource to manage ecs hpc cluster
---
**❗Notice:**
The current provider is no longer being maintained. We recommend that you use the [volcenginecc](https://registry.terraform.io/providers/volcengine/volcenginecc/latest/docs) instead.
# volcengine_ecs_hpc_cluster
Provides a resource to manage ecs hpc cluster
## Example Usage
```hcl
resource "volcengine_ecs_hpc_cluster" "foo" {
  zone_id      = "cn-beijing-b"
  name         = "acc-test-hpc-cluster"
  description  = "acc-test"
  project_name = "default"
  tags {
    key   = "tfk1"
    value = "tfv1"
  }
}
```
## Argument Reference
The following arguments are supported:
* `name` - (Required) The name of the hpc cluster.
* `zone_id` - (Required, ForceNew) The zone id of the hpc cluster.
* `description` - (Optional) The description of the hpc cluster.
* `project_name` - (Optional) The project name of the hpc cluster.
* `tags` - (Optional) Tags.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
EcsHpcCluster can be imported using the id, e.g.
```
$ terraform import volcengine_ecs_hpc_cluster.default resource_id
```

