---
subcategory: "CEN"
layout: "volcengine"
page_title: "Volcengine: volcengine_cen_service_route_entry"
sidebar_current: "docs-volcengine-resource-cen_service_route_entry"
description: |-
  Provides a resource to manage cen service route entry
---
# volcengine_cen_service_route_entry
Provides a resource to manage cen service route entry
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
```
## Argument Reference
The following arguments are supported:
* `cen_id` - (Required, ForceNew) The cen ID of the cen service route entry.
* `destination_cidr_block` - (Required, ForceNew) The destination cidr block of the cen service route entry.
* `service_region_id` - (Required, ForceNew) The service region id of the cen service route entry.
* `service_vpc_id` - (Required, ForceNew) The service VPC id of the cen service route entry.
* `description` - (Optional) The description of the cen service route entry.
* `publish_mode` - (Optional) Publishing scope of cloud service access routes. Valid values are `LocalDCGW`(default), `Custom`.
* `publish_to_instances` - (Optional) The publish instances. A maximum of 100 can be uploaded in one request. This field needs to be filled in when the `publish_mode` is `Custom`.

The `publish_to_instances` object supports the following:

* `instance_id` - (Optional) Cloud service access routes need to publish the network instance ID.
* `instance_region_id` - (Optional) The region where the cloud service access route needs to be published.
* `instance_type` - (Optional) The network instance type that needs to be published for cloud service access routes. The values are as follows: `VPC`, `DCGW`.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `creation_time` - The create time of the cen service route entry.
* `status` - The status of the cen service route entry.


## Import
CenServiceRouteEntry can be imported using the CenId#DestinationCidrBlock#ServiceRegionId#ServiceVpcId, e.g.
```
$ terraform import volcengine_cen_service_route_entry.default cen-2nim00ybaylts7trquyzt****#100.XX.XX.0/24#cn-beijing#vpc-3rlkeggyn6tc010exd32q****
```

