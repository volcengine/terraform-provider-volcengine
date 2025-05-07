---
subcategory: "PRIVATE_ZONE"
layout: "volcengine"
page_title: "Volcengine: volcengine_private_zone_record"
sidebar_current: "docs-volcengine-resource-private_zone_record"
description: |-
  Provides a resource to manage private zone record
---
# volcengine_private_zone_record
Provides a resource to manage private zone record
## Example Usage
```hcl
resource "volcengine_private_zone_record" "foo" {
  zid    = 2450000
  host   = "www"
  type   = "A"
  value  = "10.1.1.158"
  weight = 8
  ttl    = 700
  remark = "tf-test"
  enable = true
}

data "volcengine_private_zone_record_sets" "foo" {
  zid         = volcengine_private_zone_record.foo.zid
  host        = volcengine_private_zone_record.foo.host
  search_mode = "EXACT"
}

resource "volcengine_private_zone_record_weight_enabler" "foo" {
  zid            = volcengine_private_zone_record.foo.zid
  record_set_id  = [for set in data.volcengine_private_zone_record_sets.foo.record_sets : set.record_set_id if set.type == volcengine_private_zone_record.foo.type][0]
  weight_enabled = true
}
```
## Argument Reference
The following arguments are supported:
* `host` - (Required) The host of the private zone record.
* `type` - (Required) The type of the private zone record. Valid values: `A`, `AAAA`, `CNAME`, `MX`, `PTR`.
* `value` - (Required) The value of the private zone record. Record values need to be set based on the value of the `type`.
* `zid` - (Required, ForceNew) The zid of the private zone record.
* `enable` - (Optional) Whether to enable the private zone record. This field is only effected when modify this resource.
* `remark` - (Optional) The remark of the private zone record.
* `ttl` - (Optional) The ttl of the private zone record. Unit: second. Default is 600.
* `weight` - (Optional) The weight of the private zone record. This field is only effected when the `load_balance_mode` of the private zone is true and the `weight_enabled` of the record_set is true. Default is 1.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
PrivateZoneRecord can be imported using the id, e.g.
```
$ terraform import volcengine_private_zone_record.default resource_id
```

