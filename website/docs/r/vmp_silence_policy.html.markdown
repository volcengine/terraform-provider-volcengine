---
subcategory: "VMP"
layout: "volcengine"
page_title: "Volcengine: volcengine_vmp_silence_policy"
sidebar_current: "docs-volcengine-resource-vmp_silence_policy"
description: |-
  Provides a resource to manage vmp silence policy
---
# volcengine_vmp_silence_policy
Provides a resource to manage vmp silence policy
## Example Usage
```hcl
resource "volcengine_vmp_silence_policy" "example" {
  name        = "tf-acc-silence"
  description = "terraform silence policy"
  time_range_matchers {
    location = "Asia/Shanghai"
    periodic_date {
      time    = "20:00~21:12"
      weekday = "1,5"
    }
  }
  metric_label_matchers {
    matchers {
      label    = "app"
      value    = "test"
      operator = "NotEqual"
    }
    matchers {
      label    = "env"
      value    = "prod"
      operator = "Equal"
    }
  }
}
```
## Argument Reference
The following arguments are supported:
* `metric_label_matchers` - (Required) Alarm event Label matcher, allowing a maximum of 32 entries, with an "OR" relationship between different entries. Different label_matchers in the Matcher follow an "AND" relationship.
* `name` - (Required) The name of the silence policy.
* `time_range_matchers` - (Required) Alarm silence time. Case 1: Always effective. When the switch is turned on, the matched alarm events are always silenced, and only the location needs to be set. Case 2: Periodic effective. When the switch is turned on, the matched alarm events are silenced periodically, and the location and periodic_date need to be set. Case 3: Custom effective. When the switch is turned on, the matched alarm events are silenced in the specified time range, and the location and date need to be set.
* `description` - (Optional) The description of the silence policy.

The `matchers` object supports the following:

* `label` - (Required) Label.
* `value` - (Required) Label value.
* `operator` - (Optional) Operator. The optional values are as follows: Equal, NotEqual, RegexpEqual, RegexpNotEqual.

The `metric_label_matchers` object supports the following:

* `matchers` - (Required) Label matcher. Among them, each LabelMatcher array can contain a maximum of 24 items.

The `periodic_date` object supports the following:

* `day_of_month` - (Optional) Days of month, e.g. 2~3. A maximum of 10 time periods can be configured.
* `time` - (Optional) Time periods, e.g. 20:00~21:12,22:00~23:12. A maximum of 4 time periods can be configured.
* `weekday` - (Optional) Weekdays, e.g. 1,3,5. A maximum of 7 time periods can be configured.

The `time_range_matchers` object supports the following:

* `location` - (Required) Timezone, e.g. Asia/Shanghai.
* `date` - (Optional) Silence time range, like 2025-01-02 15:04~2025-01-03 14:04.
* `periodic_date` - (Optional) The cycle of alarm silence. It is used to configure alarm silence that takes effect periodically.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `auto_delete_time` - The auto delete time of the silence policy.
* `create_time` - The create time of the silence policy.
* `source` - The source of the silence policy.
* `status` - The status of the silence policy.
* `update_time` - The update time of the silence policy.


## Import
VmpSilencePolicy can be imported using the id, e.g.
```
$ terraform import volcengine_vmp_silence_policy.default resource_id
```

