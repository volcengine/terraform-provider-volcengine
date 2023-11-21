---
subcategory: "CEN"
layout: "volcengine"
page_title: "Volcengine: volcengine_cen_grant_instance"
sidebar_current: "docs-volcengine-resource-cen_grant_instance"
description: |-
  Provides a resource to manage cen grant instance
---
# volcengine_cen_grant_instance
Provides a resource to manage cen grant instance
## Example Usage
```hcl
resource "volcengine_cen_grant_instance" "foo" {
  cen_id             = "cen-2d6zdn0c1z5s058ozfcyf4lee"
  cen_owner_id       = "210000****"
  instance_type      = "VPC"
  instance_id        = "vpc-2bysvq1xx543k2dx0eeulpeiv"
  instance_region_id = "cn-beijing"
}
```
## Argument Reference
The following arguments are supported:
* `cen_id` - (Required, ForceNew) The ID of the cen.
* `cen_owner_id` - (Required, ForceNew) The owner ID of the cen.
* `instance_id` - (Required, ForceNew) The ID of the instance.
* `instance_region_id` - (Required, ForceNew) The region ID of the instance.
* `instance_type` - (Required, ForceNew) The type of the instance.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
Cen grant instance can be imported using the CenId:CenOwnerId:InstanceId:InstanceType:RegionId, e.g.
```
$ terraform import volcengine_cen_grant_instance.default cen-7qthudw0ll6jmc***:210000****:vpc-2fexiqjlgjif45oxruvso****:VPC:cn-beijing
```

