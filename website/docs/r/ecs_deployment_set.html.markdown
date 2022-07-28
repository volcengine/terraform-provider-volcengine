---
subcategory: "ECS"
layout: "volcengine"
page_title: "Volcengine: volcengine_ecs_deployment_set"
sidebar_current: "docs-volcengine-resource-ecs_deployment_set"
description: |-
  Provides a resource to manage ecs deployment set
---
# volcengine_ecs_deployment_set
Provides a resource to manage ecs deployment set
## Example Usage
```hcl
resource "volcengine_ecs_deployment_set" "default" {
  deployment_set_name = "tf-test"
  description         = "test1"
  granularity         = "host"
}
```
## Argument Reference
The following arguments are supported:
* `deployment_set_name` - (Required) The name of ECS DeploymentSet.
* `description` - (Optional) The description of ECS DeploymentSet.
* `granularity` - (Optional, ForceNew) The granularity of ECS DeploymentSet.Valid values: switch, host, rack,Default is host.
* `strategy` - (Optional, ForceNew) The strategy of ECS DeploymentSet.Valid values: Availability.Default is host.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `deployment_set_id` - The ID of ECS DeploymentSet.


## Import
ECS deployment set can be imported using the id, e.g.
```
$ terraform import volcengine_ecs_deployment_set.default i-mizl7m1kqccg5smt1bdpijuj
```

