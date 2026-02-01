---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_alarm_webhook_integration"
sidebar_current: "docs-volcengine-resource-tls_alarm_webhook_integration"
description: |-
  Provides a resource to manage tls alarm webhook integration
---
# volcengine_tls_alarm_webhook_integration
Provides a resource to manage tls alarm webhook integration
## Example Usage
```hcl
resource "volcengine_tls_alarm_webhook_integration" "foo" {
  webhook_name   = "terraform-tf-webhook"
  webhook_url    = "http://zijie.com"
  webhook_type   = "lark"
  webhook_method = "PUT"
  webhook_secret = "your secret"
  webhook_headers {
    key   = "Content-Type"
    value = "application/json"
  }
}
```
## Argument Reference
The following arguments are supported:
* `webhook_name` - (Required) The name of the webhook integration.
* `webhook_type` - (Required) The type of the webhook integration.
* `webhook_url` - (Required) The URL of the webhook.
* `webhook_headers` - (Optional) The headers of the webhook.
* `webhook_method` - (Optional) The method of the webhook.
* `webhook_secret` - (Optional) The secret of the webhook.

The `webhook_headers` object supports the following:

* `key` - (Optional) The key of the header.
* `value` - (Optional) The value of the header.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `id` - ID of the resource.
* `create_time` - The creation time of the webhook integration.
* `modify_time` - The update time of the webhook integration.


## Import
tls alarm webhook integration can be imported using the alarm_webhook_integration_id, e.g.
```
$ terraform import volcengine_tls_alarm_webhook_integration.default alarm-webhook-integration-123456
```

