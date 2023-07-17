---
subcategory: "CEN(BETA)"
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
resource "volcengine_cen_service_route_entry" "foo" {
  cen_id                 = "cen-12ar8uclj68sg17q7y20v9gil"
  destination_cidr_block = "100.64.0.0/11"
  service_region_id      = "cn-beijing"
  service_vpc_id         = "vpc-im67wjcikxkw8gbssx8ufpj8"
  description            = "test-tf"
  publish_mode           = "Custom"
  publish_to_instances {
    instance_region_id = "cn-beijing"
    instance_type      = "VPC"
    instance_id        = "vpc-2fepz36a5ra4g59gp67w197xo"
  }
  publish_to_instances {
    instance_region_id = "cn-beijing"
    instance_type      = "VPC"
    instance_id        = "vpc-im67wjcikxkw8gbssx8ufpj8"
  }
}

resource "volcengine_cen_service_route_entry" "foo1" {
  cen_id                 = "cen-12ar8uclj68sg17q7y20v9gil"
  destination_cidr_block = "100.64.0.0/10"
  service_region_id      = "cn-beijing"
  service_vpc_id         = "vpc-im67wjcikxkw8gbssx8ufpj8"
  description            = "test-tf"
  publish_mode           = "Custom"
  publish_to_instances {
    instance_region_id = "cn-beijing"
    instance_type      = "VPC"
    instance_id        = "vpc-2fepz36a5ra4g59gp67w197xo"
  }
  publish_to_instances {
    instance_region_id = "cn-beijing"
    instance_type      = "VPC"
    instance_id        = "vpc-im67wjcikxkw8gbssx8ufpj8"
  }
}

resource "volcengine_cen_service_route_entry" "foo2" {
  cen_id                 = "cen-12ar8uclj68sg17q7y20v9gil"
  destination_cidr_block = "100.64.0.0/12"
  service_region_id      = "cn-beijing"
  service_vpc_id         = "vpc-im67wjcikxkw8gbssx8ufpj8"
  description            = "test-tf"
  publish_mode           = "Custom"
  publish_to_instances {
    instance_region_id = "cn-beijing"
    instance_type      = "VPC"
    instance_id        = "vpc-2fepz36a5ra4g59gp67w197xo"
  }
  publish_to_instances {
    instance_region_id = "cn-beijing"
    instance_type      = "VPC"
    instance_id        = "vpc-im67wjcikxkw8gbssx8ufpj8"
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
CenServiceRouteEntry can be imported using the CenId:DestinationCidrBlock:ServiceRegionId:ServiceVpcId, e.g.
```
$ terraform import volcengine_cen_service_route_entry.default cen-2nim00ybaylts7trquyzt****:100.XX.XX.0/24:cn-beijing:vpc-3rlkeggyn6tc010exd32q****
```

