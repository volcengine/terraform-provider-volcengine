---
subcategory: "PRIVATELINK"
layout: "volcengine"
page_title: "Volcengine: volcengine_privatelink_vpc_endpoint_services"
sidebar_current: "docs-volcengine-datasource-privatelink_vpc_endpoint_services"
description: |-
  Use this data source to query detailed information of privatelink vpc endpoint services
---
# volcengine_privatelink_vpc_endpoint_services
Use this data source to query detailed information of privatelink vpc endpoint services
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

resource "volcengine_clb" "foo" {
  type                       = "public"
  subnet_id                  = volcengine_subnet.foo.id
  load_balancer_spec         = "small_1"
  description                = "acc-test-demo"
  load_balancer_name         = "acc-test-clb"
  load_balancer_billing_type = "PostPaid"
  eip_billing_config {
    isp              = "BGP"
    eip_billing_type = "PostPaidByBandwidth"
    bandwidth        = 1
  }
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_privatelink_vpc_endpoint_service" "foo" {
  resources {
    resource_id   = volcengine_clb.foo.id
    resource_type = "CLB"
  }
  description         = "acc-test"
  auto_accept_enabled = true
  count               = 2
}

data "volcengine_privatelink_vpc_endpoint_services" "foo" {
  ids = volcengine_privatelink_vpc_endpoint_service.foo[*].id
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) The IDs of vpc endpoint service.
* `name_regex` - (Optional) A Name Regex of vpc endpoint service.
* `output_file` - (Optional) File name where to save data source results.
* `service_name` - (Optional) The name of vpc endpoint service.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `services` - The collection of query.
    * `auto_accept_enabled` - Whether auto accept node connect.
    * `creation_time` - The create time of service.
    * `description` - The description of service.
    * `id` - The Id of service.
    * `resources` - The resources info.
        * `resource_id` - The id of resource.
        * `resource_type` - The type of resource.
        * `zone_id` - The zone id of resource.
    * `service_domain` - The domain of service.
    * `service_id` - The Id of service.
    * `service_name` - The name of service.
    * `service_resource_type` - The resource type of service.
    * `service_type` - The type of service.
    * `status` - The status of service.
    * `update_time` - The update time of service.
    * `zone_ids` - The IDs of zones.
* `total_count` - Returns the total amount of the data list.


