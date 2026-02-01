---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_alarm_content_template"
sidebar_current: "docs-volcengine-resource-tls_alarm_content_template"
description: |-
  Provides a resource to manage tls alarm content template
---
# volcengine_tls_alarm_content_template
Provides a resource to manage tls alarm content template
## Example Usage
```hcl
resource "volcengine_tls_alarm_content_template" "foo" {
  alarm_content_template_name = "test-alarm-content-template"
  need_valid_content          = true
  sms {
    content = "修改-告警策略{{Alarm}}， 告警日志项目：{{ProjectName}}， 告警日志主题：{{AlarmTopicName}}， 告警级别：{{Severity}}， 通知类型：{%if NotifyType==1%}触发告警{%else%}告警恢复{%endif%}，触发时间：{{StartTime}}， 触发条件：{{Condition}}， 当前查询结果：[{%-for x in TriggerParams-%}{{-x-}} {%-endfor-%}]， 通知内容：{{NotifyMsg}}"
    locale  = "zh-CN"
  }
  ding_talk {
    title   = "修改-告警通知"
    content = "修改-尊敬的用户，您好！\n您的账号（主账户ID：{{AccountID}} ）的日志服务{%if NotifyType==1%}触发告警{%else%}告警恢复{%endif%}\n告警策略：{{Alarm}}\n告警日志主题：{{AlarmTopicName}}\n触发时间：{{StartTime}}\n触发条件：{{Condition}}\n当前查询结果：[{%-for x in TriggerParams-%}{{-x-}} {%-endfor-%}]\n通知内容：{{NotifyMsg|escapejs}}\n日志检索详情：[查看详情]({{QueryUrl}})\n告警详情：[查看详情]({{SignInUrl}})\n\n感谢对火山引擎的支持"
    locale  = "zh-CN"
  }
  email {
    subject = "修改-告警通知"
    content = "修改-告警策略：{{Alarm}}<br> 告警日志项目：{{ProjectName}}<br>"
    locale  = "zh-CN"
  }
  lark {
    title   = "修改-告警通知"
    content = "修改-尊敬的用户，您好！\n您的账号（主账户ID：{{AccountID}} ）的日志服务{%if NotifyType==1%}触发告警{%else%}告警恢复{%endif%}\n告警策略：{{Alarm}}\n告警日志主题：{{AlarmTopicName}}\n触发时间：{{StartTime}}\n触发条件：{{Condition}}\n当前查询结果：[{%-for x in TriggerParams-%}{{-x-}} {%-endfor-%}]\n通知内容：{{NotifyMsg|escapejs}}\n日志检索详情：[查看详情]({{QueryUrl}})\n告警详情：[查看详情]({{SignInUrl}})\n\n感谢对火山引擎的支持"
    locale  = "zh-CN"
  }
}
```
## Argument Reference
The following arguments are supported:
* `alarm_content_template_name` - (Required) The name of the alarm content template.
* `ding_talk` - (Optional) The ding_talk content of the alarm content template.
* `email` - (Optional) The email content of the alarm content template.
* `lark` - (Optional) The lark content of the alarm content template.
* `need_valid_content` - (Optional) Whether to validate the content template.
* `sms` - (Optional) The sms content of the alarm content template.
* `vms` - (Optional) The vms content of the alarm content template.
* `webhook` - (Optional) The webhook content of the alarm content template.
* `wechat` - (Optional) The wechat content of the alarm content template.

The `ding_talk` object supports the following:

* `content` - (Required) The content of the ding_talk content template.
* `locale` - (Required) The locale of the ding_talk content template.
* `title` - (Required) The title of the ding_talk content template.

The `email` object supports the following:

* `content` - (Required) The content of the email content template.
* `locale` - (Required) The locale of the email content template.
* `subject` - (Required) The subject of the email content template.

The `lark` object supports the following:

* `content` - (Required) The content of the lark content template.
* `locale` - (Required) The locale of the lark content template.
* `title` - (Required) The title of the lark content template.

The `sms` object supports the following:

* `content` - (Required) The content of the sms content template.
* `locale` - (Required) The locale of the sms content template.

The `vms` object supports the following:

* `content` - (Required) The content of the vms content template.
* `locale` - (Required) The locale of the vms content template.

The `webhook` object supports the following:

* `content` - (Required) The content of the webhook content template.

The `wechat` object supports the following:

* `content` - (Required) The content of the wechat content template.
* `locale` - (Required) The locale of the wechat content template.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `alarm_content_template_id` - The ID of the alarm content template.


## Import
tls alarm content template can be imported using the alarm_content_template_id, e.g.
```
$ terraform import volcengine_tls_alarm_content_template.default alarm-content-template-123456
```

