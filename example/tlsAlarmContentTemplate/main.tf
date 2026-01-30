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
