---
subcategory: "ECS"
layout: "volcengine"
page_title: "Volcengine: volcengine_ecs_deployment_sets"
sidebar_current: "docs-volcengine-datasource-ecs_deployment_sets"
description: |-
  Use this data source to query detailed information of ecs deployment sets
---
# volcengine_ecs_deployment_sets
Use this data source to query detailed information of ecs deployment sets
## Example Usage
```hcl
resource "volcengine_ecs_deployment_set" "foo" {
  deployment_set_name = "acc-test-ecs-ds-${count.index}"
  description         = "acc-test"
  granularity         = "switch"
  strategy            = "Availability"
  count               = 3
}

data "volcengine_ecs_deployment_sets" "foo" {
  granularity = "switch"
  ids         = volcengine_ecs_deployment_set.foo[*].id
}
```
## Argument Reference
The following arguments are supported:
* `granularity` - (Optional) The granularity of ECS DeploymentSet.Valid values: switch, host, rack.
* `ids` - (Optional) A list of ECS DeploymentSet IDs.
* `name_regex` - (Optional) A Name Regex of ECS DeploymentSet.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `deployment_sets` - The collection of ECS DeploymentSet query.
    * `deployment_set_id` - The ID of ECS DeploymentSet.
    * `deployment_set_name` - The name of ECS DeploymentSet.
    * `description` - The description of ECS DeploymentSet.
    * `granularity` - The granularity of ECS DeploymentSet.
    * `strategy` - The strategy of ECS DeploymentSet.
* `total_count` - The total count of ECS DeploymentSet query.


