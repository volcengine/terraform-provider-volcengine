---
subcategory: "CEN"
layout: "volcengine"
page_title: "Volcengine: volcengine_cen_route_entry"
sidebar_current: "docs-volcengine-resource-cen_route_entry"
description: |-
  Provides a resource to manage cen route entry
---
# volcengine_cen_route_entry
Provides a resource to manage cen route entry
## Example Usage
```hcl
resource "volcengine_cen_route_entry" "foo" {
  cen_id                 = "cen-12ar8uclj68sg17q7y20v9gil"
  destination_cidr_block = "192.168.0.0/24"
  instance_type          = "VPC"
  instance_region_id     = "cn-beijing"
  instance_id            = "vpc-im67wjcikxkw8gbssx8ufpj8"
}

resource "volcengine_cen_route_entry" "foo1" {
  cen_id                 = "cen-12ar8uclj68sg17q7y20v9gil"
  destination_cidr_block = "192.168.17.0/24"
  instance_type          = "VPC"
  instance_region_id     = "cn-beijing"
  instance_id            = "vpc-im67wjcikxkw8gbssx8ufpj8"
}
```
## Argument Reference
The following arguments are supported:
* `cen_id` - (Required, ForceNew) The cen ID of the cen route entry.
* `destination_cidr_block` - (Required, ForceNew) The destination cidr block of the cen route entry.
* `instance_id` - (Required, ForceNew) The instance id of the next hop of the cen route entry.
* `instance_region_id` - (Required, ForceNew) The instance region id of the next hop of the cen route entry.
* `instance_type` - (Optional, ForceNew) The instance type of the next hop of the cen route entry.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `as_path` - The AS path of the cen route entry.
* `publish_status` - The publish status of the cen route entry.
* `status` - The status of the cen route entry.


## Import
CenRouteEntry can be imported using the CenId:DestinationCidrBlock:InstanceId:InstanceType:InstanceRegionId, e.g.
```
$ terraform import volcengine_cen_route_entry.default cen-2nim00ybaylts7trquyzt****:100.XX.XX.0/24:vpc-vtbnbb04qw3k2hgi12cv****:VPC:cn-beijing
```

