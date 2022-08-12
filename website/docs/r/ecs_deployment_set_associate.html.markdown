---
subcategory: "ECS"
layout: "volcengine"
page_title: "Volcengine: volcengine_ecs_deployment_set_associate"
sidebar_current: "docs-volcengine-resource-ecs_deployment_set_associate"
description: |-
  Provides a resource to manage ecs deployment set associate
---
# volcengine_ecs_deployment_set_associate
Provides a resource to manage ecs deployment set associate
## Example Usage
```hcl
resource "volcengine_ecs_deployment_set_associate" "default" {
  deployment_set_id = "dps-ybp1b059cb5m57n135g3"
  instance_id       = "i-ybsum2gwr6a8j7j7ak8h"
}
```
## Argument Reference
The following arguments are supported:
* `deployment_set_id` - (Required, ForceNew) The ID of ECS DeploymentSet Associate.
* `instance_id` - (Required, ForceNew) The ID of ECS Instance.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
ECS deployment set associate can be imported using the id, e.g.
```
$ terraform import volcengine_ecs_deployment_set_associate.default dps-ybti5tkpkv2udbfolrft:i-mizl7m1kqccg5smt1bdpijuj
```

