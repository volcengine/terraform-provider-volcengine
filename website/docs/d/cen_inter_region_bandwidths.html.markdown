---
subcategory: "CEN(BETA)"
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
data "volcengine_cen_inter_region_bandwidths" "foo" {
  ids = ["cirb-274q484wxao007fap8tlvl6si"]
}
```
## Argument Reference
The following arguments are supported:
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


