---
subcategory: "PRIVATE_ZONE"
layout: "volcengine"
page_title: "Volcengine: volcengine_private_zone_record_weight_enabler"
sidebar_current: "docs-volcengine-resource-private_zone_record_weight_enabler"
description: |-
  Provides a resource to manage private zone record weight enabler
---
# volcengine_private_zone_record_weight_enabler
Provides a resource to manage private zone record weight enabler
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
* `record_set_id` - (Required, ForceNew) The id of the private zone record set.
* `weight_enabled` - (Required) Whether to enable the load balance of the private zone record set.
* `zid` - (Required, ForceNew) The zid of the private zone record set.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
PrivateZoneRecordWeightEnabler can be imported using the zid:record_set_id, e.g.
```
$ terraform import volcengine_private_zone_record_weight_enabler.default resource_id
```

