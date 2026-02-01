---
subcategory: "VPC"
layout: "volcengine"
page_title: "Volcengine: volcengine_flow_log"
sidebar_current: "docs-volcengine-resource-flow_log"
description: |-
  Provides a resource to manage flow log
---
# volcengine_flow_log
Provides a resource to manage flow log
## Example Usage
```hcl
data "volcengine_zones" "foo" {
}

resource "volcengine_vpc" "foo" {
  vpc_name     = "acc-test-vpc"
  cidr_block   = "172.16.0.0/16"
  project_name = "default"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_flow_log" "foo" {
  flow_log_name        = "acc-test-flow-log"
  description          = "acc-test"
  resource_type        = "subnet"
  resource_id          = volcengine_subnet.foo.id
  traffic_type         = "All"
  log_project_name     = "acc-test-project"
  log_topic_name       = "acc-test-topic"
  aggregation_interval = 10
  project_name         = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}
```
## Argument Reference
The following arguments are supported:
* `aggregation_interval` - (Required) The aggregation interval of flow log. Unit: minute. Valid values: `1`, `5`, `10`.
* `flow_log_name` - (Required) The name of flow log.
* `log_project_name` - (Required, ForceNew) The name of log project. If there is no corresponding log project with the name, a new log project will be created. 
When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `log_topic_name` - (Required, ForceNew) The name of log topic. If there is no corresponding log topic with the name, a new log topic will be created. 
When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `resource_id` - (Required, ForceNew) The ID of resource.
* `resource_type` - (Required, ForceNew) The type of resource. Valid values: `vpc`, `subnet`, `eni`.
* `traffic_type` - (Required, ForceNew) The type of traffic. Valid values: `All`, `Allow`, `Drop`.
* `description` - (Optional) The description of flow log.
* `project_name` - (Optional) The project name of flow log.
* `tags` - (Optional) Tags.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `business_status` - The business status of flow log.
* `created_at` - The created time of flow log.
* `lock_reason` - The reason why flow log is locked.
* `log_project_id` - The ID of log project.
* `log_topic_id` - The ID of log topic.
* `status` - The status of flow log. Values: `Active`, `Pending`, `Inactive`, `Creating`, `Deleting`.
* `updated_at` - The updated time of flow log.


## Import
FlowLog can be imported using the id, e.g.
```
$ terraform import volcengine_flow_log.default resource_id
```

