---
subcategory: "KMS"
layout: "volcengine"
page_title: "Volcengine: volcengine_kms_regions"
sidebar_current: "docs-volcengine-datasource-kms_regions"
description: |-
  Use this data source to query detailed information of kms regions
---
# volcengine_kms_regions
Use this data source to query detailed information of kms regions
## Example Usage
```hcl
data "volcengine_kms_regions" "default" {
}
```
## Argument Reference
The following arguments are supported:
* `output_file` - (Optional) File name where to save data source results.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `regions` - The supported regions.
    * `region_id` - The region ID.
* `total_count` - The total count of query.


