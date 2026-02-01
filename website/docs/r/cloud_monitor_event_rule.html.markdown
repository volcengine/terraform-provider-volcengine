---
subcategory: "CLOUD_MONITOR"
layout: "volcengine"
page_title: "Volcengine: volcengine_cloud_monitor_event_rule"
sidebar_current: "docs-volcengine-resource-cloud_monitor_event_rule"
description: |-
  Provides a resource to manage cloud monitor event rule
---
# volcengine_cloud_monitor_event_rule
Provides a resource to manage cloud monitor event rule
## Example Usage
```hcl
resource "volcengine_cloud_monitor_event_rule" "foo" {
  status          = "enable"
  contact_methods = ["Phone", "TLS", "MQ"]
  event_source    = "ecs"
  level           = "notice"
  rule_name       = "tftest1"
  effective_time {
    start_time = "01:00"
    end_time   = "22:00"
  }
  event_type        = ["ecs:Disk:DiskError.Redeploy.Canceled"]
  contact_group_ids = ["1737941730782699520", "1737940985502777344"]
  filter_pattern {
    type   = ["ecs:Disk:DiskError.Redeploy.Canceled"]
    source = "ecs"
  }
  message_queue {
    instance_id = "kafka-cnoe4rfrsqfb1d64"
    vpc_id      = "vpc-2d68hz41j7qio58ozfd6jxgtb"
    type        = "kafka"
    region      = "*****"
    topic       = "tftest"
  }
  tls_target {
    project_name   = "tf-test"
    region_name_cn = "*****"
    region_name_en = "*****"
    project_id     = "17ba378d-de43-495e-8906-03ae6567b376"
    topic_id       = "7ce12237-6670-44a7-9d79-2e36961586e6"
  }
}
```
## Argument Reference
The following arguments are supported:
* `contact_methods` - (Required) Alarm notification methods. Valid value: `Phone`, `Email`, `SMS`, `Webhook`: Alarm callback, `TLS`: Log Service, `MQ`: Message Queue Kafka.
* `effective_time` - (Required) The rule takes effect at a certain time and will only be effective during this period.
* `event_source` - (Required, ForceNew) Event source.
* `filter_pattern` - (Required) Filter mode, also known as event matching rules. Custom matching rules are not currently supported.
* `level` - (Required) Severity of alarm rules. Value can be `notice`, `warning`, `critical`.
* `rule_name` - (Required) The name of the rule.
* `contact_group_ids` - (Optional) When the alarm notification method is phone, SMS, or email, the triggered alarm contact group ID.
* `description` - (Optional) The description of the rule.
* `endpoint` - (Optional) When the alarm notification method is alarm callback, it triggers the callback address.
* `event_type` - (Optional) Event type.
* `message_queue` - (Optional) The triggered message queue when the alarm notification method is Kafka message queue.
* `status` - (Optional) Rule status. `enable`: enable rule(default), `disable`: disable rule.
* `tls_target` - (Optional) The alarm method for log service triggers the configuration of the log service.

The `effective_time` object supports the following:

* `end_time` - (Required) End time for rule activation.
* `start_time` - (Required) Start time for rule activation.

The `filter_pattern` object supports the following:

* `source` - (Required, ForceNew) Event source corresponding to pattern matching.
* `type` - (Required) The list of corresponding event types in pattern matching, currently set to match any.

The `message_queue` object supports the following:

* `instance_id` - (Required) The kafka instance id.
* `region` - (Required) The region.
* `topic` - (Required) The topic name.
* `type` - (Required) The message queue type, only support kafka now.
* `vpc_id` - (Required) The vpc id.

The `tls_target` object supports the following:

* `project_id` - (Required) The project id.
* `project_name` - (Required) The project name.
* `region_name_cn` - (Required) The Chinese region name.
* `region_name_en` - (Required) The English region name.
* `topic_id` - (Required) The topic id.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.



## Import
CloudMonitorEventRule can be imported using the id, e.g.
```
$ terraform import volcengine_cloud_monitor_event_rule.default rule_id
```

