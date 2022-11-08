---
subcategory: "CEN(BETA)"
layout: "volcengine"
page_title: "Volcengine: volcengine_cen_inter_region_bandwidth"
sidebar_current: "docs-volcengine-resource-cen_inter_region_bandwidth"
description: |-
  Provides a resource to manage cen inter region bandwidth
---
# volcengine_cen_inter_region_bandwidth
Provides a resource to manage cen inter region bandwidth
## Example Usage
```hcl
resource "volcengine_cen_inter_region_bandwidth" "foo" {
  cen_id          = "cen-274vsbhwvvb407fap8sp611w7"
  local_region_id = "cn-north-3"
  peer_region_id  = "cn-zhangjiakou"
  bandwidth       = 1
}
```
## Argument Reference
The following arguments are supported:
* `bandwidth` - (Required) The bandwidth of the cen inter region bandwidth.
* `cen_id` - (Required, ForceNew) The cen ID of the cen inter region bandwidth.
* `local_region_id` - (Required, ForceNew) The local region id of the cen inter region bandwidth.
* `peer_region_id` - (Required, ForceNew) The peer region id of the cen inter region bandwidth.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `creation_time` - The create time of the cen inter region bandwidth.
* `inter_region_bandwidth_id` - The ID of the cen inter region bandwidth.
* `status` - The status of the cen inter region bandwidth.
* `update_time` - The update time of the cen inter region bandwidth.


## Import
CenInterRegionBandwidth can be imported using the id, e.g.
```
$ terraform import volcengine_cen_inter_region_bandwidth.default cirb-3tex2x1cwd4c6c0v****
```

