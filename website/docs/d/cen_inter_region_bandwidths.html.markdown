---
subcategory: "CEN"
layout: "volcengine"
page_title: "Volcengine: volcengine_cen_inter_region_bandwidths"
sidebar_current: "docs-volcengine-datasource-cen_inter_region_bandwidths"
description: |-
  Use this data source to query detailed information of cen inter region bandwidths
---
# volcengine_cen_inter_region_bandwidths
Use this data source to query detailed information of cen inter region bandwidths
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

data "volcengine_cen_inter_region_bandwidths" "foo" {
  ids = [volcengine_cen_inter_region_bandwidth.foo.id]
}
```
## Argument Reference
The following arguments are supported:
* `cen_id` - (Optional) The ID of the cen.
* `ids` - (Optional) A list of cen inter region bandwidth IDs.
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `inter_region_bandwidths` - The collection of cen inter region bandwidth query.
    * `bandwidth` - The bandwidth of the cen inter region bandwidth.
    * `cen_id` - The cen ID of the cen inter region bandwidth.
    * `creation_time` - The create time of the cen inter region bandwidth.
    * `id` - The ID of the cen inter region bandwidth.
    * `inter_region_bandwidth_id` - The ID of the cen inter region bandwidth.
    * `local_region_id` - The local region id of the cen inter region bandwidth.
    * `peer_region_id` - The peer region id of the cen inter region bandwidth.
    * `status` - The status of the cen inter region bandwidth.
    * `update_time` - The update time of the cen inter region bandwidth.
* `total_count` - The total count of cen inter region bandwidth query.


