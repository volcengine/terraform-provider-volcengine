---
subcategory: "TLS"
layout: "volcengine"
page_title: "Volcengine: volcengine_tls_alarm_webhook_integrations"
sidebar_current: "docs-volcengine-datasource-tls_alarm_webhook_integrations"
description: |-
  Use this data source to query detailed information of tls alarm webhook integrations
---
# volcengine_tls_alarm_webhook_integrations
Use this data source to query detailed information of tls alarm webhook integrations
## Example Usage
```hcl
data "volcengine_tls_alarm_webhook_integrations" "foo" {
}
```
## Argument Reference
The following arguments are supported:
* `output_file` - (Optional) File name where to save data source results.
* `webhook_id` - (Optional) The ID of the alarm webhook integration.
* `webhook_name` - (Optional) The name of the webhook integration. Fuzzy matching is supported.
* `webhook_type` - (Optional) The type of the webhook integration.

## Attributes Reference
In addition to all arguments above, the following attributes are exported:
* `integrations` - The list of alarm webhook integrations.
    * `create_time` - The creation time of the webhook integration.
    * `modify_time` - The update time of the webhook integration.
    * `webhook_headers` - The headers of the webhook.
        * `key` - The key of the header.
        * `value` - The value of the header.
    * `webhook_id` - The ID of the alarm webhook integration.
    * `webhook_method` - The method of the webhook.
    * `webhook_name` - The name of the webhook integration.
    * `webhook_secret` - The secret of the webhook.
    * `webhook_type` - The type of the webhook.
    * `webhook_url` - The URL of the webhook.
* `total_count` - The total count of alarm webhook integrations.


