---
subcategory: "VPC"
layout: "volcengine"
page_title: "Volcengine: volcengine_flow_log_active"
sidebar_current: "docs-volcengine-resource-flow_log_active"
description: |-
  Provides a resource to manage flow log active
---
# volcengine_flow_log_active
Provides a resource to manage flow log active
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

resource "volcengine_flow_log_active" "foo" {
  flow_log_id = volcengine_flow_log.foo.id
}
```
## Argument Reference
The following arguments are supported:
* `flow_log_id` - (Required, ForceNew) The ID of flow log.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `status` - The status of flow log.


## Import
FlowLogActive can be imported using the id, e.g.
```
$ terraform import volcengine_flow_log_active.default resource_id
```

