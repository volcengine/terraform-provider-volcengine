---
subcategory: "PRIVATE_ZONE"
layout: "volcengine"
page_title: "Volcengine: volcengine_private_zone"
sidebar_current: "docs-volcengine-resource-private_zone"
description: |-
  Provides a resource to manage private zone
---
# volcengine_private_zone
Provides a resource to manage private zone
## Example Usage
```hcl
resource "volcengine_private_zone" "foo" {
  zone_name         = "acc-test-pz.com"
  remark            = "acc-test-new"
  recursion_mode    = true
  intelligent_mode  = true
  load_balance_mode = true
  vpcs {
    vpc_id = "vpc-rs4mi0jedipsv0x57pf****"
  }
  vpcs {
    vpc_id = "vpc-3qdzk9xju6o747prml0jk****"
    region = "cn-shanghai"
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
* `vpcs` - (Required) The bind vpc object of the private zone. If you want to bind another account's VPC, you need to first use resource volcengine_private_zone_user_vpc_authorization to complete the authorization.
* `zone_name` - (Required, ForceNew) The name of the private zone.
* `intelligent_mode` - (Optional, ForceNew) Whether to enable the intelligent mode of the private zone.
* `load_balance_mode` - (Optional) Whether to enable the load balance mode of the private zone.
* `project_name` - (Optional) The project name of the private zone.
* `recursion_mode` - (Optional) Whether to enable the recursion mode of the private zone.
* `remark` - (Optional) The remark of the private zone.
* `tags` - (Optional) Tags.
* `vpc_trns` - (Optional) The vpc trns of the private zone. Format: trn:vpc:region:accountId:vpc/vpcId. This field is only effected when creating resource. 
When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

The `vpcs` object supports the following:

* `vpc_id` - (Required) The id of the bind vpc.
* `region` - (Optional) The region of the bind vpc. The default value is the region of the default provider config.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
PrivateZone can be imported using the id, e.g.
```
$ terraform import volcengine_private_zone.default resource_id
```

