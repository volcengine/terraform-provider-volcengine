---
subcategory: "CEN"
layout: "volcengine"
page_title: "Volcengine: volcengine_cen_service_route_entries"
sidebar_current: "docs-volcengine-datasource-cen_service_route_entries"
description: |-
  Use this data source to query detailed information of cen service route entries
---
# volcengine_cen_service_route_entries
Use this data source to query detailed information of cen service route entries
## Example Usage
```hcl
resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
  count      = 3
}

resource "volcengine_cen" "foo" {
  cen_name     = "acc-test-cen"
  description  = "acc-test"
  project_name = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_cen_attach_instance" "foo" {
  cen_id             = volcengine_cen.foo.id
  instance_id        = volcengine_vpc.foo[count.index].id
  instance_region_id = "cn-beijing"
  instance_type      = "VPC"
  count              = 3
}

resource "volcengine_cen_service_route_entry" "foo" {
  cen_id                 = volcengine_cen.foo.id
  destination_cidr_block = "100.64.0.0/11"
  service_region_id      = "cn-beijing"
  service_vpc_id         = volcengine_cen_attach_instance.foo[0].instance_id
  description            = "acc-test"
  publish_mode           = "Custom"
  publish_to_instances {
    instance_region_id = "cn-beijing"
    instance_type      = "VPC"
    instance_id        = volcengine_cen_attach_instance.foo[1].instance_id
  }
  publish_to_instances {
    instance_region_id = "cn-beijing"
    instance_type      = "VPC"
    instance_id        = volcengine_cen_attach_instance.foo[2].instance_id
  }
}

data "volcengine_cen_service_route_entries" "foo" {
  cen_id                 = volcengine_cen.foo.id
  destination_cidr_block = volcengine_cen_service_route_entry.foo.destination_cidr_block
}
```
## Argument Reference
The following arguments are supported:
* `cen_id` - (Optional) A cen ID.
* `destination_cidr_block` - (Optional) A destination cidr block.
* `output_file` - (Optional) File name where to save data source results.
* `service_region_id` - (Optional) A service region id.
* `service_vpc_id` - (Optional) A service VPC id.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `service_route_entries` - The collection of cen service route entry query.
    * `cen_id` - The cen ID of the cen service route entry.
    * `creation_time` - The create time of the cen service route entry.
    * `description` - The description of the cen service route entry.
    * `destination_cidr_block` - The destination cidr block of the cen service route entry.
    * `publish_mode` - Publishing scope of cloud service access routes. Valid values are `LocalDCGW`(default), `Custom`.
    * `publish_to_instances` - The publish instances. A maximum of 100 can be uploaded in one request.
        * `instance_id` - Cloud service access routes need to publish the network instance ID.
        * `instance_region_id` - The region where the cloud service access route needs to be published.
        * `instance_type` - The network instance type that needs to be published for cloud service access routes. The values are as follows: `VPC`, `DCGW`.
    * `service_region_id` - The service region id of the cen service route entry.
    * `service_vpc_id` - The service VPC id of the cen service route entry.
    * `status` - The status of the cen service route entry.
* `total_count` - The total count of cen service route entry.


