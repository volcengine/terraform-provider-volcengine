---
subcategory: "ESCLOUD"
layout: "volcengine"
page_title: "Volcengine: volcengine_escloud_instance"
sidebar_current: "docs-volcengine-resource-escloud_instance"
description: |-
  Provides a resource to manage escloud instance
---
# volcengine_escloud_instance
Provides a resource to manage escloud instance
## Example Usage
```hcl
resource "volcengine_escloud_instance" "foo" {
  instance_configuration {
    version            = "V7_10"
    region_id          = "cn-north-4"
    zone_id            = "cn-langfang-a"
    zone_number        = 1
    enable_https       = true
    admin_user_name    = "admin"
    admin_password     = "1qaz!QAZ"
    charge_type        = "PostPaid"
    configuration_code = "es.standard"
    enable_pure_master = false
    instance_name      = "from-tf2"
    node_specs_assigns {
      type               = "Master"
      number             = 3
      resource_spec_name = "es.x4.medium"
      storage_spec_name  = "es.volume.essd.pl0"
      storage_size       = 100
    }
    node_specs_assigns {
      type               = "Hot"
      number             = 0
      resource_spec_name = "es.x4.medium"
      storage_spec_name  = "es.volume.essd.pl0"
      storage_size       = 100
    }
    node_specs_assigns {
      type               = "Kibana"
      number             = 1
      resource_spec_name = "kibana.x2.small"
      storage_spec_name  = ""
      storage_size       = 0
    }
    subnet {
      subnet_id   = "subnet-1g0d5yqrsxszk8ibuxxzile2l"
      subnet_name = "subnet-1g0d5yqrsxszk8ibuxxzile2l"
    }
    vpc {
      vpc_id   = "vpc-3cj17x7u9bzeo6c6rrtzfpaeb"
      vpc_name = "test-1231新建"
    }
  }
}
```
## Argument Reference
The following arguments are supported:
* `instance_configuration` - (Required) The configuration of ESCloud instance.

The `instance_configuration` object supports the following:

* `admin_password` - (Required) The password of administrator account.
* `admin_user_name` - (Required, ForceNew) The name of administrator account(should be admin).
* `charge_type` - (Required, ForceNew) The charge type of ESCloud instance.
* `configuration_code` - (Required, ForceNew) Configuration code used for billing.
* `enable_https` - (Required, ForceNew) Whether Https access is enabled.
* `enable_pure_master` - (Required, ForceNew) Whether the Master node is independent.
* `node_specs_assigns` - (Required, ForceNew) The number and configuration of various ESCloud instance node.
* `region_id` - (Required, ForceNew) The region ID of ESCloud instance.
* `version` - (Required, ForceNew) The version of ESCloud instance.
* `zone_id` - (Required, ForceNew) The available zone ID of ESCloud instance.
* `zone_number` - (Required, ForceNew) The zone count of the ESCloud instance used.
* `instance_name` - (Optional) The name of ESCloud instance.
* `maintenance_day` - (Optional) The maintainable date for the instance.
* `maintenance_time` - (Optional) The maintainable time period for the instance.
* `project_name` - (Optional) The project name  to which the ESCloud instance belongs.
* `subnet` - (Optional) The ID of subnet, the subnet must belong to the AZ selected.
* `vpc` - (Optional) Information about the VPC where the instance is located.

The `node_specs_assigns` object supports the following:

* `number` - (Required, ForceNew) The number of node.
* `resource_spec_name` - (Required, ForceNew) The name of compute resource spec.
* `storage_size` - (Required, ForceNew) The size of storage.
* `storage_spec_name` - (Required, ForceNew) The name of storage spec.
* `type` - (Required, ForceNew) The type of node.

The `subnet` object supports the following:

* `subnet_id` - (Required) The ID of subnet.
* `subnet_name` - (Required) The name of subnet.

The `vpc` object supports the following:

* `vpc_id` - (Required) The ID of vpc.
* `vpc_name` - (Required) The name of vpc.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
ESCloud Instance can be imported using the id, e.g.
```
$ terraform import volcengine_escloud_instance.default n769ewmjjqyqh5dv
```

