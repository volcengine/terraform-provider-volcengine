---
subcategory: "CEN"
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
resource "volcengine_cen" "foo" {
  cen_name     = "acc-test-cen"
  description  = "acc-test"
  project_name = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_cen_bandwidth_package" "foo" {
  local_geographic_region_set_id = "China"
  peer_geographic_region_set_id  = "China"
  bandwidth                      = 5
  cen_bandwidth_package_name     = "acc-test-cen-bp"
  description                    = "acc-test"
  billing_type                   = "PrePaid"
  period_unit                    = "Month"
  period                         = 1
  project_name                   = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_cen_bandwidth_package_associate" "foo" {
  cen_bandwidth_package_id = volcengine_cen_bandwidth_package.foo.id
  cen_id                   = volcengine_cen.foo.id
}

resource "volcengine_cen_inter_region_bandwidth" "foo" {
  cen_id          = volcengine_cen.foo.id
  local_region_id = "cn-beijing"
  peer_region_id  = "cn-shanghai"
  bandwidth       = 2
  depends_on      = [volcengine_cen_bandwidth_package_associate.foo]
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

