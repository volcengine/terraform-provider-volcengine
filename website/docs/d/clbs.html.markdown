---
subcategory: "CLB"
layout: "volcengine"
page_title: "Volcengine: volcengine_clbs"
sidebar_current: "docs-volcengine-datasource-clbs"
description: |-
  Use this data source to query detailed information of clbs
---
# volcengine_clbs
Use this data source to query detailed information of clbs
## Example Usage
```hcl
data "volcengine_clbs" "default" {
  ids = ["clb-273y2ok6ets007fap8txvf6us"]
}
```
## Argument Reference
The following arguments are supported:
* `eni_address` - (Optional) The private ip address of the Clb.
* `ids` - (Optional) A list of Clb IDs.
* `load_balancer_name` - (Optional) The name of the Clb.
* `name_regex` - (Optional) A Name Regex of Clb.
* `output_file` - (Optional) File name where to save data source results.
* `project_name` - (Optional) The ProjectName of Clb.
* `tags` - (Optional) Tags.
* `vpc_id` - (Optional) The id of the VPC.

The `tags` object supports the following:

* `key` - (Required) The Key of Tags.
* `value` - (Required) The Value of Tags.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `clbs` - The collection of Clb query.
    * `business_status` - The business status of the Clb.
    * `create_time` - The create time of the Clb.
    * `deleted_time` - The expected recycle time of the Clb.
    * `description` - The description of the Clb.
    * `eip_address` - The Eip address of the Clb.
    * `eip_id` - The Eip ID of the Clb.
    * `eni_address` - The Eni address of the Clb.
    * `eni_id` - The Eni ID of the Clb.
    * `expired_time` - The expired time of the CLB.
    * `id` - The ID of the Clb.
    * `instance_status` - The billing status of the CLB.
    * `load_balancer_billing_type` - The billing type of the Clb.
    * `load_balancer_id` - The ID of the Clb.
    * `load_balancer_name` - The name of the Clb.
    * `load_balancer_spec` - The specifications of the Clb.
    * `lock_reason` - The reason why Clb is locked.
    * `master_zone_id` - The master zone ID of the CLB.
    * `modification_protection_reason` - The modification protection reason of the Clb.
    * `modification_protection_status` - The modification protection status of the Clb.
    * `overdue_reclaim_time` - The over reclaim time of the CLB.
    * `overdue_time` - The overdue time of the Clb.
    * `project_name` - The ProjectName of the Clb.
    * `reclaim_time` - The reclaim time of the CLB.
    * `remain_renew_times` - The remain renew times of the CLB. When the value of the renew_type is `AutoRenew`, the query returns this field.
    * `renew_period_times` - The renew period times of the CLB. When the value of the renew_type is `AutoRenew`, the query returns this field.
    * `renew_type` - The renew type of the CLB. When the value of the load_balancer_billing_type is `PrePaid`, the query returns this field.
    * `slave_zone_id` - The slave zone ID of the CLB.
    * `status` - The status of the Clb.
    * `subnet_id` - The subnet ID of the Clb.
    * `tags` - Tags.
        * `key` - The Key of Tags.
        * `value` - The Value of Tags.
    * `type` - The type of the Clb.
    * `update_time` - The update time of the Clb.
    * `vpc_id` - The vpc ID of the Clb.
* `total_count` - The total count of Clb query.


