---
subcategory: "ESCLOUD"
layout: "volcengine"
page_title: "Volcengine: volcengine_escloud_instance"
sidebar_current: "docs-volcengine-resource-escloud_instance"
description: |-
  Provides a resource to manage escloud instance
---
# volcengine_escloud_instance
(Deprecated! Recommend use volcengine_escloud_instance_v2 replace) Provides a resource to manage escloud instance
## Example Usage
```hcl
data "volcengine_zones" "foo" {
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet_new"
  description = "tfdesc"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_escloud_instance" "foo" {
  instance_configuration {
    version            = "V6_7"
    zone_number        = 1
    enable_https       = true
    admin_user_name    = "admin"
    admin_password     = "Password@@"
    charge_type        = "PostPaid"
    configuration_code = "es.standard"
    enable_pure_master = true
    instance_name      = "acc-test-0"
    node_specs_assigns {
      type               = "Master"
      number             = 3
      resource_spec_name = "es.x4.medium"
      storage_spec_name  = "es.volume.essd.pl0"
      storage_size       = 100
    }
    node_specs_assigns {
      type               = "Hot"
      number             = 2
      resource_spec_name = "es.x4.large"
      storage_spec_name  = "es.volume.essd.pl0"
      storage_size       = 100
    }
    node_specs_assigns {
      type               = "Kibana"
      number             = 1
      resource_spec_name = "kibana.x2.small"
    }
    subnet_id                 = volcengine_subnet.foo.id
    project_name              = "default"
    force_restart_after_scale = false
  }
}
```
## Argument Reference
The following arguments are supported:
* `instance_configuration` - (Required) The configuration of ESCloud instance.

The `instance_configuration` object supports the following:

* `admin_password` - (Required) The password of administrator account. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `admin_user_name` - (Required, ForceNew) The name of administrator account(should be admin).
* `charge_type` - (Required, ForceNew) The charge type of ESCloud instance, the value can be PostPaid or PrePaid.
* `configuration_code` - (Required, ForceNew) Configuration code used for billing.
* `enable_https` - (Required, ForceNew) Whether Https access is enabled.
* `enable_pure_master` - (Required, ForceNew) Whether the Master node is independent.
* `node_specs_assigns` - (Required) The number and configuration of various ESCloud instance node. Kibana NodeSpecsAssign should not be modified.
* `subnet_id` - (Required, ForceNew) The ID of subnet, the subnet must belong to the AZ selected.
* `version` - (Required, ForceNew) The version of ESCloud instance, the value is V6_7 or V7_10.
* `zone_number` - (Required, ForceNew) The zone count of the ESCloud instance used.
* `force_restart_after_scale` - (Optional) Whether to force restart when changes are made. If true, it means that the cluster will be forced to restart without paying attention to instance availability. Works only on modified the node_specs_assigns field.
* `instance_name` - (Optional) The name of ESCloud instance.
* `maintenance_day` - (Optional) The maintainable date for the instance. Works only on modified scenes.
* `maintenance_time` - (Optional) The maintainable time period for the instance. Works only on modified scenes.
* `project_name` - (Optional) The project name  to which the ESCloud instance belongs.
* `region_id` - (Optional) The region ID of ESCloud instance.
* `zone_id` - (Optional) The available zone ID of ESCloud instance.

The `node_specs_assigns` object supports the following:

* `number` - (Required) The number of node.
* `resource_spec_name` - (Required) The name of compute resource spec, the value is `kibana.x2.small` or `es.x4.medium` or `es.x4.large` or `es.x4.xlarge` or `es.x2.2xlarge` or `es.x4.2xlarge` or `es.x2.3xlarge`.
* `type` - (Required) The type of node, the value is `Master` or `Hot` or `Kibana`.
* `storage_size` - (Optional) The size of storage. Kibana NodeSpecsAssign should not specify this field.
* `storage_spec_name` - (Optional) The name of storage spec. Kibana NodeSpecsAssign should not specify this field.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
ESCloud Instance can be imported using the id, e.g.
```
$ terraform import volcengine_escloud_instance.default n769ewmjjqyqh5dv
```

