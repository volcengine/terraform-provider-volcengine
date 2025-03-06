---
subcategory: "RABBITMQ"
layout: "volcengine"
page_title: "Volcengine: volcengine_rabbitmq_instance"
sidebar_current: "docs-volcengine-resource-rabbitmq_instance"
description: |-
  Provides a resource to manage rabbitmq instance
---
# volcengine_rabbitmq_instance
Provides a resource to manage rabbitmq instance
## Example Usage
```hcl
data "volcengine_zones" "foo" {
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block  = "172.16.0.0/24"
  zone_id     = data.volcengine_zones.foo.zones[0].id
  vpc_id      = volcengine_vpc.foo.id
}

resource "volcengine_rabbitmq_instance" "foo" {
  zone_ids             = [data.volcengine_zones.foo.zones[0].id, data.volcengine_zones.foo.zones[1].id, data.volcengine_zones.foo.zones[2].id]
  subnet_id            = volcengine_subnet.foo.id
  version              = "3.8.18"
  user_name            = "acc-test-user"
  user_password        = "93f0cb0614Aab12"
  compute_spec         = "rabbitmq.n3.x2.small"
  storage_space        = 300
  instance_name        = "acc-test-rabbitmq"
  instance_description = "acc-test"
  charge_info {
    charge_type = "PostPaid"
  }
  project_name = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}
```
## Argument Reference
The following arguments are supported:
* `charge_info` - (Required) The charge information of the rocketmq instance.
* `compute_spec` - (Required) The compute specification of the rabbitmq instance.
* `storage_space` - (Required) The storage space of the rabbitmq instance. Unit: GiB. The valid value must be specified as a multiple of 100.
* `subnet_id` - (Required, ForceNew) The subnet id of the rabbitmq instance.
* `user_name` - (Required) The administrator name of the rabbitmq instance.
* `user_password` - (Required) The administrator password. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
* `version` - (Required, ForceNew) The version of the rabbitmq instance. Valid values: `3.8.18`, `3.12`.
* `zone_ids` - (Required, ForceNew) The zone id of the rabbitmq instance. Support specifying multiple availability zones.
* `instance_description` - (Optional) The description of the rabbitmq instance.
* `instance_name` - (Optional) The name of the rabbitmq instance.
* `project_name` - (Optional) The IAM project name where the rabbitmq instance resides.
* `tags` - (Optional) Tags.

The `charge_info` object supports the following:

* `charge_type` - (Required) The charge type of the rabbitmq instance. Valid values: `PostPaid`, `PrePaid`.
* `auto_renew` - (Optional) Whether to automatically renew in prepaid scenarios. Default is false.
* `period_unit` - (Optional) The purchase cycle in the prepaid scenario. Valid values: `Month`, `Year`. Default is `Month`.
* `period` - (Optional) Purchase duration in prepaid scenarios. When PeriodUnit is specified as `Month`, the value range is 1-9. When PeriodUnit is specified as `Year`, the value range is 1-3. Default is 1.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `account_id` - The account id of the rabbitmq instance.
* `apply_private_dns_to_public` - Whether enable the public network parsing function of the rabbitmq instance.
* `arch_type` - The type of the rabbitmq instance.
* `create_time` - The create time of the rabbitmq instance.
* `eip_id` - The eip id of the rabbitmq instance.
* `endpoints` - The endpoint info of the rabbitmq instance.
    * `endpoint_type` - The endpoint type of the rabbitmq instance.
    * `internal_endpoint` - The internal endpoint of the rabbitmq instance.
    * `network_type` - The network type of the rabbitmq instance.
    * `public_endpoint` - The public endpoint of the rabbitmq instance.
* `init_user_name` - The WebUI admin user name of the rabbitmq instance.
* `instance_status` - The status of the rabbitmq instance.
* `region_id` - The region id of the rabbitmq instance.
* `used_storage_space` - The used storage space of the rabbitmq instance. Unit: GiB.
* `vpc_id` - The vpc id of the rabbitmq instance.


## Import
RabbitmqInstance can be imported using the id, e.g.
```
$ terraform import volcengine_rabbitmq_instance.default resource_id
```

