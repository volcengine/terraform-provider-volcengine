---
subcategory: "VMP"
layout: "volcengine"
page_title: "Volcengine: volcengine_vmp_silence_policies"
sidebar_current: "docs-volcengine-datasource-vmp_silence_policies"
description: |-
  Use this data source to query detailed information of vmp silence policies
---
# volcengine_vmp_silence_policies
Use this data source to query detailed information of vmp silence policies
## Example Usage
```hcl
data "volcengine_vmp_silence_policies" "example" {
  ids  = ["ea51e747-0ead-4e09-9187-76beba6400b7"]
  name = "tf-acc-silence"
}
```
## Argument Reference
The following arguments are supported:
* `ids` - (Optional) A list of silence policy ids.
* `name` - (Optional) The name of silence policy.
* `output_file` - (Optional) File name where to save data source results.
* `sources` - (Optional) The sources of silence policy: General/LarkBot.
* `status` - (Optional) The status of silence policy: Active/Disabled/Expired.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `silence_policies` - The list of silence policies.
    * `auto_delete_time` - The auto delete time of the silence policy.
    * `create_time` - The create time of the silence policy, in RFC3339 format.
    * `description` - The description of the silence policy.
    * `id` - The id of the silence policy.
    * `name` - The name of the silence policy.
    * `source` - The source of the silence policy.
    * `status` - The status of the silence policy.
    * `time_range_matchers` - The matching time in the alert silence policy.
        * `date` - The time period for alarm silence.
        * `location` - Time zone.
        * `periodic_date` - The cycle of alarm silence.
            * `day_of_month` - Days of the month, e.g. 1,15,30.
            * `time` - Time periods, e.g. 20:00~21:12,22:00~23:12.
            * `weekday` - Weekdays, e.g. 1,3,5.
    * `update_time` - The update time of the silence policy, in RFC3339 format.
* `total_count` - The total count of query.


