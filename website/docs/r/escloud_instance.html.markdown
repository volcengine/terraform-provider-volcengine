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
    zone_number        = 1
    enable_https       = true
    admin_user_name    = "admin"
    admin_password     = "xxxx"
    charge_type        = "PostPaid"
    configuration_code = "es.standard"
    enable_pure_master = true
    instance_name      = "from-tf4"
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
      storage_spec_name  = "es.volume.essd.pl0"
      storage_size       = 0
    }
    subnet_id = "subnet-2bz9vxrixqigw2dx0eextz50p"
  }
}
```
## Argument Reference
The following arguments are supported:
* `instance_configuration` - (Required) The configuration of ESCloud instance.

The `instance_configuration` object supports the following:

* `admin_password` - (Required) The password of administrator account.
* `admin_user_name` - (Required, ForceNew) The name of administrator account(should be admin).
* `charge_type` - (Required, ForceNew) The charge type of ESCloud instance, the value can be PostPaid or PrePaid.
* `configuration_code` - (Required) Configuration code used for billing.
* `enable_https` - (Required, ForceNew) Whether Https access is enabled.
* `enable_pure_master` - (Required, ForceNew) Whether the Master node is independent.
* `node_specs_assigns` - (Required) The number and configuration of various ESCloud instance node.
* `subnet_id` - (Required, ForceNew) The ID of subnet, the subnet must belong to the AZ selected.
* `version` - (Required, ForceNew) The version of ESCloud instance, the value is V6_7 or V7_10.
* `zone_number` - (Required, ForceNew) The zone count of the ESCloud instance used.
* `force_restart_after_scale` - (Optional) Whether to force restart when changes are made. If true, it means that the cluster will be forced to restart without paying attention to instance availability.
* `instance_name` - (Optional) The name of ESCloud instance.
* `maintenance_day` - (Optional) The maintainable date for the instance. Works only on modified scenes.
* `maintenance_time` - (Optional) The maintainable time period for the instance. Works only on modified scenes.
* `project_name` - (Optional) The project name  to which the ESCloud instance belongs.
* `region_id` - (Optional) The region ID of ESCloud instance.
* `zone_id` - (Optional) The available zone ID of ESCloud instance.

The `node_specs_assigns` object supports the following:

* `number` - (Required) The number of node.
* `resource_spec_name` - (Required) The name of compute resource spec, the value is `kibana.x2.small` or `es.x4.medium` or `es.x4.large` or `es.x4.xlarge` or `es.x2.2xlarge` or `es.x4.2xlarge` or `es.x2.3xlarge`.
* `storage_size` - (Required) The size of storage.
* `storage_spec_name` - (Required) The name of storage spec.
* `type` - (Required) The type of node, the value is `Master` or `Hot` or `Kibana`.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
ESCloud Instance can be imported using the id, e.g.
```
$ terraform import volcengine_escloud_instance.default n769ewmjjqyqh5dv
```

